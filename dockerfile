# Gebruik een bestaande Go-afbeelding als basis
FROM golang:alpine

# Voeg de code toe aan de werkmap /app in de container
ADD . /app

# Stel de werkmap in als werkdirectory
WORKDIR /app

# Installeer afhankelijkheden
RUN go mod download

# Bouw de applicatie binnen de container
RUN go build -o main .

# Expose de poort waarnaar wordt geluisterd
EXPOSE 8080

# Start de applicatie als de container wordt gestart
CMD ["/app/main"]
