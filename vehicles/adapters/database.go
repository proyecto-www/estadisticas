package database

import (
	"database/sql" // add this
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq" // <------------ here
)

var (
	host     = os.Getenv("RDS_HOST")
	port     = os.Getenv("RDS_PORT")
	user     = os.Getenv("RDS_USER")
	password = os.Getenv("RDS_PASSWORD")
	dbname   = os.Getenv("RDS_DATABASE")
)

type Vehiculos struct {
	Lista []Vehiculo
}

type Vehiculo struct {
	Placa            string
	Fechahoraentrada int
	Tipodevehiculo   string
}

func GetAllVehicles() []Vehiculo {

	port, err := strconv.Atoi(port)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	res := Vehiculo{}
	var cosas []Vehiculo
	rows, err := db.Query("SELECT placa, fechahoraentrada,tipodevehiculo FROM vehiculo where sigueenuso=true")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		rows.Scan(
			&res.Placa,
			&res.Fechahoraentrada,
			&res.Tipodevehiculo,
		)
		cosas = append(cosas, res)
	}

	return cosas
}
