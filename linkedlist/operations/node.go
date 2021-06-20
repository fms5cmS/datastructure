package operations

import (
	"fmt"
	"strings"
)

type Node struct {
	Val  interface{}
	next *Node
}

func (node Node) String() string {
	ret := strings.Builder{}
	cur := &node
	for cur != nil {
		ret.WriteString(fmt.Sprintf("%v", cur.Val))
		if cur.next != nil {
			ret.WriteString(" -> ")
		}
		cur = cur.next
	}
	return ret.String()
}
