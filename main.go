package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// LEVEL 1
// func main() {
// 	scores := make(map[string]int)

// 	scores["student1"] = 10
// 	scores["student2"] = 90
// 	scores["student3"] = 85
// 	// fmt.Println(scores["student1"])

// 	// for name, score := range scores {
// 	// 	fmt.Println(name, score)

// 	// }

// 	highestname := ""
// 	highestscore := -math.MaxInt

// 	lowestname := ""
// 	lowestscore := math.MaxInt

// 	for name, score := range scores {

// 		if highestscore < score {
// 			highestscore = score
// 			highestname = name
// 		}

// 		if lowestscore > score {
// 			lowestscore = score
// 			lowestname = name

// 		}

// 	}

// 	fmt.Println(highestname, highestscore)
// 	fmt.Println(lowestname, lowestscore)

// }

// LEVEL 2
// func main() {
// 	sentence := " this is a test and  alireza have a good asss :)"
// 	words := strings.Split(sentence, " ")
// 	counts := make(map[string]int)
// 	for _, word := range words {
// 		counts[word]++
// 	}
// 	for word, count := range counts {

// 		if count > 1 {
// 			fmt.Println(word)
// 		}

// 	}
// }

// LEVEL 3

// type Addition struct {
// 	a, b int
// }

// func (a Addition) plus() int {
// 	return a.a + a.b

// }
// func calculators() func() int {
// 	add := Addition{a: 3, b: 90}
// 	return func() int {
// 		return add.plus()
// 	}
// }

// func main() {
// 	f := calculators()
// 	fmt.Println(f())
// }

// type calculator struct {
// 	a, b int
// }

// func (c calculator) add() int {
// 	return c.a + c.b

// }
// func (c calculator) sub() int {
// 	return c.a - c.b
// }
// func (c calculator) mul() int {
// 	return c.a * c.b
// }
// func (c calculator) div() int {
// 	return c.a / c.b
// }

// func calculators() {

// 	c := calculator{a: 10, b: 5}

// 	fmt.Println(c.add())
// 	fmt.Println(c.sub())
// 	fmt.Println(c.mul())
// 	fmt.Println(c.div())
// }

// func main() {
// 	calculators()
// }

// LEVEL 4

// type Student struct {
// 	ID    int
// 	NAME  string
// 	AGE   int
// 	SCORE int
// }

// var students = make(map[int]Student)

// func addStudent(s Student) string {
// 	students[s.ID] = s
// 	return "student added "
// }

// func deletStudent(ID int) string {
// 	delete(students, ID)
// 	return "student deleted "

// }

// func findStudent(ID int) string {
// 	for _, s := range students {
// 		if s.ID == ID {
// 			return s.NAME
// 		}
// 	}
// 	return ""
// }

// func topStudent() (string, int, string, int) {

// 	highestName := ""
// 	highestScore := -math.MaxInt

// 	lowestName := ""
// 	lowestScore := math.MaxInt

// 	for _, s := range students {

// 		if s.SCORE > highestScore {
// 			highestScore = s.SCORE
// 			highestName = s.NAME
// 		}

// 		if s.SCORE < lowestScore {
// 			lowestScore = s.SCORE
// 			lowestName = s.NAME
// 		}
// 	}

// 	return highestName, highestScore, lowestName, lowestScore
// }

// func duplicateStudent() []Student {
// 	newlist := []Student{}
// 	finddup := make(map[string]bool)
// 	for _, student := range students {
// 		if !finddup[student.NAME] {
// 			finddup[student.NAME] = true
// 			newlist = append(newlist, student)
// 		}

// 	}
// 	return newlist

// }

// func main() {
// 	addStudent(Student{ID: 1, NAME: "alireza", AGE: 25, SCORE: 100})
// 	addStudent(Student{ID: 3, NAME: "hidis ", AGE: 25, SCORE: 98})
// 	addStudent(Student{ID: 2, NAME: "reza  ", AGE: 87, SCORE: 76})
// 	addStudent(Student{ID: 2, NAME: "reza  ", AGE: 87, SCORE: 76})
// 	addStudent(Student{ID: 4, NAME: "meli  ", AGE: 31, SCORE: 55})

// 	deletStudent(3)

// 	fmt.Println(students)

// 	fmt.Println(findStudent(2))

// 	highName, highScore, lowName, lowScore := topStudent()

// 	fmt.Println("Top student:", highName, highScore)
// 	fmt.Println("Lowest student:", lowName, lowScore)

// 	result := duplicateStudent()
// 	for _, s := range result {
// 		fmt.Println(s.ID, s.NAME, s.AGE, s.SCORE)

// 	}
// }

// 5
// type person struct {
// 	Name string
// 	City string
// 	Age  int
// }

// func (p person) seyhelloo() {
// 	fmt.Println("helloo my name is ", p.Name, " i am ", p.Age, " yers old  and i live in ", p.City)

// }

