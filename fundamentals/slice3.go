package main

import "fmt"

type Part string

func showLine(line []Part) {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}

}

func main() {
	assemblyLine := make([]Part, 3)

	assemblyLine[0] = "peter"
	assemblyLine[1] = "scar"
	assemblyLine[2] = "jay"

	fmt.Println("------The first three parts-------")
	showLine(assemblyLine)

	assemblyLine = append(assemblyLine, "mum", "dad")
	fmt.Println("------The New parts-------")
	showLine(assemblyLine)

	assemblyLine = assemblyLine[3:]
	fmt.Println("-------Two new parts---------")
	showLine(assemblyLine)

}
