package handler

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strings"
// )

// func redirect(w http.ResponseWriter, r *http.Request) {
// 	// redirect /<key> to original URL
// 	key := r.URL.Path
// 	key = strings.Replace(key, "/", "", -1)
// 	url := db[key]

// 	io.WriteString(w, fmt.Sprintf("accessed path: %s url: %s\n", key, url))
// }
