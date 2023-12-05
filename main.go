package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Licenseplate string `json:"licenseplate"`
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
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
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
			info := db.QueryRow("SELECT firstname FROM reservering WHERE licenseplate=? AND checkout >=?", licenseplate, currentDate)
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
	http.HandleFunc("/reservering", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		case http.MethodPost:
			checkin := r.URL.Query().Get("checkin")
			checkout := r.URL.Query().Get("checkout")
			userid := r.URL.Query().Get("userid")
			housenumber := r.URL.Query().Get("housenumber")

			if checkin != "" && checkout != "" && userid != "" && housenumber != "" {
				db.QueryRow("INSERT INTO `reserveringen`.`reservering` (`checkin`, `checkout`, `userid`, `housenumber`) VALUES (?, ?, ?, ?)", checkin, checkout, userid, housenumber)
				if err != nil {
					http.Error(w, "Database error", http.StatusInternalServerError)
					return
				}
				http.Error(w, "Reservation added", http.StatusOK)
			} else {
				http.Error(w, "Please enter checkin, checkout, userid and housenumber", http.StatusNotFound)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		case http.MethodPost:
			firstname := r.URL.Query().Get("firstname")
			lastname := r.URL.Query().Get("lastname")
			email := r.URL.Query().Get("email")
			password := r.URL.Query().Get("password")
			phonenumber := r.URL.Query().Get("phonenumber")
			postalcode := r.URL.Query().Get("postalcode")
			housenumber := r.URL.Query().Get("housenumber")
			street := r.URL.Query().Get("street")
			town := r.URL.Query().Get("town")
			country := r.URL.Query().Get("country")
			birthdate := r.URL.Query().Get("birthdate")
			licenseplate := r.URL.Query().Get("licenseplate")

			if firstname != "" && lastname != "" && email != "" && password != "" && postalcode != "" && housenumber != "" && street != "" && town != "" && country != "" && birthdate != "" && licenseplate != "" {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				if err != nil {
					http.Error(w, "Failed to hash password", http.StatusInternalServerError)
					return
				}
				info := db.QueryRow("SELECT * FROM user WHERE email=?", email)
				var data Data
				err = info.Scan(&data.Email)
				if err != nil {
					if err == sql.ErrNoRows {
						http.Error(w, "user with this email already excists", http.StatusConflict)
						return
					}
				}
				info = db.QueryRow("SELECT * FROM user WHERE licenseplate=?", licenseplate)
				err = info.Scan(&data.Licenseplate)
				if err != nil {
					if err == sql.ErrNoRows {
						http.Error(w, "user with this licenseplate already excists", http.StatusConflict)
						return
					}
				}
				db.QueryRow("INSERT INTO `reserveringen`.`user` (`firstname`, `lastname`, `email`, `password`, `phonenumber`, `postalcode`, `housenumber`, `street`, `town`, `country`, `birthdate`, `licenseplate`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", firstname, lastname, email, hashedPassword, phonenumber, postalcode, housenumber, street, town, country, birthdate, licenseplate)
				if err != nil {
					http.Error(w, "Database error", http.StatusInternalServerError)
					return
				}
				http.Error(w, "User added", http.StatusOK)
			} else {
				http.Error(w, "Please enter all data necessary", http.StatusNotFound)
			}

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		password := r.URL.Query().Get("password")
		if email != "" && password != "" {
			var hashedPassword string
			err = db.QueryRow("SELECT password FROM user WHERE email=?", email).Scan(&hashedPassword)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Error(w, "email or password not valid", http.StatusNotFound)
					return
				}
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}
			err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
			if err != nil {
				http.Error(w, "email or password not valid", http.StatusNotFound)
				return
			}
			http.Error(w, "Login succesful", http.StatusOK)
		} else {
			http.Error(w, "Please enter email and password", http.StatusNotFound)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
