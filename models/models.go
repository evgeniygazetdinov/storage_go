package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func GetUser() []CustomUser {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()
	results, err := db.Exec("SELECT * FROM product")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	users := []CustomUser{}
	for results.Next() {
		var user CustomUser
		// for each row, scan into the Product struct
		err = results.Scan(&user.FirstName, &user.LastName, &user.Email, &user.IsStuff)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the product into products array
		users = append(users, user)
	}

	return users
}
