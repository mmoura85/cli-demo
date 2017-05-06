package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type PhoneBook struct {
	Phone string `json: "phone"`
	Mob   string `json: "mob"`
}

type PhoneBookListing map[string]*PhoneBook

var Pbl PhoneBookListing

func init() {

	err := Pbl.Load("./config.json")
	if err != nil {
		fmt.Println(err)
	}
}

func (pb *PhoneBookListing) Load(src string) error {

	f, err := os.OpenFile(src, os.O_RDWR|os.O_CREATE, 755)

	if err != nil {
		return err
	}

	defer f.Close()

	jsonDecoder := json.NewDecoder(f)
	err = jsonDecoder.Decode(pb)
	if err != nil {
		return err
	}
	return nil
}

func (pb *PhoneBookListing) Save() error {

	f, err := os.OpenFile("config.json", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 755)

	if err != nil {
		return err
	}

	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "\t\t")
	err = encoder.Encode(pb)

	if err != nil {
		return err
	}

	return nil
}

func (pb *PhoneBookListing) Add(name string, phoneDetails *PhoneBook) error {

	_, ok := Pbl[name]

	if ok {
		return errors.New(fmt.Sprintf("Error: Name %s already exists", name))
	}

	Pbl[name] = phoneDetails

	err := Pbl.Save()
	if err != nil {
		return err
	}

	return nil
}

func (pb *PhoneBookListing) Edit(name string, phoneDetails *PhoneBook) error {

	_, ok := Pbl[name]

	if !ok {
		return errors.New(fmt.Sprintf("Error: Name %s does not exist", name))
	}

	Pbl[name] = phoneDetails

	err := Pbl.Save()

	if err != nil {
		return err
	}

	return nil
}

func (pb *PhoneBookListing) Remove(name string) error {

	_, ok := Pbl[name]

	if !ok {
		errors.New("Error! unable to delete record as it does not exist")
	}
	delete(Pbl, name)

	err := Pbl.Save()
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%s has been deleted", name))

	return nil
}
