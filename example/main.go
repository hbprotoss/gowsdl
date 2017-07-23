package example

import (
	"fmt"
	"wsdl2go/soap/client"
	"wsdl2go/example"
)

func main() {
	var service = example.NewHelloService(
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

	response2, err2 := service.HelloList([]string{"h1", "h2"})
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(response2)

	err3 := service.HelloMulti("haha", &example.Entity{Id: 123, Message: "1231231"}, []int32{1, 2}, []int32{1, 2})
	if err3 != nil {
		fmt.Println(err3)
		return
	}
}
