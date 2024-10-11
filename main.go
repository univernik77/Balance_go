package main

import (
	"fmt"
)

func scanTransaction() float64 {
		var userTrans float64
        fmt.Println("Введите новую транзакцию: ")
		fmt.Scan(&userTrans)
		return userTrans
}

func main() {
	var userAllTrans []float64
	for {
		
		userTrans := scanTransaction()
		if userTrans == 0 {
			break
		} 
		userAllTrans = append(userAllTrans, userTrans) 
	}
	fmt.Println(userAllTrans)
}
