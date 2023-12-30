//go:build ignore

package main

import (
	"AOC/lib"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()
	part1 := solve(rows)
	fmt.Println("Part 1", part1)
}

type signal struct {
	from  string
	to    string
	value uint8
}
type module struct {
	kind         byte
	destinations []string
	isOn         bool             // applies only for FLIPFLOP module type
	inputs       map[string]uint8 // applies only to CONJUNCTION module type
	minRx        int              // applies only to "rx" module (part 2)
}
type modulesMap map[string]*module
type queue []signal

const (
	KIND_BROADCASTER = 'b'
	KIND_FLIPFLOP    = '%'
	KIND_CONJUCTION  = '&'
	BROADCASTER_NAME = "broadcaster"

	PULSE_LOW  = uint8(0)
	PULSE_HIGH = uint8(1)
)

func solve(rows []string) int {
	modules := parseInput(rows)

	low, high := 0, 0
	for i := 0; i < 1_000; i++ {
		l, h := modules.doPressButton()
		low += l
		high += h
	}

	return low * high
}

func parseInput(rows []string) modulesMap {
	modules := make(modulesMap, len(rows))
	var re = regexp.MustCompile(`(%|&|b)(.+)\s->\s(.+)`)
	for _, row := range rows {
		if strings.Index(row, BROADCASTER_NAME) == 0 {
			row = "b" + row // add "missing" type for broadcaster
		}
		matches := re.FindAllStringSubmatch(row, -1)
		destinations := strings.Split(matches[0][3], ", ")
		module := module{
			kind:         matches[0][1][0],
			destinations: destinations,
			inputs:       map[string]uint8{},
			minRx:        math.MaxInt,
		}
		modules[matches[0][2]] = &module
	}

	for name, module := range modules {
		for _, d := range module.destinations {
			if _, ok := modules[d]; ok {
				modules[d].inputs[name] = PULSE_LOW
			}
		}
	}

	return modules
}

func (modules modulesMap) doPressButton() (int, int) {
	queue := queue{}
	for _, d := range modules[BROADCASTER_NAME].destinations {
		queue = append(queue, signal{
			from:  BROADCASTER_NAME,
			to:    d,
			value: PULSE_LOW,
		})
	}

	return queue.processSignals(modules, 1, 0)
}

func (q queue) processSignals(modules modulesMap, l, h int) (int, int) {
	nextQueue := queue{}
	for _, s := range q {
		if s.value == PULSE_LOW {
			l++
		} else {
			h++
		}

		to, ok := modules[s.to]
		if !ok { // we have modules that are undefined in the map (e.g. output, rx)
			fmt.Println(s.to, l, h)
			continue
		}

		whatPulseToSendNext := PULSE_LOW
		switch {
		case to.kind == KIND_FLIPFLOP && s.value == PULSE_LOW:
			to.isOn = !to.isOn
			if to.isOn {
				whatPulseToSendNext = PULSE_HIGH
			}
			for _, d := range to.destinations {
				nextQueue = append(nextQueue, signal{from: s.to, to: d, value: whatPulseToSendNext})
			}
		case to.kind == KIND_CONJUCTION:
			to.inputs[s.from] = s.value
			for _, i := range to.inputs {
				if i == PULSE_LOW {
					whatPulseToSendNext = PULSE_HIGH
					break
				}
			}
			for _, d := range to.destinations {
				nextQueue = append(nextQueue, signal{from: s.to, to: d, value: whatPulseToSendNext})
			}
		}
	}

	if len(nextQueue) == 0 {
		return l, h
	}

	return nextQueue.processSignals(modules, l, h)
}
