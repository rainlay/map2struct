package main

import (
	"fmt"
	"github.com/rainlay/map2struct"
)

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	m := make(map[string]string, 0)
	m["id"] = "1"
	m["email"] = "foo@bar.com"
	m["phone"] = "+886565656566"
	var user User
	err := map2struct.DecodeSs(&user, m)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
}
