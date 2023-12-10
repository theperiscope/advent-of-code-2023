package main

import (
	"AOC/lib"
	"fmt"
	"slices"
	"sort"
	"strings"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()
	hands := parseRows(rows)

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], cardToInt)
	})

	sum := 0
	for i, h := range hands {
		//fmt.Println(i+1, h.Type(), h.cards, h.bid)
		sum += (len(hands) - i) * h.bid
	}
	fmt.Println(sum)

	hands2 := parseRowsPart2(rows)
	sort.Slice(hands2, func(i, j int) bool {
		return compareHands(hands2[i], hands2[j], cardToInt2)
	})
	sum2 := 0
	for i, h := range hands2 {
		//fmt.Println(i+1, h.Type(), h.cards, h.bid)
		sum2 += (len(hands2) - i) * h.bid
	}
	fmt.Println(sum2)
}

type hand struct {
	cards string
	count map[int][]string
	bid   int
}

func (h *hand) Type() int {
	if _, ok := h.count[5]; ok {
		return 1
	} else if _, ok := h.count[4]; ok && len(h.count[1]) == 1 {
		return 2
	} else if _, ok := h.count[3]; ok {
		if _, ok2 := h.count[2]; ok2 {
			return 3
		} else if len(h.count[1]) == 2 {
			return 4
		} else {
			panic("unexpected")
		}
	} else if _, ok := h.count[2]; ok {
		if len(h.count[2]) == 2 && len(h.count[1]) == 1 { // two pairs
			return 5
		} else if len(h.count[2]) == 1 && len(h.count[1]) == 3 { // one pair
			return 6 // one pair
		} else {
			panic("unexpected")
		}
	} else if len(h.count[1]) == 5 {
		return 7 // all different
	}
	panic("unexpected")
}

func compareHands(a, b hand, cardToInt func(string) int) bool {
	aType := a.Type()
	bType := b.Type()
	if aType == bType {
		for i := 0; i < len(a.cards); i++ {
			if a.cards[i] == b.cards[i] {
				continue
			}
			return cardToInt(string(a.cards[i])) < cardToInt(string(b.cards[i]))
		}
		panic("unexpected")
	}

	return aType < bType
}

func parseRows(rows []string) []hand {
	result := []hand{}
	for _, row := range rows {
		items := strings.Split(row, " ")
		x := map[string]int{}
		for _, r := range items[0] {
			x[string(r)]++
		}
		xx := lib.Invert(x)

		result = append(result, hand{cards: items[0], count: xx, bid: lib.Atoi(items[1])})
	}
	return result
}

func parseRowsPart2(rows []string) []hand {
	result := []hand{}
	for _, row := range rows {
		items := strings.Split(row, " ")
		x := map[string]int{}
		for _, r := range items[0] {
			x[string(r)]++
		}

		// how many Js?
		j := x["J"]
		if j > 0 {
			delete(x, "J")
		}

		xx := lib.Invert(x)

		// for JJJJJ (special case) len(xx) will be 0 because we removed all cards from x
		if len(xx) > 0 {
			largestKey := slices.Max(lib.Keys(xx))
			if len(xx[largestKey]) > 1 {
				selectedIndex := 0
				for i := 1; i < len(xx[largestKey]); i++ {
					if cardToInt2(xx[largestKey][i]) < cardToInt2(xx[largestKey][selectedIndex]) {
						selectedIndex = i
					}
				}
				x[xx[largestKey][selectedIndex]] += j
				xx = lib.Invert(x)
			} else {
				selectedIndex := 0
				x[xx[largestKey][selectedIndex]] += j
				xx = lib.Invert(x)
			}
		} else {
			xx = map[int][]string{5: []string{"J"}}
		}

		result = append(result, hand{cards: items[0], count: xx, bid: lib.Atoi(items[1])})
	}
	return result
}

func cardToInt(card string) int {
	m := map[string]int{"A": 1, "K": 2, "Q": 3, "J": 4, "T": 5, "9": 6, "8": 7, "7": 8, "6": 9, "5": 10, "4": 11, "3": 12, "2": 13}
	return m[card]
}

func cardToInt2(card string) int {
	m := map[string]int{"A": 1, "K": 2, "Q": 3, "J": 14, "T": 5, "9": 6, "8": 7, "7": 8, "6": 9, "5": 10, "4": 11, "3": 12, "2": 13}
	return m[card]
}
