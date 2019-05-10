package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Bro struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
}

func main() {
	db, _ := gorm.Open("sqlite3", "./gorm.db")
	defer db.Close()

	db.AutoMigrate(&Bro{})
	p1 := Bro{FirstName: "Bro"}
	db.Create(&p1)
	var p3 Bro
	db.First(&p3)
	fmt.Println(p1.FirstName)
	fmt.Println(p3.FirstName)
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World",
	// 	})
	// })
	// r.Run()
}
