package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

//Bro comment for the linter to stop yelling
type Bro struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	db, err = gorm.Open("sqlite3", "./gorm.db")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	db.AutoMigrate(&Bro{})

	r := gin.Default()

	r.GET("/bros/", GetBros)

	r.POST("/bros/", CreateBro)

	r.Run(":8080")
}

//GetBros Comment for linter
func GetBros(c *gin.Context) {
	var bros []Bro
	if err := db.Find(&bros).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, bros)
	}
}

//CreateBro here
func CreateBro(c *gin.Context) {
	var bro Bro
	c.BindJSON(&bro)
	db.Create(&bro)
	c.JSON(200, bro)
}
