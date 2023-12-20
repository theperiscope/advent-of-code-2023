//go:build ignore

package main

import (
	"AOC/interval"
	"AOC/lib"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	ACTION_WORKFLOW string = "WORKFLOW"
	ACTION_ACCEPT   string = "A"
	ACTION_REJECT   string = "R"
)

func main() {
	lib.AssertArgs()
	rows := lib.AssertInput()
	emptyRowIndex := slices.IndexFunc(rows, func(r string) bool { return r == "" })
	workflowStrings := rows[0:emptyRowIndex]
	ratingStrings := rows[emptyRowIndex+1:]
	_, _ = workflowStrings, ratingStrings
	workflows := map[string]workflow{}
	for _, w := range workflowStrings {
		wf := parseWorkflow(w)
		fmt.Println(wf)
		workflows[wf.name] = wf
	}

	ratings := []rating{}
	for _, r := range ratingStrings {
		ratings = append(ratings, parseRating(r))
	}
	fmt.Println(ratings)
	//fmt.Println(workflowMap)

	sum := 0
	w := workflows["in"]
	for _, rating := range ratings {
		fmt.Println(rating)
		found := ""
		for found == "" {
			applicableRule := rule{}
			for _, rr := range w.rules {
				if !rr.isApplicable(rating) {
					continue
				}
				applicableRule = rr
				//fmt.Println("Applicable rule: ", applicableRule)
				break
			}
			if applicableRule.act.kind() == ACTION_ACCEPT || applicableRule.act.kind() == ACTION_REJECT {
				found = applicableRule.act.kind()
				break
			}
			w = workflows[applicableRule.act.kind()]
		}
		if found == ACTION_ACCEPT {
			sum += rating.Sum()
		}
		//fmt.Println("- - -")
		w = workflows["in"]
	}
	fmt.Println("Sum", sum)

	part2(workflows)
}

func parseRating(s string) rating {
	var re = regexp.MustCompile(`(?m)\d+`)
	m, _ := lib.Convert(re.FindAllString(s, -1), strconv.Atoi)
	return rating{x: m[0], m: m[1], a: m[2], s: m[3]}
}

func parseWorkflow(s string) workflow {
	ci := strings.Index(s, "{")
	name := s[0:ci]
	stringRules := strings.Split(getStringBetweenBraces(s), ",")
	var re = regexp.MustCompile(`(?m)(x|m|a|s)(<|>)(\d+)\:(A|R|[a-z]+)`)
	rules := []rule{}
	for _, r := range stringRules {
		matches := re.FindAllStringSubmatch(r, -1)
		if len(matches) > 0 {
			for _, match := range matches {
				//fmt.Println(match, "found at index", i)
				rr := rule{variable: match[1], op: match[2], value: lib.Atoi(match[3])}
				switch match[4] {
				case "A":
					rr.act = acceptAction{}
				case "R":
					rr.act = rejectAction{}
				default:
					rr.act = workflowAction{workflowName: match[4]}
				}
				rules = append(rules, rr)
			}
		} else { // no match (for last rule)
			rr := rule{}
			switch r {
			case "A":
				rr.act = acceptAction{}
			case "R":
				rr.act = rejectAction{}
			default:
				rr.act = workflowAction{workflowName: r}
			}
			rules = append(rules, rr)
		}
	}
	return workflow{name: name, rules: rules}
}

