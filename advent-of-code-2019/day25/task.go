package day25

import (
	"fmt"
	"slices"
	"strings"

	intcodecomputer "github.com/sogard-dev/advent-of-code-2019/intcode"
	"github.com/sogard-dev/advent-of-code-2019/utils"
)

var DIRECTIONS = [...]string{"east", "south", "west", "north"}

type grid struct {
	connections map[*node]map[string]*node
	nodes       map[string]*node
}

func (g *grid) getNode(s string) *node {
	if n, exists := g.nodes[s]; exists {
		return n
	}
	n := node{
		name:  s,
		items: []string{},
	}
	g.nodes[s] = &n
	return &n
}

func newGrid() grid {
	return grid{
		nodes:       map[string]*node{},
		connections: map[*node]map[string]*node{},
	}
}

func (g *grid) addConnection(from *node, to *node, d string) {
	if _, exist := g.connections[from]; !exist {
		g.connections[from] = map[string]*node{}
	}
	g.connections[from][d] = to
}

type node struct {
	name         string
	visited      bool
	isCheckpoint bool
	isPressure   bool
	items        []string
}

func (g grid) bfs(me *node, found func(*node) bool) []string {
	open := []*node{me}
	closed := map[string]*node{me.name: me}
	prev := map[string]string{}

	for len(open) > 0 {
		elem := open[0]
		open = open[1:]

		if found(elem) {
			backtrack := []string{}
			if elem == me {
				return backtrack
			}
			to := elem.name
			for {
				from := prev[to]
				for k, v := range g.connections[g.getNode(from)] {
					if v.name == to {
						backtrack = append(backtrack, k)
					}
				}
				to = from
				if to == me.name {
					slices.Reverse(backtrack)
					return backtrack
				}
			}
		}

		for _, n := range g.connections[elem] {
			if _, exist := closed[n.name]; !exist {
				closed[n.name] = n
				prev[n.name] = elem.name
				open = append(open, n)
			}
		}
	}
	return []string{}
}

func consume(s string, g *grid, lastCommand string, prevNode *node) *node {
	name := ""
	for _, line := range strings.Split(s, "\n") {
		if strings.Contains(line, "==") {
			name = line
		}
	}
	println("Moved", lastCommand, "to", name)

	me := prevNode
	if len(name) > 0 {
		me = g.getNode(name)
		me.isCheckpoint = strings.Contains(s, "verify your identity")
		me.visited = true
		if prevNode != nil && prevNode != me {
			g.addConnection(prevNode, me, lastCommand)
		}

		if prevNode == me {
			return nil
		}
	}

	if strings.Contains(s, "== Pressure-Sensitive Floor ==") {
		if strings.Contains(s, "Alert! Droids on this ship are lighter than the detected value!") {
			//fmt.Println("Too light")
		} else if strings.Contains(s, "Droids on this ship are heavier than the detected value!") {
			//fmt.Println("Too fat")
		} else {
			fmt.Println(s)
		}
	}

	blocks := strings.Split(s, "\n\n")
	for _, b := range blocks {
		if strings.Contains(b, "Doors here lead:") {
			for _, d := range DIRECTIONS {
				if strings.Contains(b, "- "+d) {
					if _, exists := g.connections[me][d]; !exists {
						n := g.getNode("UNKNOWN")
						n.isPressure = me.isCheckpoint
						g.addConnection(me, n, d)
					}
				}
			}
		} else if strings.Contains(b, "Items here:") && len(me.items) == 0 {
			for _, i := range strings.Split(b, "\n")[1:] {
				iName := strings.ReplaceAll(i, "- ", "")
				me.items = append(me.items, iName)
				println("New item found: ", iName)
			}
		}
	}

	//g.print(me.p)

	return me
}

func action(g *grid, me *node, itemsToTake []string) string {
	for _, itemToTake := range itemsToTake {
		if slices.Contains(me.items, itemToTake) {
			me.items = slices.DeleteFunc(me.items, func(e string) bool {
				return e == itemToTake
			})

			return "take " + itemToTake
		}
	}

	if path := g.bfs(me, func(n *node) bool { return !n.visited && !n.isPressure }); len(path) > 0 {
		return path[0]
	}

	if path := g.bfs(me, func(n *node) bool { return n.isPressure }); len(path) > 0 {
		return path[0]
	}

	return ""
}

func part1(input string, itemsToTake []string) int {
	icc := intcodecomputer.New(utils.GetAllNumbers(input))

	instructions := []rune("")
	ptr := -1
	setInput := func(s string) {
		println("Command given: ", s)

		instructions = []rune(s + "\n")
		ptr = -1
	}

	grid := newGrid()
	var me *node
	var prevMe *node
	var lastCommand string

	buffer := ""
	icc.SetOutputter(func(a int) {
		buffer += string(rune(a))
		if strings.Contains(buffer, "Command?") {
			prevMe = me
			me = consume(buffer, &grid, lastCommand, prevMe)
			if me == nil {
				icc.Stop()
				return
			}
			lastCommand = action(&grid, me, itemsToTake)
			if lastCommand == "" {
				icc.Stop()
			} else {
				setInput(lastCommand)
				buffer = ""
			}
		}
	})

	icc.SetInputter(func() int {
		ptr++
		return int(instructions[ptr])
	})

	icc.ExecuteUntilHalt()
	return 0
}
