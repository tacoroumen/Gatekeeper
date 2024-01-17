# Uitleg code Applicatie laag 
  
## In dit Bestand zal ik bij elke code een stukje informatie geven over wat deze code doet en waarvoor het gebruikt wordt.  
  
### **[Gatekeeper/main.go](gatekeeper/main.go)**  
Er wordt een Config struct gedefinieerd om configuratiegegevens op te slaan.  
Er is ook een Payload struct gedefinieerd om gegevens voor een HTTP-verzoek te bevatten.  
  
1. Het programma begint met het openen van logbestanden voor fouten en toegang.  
2. Het controleert of er een kentekenplaat is opgegeven als argument bij het uitvoeren van het programma. Zo niet, wordt het programma beëindigd.  
3. Het leest de configuratiegegevens uit een JSON-bestand en decodeert deze in een slice van Config structs.  
4. Het maakt een HTTP-verzoek voor elke configuratie en voegt het kentekenplaat als een queryparameter toe.  
5. Het verwerkt de JSON-respons van het verzoek en krijgt de naam van de gebruiker.
6. Als er geen gebruikersnaam is gevonden, wordt het programma beëindigd.  
7. Het schrijft toegangsinformatie naar het logbestand en stuurt een HTTP-verzoek naar een API met behulp van de geconfigureerde URL en authenticatiegegevens.  
8. Het controleert de huidige tijd en geeft berichten weer op basis van de configuratiegegevens.  
9.Het programma wordt herhaald voor elke configuratie in de slice.  
  
Kortom, deze code is een programma die toegang verleent tot een parkeerplaats op basis van een kentekenplaat, interactie heeft met een externe API en logboeken bijhoudt.  
    
### **[Gatekeeper/gatekeeper.py](gatekeeper/gatekeeper.py)**  
De code importeert de vereiste bibliotheken: cv2, pytesseract en subprocess.  
  
1. Er wordt een videostream van de webcam gestart met behulp van de cv2.VideoCapture(0)-functie.  
2. Er wordt een oneindige lus gestart om continu frames van de videostream te lezen.  
3. Voor elk frame wordt het volgende gedaan:  
    &nbsp;&nbsp;&nbsp;&nbsp;3.01 Het huidige frame wordt gelezen met behulp van cap.read().  
    &nbsp;&nbsp;&nbsp;&nbsp;3.02 Het frame wordt omgezet naar grijswaarden met behulp van cv2.cvtColor().  
    &nbsp;&nbsp;&nbsp;&nbsp;3.03 Er wordt een Gaussische vervaging toegepast op het grijswaardenbeeld met cv2.GaussianBlur().  
    &nbsp;&nbsp;&nbsp;&nbsp;3.04 Randen worden gedetecteerd met behulp van de Canny-randdetectiemethode met cv2.Canny().  
    &nbsp;&nbsp;&nbsp;&nbsp;3.05 Contouren worden gevonden in de randafbeelding met behulp van cv2.findContours().  
    &nbsp;&nbsp;&nbsp;&nbsp;3.06 De contouren worden gefilterd op basis van grootte om mogelijke kentekenplaten te vinden.  
    &nbsp;&nbsp;&nbsp;&nbsp;3.07 Bounding boxen worden getekend rond de gevonden kentekenplaten met behulp van cv2.rectangle().  
    &nbsp;&nbsp;&nbsp;&nbsp;3.08 PyTesseract wordt gebruikt voor optische tekenherkenning (OCR) op de individuele kentekenplaatgebieden om de tekens op het kenteken te lezen.  
    &nbsp;&nbsp;&nbsp;&nbsp;3.09 Als een kenteken precies 6 tekens bevat, wordt een terminalopdracht uitgevoerd met het kenteken als argument.  
    &nbsp;&nbsp;&nbsp;&nbsp;3.10 Het bewerkte frame wordt weergegeven in een venster met cv2.imshow().  
    &nbsp;&nbsp;&nbsp;&nbsp;3.11 Als de 'q'-toets wordt ingedrukt, wordt de lus beëindigd.  
4. Nadat de lus is beëindigd, worden de video-opnamebron vrijgegeven en de vensters gesloten met cap.release() en cv2.destroyAllWindows(). 
   
In het kort, deze code maakt gebruik van computer vision-technieken om kentekenplaten te detecteren en te lezen van een videostream van de webcam, en voert vervolgens een terminalopdracht uit met het gedetecteerde kenteken als argument.  
  
