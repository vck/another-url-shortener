/*
yet another url shortener built with Go
*/

package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func genRandom(n int) string {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	temp := make([]string, n, n)

	for i := 0; i < n; i++ {
		temp[i] = s[rand.Intn(len(s))]
	}
	rword := strings.Join(temp[:], "")
	return rword
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	db := make(map[string]string)

	h1 := func(w http.ResponseWriter, r *http.Request) {

		for k, v := range r.URL.Query() {
			if k == "url" {
				url := v[0]
				key := genRandom(20)
				db[key] = url
				io.WriteString(w, fmt.Sprintf("log  --> %s\n", db))
				io.WriteString(w, fmt.Sprintf("%s --> %s\n", key, url))
			}
		}
	}

	redirect := func(w http.ResponseWriter, r *http.Request) {

		// redirect /<key> to original URL
		key := r.URL.Path
		key = strings.Replace(key, "/", "", -1)
		url := db[key]

		io.WriteString(w, fmt.Sprintf("accessed path: %s url: %s\n", key, url))
	}

	http.HandleFunc("/sort", h1)
	http.HandleFunc("/", redirect)

	log.Println("Start and Running at Port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
