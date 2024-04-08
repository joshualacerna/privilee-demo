package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		out, err := exec.Command("/bin/sh", "zone.sh").Output()
		// return fmt.Println(string(output))
		if err != nil {
			log.Fatal(err)
		}
		if len(out) == 0 {
			log.Fatal("no output was returned from script")
		}
		return c.String(http.StatusOK, "Hello, Privilee!!! \nfrom region "+string(out))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
