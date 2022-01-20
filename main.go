package main

import (
	"controllers"
	"fmt"
	"models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
