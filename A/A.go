package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Посередине озера плавает плот, имеющий форму прямоугольника.
Стороны плота направлены вдоль параллелей и меридианов.
Введём систему координат, в которой ось OX направлена на восток,
а ось ОY – на север. Пусть юго-западный угол плота имеет координаты (x1, y1)
северо-восточный угол – координаты (x2, y2)

Пловец находится в точке с координатами (x, y).
Определите, к какой стороне плота (северной, южной, западной или восточной)
или к какому углу плота (северо-западному, северо-восточному,
юго-западному, юго-восточному) пловцу нужно плыть, чтобы
как можно скорее добраться до плота.
Ввод:
-1
-2
5
3
-4
6
Вывод:
NW
*/

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inp [6]int

	for i := 0; i < 6; i++ {
		scanner.Scan()
		input := scanner.Text()
		num, _ := strconv.Atoi(input)
		inp[i] = num
	}

	x1, y1, x2, y2, x, y := inp[0], inp[1], inp[2], inp[3], inp[4], inp[5]

	first := ""
	second := ""
	if x < x1 {
		second = "W"
	} else if x > x2 {
		second = "E"
	}
	if y < y1 {
		first = "S"
	} else if y > y2 {
		first = "N"
	}

	fmt.Println(first + second)
}
