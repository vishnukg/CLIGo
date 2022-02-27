package main

import (
	"bufio"
	"io"
)

func main() {

	array := [5]int{10, 20, 30, 40, 50}

	for _, v := range array {
		println(v)
	}

	var array1 [3]*string
	array2 := [3]*string{new(string), new(string), new(string)}

	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"

	array1 = array2

	for _, v := range array1 {
		println(*v)
	}

	slice := []int{10, 20, 30, 40, 50}

	newSlice := slice[1:3]

	println(len(newSlice), len(slice))

	// lines := flag.Bool("l", false, "Count lines")
	// flag.Parse()

	// fmt.Println(count(os.Stdin, *lines))

}

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
