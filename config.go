package main

import (
	"errors"
)

type PhoneBook struct {
	Name  string `json: "name"`
	Phone int    `json: "phone"`
	Mob   int    `json: "mob"`
}

func init() errors {

}

func (pb *PhoneBook) Load(src) errors {

}

func (pb *PhoneBook) Save() errors {

}