// func main() {
// 	newperson := person{
// 		Name: "John",
// 		City: "Amsterdam",
// 		Age:  25,
// 	}
// 	newperson.seyhelloo()
// }

// 6
// func main() {
// 	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	fmt.Println(numbers)

// 	for _, v := range numbers {
// 		if v > 5 {
// 			fmt.Println(v)

// 		}

// 	}

// }

// 7
// func divide(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("this number is invalid")

// 	}
// 	return a / b, nil
// }

// func main() {
// 	result, err := divide(10, 5)
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return
// 	}
// 	fmt.Println("result", result)
// }

// 8
// func Person(name string, age int) string {
// 	return fmt.Sprintf("My name is %s and I'm %d years old", name, age)
// }

// func main() {
// 	result := Person("Alice", 25)
// 	fmt.Println(result)
// }

// 9
// func printnumber() {
// 	for i := 0; i <= 20; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i, " this number is even ")
// 		} else {
// 			fmt.Println(i, " its odd ")
// 		}
// 	}
// }

// func main() {
// 	printnumber()
// }

// 10
// func GradeChecker(i int) {
// 	if i == 100 {
// 		fmt.Println(i, "this grade is A+")
// 	} else if i >= 90 {
// 		fmt.Println(i, "this grade is A")
// 	} else if i >= 80 {
// 		fmt.Println(i, "this grade is B")
// 	} else if i >= 70 {
// 		fmt.Println(i, "this grade is C")
// 	} else {
// 		fmt.Println(i, "this student is fail")
// 	}
// }

// func main() {
// 	GradeChecker(76)
// }

//11

// func DayOfWeek(i int) {
// 	switch i {
// 	case 1:
// 		fmt.Printf("Sunday")
// 	case 2:
// 		fmt.Printf("monday")
// 	case 3:
// 		fmt.Printf("tuesday")
// 	default:
// 		fmt.Println("Invalid day")
// 	}
// }

// func main() {
// 	DayOfWeek(5)

// }

// 12
// func counter() {
// 	words := []string{"go", "java", "go", "python", "go"}
// 	counts := make(map[string]int)

// 	for _, word := range words {
// 		counts[word]++
// 	}

// 	for word, count := range counts {
// 		fmt.Println(word, count)
// 	}
// }

// func main() {
// 	counter()
// }

// 13
// func main() {
// 	x := 45
// 	p := &x
// 	*p = 100

//		fmt.Println(x)
//		fmt.Println(p)
//	}
// func main() {
// 	a := 10
// 	p := &a
// 	*p = 500
// 	fmt.Println(a)
// 	fmt.Println(p)
// }

// func main() {
// 	a := 40
// 	b := 60

// 	p := &a

// 	*p = 100

// 	fmt.Println("a =", a)
// 	fmt.Println("b =", b)

// 	p = &b

// 	*p = 200

// 	fmt.Println("a =", a)
// 	fmt.Println("b =", b)
// }

// func main() {
// 	a := 40
// 	b := 60

// 	fmt.Println("Before swap:")
// 	fmt.Println("a =", a)
// 	fmt.Println("b =", b)

// 	swap(&a, &b)

// 	fmt.Println("After swap:")
// 	fmt.Println("a =", a)
// 	fmt.Println("b =", b)
// }

// func swap(x *int, y *int) {
// 	temperary := *x
// 	*x = *y
// 	*y = temperary
// }

// 14
// type Geometry interface {
// 	Area() float64
// }

// type Rectangle struct {
// 	length float64
// 	width  float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.length * r.width
// }

// func main() {
// 	rect := Rectangle{length: 10, width: 5}

// 	fmt.Println("Rectangle Area:", rect.Area())

// 	var g Geometry
// 	g = rect

// 	fmt.Println("Shape (Rectangle) Area via interface:", g.Area())
// }

// 15
//divide(a, b int) (int, error)

// func calculators(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("can't work with 0")

// 	} else {
// 		return a / b, nil
// 	}

// }
// func main() {
// 	result, err := calculators(5, 3)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 		return
// 	}

// 	fmt.Println("result:", result)
// }

// 16

// type test struct {
// 	message string
// 	code    int
// }

// func (e test) Error() string {
// 	return e.message
// }

// func main() {
// 	err := test{
// 		message: "invalid email address",
// 		code:    400,
// 	}

// 	fmt.Println(err)
// }

// 17
// func printNumber(n int) {
// 	fmt.Println(n)
// }

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		go printNumber(i)
// 	}

// 	time.Sleep(1 * time.Second)
// }

// 18
// func massage(ch chan string) {
// 	ch <- " hello from goroutine"
// }
// func main() {
// 	ch := make(chan string)

// 	go massage(ch)
// 	msg := <-ch
// 	fmt.Println(msg)
// }

// 19
// func sendnumber(ch chan int) {
// 	for i := 1; i <= 5; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }
// func recivenumber(ch chan int) {
// 	for v := range ch { //v:= <-ch
// 		fmt.Println(v)

// 	}

// }

