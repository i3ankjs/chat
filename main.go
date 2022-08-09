package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createListener() (l net.Listener, close func()) {
	l, err := net.Listen("tcp", ":2222")
	if err != nil {
		panic(err)
	}
	return l, func() {
		_ = l.Close()
	}
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func OrderIdHandler(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf(id)
	c.JSON(http.StatusOK, books)
}

func main() {
	// db, err = gorm.Open("postgres", "host=gksoftware02.cgpbwmgdgyz2.ap-southeast-1.rds.amazonaws.com port=5432 user=root dbname=gk_dev_db sslmode=disable password=55k3fhGgsbKRMTvR")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	r := gin.New()
	r.GET("/order/:id", OrderIdHandler)
	http.ListenAndServe(":2222", r)
}
