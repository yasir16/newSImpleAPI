package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type response struct {
	Result interface{} `json:"result"`
}

func binerToDecimal(biner int) int {
	numDecimal, index, tmp := 0, 0, 0
	for biner != 0 {
		tmp = biner % 10
		biner = biner / 10
		numDecimal = numDecimal + tmp*int(math.Pow(2, float64(index)))
		index++
	}
	return numDecimal
}

func controllerBinerToDecimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	binerStr := vars["biner"]

	biner, _ := strconv.ParseInt(binerStr, 10, 64)

	// fmt.Fprint(w, binerToDecimal(int(biner)))
	result := response{
		Result: binerToDecimal(int(biner)),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}

func controllerDecimalToBiner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	decimalStr := vars["decimal"]

	decimal64, _ := strconv.ParseInt(decimalStr, 10, 64)

	result := response{
		Result: strconv.FormatInt(decimal64, 2),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	// fmt.Fprint(w, strconv.FormatInt(decimal64, 2))

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/binerToDecimal/{biner}", controllerBinerToDecimal).Methods("GET")
	router.HandleFunc("/decimalToBiner/{decimal}", controllerDecimalToBiner).Methods("GET")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
