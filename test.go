package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
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

func (u *User) Login(password string) bool {

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

// Method 1 :UUID /  id = string
// ID         string
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

//Method 1 : UUID
// Add creates a new Todo item and appends it to the list.
// It generates a unique UUID as the ID.
// func (t *TodoList) Add(title string) {
// 	todo := Todo{
// 		ID:    uuid.New().String(),
// 		Title: title,
// 	}

// 	t.Data = append(t.Data, todo)
// }

func (t *TodoList) Add(title string) Todo {
	todo := Todo{
		ID:    len(t.Data) + 1,
		Title: title,
	}

	t.Data = append(t.Data, todo)
	return todo
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

func (t *TodoList) Edit(ID int, NewTitle string) (bool, string) {
	for i := range t.Data {
		if t.Data[i].ID == ID {
			t.Data[i].Title = NewTitle
			return true, "edited"

		}
		if t.Data[i].Done {
			return true, "done"
		}

		if t.Data[i].InProgress {
			return true, "inProgress"
		}
	}
	return false, "not found"
}

// Method 1 : UUID / ID:%s
// fmt.Printf("ID:%s | Title:%s | InProgress:%t | Done:%t\n",
func (t *TodoList) Print() {
	for _, item := range t.Data {
		fmt.Printf("ID:%d | Title:%s | InProgress:%t | Done:%t\n",
			item.ID, item.Title, item.InProgress, item.Done)
	}
}

func (t *TodoList) Load() error {
	data, err := os.ReadFile("output.txt")
	if err != nil {
		return err
	}

	return json.Unmarshal(data, t)

}

func (t *TodoList) Save() error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = os.WriteFile("output.txt", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user := User{
		Username: username,
		Password: password,
	}
	fmt.Println("User created:", user)

	list := TodoList{}
	list.Load()
	fmt.Println("1) Add")
	fmt.Println("2) Remove")
	fmt.Println("3) Edit")
	fmt.Println("4) showlist")
	for {

		var choice int
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var title string
			fmt.Print("Enter title: ")
			fmt.Scan(&title)
			list.Add(title)
			list.Save()
			list.Print()

		case 2:
			var id int
			fmt.Print("Enter id: ")
			fmt.Scan(&id)
			list.Remove(id)
			list.Save()
			list.Print()
		case 3:
			var id int
			fmt.Print("Enter id: ")
			fmt.Scan(&id)
			var title string
			fmt.Scan(&title)

			list.Edit(id, title)
			list.Save()
			list.Print()
		case 4:
			list.Print()
			if len(list.Data) == 0 {
				fmt.Println("List is empty")

			}
		default:
			fmt.Println("Invalid choice")

			return

		}
	}

}
