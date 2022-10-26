package handlers

import "fmt"

func loop() {
	var fruits = [4]string{"avocado", "apple", "banana", "cherry"}

	for _, fruit := range fruits {
		//fmt.Printf("%d\n", i)
		fmt.Printf("%s\n", fruit)
	}

	bands := make([]string, len(fruits))
	bands[0] = "jannabi"
	bands[1] = "pptnz"
	bands[2] = "reality club"

	for i := 0; i < len(bands); i++ {
		fmt.Printf("%d : %s\n", i, bands[i])
	}

	bands[1] = "peppertones"
	fmt.Printf("%d : %s\n", 1, bands[1])
}
