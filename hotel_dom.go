package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string
	Password string
	Age      int
	Phone    int
	ID       int
}

type Hotel struct {
	HotelName string
	RoomName  string
	RoomType  string
	ID        int
	Star      int
	Price     float64
}

type HotelService interface {
	AddRoom(room Hotel) error
	RemoveRoom(id int) error
	UpdateRoom(room Hotel) error
	ListRooms() ([]Hotel, error)
}

type UserService interface {
	SelectRoom(id int) (Hotel, error)
	RemoveRoom(id int) error
	BookRoom(id int) error
	ListRooms() ([]Hotel, error)
}

func (u *User) SignUp(Password string, Name string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	u.Name = Name

}

func (u *User) SingIn(Password string, Name string) bool {
	if u.Name != Name {
		if u.Name == Name {
			fmt.Println("Wrong name")
			return false

		}

	}

	if bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(Password)) == nil {
		fmt.Println("Login Successful")
		return true
	} else {
		fmt.Println("Wrong Password")
		return false

	}

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Your Name ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Println("Enter Your Phonenumber ")
	phonenumber, _ := reader.ReadString('\n')
	phonenumber = strings.TrimSpace(phonenumber)
	phone, _ := strconv.Atoi(phonenumber)

	fmt.Println("Enter Your Password ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user := User{
		Name:     username,
		Phone:    phone,
		Password: password,
	}
	fmt.Printf("User created: Name: %s | Phone: %d | Password: %s\n", user.Name, user.Phone, user.Password)
}
