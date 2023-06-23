package user

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

type User struct {
	App         string `json:"app"`
	UserId      string `json:"userid"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

var file *os.File
var users []User

func Init() {
	var err error
	os.Mkdir("files", 0644)
	file, err = os.Create("files/pass.json")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

// Returns user data matching to user
func Get(user *User) (*User, error) {
	defer file.Close()
	err := read(os.O_RDONLY)
	if err != nil {
		return nil, err
	}

	for _, val := range users {
		if user.App == val.App {
			return &val, nil
		}
	}

	return nil, errors.New("no such user exists")

}

func Add(user *User) error {

	err := read(os.O_RDWR)
	if err != nil {
		return err
	}

	users = append(users, *user)
	err = write()
	if err != nil {
		return err
	}
	return nil
}

// Reads the file and populates users
func read(mode int) error {

	var err error
	file, err = os.OpenFile("files/pass.json", mode, 0644)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &users)

	if err != nil {
		return err
	}

	return nil
}

func write() error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	file.Truncate(0)
	file.Seek(0, 0)
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
