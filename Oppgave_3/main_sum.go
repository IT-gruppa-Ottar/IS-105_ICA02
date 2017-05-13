package main

import (
  "fmt"
  "os"
  "strconv"
)

func main(){
	if len(os.Args) <= 2 {
		fmt.Println("Ikke nokk argumenter!")
		fmt.Println()
		os.Exit(2)
	}

	arg1 := os.Args[1]
    	arg2 := os.Args[2]

    	argInput1, err := strconv.ParseFloat(arg1, 64)
    	argInput2, err := strconv.ParseFloat(arg2, 64)

    	if err != nil {
		fmt.Println()
		fmt.Println(err)
        	fmt.Println()
        	fmt.Println("Eksempel på hvordan å kjøre denne filen:")
        	fmt.Println("go run FILNAVN NUMBER_1 NUMBER_2")
        	fmt.Println()
        	fmt.Println("Programmet tar heller ikke imot bokstaver")
		os.Exit(2)
    	} else {
		fmt.Println()
		fmt.Println("Summen av " + arg1 + " + " + arg2 + " er:")
		fmt.Println(argInput1 + argInput2)
		fmt.Println()
	}
}
