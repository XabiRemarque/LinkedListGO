package nodeandlistops

import "fmt"

func (list *LinkedList) NodeOutput() {
	index := 1
	if list.head == nil {
		return
	}
	current := list.head
	for current != nil {
		current.String(index)
		current = current.next
		index++
	}
}

func (n *Node) String(index int) {
	fmt.Printf("Node %d:\n  ID: %d\n  Company: %s\n  IP: %s\n  Port: %d\n\n", index, n.data.id, n.data.company, n.data.socket.ip, n.data.socket.port)
}
