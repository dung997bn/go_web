package main

func basic() {

}

// import (
// 	"fmt"
// 	"net/http"
// )

// type login int
// type welcom int

// func (l login) ServeHTTP(w http.ResponseWriter, t *http.Request) {
// 	fmt.Fprintln(w, "On login page")
// }
// func (w welcom) ServeHTTP(w1 http.ResponseWriter, t *http.Request) {
// 	fmt.Fprintln(w1, "On welcome page")
// }

// func mywelcom(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "On welcome page")
// }

// func mylogin(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, `
// 	// <html>
// 	// <head>Login</head>
// 	// <body>
// 	// <h1>Please enter username & password</h1>
// 	// </body>
// 	// </html>
// 	// `)
// 	// fmt.Fprintf(w, "My request: %+v\n", r)
// 	if r.Method == "GET" {
// 		fmt.Fprintln(w, "Using GET for login endpoint")
// 	}
// 	if r.Method == "POST" {
// 		fmt.Fprintln(w, "Using POST for login endpoint")
// 	}
// 	fmt.Fprintln(w, "On login page")
// }

// func main() {
// 	//way1
// 	http.HandleFunc("/login", mylogin)
// 	http.HandleFunc("/welcome/", mywelcom)

// 	//way2
// 	// http.Handle("/login/", http.HandlerFunc(mylogin))
// 	// http.Handle("/welcome", http.HandlerFunc(mywelcom))

// 	//way 3
// 	// var i login
// 	// var j welcom
// 	// http.Handle("/login", i)
// 	// http.Handle("/welcome", j)

// 	fmt.Println("Listening on port 8000...")
// 	http.ListenAndServe("localhost:8000", nil)
// }
