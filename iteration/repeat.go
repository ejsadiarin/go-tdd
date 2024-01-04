package iteration

import "fmt"

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		// repeated = repeated + character
		repeated += character
	}
	return repeated
}

// shows for loop variants in Go
func forVariants() {
	// like a while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// normal for (C-like syntax)
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// normal with break statement: shortcircuits the for loop
	for {
		fmt.Println("loop")
		break
	}

	// normal with continue statement: skips specific iteration and "continues" the for loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