func getStringBetweenBraces(s string) string {
	regex := regexp.MustCompile(`\{([^}]+)\}`)
	match := regex.FindStringSubmatch(s)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

type workflow struct {
	name  string
	rules []rule
}

type rule struct {
	variable string
	op       string
	value    int
	act      action
}

type rating struct {
	x, m, a, s int
}

type action interface {
	// return "A", "R", or another workflow name
	kind() string
}

type workflowAction struct{ workflowName string }
type acceptAction struct{}
type rejectAction struct{}

func (w workflow) String() string {
	return fmt.Sprintf("%s\t%v", w.name, w.rules)
}

func (r rating) Sum() int { return r.x + r.m + r.a + r.s }

func (r rating) String() string {
	return fmt.Sprintf("x=%d,m=%d,a=%d,s=%d", r.x, r.m, r.a, r.s)
}

func (r rule) String() string {
	if r.variable == "" {
		return r.act.kind()
	}

	return fmt.Sprintf("%s%s%d:%s", r.variable, r.op, r.value, r.act.kind())
}

func (r rule) isApplicable(theRating rating) bool {
	if r.variable == "" {
		return true
	}

	switch {
	case r.variable == "x" && r.op == "<":
		return theRating.x < r.value
	case r.variable == "m" && r.op == "<":
		return theRating.m < r.value
	case r.variable == "a" && r.op == "<":
		return theRating.a < r.value
	case r.variable == "s" && r.op == "<":
		return theRating.s < r.value
	case r.variable == "x" && r.op == ">":
		return theRating.x > r.value
	case r.variable == "m" && r.op == ">":
		return theRating.m > r.value
	case r.variable == "a" && r.op == ">":
		return theRating.a > r.value
	case r.variable == "s" && r.op == ">":
		return theRating.s > r.value
	default:
		panic("oh nooo")
	}
}

func (a workflowAction) kind() string { return a.workflowName }
func (a acceptAction) kind() string   { return ACTION_ACCEPT }
func (a rejectAction) kind() string   { return ACTION_REJECT }

type workItem struct {
	workflow       string
	possibleValues map[string]*interval.Interval
}

func part2(workflows map[string]workflow) {
	q := []workItem{{
		workflow: "in",
		possibleValues: map[string]*interval.Interval{
			"x": {Start: 1, End: 4000},
			"m": {Start: 1, End: 4000},
			"a": {Start: 1, End: 4000},
			"s": {Start: 1, End: 4000},
		}}}
	sum := int64(0)
	for len(q) > 0 {
		item := q[0]
		q = lib.Remove(q, 0)

		for _, r := range workflows[item.workflow].rules {
			result, remainder := processRule(item.possibleValues, r)
			fmt.Println(result)
			fmt.Println(remainder)
			fmt.Println("===")
			if r.act.kind() == ACTION_ACCEPT {
				sum += result["x"].Len() * result["m"].Len() * result["a"].Len() * result["s"].Len()
				fmt.Println("[A]", result["x"].Len()*result["m"].Len()*result["a"].Len()*result["s"].Len())
				item.possibleValues = remainder
			} else if r.act.kind() == ACTION_REJECT {
				fmt.Println("[R]", result)
				item.possibleValues = remainder
			} else {
				q = append(q, workItem{workflow: r.act.kind(), possibleValues: result})
				item.possibleValues = remainder
			}
			if len(remainder) == 0 {
				break
			}
		}
	}
	fmt.Println("Part 2", sum)
}

func processRule(intervals map[string]*interval.Interval, r rule) (result, remainder map[string]*interval.Interval) {
	if r.variable != "" {
		i := intervals[r.variable]
		result, remainder = make(map[string]*interval.Interval), make(map[string]*interval.Interval)
		for k, v := range intervals {
			result[k] = v.Clone()
			remainder[k] = v.Clone()
		}
		switch r.op {
		case "<":
			split := i.Split(int64(r.value))
			result[r.variable] = &split[0]
			remainder[r.variable] = &split[1]
			return result, remainder
		case ">":
			split := i.Split(int64(r.value) + 1)
			result[r.variable] = &split[1]
			remainder[r.variable] = &split[0]
			return result, remainder
		default:
			panic("oh no")
		}
	}
	result = make(map[string]*interval.Interval)
	for k, v := range intervals {
		result[k] = v.Clone()
	}
	return result, make(map[string]*interval.Interval)
}
