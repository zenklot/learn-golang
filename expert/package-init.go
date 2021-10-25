package main

import (
	"fmt"
	"learn-golang/expert/database"
	_ "learn-golang/expert/helper" // Blank Identifier
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
