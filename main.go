package main

import (
	"hacktiv8-go/router"
)

func main() {
	router.NewRouter().Run(":8080")
}
