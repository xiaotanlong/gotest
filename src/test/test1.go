package main

import (
	"fmt"
	"os"
)

func main() {
	var JAVAHOME string
	JAVAHOME = os.Getenv("JAVA_HOME")
	fmt.Println(JAVAHOME)


}