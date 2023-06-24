package user

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/sum28it/pass-manager/pkg/auth"
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

// ===========================================================================================
// initializing application
func Init(secret string) error {

	var err error

	// files directory holds the env file and passwords
	os.Mkdir("files", 0644)

	// Creating pass.json and .env files
	file, err = os.Create("files/pass.json")
	if err != nil {
		return err
	}
	file.Close()
	// .env file holds the user secret and salt
	file, err = os.Create("files/.env")
	if err != nil {
		return err
	}

	// Creating a salt to append with passwords
	rand.New(rand.NewSource(time.Now().Unix()))
	salt := []byte{}
	for i := 0; i < 16; i++ {
		salt = append(salt, byte(rand.Int()%256))
	}

	// Generate a new key from the secret and salt to be used for encryption
	// key := pbkdf2.Key([]byte(secret), salt, 4096, 32, sha256.New)

	h := sha256.New()
	h.Write([]byte(secret))

	// Write the salt and
	file.WriteString(fmt.Sprintf("SALT=%x\n", string(salt)))
	file.WriteString(fmt.Sprintf("HASHED_SECRET=%x", string(h.Sum(nil))))

	return nil
}

// ==========================================================================================
// Returns user data matching to user
func Get(user *User, secret string) (*User, error) {

	if err := auth.Authenticate(secret); err != nil {
		return nil, err
	}

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

func Add(user *User, secret string) error {

	if err := auth.Authenticate(secret); err != nil {
		return err
	}

	err := read(os.O_RDWR)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return err
		}
	}

	users = append(users, *user)
	err = write()
	if err != nil {
		return err
	}
	return nil
}

// Reads the users from the file
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

// Writing a new user data to file
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
