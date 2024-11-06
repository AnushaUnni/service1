package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// API endpoint that accepts two integers as query parameters and returns their sum
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		aStr := r.URL.Query().Get("a")
		bStr := r.URL.Query().Get("b")

		// Convert query parameters to integers
		a, err := strconv.Atoi(aStr)
		if err != nil {
			http.Error(w, "Invalid parameter 'a', must be an integer", http.StatusBadRequest)
			return
		}

		b, err := strconv.Atoi(bStr)
		if err != nil {
			http.Error(w, "Invalid parameter 'b', must be an integer", http.StatusBadRequest)
			return
		}

		// Compute the sum
		sum := a + b

		// Write the response
		fmt.Fprintf(w, "The sum of %d and %d is %d", a, b, sum)
	})

	port := 8080
	fmt.Printf("Microservice running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
