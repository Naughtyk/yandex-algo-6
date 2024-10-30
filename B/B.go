package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Как известно, осенью и зимой светает поздно,
и так хочется утром ещё хоть немного поспать,
а не идти в школу! Некоторые школьники готовы
даже одеваться, не открывая глаз, лишь бы отложить
момент пробуждения. Вот и Саша решил, что майку и
носки он вполне может вытащить из шкафа на ощупь с
закрытыми глазами и только потом включить свет и одеться.
В шкафу у Саши есть два ящика. В одном из них лежит AA синих и BB красных маек,
в другом — CC синих и DD красных пар носков.
Саша хочет, чтобы и майка, и носки были одного цвета.
Он вслепую вытаскивает MM маек и NN пар носков.
В первое же утро Саша задумался, какое минимальное суммарное количество
предметов одежды (M+N) он должен вытащить, чтобы среди них гарантированно оказались
майка и носки одного цвета. Какого именно цвета окажутся предметы одежды, для Саши совершенно неважно.

Ввод:
6
2
7
3

Вывод:
3 4
*/

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inp [4]int

	for i := 0; i < 4; i++ {
		scanner.Scan()
		input := scanner.Text()
		num, _ := strconv.Atoi(input)
		inp[i] = num
	}

	A, B, C, D := inp[0], inp[1], inp[2], inp[3]

	fmt.Println(f(A, B, C, D))
}

func f(A, B, C, D int) string {
	if A == 0 {
		return fmt.Sprintf("%d %d", 1, C+1)
	} else if C == 0 {
		return fmt.Sprintf("%d %d", A+1, 1)
	} else if B == 0 {
		return fmt.Sprintf("%d %d", 1, D+1)
	} else if D == 0 {
		return fmt.Sprintf("%d %d", B+1, 1)
	}

	x1 := B + D
	x2 := A + C
	x3 := max(A, B)
	x4 := max(C, D)
	m := min(x1, x2, x3, x4)
	if x1 == m {
		return fmt.Sprintf("%d %d", B+1, D+1)
	}
	if x2 == m {
		return fmt.Sprintf("%d %d", A+1, C+1)
	}
	if x3 == m {
		return fmt.Sprintf("%d %d", max(A, B)+1, 1)
	}
	if x4 == m {
		return fmt.Sprintf("%d %d", 1, max(C, D)+1)
	}
	return ""
}
