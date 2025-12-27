package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jhands0/project_euler/solutions"
)

func main() {
	problem, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])

	switch problem {
	case 0:
		fmt.Println(solutions.Sol0(n))
	case 1:
		fmt.Println(solutions.Sol1(n))
	case 2:
		fmt.Println(solutions.Sol2(n))
	case 4:
		fmt.Println(solutions.Sol4(n))
	case 5:
		fmt.Println(solutions.Sol5(n))
	case 6:
		fmt.Println(solutions.Sol6(n))
	default:
		log.Fatal("No arguments provided.")
	}
}
