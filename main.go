package main

import (
	"fmt"
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/load", load)
	e.Logger.Fatal(e.Start(":80"))
}

func load(c echo.Context) error {
	client := resty.New()
	resp, err := client.R().
		Get("https://api.code.gov/repos?api_key=IiLkBh9KDgZEhlJ6JghDKjKqWKklAZemuMBfxNow")
	if err != nil {
		fmt.Println("Data retrieval error")
		return err
	}
	// fmt.Println("Body	:\n", resp)
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		fmt.Println(jsonparser.GET())
	})
	db, err := gorm.Open("mysql", "root:supersecret@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")
	defer db.Close()
	if err!=nil{
	log.Println(“Connection Failed to Open”)
	} 
	log.Println(“Connection Established”)
	return nil
}
