package main

import (
	"database/sql"
)

type Product struct {
	Id       int
	Name     string
	Category int
	Price    float64
}

func queryDatabase(db *sql.DB) []Product {

	procucts := []Product{}

	rows, err := db.Query("SELECT * FROM Products")

	if err == nil {

		// fmt.Printf("Rows: %#v \n", rows)

		for rows.Next() {

			p := Product{}

			scanErr := rows.Scan(&p.Id, &p.Name, &p.Category, &p.Price)

			if scanErr == nil {

				procucts = append(procucts, p)

			}

		}

	} else {
		Printfln("Error: %v", err)

	}

	return procucts

}

func main() {

	// listDrivers()

	db, err := openDatabase()

	if err == nil {
		products := queryDatabase(db)

		for i, p := range products {

			Printfln("#%v: %v", i, p)
		}

		db.Close()

	} else {
		panic(err)

	}

}
