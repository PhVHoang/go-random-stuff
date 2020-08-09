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

func NewProposer(id int, val string, nt NodeNetwork, acceptors ...int) *proposer {
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
	// Stage 1: Proposer sends prepare message to acceptor to reach accept from majority
	for !p.majorityReached() {
		log.Println("[Proposer:Prepare]")
		outMsgs := p.prepare()

	}
}

// After receiving the promise from acceptor and reach majority
// proposer will propose values to those acceptors and let them know the consensus is already ready
func (p *proposer) Propose() []message {
	sendMsgCount := 0
	var msgList []message
	log.Println("proposer: propose msg: ", len(p.acceptors))
	for acceptedId, acceptedMsg := range p.acceptors {
		log.Println("Checking promise id: ", acceptedMsg.getProposeSeq(), p.getProposeNum())
		if acceptedMsg.getProposeSeq() == p.getProposeNum() {
			msg := message{from: p.id, to: acceptedId, typ: Propose, seq: p.getProposeNum()}
			msg.val = p.proposeVal
			log.Println("ProposedVal: ", msg.val)
			msgList = append(msgList, msg)
		}
		sendMsgCount++
		if sendMsgCount > p.majority() {
			break
		}
	}
	log.Println("Proposer proposed this mesage list: ", msgList)
	return msgList
}
