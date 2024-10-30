package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Ввод:
4
.##.
.##.
.##.
....

Вывод:
I
*/

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Считываем первое число N - размерность матрицы
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Инициализируем матрицу размером NxN
	matrix := make([]string, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		matrix[i] = line
	}

	fmt.Println(f(matrix))
}

func f(matrix []string) string {
	templates := [6][]string{
		{"#"},
		{"###", "#.#", "###"},
		{"##", "#.", "##"},
		{"#.", "##"},
		{"#.#", "###", "#.#"},
		{"###", "#.#", "###", "#.."},
	}

	dict := map[int]string{0: "I", 1: "O", 2: "C", 3: "L", 4: "H", 5: "P"}
	matrix = compress(&matrix)
	matrix = invert(&matrix)
	matrix = compress(&matrix)
	matrix = invert(&matrix)
	for i, template := range templates {
		if eq(&matrix, &template) {
			return dict[i]
		}
	}
	return "X"
}

func eq(m1 *[]string, m2 *[]string) bool {
	if len(*m1) != len(*m2) || len((*m1)[0]) != len((*m2)[0]) {
		return false
	}
	for i := range *m1 {
		for j := range (*m1)[i] {
			if (*m1)[i][j] != (*m2)[i][j] {
				return false
			}
		}
	}
	return true
}

func compress(matrix *[]string) []string {
	ans := []string{(*matrix)[0]}
	for _, now := range (*matrix)[1:] {
		if now != ans[len(ans)-1] {
			ans = append(ans, now)
		}
	}
	if len(ans) > 1 && ans[0] == strings.Repeat(".", len(ans[0])) {
		ans = ans[1:]
	}

	if len(ans) > 1 && ans[len(ans)-1] == strings.Repeat(".", len(ans[0])) {
		ans = ans[:len(ans)-1]
	}
	return ans
}

func invert(matrix *[]string) []string {
	ans := [][]string{}
	for range len((*matrix)[0]) {
		ans = append(ans, []string{})
	}
	for _, now := range *matrix {
		for idx, item := range now {
			ans[idx] = append(ans[idx], string(item))
		}
	}
	ret := []string{}
	for idx := range ans {
		ret = append(ret, strings.Join(ans[idx], ""))
	}

	return ret
}
