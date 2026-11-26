package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var btn_count int

func click(w http.ResponseWriter, r *http.Request) {
	if btn_count < 50 {
		btn_count += 1
	} else if btn_count >= 50 && btn_count <= 149 {
		btn_count += 2
	} else if btn_count >= 150 && btn_count <= 299 {
		btn_count += 4
	} else if btn_count >= 300 && btn_count <= 699 {
		btn_count += 6
	} else if btn_count >= 700 && btn_count <= 999 {
		btn_count += 10
	} else if btn_count >= 1000 && btn_count <= 1500 {
		btn_count += 11
	} else {
		btn_count += 12
	}

	fmt.Fprint(w, btn_count)
}

func restore(w http.ResponseWriter, r *http.Request) {
	btn_count = 0
	fmt.Fprint(w, btn_count)
}

func loadClicks(w http.ResponseWriter, r *http.Request) {
	btnCountStr := strconv.Itoa(btn_count)
	fmt.Fprint(w, btnCountStr)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/click", click)
	http.HandleFunc("/restore", restore)
	http.HandleFunc("/loadClicks", loadClicks)

	fmt.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
