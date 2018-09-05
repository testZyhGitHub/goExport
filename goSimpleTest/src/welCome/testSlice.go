package main
import "fmt"

func main() {
	var values []int
	
	values = make([]int, 5)
	
	for i, _ := range values {
		values[i] = i + 1
	}
	
	for i, item := range values {
		fmt.Printf("values[%d]=%d\n", i, item)
	}
}

