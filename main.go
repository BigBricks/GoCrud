package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Bro comment for the linter to stop yelling
type Bro struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Bro{})
	b1 := Bro{FirstName: "Bro", LastName: "Cephus"}
	db.Create(&b1)
	r := gin.Default()
	r.GET("/", GetBros)
	r.Run(":8080")
}

//GetBros func
func GetBros(c *gin.Context) {
	var bros []Bro
	if len(bros) <= 0 {
		c.AbortWithStatus(404)
		fmt.Println("No Bros Found, Sorry Bro")
	} else {
		c.JSON(200, bros)
	}
}
