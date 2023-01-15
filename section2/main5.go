package main

import "fmt"

func main() {
	arr := []string {"1", "2", "3"}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := range arr {
		fmt.Println(i)
	}

	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "YOYOYO"
	mymap["age"] = 33

	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v\n", k, v)
	}
}