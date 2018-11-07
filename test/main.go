package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("1+1 =", 1+1)

	var d int = 3
	fmt.Println(d)

	// := is used both for declaration and assignment, here the compiler will infer the type of f to be string
	// this will be equivalent to var f string = "short"
	f := "short"
	fmt.Println(f)

	e := 5
	fmt.Println(e)

	i := 5
	for i < 10 {
		fmt.Println("Current number is ", i)
		i = i + 1
	}

	// arrays
	var a [5] int
	fmt.Println("emp", a)
	a[4] = 100
	fmt.Println("set: ", a)

	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl", b)

	var twoD [2][3]int
	fmt.Println("twoD", twoD)

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[0]); j++ {
			twoD[i][j] = 3
		}
	}
	fmt.Println("twoD", len(twoD[0]), len(twoD), twoD)

	// slices
	s := make([]string, 4)
	fmt.Println("emp:", s)
	s = append(s, "d")
	fmt.Println(s)

	t := []string{"r", "b", "d"}
	fmt.Println(t[:1])
}