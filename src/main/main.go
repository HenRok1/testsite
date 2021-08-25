package main

import (
	"fmt"
	"main/router"
)

//main func
func main(){
	fmt.Println("Welcome to the server!")

	e:= router.New()

	e.Start(":8000")
}

