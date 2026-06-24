package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var users = map[string]User{}

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

func SignUp(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	password := r.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintln(w, "error")
		return
	}

	users[name] = User{
		Name:     name,
		Password: string(hashedPassword),
	}

	fmt.Fprintln(w, "User Created")

}

func SignIn(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")

	user, ok := users[name]
	if !ok {
		fmt.Fprintln(w, "User Not Found")
		return
	}

	password := r.FormValue("password")
	if bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password)) == nil {
		fmt.Fprintln(w, "Login Successful")
		return
	} else {
		fmt.Println("Wrong Password")
		return

	}

}

func main() {

	http.HandleFunc("/signup", SignUp)
	http.HandleFunc("/signin", SignIn)

	http.ListenAndServe(":8080", nil)
}
