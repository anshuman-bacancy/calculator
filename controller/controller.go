package controller

import (
	"calculator/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Solve computes the values
func Solve(res http.ResponseWriter, req *http.Request) {
	var data models.Data
	json.NewDecoder(req.Body).Decode(&data)

	fmt.Println(data)
	var op1, op2, result float64
	op1, _ = strconv.ParseFloat(data.Operator1, 64)
	op2, _ = strconv.ParseFloat(data.Operator2, 64)

	switch data.Operand {
	case "+":
		result = op1 + op2
	case "*":
		result = op1 * op2
	case "/":
		result = op1 / op2
	case "-":
		result = op1 - op2
	}

	fmt.Println(result)
	json.NewEncoder(res).Encode(struct {
		Result string `json:"result"`
	}{fmt.Sprint(result)})
}
