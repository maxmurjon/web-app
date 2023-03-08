package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Instagram string `json:"instagram"`
	Facebook  string `json:"facebook"`
}

func main() {
	jsonFile, err := os.Open("user.json")

	defer jsonFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File ochildi")

	var users Users

	byteValue, err := ioutil.ReadAll(jsonFile)
	
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteValue, &users)
	for _,u:=range users.Users{
		PirntUser(&u)
	}
}
func PirntUser(u *User){
	fmt.Printf("Name: %s\n",u.Name)
	fmt.Printf("Type: %s\n",u.Type)
	fmt.Printf("Age: %d\n",u.Age)
	fmt.Printf("Social ins: %s and FB %s\n",u.Social.Instagram,u.Social.Facebook)
}


