// func (rt RecType) ProcMethod(p Tplayload, r *Treply) error

package typeSharing

import (
	"fmt"
)

type Student struct {
  Id int
  FirstName, LastName string
}

func (s Student) FullName() string {
  return s.FirstName + " " + s.LastName
}

type College struct {
  database map[int]Student
}

// The second argument of this method should be a pointer. This argument will be used
// to return a result by overriding its value
func (c *College) Add(payload Student, reply *Student) error {
  if _, ok := c.database[payload.Id]; ok{
    return fmt.Errorf("student with id '%d' already exists", payload.Id)
  }

  c.database[payload.Id] = payload
  *reply = payload
  return nil
}

func NewCollege() *College {
  return &College{
    database: make(map[int]Student),
  }
}


