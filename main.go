package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Map struct {
	Right, Left string
}

func printMap(m *map[string]Map) {
	for k, v := range *m {
		fmt.Printf("Key: %s, Left: %v, Right: %v\n", k, v.Left, v.Right)
	}
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

func day8part1() {
	var command string
	m := make(map[string]Map)
	readMap(&m, &command, "map.txt")
	index := resolvePuzzle(&m, &command)
	println("Index: ", index)
}

////////////////////////////////////////////////////////////////////////
//							  PART 2								 //
///////////////////////////////////////////////////////////////////////

func resolvePuzzlePart2(m *map[string]Map, command *string) int {

	startKey := make([]string, 0, 3)
	endKey := make([]string, 0, 3)
	var endBool bool = false
	index := 0

	//START POINT
	for key := range *m {
		if strings.HasSuffix(key, "A") { //ENDS WITH A
			startKey = append(startKey, key)
		}
	}

	for key := range *m {
		if strings.HasSuffix(key, "Z") { //ENDS WITH A
			endKey = append(endKey, key)
		}
	}

	fmt.Println(startKey)
	fmt.Println(endKey)
	time.Sleep(10 * time.Second)

	//start puzzle
	for !endBool {
		for _, char := range *command {
			fmt.Println("", startKey)
			tempBool := true
			for _, currentKey := range startKey {
				if !strings.HasSuffix(currentKey, "Z") {
					tempBool = false
				}
			}
			time.Sleep(1 * time.Second)
			endBool = tempBool

			if endBool {
				break
			}

			if string(char) == "L" {
				for i := 0; i < len(startKey); i++ {
					startKey[i] = (*m)[startKey[i]].Left
				}

			} else {
				for i := 0; i < len(startKey); i++ {
					startKey[i] = (*m)[startKey[i]].Right
				}
			}

			index++
			fmt.Printf("| [%d]: %c \n", index, char)
		}
	}
	return index

}

func day8part2() {
	var command string
	m := make(map[string]Map)
	readMap(&m, &command, "map.txt")
	index := resolvePuzzlePart2(&m, &command)
	println("Index: ", index)
}

func main() {
	//day8part1()
	day8part2()

}
