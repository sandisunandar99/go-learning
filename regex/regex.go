package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "nana doro doremi fasole dodoku"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		panic(err.Error())
	}

	var matches = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", matches)

}
