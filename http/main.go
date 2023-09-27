package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Recover())

	// Routes
	e.GET("/pinghttp", httpHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}
func httpHandler(c echo.Context) error {
	//resp, err := http.Get("https://dummyjson.com/products")
	//resp, err := http.Get("https://www.google.com/")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("error encountered ", err)
		return err
	}
	resp.Body.Close()
	fmt.Println("the response body was", resp.Body)
	return nil
}
