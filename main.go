package main

import (
	"fmt"
	"github.com/laioncorcino/students/database"
	"github.com/laioncorcino/students/route"
)

func main() {
	fmt.Println("Iniciando servidor Rest em Go ..")
	database.ConnectDatabase()
	route.HandleRequest()
}
