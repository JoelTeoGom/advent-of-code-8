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
		// fmt.Println(*command)
	}

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		lines := make([]string, 0)
		for _, char := range line {
			// fmt.Printf("[%d]: %s\n", i, string(char))
			lines = append(lines, string(char))
		}

		// Initialize slices with capacity for 3 elements
		key := make([]string, 0, 3)
		left := make([]string, 0, 3)
		right := make([]string, 0, 3)

		var leftString, rightString, keyString string
		// Verify to avoid panicking
		if len(lines) >= 16 {

			key = append(key, lines[0], lines[1], lines[2])
			left = append(left, lines[7], lines[8], lines[9])
			right = append(right, lines[12], lines[13], lines[14])

			// Convert slices to strings
			leftString = strings.Join(left, "") // Using ", " as a separator
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

func resolvePuzzle(m *map[string]Map, command *string) int {
	var endBool bool = true
	currentKey := "AAA"
	index := 0

	for endBool {
		for _, char := range *command {
			fmt.Println(currentKey)

			if currentKey == "ZZZ" {
				endBool = false
				break
			}
			if string(char) == "L" {
				currentKey = (*m)[currentKey].Left
			} else {
				currentKey = (*m)[currentKey].Right
			}
			index++
			fmt.Printf("| [%d] \n", index)
		}
	}
	return index
}

func resolvePuzzlePart2(m *map[string]Map, command *string) int {

}

func day8part1() {

	var command string
	m := make(map[string]Map)

	readMap(&m, &command, "map.txt")

	/*for k, v := range m {
		fmt.Printf("Key: %s, Left: %v, Right: %v\n", k, v.Left, v.Right)
	}*/
	index := resolvePuzzle(&m, &command)
	println("Index: ", index)
}

func day8part2() {
	var command string
	m := make(map[string]Map)

	readMap(&m, &command, "map.txt")

}

func main() {
	//day8part1()
	day8part2()

}
