package main
import "fmt"

func check(index int, solution [8]int) bool {
	for i:=0; i< index; i++ {
		if (solution[i] == solution[index] || solution[i] + (index - i) == solution[index] || solution[i] - (index - i) == solution[index]){
			return false
		}
	}
	return true
}

func eightqueens(index int, solution[8]int) {
	if (index == 8 ) {
		for i:= 0; i<8;i++ {
			fmt.Print(solution[i])
		}
		fmt.Println()
	} else {
		for  i:=1;i<9;i++ {
			solution[index] = i
			if (check(index,solution)) {
				eightqueens(index+1,solution)
			}
			solution[index] = 0
		}
	}
}

func main() {
	solution := [8]int{}
	index := 0

	eightqueens(index, solution)
}