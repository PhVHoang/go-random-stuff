package paxos

import "log"

/*
  NewAcceptor creates a new acceptor and also assigns learning IDs into this
  new acceptor
*/
func NewAcceptor(id int, nt NodeNetwork, learners ...int) acceptor {
    newAcceptor := acceptor{id: id, nt: nt}
    newAcceptor.learners = learners
    return newAcceptor
}

type acceptor struct {
    id int
    learners []int
    acceptMsg message
    promiseMsg message
    nt NodeNetwork
}

// Process logics for acceptor
func (a *acceptor) run() {
    for {
      m := a.nt.recev()
      if m == nil {
          continue
      }

      switch m.typ {
        case Prepare:
          promiseMsg := a.recevPrepare(*m)
          a.nt.send(*promiseMsg)
          continue
        case Propose:
          accepted := a.recevPropose(*m)
          if accepted {
            for _, lId := range a.learners {
                m.from = a.id
                m.to = lId
                m.typ = Accept
                a.nt.send(*m)
            }
          }
        default:
          log.Fatalln("Unsupported message type in acceptor Id: ". a.id)
      }
    }
    log.Println("acceptor: ", a.id, " leave.")
}

/*
  After receiving the prepare message, acceptors will check the prepare number and
  return this message if it is the bigest one
*/
func (a * acceptor) recevPrepare(prepareMsg message) *message {
    if a.promiseMsg.getProposeSeq() >= prepareMsg.getProposeSeq() {
        log.Println("ID: ", a.id, "has already accepted the biggest message")
        return nil
    }
    log.Println("ID: ", a.id, " Promise")
    prepareMsg.to = prepareMsg.from
    prepareMsg.typ = Promise
    a.acceptMsg = prepareMsg
    return &prepareMsg
}

func (a *acceptor) recevPropose(proposeMsg message) bool {
    log.Println("accept:check propose. ", a.acceptMsg.getProposeSeq(), proposeMsg.getProposeSeq())
    if a.acceptMsg.getProposeSeq() != proposeMsg.getProposeSeq() {
      log.Println("ID: ", a.id, " acceptor doesn't take this propose message: ", proposeMsg.val)
      return false
    }
    log.Println("ID: ", a.id, " accepted")
    return true
}

