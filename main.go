package main

import (
	"fmt"
	"gowsdl/soap/client"
	"gowsdl/temp"
)

func main() {
	var service = temp.NewHelloService(
		"http://127.0.0.1:8080/ws/hello",
		&client.SecurityAuth{
			Username: "server",
			Password: "serverpass",
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
