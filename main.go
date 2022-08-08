package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello! I can solve your math problem ;)")
}

//The Scanner funcion reads an argument from the command line
//as a strinfg and splits it into two arguments and an operator
func Scanner1(s string) (float64, string, float64) {
	str := os.Args[1]
	operFlag := false
	arg1 := ""
	arg2 := ""
	oper := ""
	err1 := true
	for _, el := range s {
		if (el >= '0' && el <= '9') && !operFlag {
			arg1 += string(el)
		} else if MathOper(el) {
			oper += string(el)
			err1 = false
			operFlag = true
		} else if (el >= '0' && el <= '9') && operFlag {
			arg2 += string(el)
		}
	}
	if err1 {
		fmt.Println("Please input correct math operator such as: + - * / %")
		return
	}
}

//MathOper function checks if a given rune is a math operator
func MathOper(r rune) bool {
	operArr := "+-*/%"
	for _, el := range operArr {
		if el == r {
			return true
		}
	}
	return false
}
