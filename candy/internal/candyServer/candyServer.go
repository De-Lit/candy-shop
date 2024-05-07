package candyserver

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	thanks = "Thank you!"
	port   = ":3333"
)

var priceList = map[string]uint{
	"CE": 10, // Cool Eskimo
	"AA": 15, // Apricot Aardvark
	"NT": 17, // Natural Tiger
	"DE": 21, // Dazzling Elderberry
	"YR": 23, // Yellow Rambutan
}

type CandyOrder struct {
	Money      uint   `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount uint   `json:"candyCount"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Run() {
	server := getServer()
	http.HandleFunc("/buy_candy", buyCandy)
	log.Fatal(server.ListenAndServe())
}

func getServer() *http.Server {
	config := &tls.Config{}

	server := &http.Server{
		Addr:      port,
		TLSConfig: config,
	}
	return server
}

func buyCandy(w http.ResponseWriter, r *http.Request) {
	var order CandyOrder
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		errorResponse := ErrorResponse{Error: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Check if candy is not in price-list
	if _, ok := priceList[order.CandyType]; !ok {
		errorResponse := ErrorResponse{Error: fmt.Sprintf("%s is not in price-list!", order.CandyType)}
		w.WriteHeader(http.StatusPaymentRequired)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Check if there is enough money for the purchase
	price := calculatePrice(order.CandyType, order.CandyCount)
	if order.Money < price {
		errorResponse := ErrorResponse{Error: fmt.Sprintf("You need %d more money!", price-order.Money)}
		w.WriteHeader(http.StatusPaymentRequired)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Process the purchase
	// ...

	// Return a successful response
	response := map[string]interface{}{
		"thanks": thanks,
		"change": order.Money - price,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func calculatePrice(candyType string, candyCount uint) uint {
	// Calculate the price based on the candy type and quantity
	var price uint = 0
	price = priceList[candyType] * candyCount
	return price
}
