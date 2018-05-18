package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Thanos struct {
	PageTitle string
	Life      string
}

func doesThanosSpareYou(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	q := rand.Intn(30)
	if q%13 == 0 {
		fmt.Fprintf(w, "You were spared by Thanos.")
	} else {
		fmt.Fprintf(w, "You were slain by Thanos.")
	}
	//fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

func main() {
	fmt.Println("Thanos has snapped is fingers")
	http.HandleFunc("/", doesThanosSpareYou)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
