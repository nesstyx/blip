package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	saveClicks()
	fmt.Fprint(w, btn_count)
}

func restore(w http.ResponseWriter, r *http.Request) {
	btn_count = 0
	saveClicks()
	fmt.Fprint(w, btn_count)
}

func loadClicks(w http.ResponseWriter, r *http.Request) {
	btnCountStr := strconv.Itoa(btn_count)
	fmt.Fprint(w, btnCountStr)
}

func saveClicks() {
	os.WriteFile("count.txt", []byte(strconv.Itoa(btn_count)), 0644)
}

func loadClicksFile() error {
	data, err := os.ReadFile("count.txt")

	if err != nil {
		return fmt.Errorf("С файлом count.txt что то не так..", err)
	}
	n, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return fmt.Errorf("С файлом count.txt что то не так..", err)
	}
	btn_count = n
	return nil
}

func main() {
	loadClicksFile()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/click", click)
	http.HandleFunc("/restore", restore)
	http.HandleFunc("/loadClicks", loadClicks)

	fmt.Println("Сервер запущен: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Не удалось запустить сервер: ", err)
	}
}
