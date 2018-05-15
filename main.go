package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type Thanos struct {
	PageTitle string
	Life      string
}

func doesThanosSpareYou(w http.ResponseWriter, r *http.Request) {
	var message string
	rand.Seed(time.Now().UnixNano())
	q := rand.Intn(30)
	if q%13 == 0 {
		message = "You were spared by Thanos."
	} else {
		message = "You were slain by Thanos."
	}
	thanos := Thanos{PageTitle: "Did Thanos kill me", Life: message}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, thanos)
}

func main() {
	fmt.Println("Thanos has snapped is fingers")
	http.HandleFunc("/", doesThanosSpareYou)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
