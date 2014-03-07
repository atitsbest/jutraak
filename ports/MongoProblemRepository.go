package ports

import (
  "dccs/jutraak/bugtracking/entities"

  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type MongoProblemRepository struct {
  connectionString string
}

/**
  CTR
 */
func NewMongoProblemRepository(connectionString string) *MongoProblemRepository {
  return &MongoProblemRepository{
    connectionString: connectionString,
  } 
}

/**
  Neues Problem einfügen.
 */
func (self *MongoProblemRepository) Insert(problem *entities.Problem) error {
  session, err := mgo.Dial(self.connectionString)
  if err != nil { return err }
  defer session.Close()

  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)

  c := session.DB("jutraak_test").C("problems")
  // Neue Id erstellen.
  problem.Id = string(bson.NewObjectId())
  err = c.Insert(problem)
  if err != nil { 
    // Bei einem Fehler müssen wir die Id wieder löschen.
    problem.Id = ""
    return err; 
  }

  return nil;
}
