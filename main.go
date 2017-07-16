package main

import (
	"fmt"
	"gowsdl/enterprise"
	"gowsdl/soap/client"
)

func main() {
	var service = enterprise.NewHelloService(
		"http://127.0.0.1:8080/ws/hello",
		&client.SecurityAuth{
			Username: "client",
			Password: "GT666lucknumber",
			Type:     "PasswordText",
		},
	)
	response, err := service.Hello("haha")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
