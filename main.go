package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", mainP)
	m.HandleFunc("/hello", helloPage)
	log.Fatal(http.ListenAndServe(":8081", m))
}

func mainP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
	fmt.Println("1")

}

func helloPage(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем CORS заголовки
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	type incomingData struct {
		Name string `json:"name"`
	}
	type outgoingData struct {
		Message string `json:"message"`
	}
	var incV incomingData
	buff := r.Body
	fmt.Println(buff)
	decoder := json.NewDecoder(buff)
	decoder.Decode(&incV)
	//fmt.Fprintf(w, "hello %s", incV.Name)
	// Создаем JSON ответ
	//response := map[string]string{
	//	"message": fmt.Sprintf("Привет, %s!", incV.Name),
	//}
	response := outgoingData{"Hello " + incV.Name}

	// Кодируем и отправляем ответ
	json.NewEncoder(w).Encode(response)

}
