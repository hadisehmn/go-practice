func main() {

	http.HandleFunc("/signup", SignUp)
	http.HandleFunc("/signin", SignIn)

	http.ListenAndServe(":8080", nil)
}
