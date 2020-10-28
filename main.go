/*
yet another url shortener built with Go
*/

package main

import (
	"io"
	"log"
    "fmt"
	"net/http"
    "math/rand"
    "strings"
)

func genRandom(n int)string{
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
	//db := make(map[string]string)

    h1 := func(w http.ResponseWriter, _ *http.Request) {
        io.WriteString(w, fmt.Sprintf("%s\n", genRandom(200)))
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
