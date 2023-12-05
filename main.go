package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	//"time"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	Name string `json:"name"`
}

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

func getconfig() (string, string, string, string, string) {
	// Read the content of the aconfig.json file
	data, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config.json:", err)
		return "", "", "", "", ""
	}

	// Parse the JSON data into the Config struct
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return "", "", "", "", ""
	}
	return config.Username, config.Password, config.Ip, config.Port, config.Database
}

func main() {
	//currentTime := time.Now()
	//currentDate := currentTime.Format("2006-01-02")
	username, password, ip, port, database := getconfig()
	connectionstring := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=true", username, password, ip, port, database)
	db, err := sql.Open("mysql", connectionstring)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/licenseplate", func(w http.ResponseWriter, r *http.Request) {
		licenseplate := r.URL.Query().Get("licenseplate")
		if licenseplate != "" {
			info := db.QueryRow("SELECT firstname FROM reservering WHERE licenseplate=?", licenseplate)
			var data Data
			err = info.Scan(&data.Name)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Error(w, "licenseplate or date not valid", http.StatusNotFound)
					return
				}
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
			return
		} else {
			http.Error(w, "Please enter an licenseplate", http.StatusNotFound)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
