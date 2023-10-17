package main

import (
	"encoding/json"
	"fmt"
	"go-json/entity"
)

func main() {

	myBalance := `{"id" : "balance-1", "balance" : 1000}`

	decodeToStruct(myBalance)
	decodeToMap(myBalance)

	var dataJson = `[
    {"id": "balance-1", "balance": 1000},
    {"id": "balance-2", "balance": 1200}
]`
	decodeToArrayOfObj(dataJson)

}

func decodeToStruct(jsonStr string) {
	var data entity.Balance
	err := json.Unmarshal([]byte(jsonStr), &data)

	if err != nil {
		panic(err)
	}

	fmt.Println("Id Balance :", data.Id)
	fmt.Println("Total Balance:", data.Balance)
}

func decodeToMap(jsonStr string) {

	var data map[string]any // atau interface{}

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)

}

func decodeToArrayOfObj(strObj string) {

	var temporaryJson []entity.Balance

	err := json.Unmarshal([]byte(strObj), &temporaryJson)
	if err != nil {
		panic(err)
	}

	fmt.Println(temporaryJson)
}
