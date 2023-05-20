package functions

import "fmt"

func CalculateTaxes(revenue, deductions, credits float64) float64 {
	// Here, "fmt.Println()" function will be called at the very end of "calculateTaxes" function.
	// After all the calculations are finished, then only we will call "fmt.Println()"
	defer fmt.Println("Taxes Calculated!")
	taxRate := .06143
	fmt.Println("Calculating Taxes")

	if deductions == 0 || credits == 0 {
		return revenue * taxRate
	}

	taxValue := (revenue - (deductions * credits)) * taxRate
	if taxValue >= 0 {
		return taxValue
	} else {
		return 0
	}
}
