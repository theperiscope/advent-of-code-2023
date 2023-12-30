//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
	"strings"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()
	part2 := Solve(rows)
	fmt.Println("Part 2", part2)
}

type Modules map[string]*Module
type Module struct {
	Type         ModuleType
	State        bool
	Destinations []string
	Inputs       map[string]Pulse
}
type Signal struct {
	From  string
	To    string
	Pulse Pulse
}
type ModuleType rune
type Pulse byte
type Queue []Signal

const (
	BROADCASTER = ModuleType('b')
	FLIPFLOP    = ModuleType('%')
	CONJUCTION  = ModuleType('&')

	LOW  = Pulse(0)
	HIGH = Pulse(1)
)

func Solve(input []string) int {
	modules := parseInput(input)

	watch := map[string]bool{}
	for d := range modules["rx"].Inputs {
		for destination := range modules[d].Inputs {
			watch[destination] = false
		}
	}

	results := []int{}
	for i := 1; len(results) < len(watch); i++ {
		signals := modules.PressButton()
		for _, signal := range signals {
			if signal.Pulse == HIGH {
				if found, ok := watch[signal.From]; ok && !found {
					results = append(results, i)
					watch[signal.From] = true
				}
			}
		}
	}
	fmt.Println(results)
	return getLeastCommonMultiple(results)
}

func parseInput(lines []string) Modules {
	modules := make(Modules, len(lines))
	modules["rx"] = &Module{
		Inputs: map[string]Pulse{},
	}

	for _, line := range lines {
		line = strings.ReplaceAll(line, " ", "")
		parts := strings.Split(line, "->")
		destinations := strings.Split(parts[1], ",")
		module := Module{
			Type:         ModuleType(parts[0][0]),
			Destinations: destinations,
			Inputs:       map[string]Pulse{},
		}
		key := parts[0][1:]
		if module.Type == BROADCASTER {
			key = parts[0]
		}
		modules[key] = &module
	}

	for key, module := range modules {
		for _, destination := range module.Destinations {
			modules[destination].Inputs[key] = LOW
		}
	}

	return modules
}

func (modules Modules) PressButton() []Signal {
	queue := Queue{}
	for _, d := range modules["broadcaster"].Destinations {
		queue = append(queue, Signal{
			From:  "broadcaster",
			To:    d,
			Pulse: LOW,
		})
	}

	return queue.Read(modules, []Signal{})
}

func (queue Queue) Read(modules Modules, signals []Signal) []Signal {
	newQueue := Queue{}
	for _, signal := range queue {
		signals = append(signals, signal)
		to, ok := modules[signal.To]
		if !ok {
			continue
		}

		hasPulse := false
		pulse := LOW
		if to.Type == FLIPFLOP && signal.Pulse == LOW {
			to.State = !to.State
			hasPulse = true
			if to.State {
				pulse = HIGH
			}
		} else if to.Type == CONJUCTION {
			to.Inputs[signal.From] = signal.Pulse
			hasPulse = true
			for _, input := range to.Inputs {
				if input == LOW {
					pulse = HIGH
					break
				}
			}
		}

		if hasPulse {
			for _, destination := range to.Destinations {
				newQueue = append(newQueue, Signal{
					From:  signal.To,
					To:    destination,
					Pulse: pulse,
				})
			}
		}
	}

	if len(newQueue) == 0 {
		return signals
	}

	return newQueue.Read(modules, signals)
}

func getLeastCommonMultiple(numbers []int) int {
	lcm := numbers[0]
	for i := 0; i < len(numbers); i++ {
		num1 := lcm
		num2 := numbers[i]
		gcd := 1
		for num2 != 0 {
			temp := num2
			num2 = num1 % num2
			num1 = temp
		}
		gcd = num1
		lcm = (lcm * numbers[i]) / gcd
	}

	return lcm
}

// From exp/maps package
func mapValues[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
