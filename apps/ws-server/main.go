package main

import (
	"fmt"
	"github.com/itsindigo/scaling-ws/ws-server/internal/thing"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, thing.PrintThing())

}

func main() {
	http.HandleFunc("/", helloHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
