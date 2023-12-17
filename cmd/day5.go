//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()

	currentInputSectionNumber := 1
	seeds := []int{}
	_ = seeds
	remaps := make([][]rule, 0)
	fmt.Println("Section", currentInputSectionNumber)
	for _, row := range rows {
		if len(row) == 0 {
			currentInputSectionNumber++
			fmt.Println()
			fmt.Println("Section", currentInputSectionNumber)
			continue
		}

		if currentInputSectionNumber == 1 { // seeds
			input := strings.Split(row[strings.Index(row, ":")+2:], " ")
			data, err := lib.Convert[string, int](input, strconv.Atoi)
			if err != nil {
				panic(err)
			}
			fmt.Println(data)
			seeds = data
		} else {
			if row[len(row)-1] == ':' { // skip map name row, add new remap element
				remaps = append(remaps, make([]rule, 0))
				continue
			}
			input := strings.Split(row, " ")
			data, err := lib.Convert[string, int](input, strconv.Atoi)
			if err != nil {
				panic(err)
			}
			fmt.Println(data)
			r := rule{target: data[0], source: data[1], count: data[2]}
			remaps[currentInputSectionNumber-2] = append(remaps[currentInputSectionNumber-2], r)
		}
	}

	locations := []int{}
	for _, seed := range seeds {
		s := seed
		for _, rules := range remaps {
			for _, r := range rules {
				var remapped bool
				s, remapped = r.remap(s)
				if remapped {
					break
				}
			}
		}
		fmt.Println("Seed", seed, "Location", s)
		locations = append(locations, s)
	}

	fmt.Println(locations)
	fmt.Println(lib.Min(locations...))

	locations = []int{}
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j <= seeds[i]+seeds[i+1]-1; j++ {
			seed := j
			s := seed
			for _, rules := range remaps {
				for _, r := range rules {
					var remapped bool
					s, remapped = r.remap(s)
					if remapped {
						break
					}
				}
			}
			//fmt.Println("Seed", seed, "Location", s)
			locations = append(locations, s)
		}
	}

	//fmt.Println(locations)
	fmt.Println(lib.Min(locations...))
}

type rule struct {
	target int
	source int
	count  int
}

func (r *rule) remap(n int) (int, bool) {
	if n >= r.source && n <= r.source+r.count-1 {
		remapBy := r.target - r.source
		return n + remapBy, true
	}
	return n, false
}
