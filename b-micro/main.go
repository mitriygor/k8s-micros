package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintln(w, "B-Micro: Hello B - 008")
		} else {
			fmt.Fprintln(w, "B-Micro: Error")
		}
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("B-Micro; Error listen")
	} else {
		fmt.Println("B-Micro: launched")
	}
}
