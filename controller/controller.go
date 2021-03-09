package controller

import (
	"calculator/models"
	"calculator/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Solve computes the values
func Solve(res http.ResponseWriter, req *http.Request) {
	var data models.Data
	json.NewDecoder(req.Body).Decode(&data)

	var op1, op2, result float64
	op1, _ = strconv.ParseFloat(data.Operand1, 64)
	op2, _ = strconv.ParseFloat(data.Operand2, 64)

	switch data.Operator {
	case "+":
		result = op1 + op2
	case "*":
		result = op1 * op2
	case "/":
		result = op1 / op2
	case "-":
		result = op1 - op2
	}

	services.Save(data, result)
	json.NewEncoder(res).Encode(struct{ Result string `json:"result"`}{fmt.Sprint(result)})
}
