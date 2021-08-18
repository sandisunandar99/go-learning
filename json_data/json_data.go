package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	fulname string `json:"name"`
	age     int    `json:"age"`
}

func main() {
	var jsonString = `{"name":"John","age":30}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Username :", data.fulname)
	fmt.Println("Age :", data.age)
}
