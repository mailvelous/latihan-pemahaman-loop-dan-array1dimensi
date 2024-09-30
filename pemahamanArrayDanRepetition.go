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
	fmt.Println()
	caseH()
	fmt.Println()
	caseI()

}

func caseA() {
	var number int
	fmt.Print("Enter a number: ")  // Prompt the user to enter a number
	fmt.Scan(&number) 

	if number < 0 {
		fmt.Println("Negative")
	} else if number > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Zero")
	}
}

func caseB() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, -10}
	positive := 0
	negative := 0
	zero := 0

	for _, number := range numbers {
		if number > 0 {
			positive++
		} else if number < 0 {
			negative++
		} else {
			zero++
		}
	}

	fmt.Println("Positive:", positive)
	fmt.Println("Negative:", negative)
	fmt.Println("Zero:", zero)
}

func caseC() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, -10}
	max := 0

	for i, number := range numbers {
		if number > max {
			max = number
			
		}
		fmt.Println("Max:", max, "at index", i)
	}
}

func caseD() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, -10}
	total := 0

	for _, number := range numbers {
		total += number
		
	}
	fmt.Println("Total:", total)

}

func caseE() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total := 0
	totalElement := 0

	for _, number := range numbers {
		total += number
		totalElement++
	}
	totalElements := totalElement + 1
	average := total/totalElements
	fmt.Println("Average:", average)
}


func caseF() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	max := numbers[0]
	min := numbers[0]

	for _, number := range numbers {
		if number > max {
			max = number
			
		}
		
	}

	for _, number := range numbers {
		if number < min {
			min = number
			
		}
		
	}
	fmt.Println("Max:", max, "at index", )
	fmt.Println("Max:", min, "at index", )

	selisih := max - min
	fmt.Println("Selisih:", selisih)


}

func caseG() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 19, 10}
	counter := 0

    n := len(numbers)
    for i := 0; i < n-1; i++ {
		counter++
        for j := 0; j < n-i-1; j++ {
            if numbers[j] > numbers[j+1] {
                // Swap arr[j] and arr[j+1]
                numbers[j] = numbers[j+1] 
				numbers[j+1] = numbers[j]
            }
        }
    }
	fmt.Println("After sorting:", numbers, counter)
}

func caseH() {
	var char = [10]string{"p", "1", "2", "3", "4", "5", "6", "7", "p", "9"}

	for _, character := range char {
		if character == "p" {
			fmt.Print("True")
		} else {
			fmt.Print("False")
		}
	}
}

func caseI() {
	var char = [10]string{"p", "1", "2", "3", "4", "5", "6", "7", "p", "9"}
	counter := 0

	for _, character := range char {
		if character == "p" {
			counter++
		} else {

		}
	}
	fmt.Printf("%d", counter)

}