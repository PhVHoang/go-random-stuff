package paxos

type Network struct {
	recvQueue map[int]chan message
}

type NodeNetwork struct {
	id  int
	net *Network
}

func CreeateNetwork(nodes ...int) *Network {
	nt := network{recvQueue: make(map[int]chan message, 0)}
	for _, node := range nodes {
		nt.recvQueue[node] = make(chan message, 1024)
	}
	return &nt
}

func (n *Network) GetNodeNetwork(id int) NodeNetwork {
	return NodeNetwork{id: id, net: n}
}

func (n *Network) sendTo(m message) {

}
