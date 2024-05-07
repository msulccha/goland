/*
	package main

import (

	"fmt"
	"math"

)

func main() {

	//investmentAmount, year, expectedReturnRate := 1000.0, 10.0, 5.5
	const inflationRate = 2.5
	var investmentAmount float64
	var year float64
	expectedReturnRate := 5.5

	fmt.Print("Ingrese el monto de inversión: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Ingrese la tasa: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Ingrese los años: ")
	fmt.Scan(&year)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, year)

	futureRealValue := futureValue * math.Pow(1+inflationRate/100, year)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
*/
package main

import (
	"fmt"
	"os"
)

func main() {

	revenue := InputValuesAndText("Ingresos: ")
	expenses := InputValuesAndText("Gastos: ")
	taxRate := InputValuesAndText("Impuestos: ")

	profit, ebt, ratio := OutputEvaluation(revenue, expenses, taxRate)

	writeBalanceTofile(profit, ebt, ratio)

	fmt.Println("tus ganancias despues de impuestos es: ", profit)
	fmt.Println("tus utilidad neta antes de impuestos es: ", ebt)
	fmt.Printf("el ratio es: %.3f\n", ratio)

	fmt.Println("Gracias")

}

func InputValuesAndText(infoText string) float64 {

	var value float64

	fmt.Print(infoText)
	fmt.Scan(&value)

	if value <= 0 {
		panic("Valor ingresado inválido, el valor debe ser diferente de 0 y no negativo")

	} else {
		return value
	}
}

func OutputEvaluation(revenue, expenses, taxRate float64) (float64, float64, float64) {

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return profit, ebt, ratio

}

const calculateValue = "calculate.txt"

func writeBalanceTofile(value1, value2, value3 float64) {

	calculateText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", value1, value2, value3)
	os.WriteFile(calculateValue, []byte(calculateText), 0644)
}
