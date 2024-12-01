package year2024

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	s "strings"
)

type Day01 struct{}

func (p Day01) PartA(lines []string) any {
	var a []int
	var b []int

	fmt.Println("len:", len(lines))

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		var ab = s.Split(lines[i], "   ")
		var i, _ = strconv.Atoi(ab[0])
		a = append(a, i)
		i, _ = strconv.Atoi(ab[1])
		b = append(b, i)
	}
	slices.Sort(a)
	slices.Sort(b)
	var total int = 0
	for i := 0; i < len(a); i++ {
		total += int(math.Abs(float64(a[i] - b[i])))
	}
	return total
}

func (p Day01) PartB(lines []string) any {
	var a []int
	var b []int

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		var ab = s.Split(lines[i], "   ")
		var i, _ = strconv.Atoi(ab[0])
		a = append(a, i)
		i, _ = strconv.Atoi(ab[1])
		b = append(b, i)
	}

	var total int = 0
	m := make(map[int]int)
	for i := 0; i < len(b); i++ {
		m[b[i]]++
	}
	for i := 0; i < len(a); i++ {
		total += a[i] * m[a[i]]
	}
	return total
}