### **[config.json](gatekeeper/config.json)**  
`"Morning_start_time:"` De starttijd van de ochtendperiode (7 uur).  
`"Noon_start_time:"` De starttijd van de middagperiode (12 uur).  
`"Evening_start_time:"` De starttijd van de avondperiode (18 uur).  
`"No_parking_acces_start_time:"` De starttijd waarop geen toegang tot de parkeerfaciliteiten is toegestaan (23 uur).    
`"Morning_message:"` Het bericht dat wordt weergegeven tijdens de ochtendperiode ("Good morning").  
`"Noon_message:"` Het bericht dat wordt weergegeven tijdens de middagperiode ("Good afternoon").  
`"Evening_message:"` Het bericht dat wordt weergegeven tijdens de avondperiode ("Good evening").  
`"No_parking_acces_message:"` Het bericht dat wordt weergegeven wanneer de parkeerfaciliteiten gesloten zijn ("Sorry, our parking facilities are currently closed.").  
`"Technical_dificulties:"` Het bericht dat wordt weergegeven wanneer er technische problemen zijn ("Sorry, we are currently experiencing technical difficulties").  
`"Welcome_message:"` Het welkomstbericht dat wordt weergegeven ("Welcome at Holiday parks.").  
`"Not_allowed:"` Het bericht dat wordt weergegeven wanneer een kentekenplaat niet is toegestaan ("License plate not permitted!").    
  
Deze configuratiegegevens worden gebruikt in de code om berichten weer te geven en communicatie met de API uit te voeren.  
  
### **[ESP Gatekeeper](esphome32/gatekeeper.yaml)**  
Deze code is een configuratiebestand voor het ESPHome-platform, waarmee je ESP32-microcontrollers kunt programmeren en configureren.  
  
`esphome:` Dit geeft het begin van de ESPHome-configuratie aan en definieert de naam van het apparaat (gatekeeper).  
`logger:` Hier wordt de logging geconfigureerd. Met level: VERY_VERBOSE worden zeer gedetailleerde logberichten ingeschakeld.  
`esp32:` Hier worden de instellingen voor de ESP32-microcontroller gespecificeerd, zoals het gebruikte board (esp32doit-devkit-v1) en het framework (arduino).  
`wifi:` Dit configureert de wifi-instellingen voor het apparaat, waarbij de ssid (netwerknaam) en het wachtwoord worden gelezen uit geheime waarden (secrets).  
`ota:` Dit stelt over-the-air (OTA) updates in en specificeert het wachtwoord voor beveiligde updates, ook gelezen uit geheime waarden.  
`servo:` Hier wordt een servo-motor geconfigureerd met een ID (my_servo) en een uitvoerkanaal (pwm_out).  
`output:` Dit definieert de uitvoerinstellingen voor het apparaat, in dit geval een ledc-uitvoer (pwm_out) op pin 25 met een frequentie van 50 Hz.  
`sensor:` Hier wordt een sensor geconfigureerd om de uptime (tijd dat het apparaat actief is) bij te houden.  
`web_server:` Dit configureert een webserver op poort 80 en stelt authenticatie in met een gebruikersnaam en wachtwoord, gelezen uit geheime waarden. Ook word met de web_server de ESPHome RestAPI aangezet.  
`switch:` Hier worden schakelaars (actuatoren) geconfigureerd. Er is een GPIO-schakelaar (gate) op pin 26 die wordt gebruikt om een poort te activeren. Er zijn twee acties gedefinieerd voor het in- en uitschakelen van de poort, waarbij ook een servo-motor wordt aangestuurd en een vertraging wordt toegepast. Er is ook een herstartschakelaar (reboot) geconfigureerd.  

Kort samengevat, deze code configureert een ESP32-apparaat met wifi-connectiviteit, een servo-motor, een sensor voor het bijhouden van de uptime en een api met authenticatie. Er zijn ook schakelaars geconfigureerd om een poort te activeren en het apparaat te herstarten.  
  
### **[Gatekeeper API](gatekeeper%20API/main.go)**  
Deze code is een HTTP-server die luistert op poort 8080 en communiceert met een MySQL-database. Hier is een uitleg van de code:
  
