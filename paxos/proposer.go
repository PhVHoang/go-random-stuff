package paxos

import "log"

type proposer struct {
	id         int
	seq        int
	proposeNum int
	proposeVal int
	acceptors  map[int]message
	nt         NodeNetwork
}

func NewProposer(id int, val string, nt NodeNetwork, acceptors ...int) {
	pro := proposer{id: id, proposeVal: val, seq: 0, nt: nt}
	pro.acceptors = make(map[int]message, len(acceptors))
	log.Println("proposer has ", len(acceptors), " acceptors, val:", pro.proposeVal)
	for _, acceptor := range acceptors {
		pro.acceptors[acceptor] = message{}
	}
	return &pro
}

// Detail process for Proposer
func (p *proposer) run() {
	log.Println("Proposer start running... val:", p.proposeVal)
	// Stage 1: Proposer sends preapare message to acceptor to reach accept from majority
}
