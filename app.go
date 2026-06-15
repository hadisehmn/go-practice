
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)




type User struct {
	Name     string
	email    string
	Password string
}

func (u *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) Login() bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// if username != u.Name {
	// 	fmt.Println("User not found")
	// 	return false
	// }

	if bcrypt.CompareHashAndPassword(
		[]byte(u.Password),
		[]byte(password),
	) == nil {
		fmt.Println("Login successful")
		return true
	} else {
		fmt.Println("Wrong password")
		return false
	}
}

type Todo struct {
	Title      string
	ID         int
	InProgress bool
	Done       bool
}

type TodoList struct {
	Data []Todo
}

type Todos interface {
	Add(title string)
	Remove(id int)
	Edit(id int, title string)
}

func (t *Todo) Status() string {
	if t.Done {
		return "done"

	} else if t.InProgress {
		return " inprogress"

	} else {
		return "no status"
	}
}

func (t *TodoList) Add(title string) {
	t.Data = append(t.Data, Todo{
		ID:    len(t.Data),
		Title: title,
	})
}

func (t *TodoList) Remove(id int) {
	for i := range t.Data {
		if t.Data[i].ID == id {
			t.Data = append(t.Data[:i], t.Data[i+1:]...)
			fmt.Println(t.Data)

			return
		}
	}
}

func (t *TodoList) Edit(ID int, NewTitle string) {
	for i := range t.Data {
		if t.Data[i].ID == ID {
			t.Data[i].Title = NewTitle
			return

		}

	}
}

func main() {

	user := User{
		Name: "alireza",
	}

	err := user.HashPassword("1234")
	if err != nil {
		return
	}
	if !user.Login() {
		return
	}

	// fmt.Println("Hashed password:", user.Password)
	// user.Login()

	list := TodoList{}

	list.Add("GO GYM")
	list.Add("drink water")
	list.Add("take shower ")

	list.Edit(2, "studying go ")

	fmt.Println(list.Data)

	fmt.Println("Before remove:", list.Data)

	list.Remove(1)

	fmt.Println("After remove:", list.Data)

	list.Data[0].InProgress = true
	list.Data[1].Done = true
	for _, t := range list.Data {
		fmt.Println(t.Status())
	}

}
