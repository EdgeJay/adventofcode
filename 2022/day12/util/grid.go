package util

import (
	"fmt"
	"strings"
)

const runeOffset = 96

type Node struct {
	Name      string
	Row       int
	Col       int
	Elevation int
	Grid      *Grid
}

func (n *Node) Print() string {
	return fmt.Sprintf("node(%d, %d)(%d)", n.Col, n.Row, n.Elevation)
}

func NewNode(grid *Grid) *Node {
	return &Node{
		Elevation: 1,
		Grid:      grid,
	}
}

type Grid struct {
	Nodes         [][]*Node
	StartLocation *Node
	EndLocation   *Node
}

func (g *Grid) PrintNodes() string {
	output := "[\n"
	for _, ln := range g.Nodes {
		arr := make([]string, len(ln))
		for col, node := range ln {
			arr[col] = node.Name
		}
		output += "\t" + strings.Join(arr, ",\t") + "\n"
	}
	return output + "]"
}

func (g *Grid) GetAdjacentNodes(node *Node) []*Node {
	arr := make([]*Node, 0)

	adjNode := g.GetNodeAt(node.Col-1, node.Row)
	if adjNode != nil {
		arr = append(arr, adjNode)
	}

	adjNode = g.GetNodeAt(node.Col+1, node.Row)
	if adjNode != nil {
		arr = append(arr, adjNode)
	}

	adjNode = g.GetNodeAt(node.Col, node.Row-1)
	if adjNode != nil {
		arr = append(arr, adjNode)
	}

	adjNode = g.GetNodeAt(node.Col, node.Row+1)
	if adjNode != nil {
		arr = append(arr, adjNode)
	}

	return arr
}

func (g *Grid) GetNodeAt(col, row int) *Node {
	if row >= 0 && row < len(g.Nodes) {
		if col >= 0 && col < len(g.Nodes[row]) {
			return g.Nodes[row][col]
		}
	}
	return nil
}

func (g *Grid) FindShortestPath(start, end *Node) int {
	context := &SearchContext{
		RootNode:   start,
		TargetNode: end,
	}

	fmt.Printf("Finding shortest path from (%d, %d) to (%d, %d)\n", start.Col, start.Row, end.Col, end.Row)

	steps := BFS(context)

	return steps
}

func (g *Grid) GetShortestPathSteps() int {
	return g.FindShortestPath(g.StartLocation, g.EndLocation)
}

func MakeGrid(input []string) *Grid {
	grid := &Grid{}
	nodes := make([][]*Node, len(input))
	grid.Nodes = nodes

	for idx, ln := range input {
		runes := []rune(ln)
		nodes[idx] = make([]*Node, len(runes))
		for n, c := range runes {
			node := NewNode(grid)
			node.Row = idx
			node.Col = n

			if c == 'S' {
				node.Name = "node(start)"
				node.Elevation = 1
				grid.StartLocation = node
			} else if c == 'E' {
				node.Name = "node(end)"
				node.Elevation = 26
				grid.EndLocation = node
			} else {
				node.Elevation = int(c) - runeOffset
				node.Name = fmt.Sprintf("node(%d)", node.Elevation)
			}

			nodes[idx][n] = node
		}
	}

	return grid
}

/*
func printNodes(nodes []*Node) string {
	str := ""
	for _, node := range nodes {
		str += node.Print() + "\t"
	}
	return str
}
*/
