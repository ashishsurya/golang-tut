package main

import "fmt"


func birthday(age *int) {
	*age = *age + 1
}

func main()  {
	age := 1
	fmt.Println(age)
	birthday(&age)
	fmt.Println(age)
}