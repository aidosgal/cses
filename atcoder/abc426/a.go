package main

import "fmt"

func main() {
	mp := map[string]int{"Ocelot": 0, "Serval": 1, "Lynx": 2,}

	var x, y string
	fmt.Scan(&x, &y)
	if (mp[x] >= mp[y]) {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}
