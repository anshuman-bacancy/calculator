package services

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
	"calculator/models"
)


func Save(computeData models.Data, result float64) {
	f, err := os.OpenFile("static/history.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	    log.Fatal(err)
	}
 	jsonToSave := struct{
		Op1 string `json:"Operator1"`
		Operand string `json:"Operand"`
		Op2 string `json:"Operator2""`
		Result float64 `json:"Result"`

	}{computeData.Operand1, computeData.Operator, computeData.Operand2, result,}

	fmt.Println("anonymous struct", jsonToSave)
	jsonFile, _ := json.MarshalIndent(jsonToSave, "", " ") 
	f.Write(jsonFile)
}



