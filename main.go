package main

import (
	"fmt"
)

func scanTransaction() float64 {
		var userTrans float64
        fmt.Print("Введите новую транзакцию (n для выхода): ")
		fmt.Scan(&userTrans)
		return userTrans
}

func sumTrans(userAllTrans []float64) float64 {
	var sum float64
	for _, trans := range userAllTrans {
		sum += trans
	}
	return sum
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
	sum := sumTrans(userAllTrans)
	fmt.Printf("Итоговый баланс всех транзакций: %.2f\n", sum)
}
