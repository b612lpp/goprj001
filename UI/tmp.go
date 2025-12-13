package main

import "fmt"

func main() {

	i := 42
	q := &i

	fmt.Println(i, &i, *q, q)

}
