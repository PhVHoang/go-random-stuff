package paxos

import "log"

type learner struct {
	id           int
	acceptedMsgs map[int]message
	nt           NodeNetwork
}

func NewLearner(id int, nt NodeNetwork, acceptorIds ...int) *learner {
	newLearner := &learner{id: id, nt: nt}
	newLearner.acceptedMsgs = make(map[int]message)
	for _, acceptId := range acceptorIds {
		newLearner.acceptedMsgs[acceptId] = message{}
	}
	return newLearner
}

func (l *learner) run() string {
	for {
		m := l.nt.recev()
		if m == nil {
			continue
		}
		log.Println("Learner: recev msg: ", *m)

	}
}
