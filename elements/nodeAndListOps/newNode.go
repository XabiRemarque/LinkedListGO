package nodeandlistops

import socketmaker "LinkedList/elements/socketMaker"

type Data struct {
	id      int
	company string
	socket  Socket
}

type Socket struct {
	ip   string
	port int
}

type Node struct {
	data Data
	next *Node
}

func newNode(company string, ID int) *Node {
	socket := newSocket(socketmaker.RandomSocket())

	node := &Node{
		data: Data{
			id:      ID,
			company: company,
			socket:  socket,
		},
		next: nil,
	}

	return node
}

func newSocket(ip string, port int) Socket {
	return Socket{
		ip:   ip,
		port: port,
	}
}
