package main

import "fmt"

func main() {
	fmt.Println("........Golang Microservice with........")

	server := NewAPIServer(":3000")
	server.Run()
}
