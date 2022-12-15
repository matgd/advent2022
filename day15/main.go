package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/matgd/advent2022/utils"
)

type Field struct {
	x, y          int
	empty         bool
	sensor        bool
	beacon        bool
	closestBeacon *Field
}

type Map struct {
	fields                map[int]map[int]*Field
	nonBeaconFieldsCoords map[int][]int
}

func (m *Map) PrintExampleBoard() {
	for y := -2; y <= 22; y++ {
		// Print 'y' aliged to the right 2 places
		if y >= 0 && y < 10 {
			fmt.Print(" ")
		}
		fmt.Print(y, " ")

		for x := -2; x <= 25; x++ {
			if field, ok := m.fields[y][x]; ok {
				if field.sensor {
					fmt.Print("S")
				} else if field.beacon {
					fmt.Print("B")
				} else if field.empty {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func Day15(input [][4]int) {
	fieldMap := Map{make(map[int]map[int]*Field), make(map[int][]int)}

	for _, line := range input {
		beacon := Field{line[2], line[3], false, false, true, nil}
		sensor := Field{line[0], line[1], false, true, false, &beacon}

		if _, ok := fieldMap.fields[beacon.y]; !ok {
			fieldMap.fields[beacon.y] = make(map[int]*Field)
		}
		fieldMap.fields[beacon.y][beacon.x] = &beacon

		if _, ok := fieldMap.fields[sensor.y]; !ok {
			fieldMap.fields[sensor.y] = make(map[int]*Field)
		}
		fieldMap.fields[sensor.y][sensor.x] = &sensor

		// Distance between sensor and beacon
		distance := int(math.Abs(float64(beacon.x-sensor.x)) + math.Abs(float64(beacon.y-sensor.y)))

		distance_offset := 0
		for y := sensor.y - distance; y <= sensor.y+distance; y++ {
			if y != 2000000 {
				if y < sensor.y {
					distance_offset++
				} else {
					distance_offset--
				}
				continue
			}
			for x := sensor.x - distance_offset; x <= sensor.x+distance_offset; x++ {
				var field *Field = nil
				if x == sensor.x && y == sensor.y {
					field = &sensor
				}
				if x == beacon.x && y == beacon.y {
					field = &beacon
				}
				if field == nil {
					field = &Field{x, y, true, false, false, nil}
				}
				if _, ok := fieldMap.fields[y]; !ok {
					fieldMap.fields[y] = make(map[int]*Field)
				}
				if _, ok := fieldMap.fields[y][x]; !ok {
					fieldMap.fields[y][x] = field
				}
				if !field.beacon {
					fieldMap.nonBeaconFieldsCoords[y] = append(fieldMap.nonBeaconFieldsCoords[y], x)
				}
			}
			if y < sensor.y {
				distance_offset++
			} else {
				distance_offset--
			}
		}
		// fieldMap.PrintExampleBoard()
	}
	// Remove duplicates from nonBeaconFieldsCoords[10]
	y_nonBeaconFieldsCoords := make(map[int]interface{})
	for _, x := range fieldMap.nonBeaconFieldsCoords[2000000] {
		y_nonBeaconFieldsCoords[x] = nil
	}
	sum_y := 0
	// fmt.Println(y10_nonBeaconFieldsCoords)
	for range y_nonBeaconFieldsCoords {
		sum_y++
	}
	fmt.Println(sum_y)
}

func main() {
	useExampleInput := false

	inputFile := "input.txt"
	if useExampleInput {
		inputFile = "input_example.txt"
	}

	lines := utils.ReadFileLines(inputFile)
	cleanInput := make([][4]int, 0)
	for _, line := range lines {
		line = strings.Replace(line, "Sensor at", "", 1)
		line = strings.Replace(line, " closest beacon is at ", "", 1)
		line = strings.ReplaceAll(line, "x=", "")
		line = strings.ReplaceAll(line, " y=", "")
		line = strings.ReplaceAll(line, ":", ",")
		line = strings.ReplaceAll(line, " ", "")
		line = strings.TrimSpace(line)

		// To integers
		lineInts := [4]int{0, 0, 0, 0}
		for i, s := range strings.Split(line, ",") {
			num, _ := strconv.Atoi(s)
			lineInts[i] = num
		}
		cleanInput = append(cleanInput, lineInts)
	}
	Day15(cleanInput)
}
