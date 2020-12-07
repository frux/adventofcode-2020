package day7

import (
	"adventofcode-2020/input"
	"fmt"
	"strconv"
	"strings"
)

const INPUT = "./day7/input.txt"
const MY_BAG_COLOR = "shiny gold"

func Part1() string {
	entries := input.ReadStrings(INPUT)

	bags := parseRules(entries)
	parents := flattenParents(getBag(bags, MY_BAG_COLOR).Parents)

	counted := map[string]bool{}
	count := 0
	for _, parent := range parents {
		if counted[parent.Color] {
			continue
		}
		counted[parent.Color] = true
		count++
	}

	return fmt.Sprint(count)
}
func Part2() string {
	entries := input.ReadStrings(INPUT)

	bags := parseRules(entries)
	count := countChildren(getBag(bags, MY_BAG_COLOR))

	return fmt.Sprint(count)
}

func flattenParents(bags []*Bag) []*Bag {
	parents := []*Bag{}
	bagsToCount := make([]*Bag, len(bags))
	copy(bagsToCount, bags)

	for i := 0; i < len(bagsToCount); i++ {
		child := bagsToCount[i]
		parents = append(parents, child)
		bagsToCount = append(bagsToCount, child.Parents...)
	}

	return parents
}

func countChildren(parent *Bag) int {
	count := 0

	for _, child := range parent.Children {
		count += child.Capacity * (1 + countChildren(child.Bag))
	}

	return count
}

type Bags = map[string]*Bag
type Bag struct {
	Color    string
	Children []ChildBag
	Parents  []*Bag
}
type ChildBag struct {
	Bag      *Bag
	Capacity int
}

func parseRules(rules []string) Bags {
	bags := make(Bags)

	for _, rule := range rules {
		parts := strings.Split(rule[0:len(rule)-1], " bags contain ")
		bagColor := parts[0]
		bag := getBag(bags, bagColor)

		if parts[1] == "no other bags" {
			continue
		}

		children := strings.Split(parts[1], ", ")

		for _, child := range children {
			parts = strings.Split(child, " ")
			capacity, _ := strconv.Atoi(parts[0])
			childColor := strings.Join(parts[1:len(parts)-1], " ")
			bag.addChild(&bags, childColor, capacity)
		}
	}

	return bags
}

func (bag *Bag) addChild(bags *Bags, childColor string, capacity int) {
	childBag := getBag(*bags, childColor)
	bag.Children = append(
		bag.Children,
		ChildBag{
			Bag:      childBag,
			Capacity: capacity,
		},
	)
	childBag.Parents = append(childBag.Parents, bag)
}

func getBag(bags Bags, color string) *Bag {
	if bags[color] == nil {
		bags[color] = createBag(color)
	}

	return bags[color]
}

func createBag(color string) *Bag {
	return &Bag{
		Color: color,
	}
}
