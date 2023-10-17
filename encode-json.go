package main

import (
	"encoding/json"
	"fmt"
	"go-json/entity"
)

func main() {

	myFirstBalance := entity.Balance{
		Id:      "balance-1",
		Balance: 1000,
	}
	mySecondBalance := entity.Balance{
		Id:      "balance-2",
		Balance: 1500,
	}

	data := []entity.Balance{myFirstBalance, mySecondBalance}

	dataByte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataByte))
}
