package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Println("http handler called")
	resp, err := http.Get("http://dummyjson.com/products/1")
	//resp, err := http.Get("https://www.google.com")
	//resp, err := http.Get("https://www.facebook.com")
	if err != nil {
		fmt.Println("error while making http call internally", err)
		return nil
	}
	defer resp.Body.Close()
	ans, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error while reading the response body", err)
		return nil
	}
	if ans != nil {
		fmt.Println("body is sucessfullly received to the GinApp during internal httpCall", string(ans))
	} else {
		fmt.Println("no body recieved to ginAppp during internal httpCalll")
	}
	return nil
}
