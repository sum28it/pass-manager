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
	ModifiedAt  string `json:"modifiedAt"`
}

var file *os.File

var (
	Dir string = "C:\\users\\user\\"
)

const (
	localDir string = "password-manager-data\\"
	dataFile string = "users.json"
	envFile  string = "keys.env"
)

// ===========================================================================================
// initializing application
func Init(secret string) (string, error) {

	// Check if app is already initialized
	if auth.IsInit(Dir + localDir + envFile) {
		return "", errors.New("app already initialized")
	}
	var err error

	// files directory holds the env file and passwords
	err = os.Mkdir(Dir+localDir, 0644)
	if err != nil {
		return "", err
	}
	// Creating pass.json and .env files
	file, err = os.Create(Dir + localDir + dataFile)
	if err != nil {
		return "", err
	}
	file.Close()
	// .env file holds the user secret and salt
	file, err = os.Create(Dir + localDir + envFile)
	if err != nil {
		return "", err
	}

	// Creating a salt to append with passwords
	rand.New(rand.NewSource(time.Now().Unix()))
	salt := []byte{}
	for i := 0; i < 16; i++ {
		salt = append(salt, byte(rand.Int()%256))
	}

	h := sha256.New()
	h.Write([]byte(secret))

	// Write the salt and
	file.WriteString(fmt.Sprintf("SALT=%x\n", string(salt)))
	file.WriteString(fmt.Sprintf("HASHED_SECRET=%x", string(h.Sum(nil))))
	file.Close()

	return Dir + localDir, nil
}
func Add(user *User, secret string) error {

	var users []User
	defer file.Close()
	if err := auth.Authenticate(secret, Dir+localDir+envFile); err != nil {
		return err
	}

	users, err := read(os.O_RDWR)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New("error reading file")
		}
	}

	// Encryption Here
	user.Password, err = encrypt(secret, user.Password)

	if err != nil {
		fmt.Println(err)
		return errors.New("error encrypting password")
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	user.ModifiedAt = now
	users = append(users, *user)
	err = write(users)
	if err != nil {
		return err
	}
	return nil
}

// Reads the users from the file
func read(mode int) ([]User, error) {

	var err error
	file, err = os.OpenFile(Dir+localDir+dataFile, mode, 0644)
	if err != nil {
		return nil, err
	}

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
	file.Truncate(0)
	file.Seek(0, 0)
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func Reset(secret string) error {

	// Authenticate user before resetting
	err := auth.Authenticate(secret, Dir+localDir+envFile)
	if err != nil {
		return err
	}

	err = os.RemoveAll(Dir + localDir)
	if err != nil {
		return errors.New("error deleting files")
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
