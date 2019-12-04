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

}
