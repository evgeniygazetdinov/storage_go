package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var Albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type CustomUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     int    `json:"email"`
	IsStuff   string `json:"is_stuff"`
}
type BaseUser struct {
	ID   int
	Name string
	Age  int
}

func GetUser(id int) []BaseUser {

	db, err := sql.Open("mysql", "example"+":"+"secret2"+"@tcp(127.0.0.1:3306)/"+"stage")

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()
	SimpleUser, err := db.Query("SELECT * FROM users where id=?", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	products := []BaseUser{}
	for SimpleUser.Next() {
		var prod BaseUser
		// for each row, scan into the Product struct
		err = SimpleUser.Scan(&prod.ID, &prod.Name, &prod.Age)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the product into products array
		products = append(products, prod)
	}

	return products

}
