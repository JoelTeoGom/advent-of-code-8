package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Map struct {
	Right, Left string
}

func readMap(m *map[string]Map, command *string, filename string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		*command = scanner.Text()
		fmt.Println(*command)
	}

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		lines := make([]string, 0)
		for i, char := range line {
			fmt.Printf("[%d]: %s\n", i, string(char))
			lines = append(lines, string(char))
		}

		key := make([]string, 0)   // Inicializa key
		right := make([]string, 0) // Inicializa right
		left := make([]string, 0)  // Inicializa left

		var leftString, rightString, keyString string
		// verify to avoid PANIQUING
		if len(lines) >= 16 {
			key = append(key, lines[0])
			key = append(key, lines[1])
			key = append(key, lines[2])

			left = append(left, lines[7])
			left = append(left, lines[8])
			left = append(left, lines[9])

			right = append(right, lines[12])
			right = append(right, lines[13])
			right = append(right, lines[14])

			// Convertir slices a strings
			leftString = strings.Join(left, "") // Usando ", " como separador
			rightString = strings.Join(right, "")
			keyString = strings.Join(key, "")

		}

		if len(key) > 0 && len(left) > 0 && len(right) > 0 {
			(*m)[keyString] = Map{Left: leftString, Right: rightString}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}

func main() {

	var command string
	m := make(map[string]Map)

	readMap(&m, &command, "map.txt")
	// Imprimir el mapa
	for k, v := range m {
		fmt.Printf("Clave: %s, Left: %v, Right: %v\n", k, v.Left, v.Right)
	}

}
