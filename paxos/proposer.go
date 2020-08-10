package paxos

import "log"

type proposer struct {
	id         int
	seq        int
	proposeNum int
	proposeVal string
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

func (p *proposer) prepare() []message {
	p.seq++
	sendMsgCount := 0
	var msgList []message
	log.Println("proposer: propose msg: ", len(p.acceptors))
	for acceptedID, _ := range p.acceptors {
		msg := message{from: p.id, to: acceptedID, typ: Prepare, seq: p.getProposeNum(), val: p.proposeVal}
		msgList = append(msgList, msg)
		if sendMsgCount > p.majority() {
			break
		}
	}
	return msgList
}

func (p *proposer) checkRecvPromise(promise message) {
	previousMessage := p.acceptors[promise.from]
	log.Println(" prevMsg: ", previousMessage, " promiseMsg: ", promise)
	log.Println("checking on: ", previousMessage.getProposeSeq(), " and ", promise.getProposeSeq())
	if previousMessage.getProposeSeq() < promise.getProposeSeq() {
		log.Println("Proposer: ", p.id, " got a new promise : ", promise)
		p.acceptors[promise.from] = promise
		if promise.getProposeSeq() > p.getProposeNum() {
			p.proposeNum = promise.getProposeSeq()
			p.proposeVal = promise.getProposeVal()
		}
	}
}

// After receiving the promise from acceptor and reach majority
// proposer will propose values to those acceptors and let them know the consensus is already ready
func (p *proposer) propose() []message {
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

func (p *proposer) majority() int {
	return len(p.acceptors)/2 + 1
}

func (p *proposer) getReceivedPromiseCount() int {
	recvCount := 0
	for _, acceptedMsg := range p.acceptors {
		log.Println("Proposer has total: ", len(p.acceptors), " acceptors ",
			acceptedMsg, " current Num: ", p.getProposeNum(), " msgNum ", acceptedMsg.getProposeSeq())
		if acceptedMsg.getProposeSeq() == p.getProposeNum() {
			log.Println("Recv ++", recvCount)
		}
	}
	log.Println("Current proposer received ", recvCount, " promise count")
	return recvCount
}

func (p *proposer) majorityReached() bool {
	return p.getReceivedPromiseCount() > p.majority()
}

func (p *proposer) getProposeNum() int {
	p.proposeNum = p.seq<<4 | p.id
	return p.proposeNum
}
