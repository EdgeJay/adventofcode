package util

import (
	"github.com/EdgeJay/adventofcode/common/utils/data"
)

type NodeQueueItem struct {
	Node  *Node
	Steps int
}

type SearchContext struct {
	RootNode   *Node
	TargetNode *Node
	Explored   *map[[2]int]bool
	Queue      *data.Queue[NodeQueueItem]
}

func setNodeAsExplored(node *Node, exploredMap *map[[2]int]bool) {
	(*exploredMap)[[2]int{node.Col, node.Row}] = true
}

func isNodeExplored(node *Node, exploredMap *map[[2]int]bool) bool {
	return (*exploredMap)[[2]int{node.Col, node.Row}]
}

func BFS(context *SearchContext) int {
	context.Explored = &map[[2]int]bool{}
	context.Queue = &data.Queue[NodeQueueItem]{}

	context.Queue.In(&NodeQueueItem{
		Node:  context.RootNode,
		Steps: 0,
	})

	for !context.Queue.IsEmpty() {
		queueItem := context.Queue.Out()
		node := queueItem.Node

		if isNodeExplored(node, context.Explored) {
			continue
		}
		setNodeAsExplored(node, context.Explored)

		if node == context.TargetNode {
			return queueItem.Steps
		}

		adjNodes := node.Grid.GetAdjacentNodes(node)

		if len(adjNodes) == 0 {
			break
		}

		for _, n := range adjNodes {
			if node.Elevation == n.Elevation || node.Elevation+1 == n.Elevation || node.Elevation > n.Elevation {
				context.Queue.In(&NodeQueueItem{
					Node:  n,
					Steps: queueItem.Steps + 1,
				})
			}
		}
	}

	return 0
}
