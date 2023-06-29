// Package user defines the user type and provides access to the stored data

package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

//CURRENT TASK: Refactor to support linux path names

type User struct {
	App         string `json:"app"`
	UserId      string `json:"userid"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Description string `json:"description"`
	ModifiedAt  string `json:"modifiedAt"`
}

var Dir string

func init() {
	Dir, _ = os.UserConfigDir()
	Dir = Dir + "\\"
}

const (
	localDir string = "password-manager-data\\"
	dataFile string = "users.json"
	envFile  string = "keys.env"
)

func (u User) PrintLong() string {
	return fmt.Sprintf("{\nApp: %s\nUserId: %s\nEmail: %s\nPassword: %s\nDescription: %s\nModifiedAt: %s\n}", u.App, u.UserId, u.Email, u.Password, u.Description, u.ModifiedAt)
}

func (u User) Print() string {
	return fmt.Sprintf("App: %s\tPassword: %s", u.App, u.Password)
}

// Reads the users from the file
func read(mode int) ([]User, error) {

	var file *os.File
	var err error
	file, err = os.OpenFile(Dir+localDir+dataFile, mode, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var users []User
	err = json.Unmarshal(data, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Writing a new user data to file
func write(users []User) error {

	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(Dir+localDir+dataFile, os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("error opening file for writing")
	}
	defer file.Close()

	file.Truncate(0)
	file.Seek(0, 0)
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) match(user User) bool {

	if (u.App != "") && (u.App != user.App) {
		return false
	}
	if (u.Email != "") && (u.Email != user.Email) {
		return false
	}
	if (u.UserId != "") && (u.UserId != user.UserId) {
		return false
	}

	return true
}
