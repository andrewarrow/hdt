package main

import "fmt"

//import "github.com/andrewarrow/hdt/client"
import "github.com/andrewarrow/hdt/server"

func main() {
	fmt.Println("starting")
	server.Start()
}