`package main:` Dit geeft aan dat dit het hoofdpakket van de Go-code is.  
`import (...):` Hier worden de vereiste pakketten geïmporteerd voor de code, waaronder de MySQL-driver.  
`type Data struct { ... }:` Dit definieert een struct genaamd "Data" met twee velden: "Naam" en "Checkout". Deze struct wordt gebruikt om gegevens te coderen en te decoderen in JSON-indeling.  
`func main() { ... }:` Dit is de hoofdfunctie van het programma.  
`currentTime := time.Now():` Dit verkrijgt de huidige tijd.  
`currentDate := currentTime.Format("2006-01-02"):` Dit formatteert de huidige tijd als een datum in het formaat "YYYY-MM-DD".  
`db, err := sql.Open("mysql", "Fonteyn:P@ssword@tcp(reserveringen.mysql.database.azure.com:3306)/klanten?tls=true"):` Dit opent een verbinding met de MySQL-database. De gebruikersnaam, wachtwoord, host en databasegegevens worden opgegeven in de verbindingsreeks.  
`defer db.Close():` Dit zorgt ervoor dat de databaseverbinding wordt gesloten wanneer de functie eindigt.  
`http.HandleFunc("/nummerplaat", func(w http.ResponseWriter, r *http.Request) { ... }):` Dit definieert een HTTP-handler voor het pad "/nummerplaat". Het verwerkt inkomende HTTP-verzoeken naar dit pad.  
`licenseplate := r.URL.Query().Get("licenseplate"):` Dit haalt de queryparameter "licenseplate" op uit het HTTP-verzoek.  
`if licenseplate != "" { ... }:` Dit controleert of er een licentieplaatwaarde is opgegeven in het verzoek.  
`info := db.QueryRow("SELECT name, checkout FROM reservering WHERE kenteken=? AND checkout >=?", licenseplate, currentDate):` Dit voert een SQL-query uit op de database om informatie op te halen op basis van de licentieplaatwaarde en de huidige datum.  
`var data Data:` Dit declareert een variabele "data" van het type "Data".  
`err = info.Scan(&data.Naam, &data.Checkout):` Dit scant de queryresultaten en vult de velden van de "data" struct met de corresponderende waarden.  
`w.Header().Set("Content-Type", "application/json"):` Dit stelt de HTTP-header in om aan te geven dat het antwoord een JSON-indeling heeft.  
`json.NewEncoder(w).Encode(data):` Dit encodeert de "data" struct naar JSON en schrijft het antwoord terug naar de HTTP-respons.  
`log.Fatal(http.ListenAndServe(":8080", nil)):` Dit start de HTTP-server op poort 8080 en logt een fout als de server niet kan worden gestart. 
   
Kort samengevat, deze code implementeert een HTTP-server in Go die verbinding maakt met een MySQL-database en informatie opvraagt op basis van een opgegeven licentieplaatwaarde. Het retourneert de gevonden gegevens in JSON-indeling   

### **[Gatekeeper API/dockerfile](gatekeeper%20API/dockerfile)**  
Deze code is een Dockerfile, een configuratiebestand dat wordt gebruikt om een Docker-container te bouwen.    
  
`FROM golang:alpine:` Dit geeft aan dat we de officiële Golang Docker-image willen gebruiken als basis voor onze container. Specifiek wordt de "alpine" versie van de image gebruikt, wat een lichtgewicht versie is gebaseerd op het Alpine Linux-besturingssysteem.  
`ADD . /app:` Dit voegt alle bestanden en mappen in de huidige directory toe aan de map "/app" in de container. Dit omvat waarschijnlijk de Go-code die we willen uitvoeren.  
`WORKDIR /app:` Dit stelt de werkdirectory in voor alle volgende instructies in de Dockerfile op "/app". Dit betekent dat alle verdere opdrachten zullen plaatsvinden binnen de map "/app" in de container.  
`RUN go mod download:` Dit voert het commando "go mod download" uit binnen de container. Dit commando downloadt alle afhankelijkheden die zijn gespecificeerd in het go.mod-bestand van de Go-code.  
`RUN go build -o main .:` Dit voert het commando "go build -o main ." uit binnen de container. Dit bouwt de Go-applicatie en produceert een uitvoerbaar bestand genaamd "main".  
`EXPOSE 8080:` Dit geeft aan dat de container luistert op poort 8080. Het is belangrijk op te merken dat dit alleen een documentatiefunctie heeft en de werkelijke poortkoppeling moet worden gedaan bij het uitvoeren van de container.  
`CMD ["/app/main"]:` Dit stelt het standaardcommando in dat wordt uitgevoerd wanneer de container wordt gestart. In dit geval wordt het uitvoerbare bestand "main" uitgevoerd binnen de "/app" map.  
  
Samengevat, deze Dockerfile bouwt een Docker-container op basis van de Golang-image. Het voegt de Go-code toe aan de container, installeert de afhankelijkheden, bouwt de applicatie en stelt de container in om de applicatie uit te voeren wanneer deze wordt gestart op poort 8080.
