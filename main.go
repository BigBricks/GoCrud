package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&Bro{})
	r := gin.Default()
	r.GET("/", GetBros)

	r.Run(":8080")
}

//Bro comment for the linter to stop yelling
type Bro struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//dbCon
// func dbCon() *gorm.DB {
// 	var db, err = gorm.Open("sqlite3", "./gorm.db")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer db.Close()
// 	return db
// }

//GetBros func
func GetBros(c *gin.Context) {
	var bros []Bro
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.Find(&bros)
	if len(bros) <= 0 {
		c.AbortWithStatus(404)
		fmt.Println("L")
	} else {
		c.JSON(200, bros)
	}
}
