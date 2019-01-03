package piscine

import "github.com/01-edu/z01"

func show3(n int, table [9]int, tmax [9]int) {
	i := 0
	for i < n {
		printDigit(table[i])
		i++
	}
	if table[0] != tmax[0] {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

func printComb1() {
	table := [9]int{0}
	tmax := [9]int{9}
	for table[0] <= tmax[0] {
		show3(1, table, tmax)
		table[0]++
	}
}

func PrintCombn(n int) {
	table := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	tmax := [9]int{}

	if n == 1 {
		printComb1()
	} else {
		i := n - 1
		j := 9
		for i >= 0 {
			tmax[i] = j
			i--
			j--
		}
		i = n - 1
		for table[0] < tmax[0] {
			if table[i] != tmax[i] {
				show3(n, table, tmax)
				table[i]++
			}
			if table[i] == tmax[i] {
				if table[i-1] != tmax[i-1] {
					show3(n, table, tmax)
					table[i-1]++
					j = i
					for j < n {
						table[j] = table[j-1] + 1
						j++
					}
					i = n - 1
				}
			}
			for table[i] == tmax[i] && table[i-1] == tmax[i-1] && i > 1 {
				i--
			}
		}
		show3(n, table, tmax)
	}
	z01.PrintRune('\n')
}