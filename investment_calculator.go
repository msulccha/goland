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
)

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Ingresos: ")
	fmt.Scan(&revenue)

	fmt.Print("Gastos: ")
	fmt.Scan(&expenses)

	fmt.Print("Tasa de Impuestos: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses

	earningAfterTax := earningBeforeTax * (1 - taxRate/100)

	ratio := earningBeforeTax / earningAfterTax

	fmt.Print("EBT: ")
	fmt.Println(earningBeforeTax)

	fmt.Print("profit: ")
	fmt.Println(earningAfterTax)

	fmt.Print("ratio: ")
	fmt.Println(ratio)
}
