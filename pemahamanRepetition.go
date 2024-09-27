package main

import (
	"fmt"
)

func main() {


	caseA()
	fmt.Println()
	caseB()
	fmt.Println()
	caseC()
	fmt.Println()
	caseD()
	fmt.Println()
	caseE()
	fmt.Println()
	caseF()
	fmt.Println()
	caseG()


}

func caseA () {
	n := 10

	for i := 1; i <= n; i++ {
		fmt.Print(" ")
		fmt.Printf("%d",i)
	}


}

func caseB () {
	n := 5

	for i := 1; i <= n; i++ {
		
		fmt.Printf("%d",i)
		fmt.Print(" ")
		fmt.Print("*")
		fmt.Print(" ")
	}
}

func caseC() {
	n := 20

	for i := 1; i <= n; i++ {
		if i % 2 == 1 {
			fmt.Print(" ")
			fmt.Print(i)
		}

	}
}


func caseD() {
	n := 10

	for i := 1; i <= n; i++ {
		fmt.Print(i*2)
		fmt.Print(" ")
	}
}

func caseE() {
	n := 30

	for i := 0; i <= n; i++ {
	
		if i % 3 == 0 {
			fmt.Print(i)
		    fmt.Print(" ")
		}
	}
}

func caseF() {
	
	j := 0
	k := 1
	n := 1


	for i := 1; i <= 10; i++ {
		j = j + k
		k = j - k
		n = n + j

		fmt.Printf("%d", n)
		fmt.Print(" ")



	}

}


/*************  ✨ Codeium Command ⭐  *************/
// caseG prints the numbers from 1 to 10 separated by spaces.
/******  b01d7558-dddf-4757-b05e-632d3ed88ed2  *******/
func caseG() {
	j := 1
	for i := 1; i <= 5; i++ {

		j = i*2

		fmt.Print(i)
		fmt.Print(" ")
		fmt.Print(j)
		fmt.Print(" ")
	}
}