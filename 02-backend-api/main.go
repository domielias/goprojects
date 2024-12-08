package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Operation struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type RequestItems struct {
	Items []int `json:"items"`
}

func (operation Operation) add() int {
	return operation.Number1 + operation.Number2
}

func (operation Operation) subtract() int {
	return operation.Number1 - operation.Number2
}

func (operation Operation) multiply() int {
	return operation.Number1 * operation.Number2
}

func (operation Operation) divide() int {
	return operation.Number1 / operation.Number2
}

func (items RequestItems) sum() int {
	sum := 0
	for _, num := range items.Items {
		sum += num
	}
	return sum
}

func main() {
	mux := http.NewServeMux()

	addHandler := http.HandlerFunc(add)
	mux.Handle("/add", middleware(addHandler))
	mux.Handle("/subtract", middleware(http.HandlerFunc(subtract)))
	mux.Handle("/multiply", middleware(http.HandlerFunc(multiply)))
	mux.Handle("/divide", middleware(http.HandlerFunc(divide)))
	mux.Handle("/sum", middleware(http.HandlerFunc(sum)))
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing the following path: ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func sum(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var items RequestItems
		err := json.NewDecoder(r.Body).Decode(&items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		result := map[string]int{"result": items.sum()}
		js, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var op Operation
		err := json.NewDecoder(r.Body).Decode(&op)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		result := map[string]int{"result": op.add()}
		js, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func subtract(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var op Operation
		err := json.NewDecoder(r.Body).Decode(&op)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		result := map[string]int{"result": op.subtract()}
		js, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func multiply(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var op Operation
		err := json.NewDecoder(r.Body).Decode(&op)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		result := map[string]int{"result": op.multiply()}
		js, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func divide(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var op Operation
		err := json.NewDecoder(r.Body).Decode(&op)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		result := map[string]int{"result": op.divide()}
		js, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
