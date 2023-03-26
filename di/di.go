package main

import (
	"fmt"
	"io"
	"net/http"
)

// 引数に interface を指定することで特定の型ではなく interface を実装した何を受け取れるようになる
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "Chinatsu")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreetingHandler))
}
