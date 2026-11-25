package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var btn_count int

func click(w http.ResponseWriter, r *http.Request) {
	btn_count += 1
	fmt.Fprint(w, btn_count)
}

func restore(w http.ResponseWriter, r *http.Request) {
	btn_count = 0
	fmt.Fprint(w, btn_count)
}

func addclicks(w http.ResponseWriter, r *http.Request) {
	btn_count = btn_count + 100 // or btn_count += 100
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
	http.HandleFunc("/addclicks", addclicks)
	http.HandleFunc("/loadClicks", loadClicks)

	fmt.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
