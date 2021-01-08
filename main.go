package main

func main() {
	
}

// func main() {
// 	// Hello world, the web server
// 	r := mux.NewRouter()
// 	helloHandler := func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Hello, world!\n")
// 	}

// 	r.HandleFunc("/hello", helloHandler)

// 	srv := &http.Server{
// 		Handler: r,
// 		Addr:    "127.0.0.1:8088",
// 		// Good practice: enforce timeouts for servers you create!
// 		WriteTimeout: 15 * time.Second,
// 		ReadTimeout:  15 * time.Second,
// 	}

// 	log.Fatal(srv.ListenAndServe())
// }

// func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	numberString := vars["number"]
// 	n, err := strconv.Atoi(numberString)
// 	if err != nil {
// 		fmt.Fprint(w, err.Error())
// 		return
// 	}

// 	s := fizzbuzz.Say(n)
// 	fmt.Fprint(w, s)
// }

// func main() {
// println(swap("a", "b"))
// s := fmt.Sprintf("%d%s", 4, "test")
// fmt.Println(s)

// // var n int
// // var err error
// // n, err = strconv.Atoi("15")
// // if err != nil {
// // 	log.Fatal(err)
// // }

// // fmt.Println("convert to", n)

// if n, err := strconv.Atoi("15"); err != nil {
// 	log.Fatal(err)
// } else {
// 	fmt.Println("convert to", n)
// }

// println(math.Power(2, 3))
// }

// func swap(a,b string) (string, string) {
// 	return b, a
// }

// func squareArea(a int) int {
// 	return a*a
// }
