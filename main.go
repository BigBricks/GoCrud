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
	//GET ALL
	r.GET("/bros/", GetBros)
	//GET BY ID
	r.GET("/bros/:id", GetBroID)
	//CREATE
	r.POST("/bros/", CreateBro)
	//UPDATE
	r.PUT("/bros/:id", UpdateBro)
	//DELETE
	r.DELETE("/bros/:id", DeleteBro)
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

//GetBroID self explanatory title
func GetBroID(c *gin.Context) {
	id := c.Params.ByName("id")
	var bro Bro
	if err := db.Where("id=?", id).First(&bro).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, bro)
	}
}

//CreateBro here
func CreateBro(c *gin.Context) {
	var bro Bro
	c.BindJSON(&bro)
	db.Create(&bro)
	c.JSON(200, bro)
}

//UpdateBro here
func UpdateBro(c *gin.Context) {
	id := c.Params.ByName("id")
	var bro Bro
	if err := db.Where("id=?", id).First(&bro).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&bro)
	db.Save(&bro)
	c.JSON(200, bro)
}

//DeleteBro here
func DeleteBro(c *gin.Context) {
	id := c.Params.ByName("id")
	var bro Bro
	x := db.Where("id=?", id).Delete(&bro)
	fmt.Println(x)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
