package main

import "fmt"

func main() {
	name := "Georgeous"
	fmt.Println(name)

	var (
		age       = 42
		hairstyle = 1
	)

	fmt.Println(age + hairstyle)

	const fingerprint string = "asodhbfauiydfhoaiemfkjhandflkjan"
	d := 32

	fmt.Println(d)

	for i := 0; i < 10; i++ {
		fmt.Println(d)
		d += 1
	}

	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		} else if n%2 != 0 {
			fmt.Println("The number ", n, "is even")
		} else {
			fmt.Println("there is a bug in the matrix")
		}
	}

	if len(fingerprint) > 0 {
		fmt.Println("The guy has a fingerprint")
	}

	var pets [5]int

	fmt.Println(pets)

	for i := 0; i < len(pets); i++ {
		pets[i] = i
	}

	fmt.Println(pets)

	var tower [2][2]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			tower[i][j] = i + j
		}
	}

	fmt.Println(tower)

	// Slices are dynamical arrays
	slice := make([]int, 3)
	slice[0] = 0
	slice[1] = 1
	slice[2] = 2
	fmt.Println("Basic slice: ", slice)

	// Making an existing slice overwrites its elements with 0
	slice = make([]int, 5)
	slice[3] = 3
	slice[4] = 4
	fmt.Println("New make, before append: ", slice)

	slice = append(slice, 3)
	fmt.Println("After append: ", slice)

	copiedSlice := make([]int, len(slice))
	fmt.Println("Copied slice after make: ", copiedSlice)
	copy(copiedSlice, slice)
	fmt.Println("Copied slice after copy: ", copiedSlice)

	fmt.Println("Last three elements: ", copiedSlice[2:6])

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	// The length of inner slices varies, unlike with multi-dimensional arrays
	fmt.Println("2d: ", twoD)

}
