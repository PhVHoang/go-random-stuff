package paxos

import (
	"log"
	"time"
)

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

func (n *Network) SendTo(m message) {
	log.Print("Send msg from: ", m.from, " send to ", m.to, " val: ", m.val, " typ: ", m.typ)
	n.recvQueue[m.to] <- m
}

func (n *NetWork) RecvFrom(id *int) *message {
	select {
	case retMsg := <-n.recvQueue[id]:
		log.Println("Recev msg from: ", retMsg.from, " send to ", retMsg.to, " val: ", retMsg.val, " typ: ", retMsg.typ)
		return &retMsg
	case <-time.After(time.Second):
		return nil
	}
}

func (n *NodeNetwork) send(m message) {
	n.net.SendTo(m)
}

func (n *NodeNetwork) recev() *message {
	return n.net.RecvFrom(n.id)
}
