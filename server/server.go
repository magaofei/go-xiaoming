package main

import (
	"fmt"
	"log"
	"net/http"
)



func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "asdasdsds handle %s, %s", r.URL.Path, r.URL.Path[1:])
}

// 方法的返回值在后边
// 方法的首字母来确定作用域
func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/golang", index)

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalf("sds", err)
	}
}