// func main() {
// 	ch := make(chan int)
// 	go sendnumber(ch)
// 	recivenumber(ch)
// }

// 20
// func main() {
// 	ch := make(chan int, 2)

// 	ch <- 1
// 	ch <- 2
// 	ch <- 3

// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }
// 2 ---> output : exit status 2
// 3 or more ---> output : true

// 21
// func number(n int, wg *sync.WaitGroup) {

// 	defer wg.Done()
// 	fmt.Println("this number is", n)
// }
// func main() {
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)

// 		go number(i, &wg)
// 	}

// 	wg.Wait()

// }

// 22
// func worker(done chan bool, id int) {
// 	fmt.Println("worker", id, "is done")
// 	done <- true
// }
// func main() {

// 	ch := make(chan bool)
// 	go worker(ch, 1)
// 	go worker(ch, 2)
// 	go worker(ch, 3)

// 	<-ch
// 	<-ch
// 	<-ch
// }

// 23

// func message(ch chan string) {
// 	fmt.Println("message sent")

// 	ch <- "hello from goroutine"
// }

// func main() {
// 	ch := make(chan string)

// 	go message(ch)

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("received:", msg)
// 	}
// }

// 24 e.g
// func worker(ch chan string, msg string, delay time.Duration) { //time.Duratio =  type of 3*time.Second
// 	time.Sleep(delay)
// 	ch <- msg
// }

// func main() {
// 	ch := make(chan string)

// 	// goroutine 1
// 	go worker(ch, "sent after 3 seconds", 3*time.Second)

// 	// goroutine 2
// 	go worker(ch, "sent after 1 second", 1*time.Second)

// 	// receive messages
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// 25?
// func worker(ch chan string, msg string, delay time.Duration) {
// 	time.Sleep(delay)
// 	ch <- msg
// }

// func main() {
// 	ch := make(chan string)

// 	go worker(ch, "hello", 1*time.Second)

// 	select {
// 	case msg := <-ch:
// 		fmt.Println(msg)

// 	case <-time.After(2 * time.Second):
// 		fmt.Println("timeout")
// 	}
// }

//leecode
//26

// func main() {
// 	nums := [7]int{1, 2, 3, 4, 5}
// 	for i, num := range nums {
// 		if num+2 == 6 {
// 			fmt.Println("equal", i)

//			} else {
//				fmt.Println("not equal")
//			}
//		}
//	}
// func main() {
// 	nums := []int{1, 2, 3, 4, 5}
// 	for i, num1 := range nums {
// 		for j, num2 := range nums {
// 			if num1+num2 == 6 {
// 				fmt.Println("equal ", i, j)

// 			} else {
// 				fmt.Println("not equal")
// 			}
// 		}

//		}
//	}
//
// 27

// func roman() map[string]int {
// 	m := map[string]int{
// 		"I": 1,
// 		"V": 5,
// 		"X": 10,
// 	}
// 	return m
// }

// func main() {
// 	m := roman()

// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}
// }

// func main() {
// 	m := map[string]int{
// 		"I": 1,
// 		"V": 5,
// 		"X": 10,
// 	}

// 	fmt.Println(m["I"]) // 1
// 	fmt.Println(m["V"]) // 5
// }

// 28
// func main() {
// 	m1 := []string{"apple", "banana"}
// 	m2 := []string{"kiwi"}
// 	marge := append(m1, m2...)
// 	fmt.Println(marge)
// }

// func main() {
// 	m := map[string][]string{
// 		"fruits": {"apple", "banana"},
// 		"cars":   {"BMW", "Tesla"},
// 	}

// 	var result strings.Builder

// 	for _, v := range m {
// 		for _, item := range v {
// 			result.WriteString(item)
// 			result.WriteString(" ")
// 		}
// 	}

// 	fmt.Println(result.String())
// }
//29

// func main() {
// 	numbers := []int{1, 2, 3, 3, 3, 5, 6, 6, 7, 8, 7, 9}

// 	exists := make(map[int]bool)
// 	result := []int{}

// 	for _, v := range numbers {
// 		if !exists[v] {
// 			exists[v] = true
// 			result = append(result, v)
// 		}
// 	}

// 	fmt.Println(result)
// }

// 30

// type Color struct {
// 	Name  string
// 	Value int
// }

// func sortColor(colors *map[string]int) {
// 	var items []Color

// 	for name, value := range *colors {
// 		items = append(items, Color{
// 			Name:  name,
// 			Value: value,
// 		})
// 	}

// 	sort.Slice(items, func(i, j int) bool {
// 		return items[i].Value < items[j].Value
// 	})

// 	fmt.Print("[")
// 	for _, item := range items {
// 		fmt.Printf("%d %s ", item.Value, item.Name)
// 	}
// 	fmt.Println("]")
// }

// func main() {
// 	colors := map[string]int{
// 		"red":   1,
// 		"black": 2,
// 		"blue":  3,
// 	}

// 	sortColor(&colors)
// }

// 31
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
