package wrr

import "fmt"

type Node struct {
	Name string
	Current int
	Weight int
}

// 算法执行2步，选择出1个当前节点。
//
// 每个节点，用它们的当前值加上它们自己的权重。
// 选择当前值最大的节点为选中节点，并把它的当前值减去所有节点的权重总和。
func Pick(nodes []*Node) (best *Node) {
	if len(nodes) == 0 {
		return
	}
	total := 0
	for _, node := range nodes {
		if node == nil {
			continue
		}
		total += node.Weight
		node.Current+=node.Weight
		if best == nil || node.Current > best.Current {
			best = node
		}
	}
	if best == nil {
		return
	}
	best.Current -= total
	return
}

func example() {
	nodes := []*Node{
		&Node{"a", 0, 5},
		&Node{"b", 0, 1},
		&Node{"c", 0, 1},
	}

	for i := 0; i < 7; i++ {
		best := Pick(nodes)
		if best != nil {
			fmt.Println(best.Name)
		}
	}
}