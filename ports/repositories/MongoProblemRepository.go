package repositories

import (
    . "github.com/atitsbest/jutraak/bugtracking/domain/entities"
    "log"
    "os"

    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

type MongoProblemRepository struct {
    connectionString string
}

// CTR
func NewMongoProblemRepository(connectionString string) *MongoProblemRepository {
    return &MongoProblemRepository{
        connectionString: connectionString,
    }
}

// Neues Problem einfügen.
func (self *MongoProblemRepository) Insert(problem *Problem) error {
    session, err := mgo.Dial(self.connectionString)
    if err != nil {
        return err
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("jutraak_test").C("problems")
    // Neue Id erstellen.
    problem.Id = NewProblemId()
    err = c.Insert(problem)
    if err != nil {
        // Bei einem Fehler müssen wir die Id wieder löschen.
        problem.Id = ""
        return err
    }

    return nil
}

func (self *MongoProblemRepository) Update(problem *Problem) error {
    session, err := mgo.Dial(self.connectionString)
    if err != nil {
        return err
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("jutraak_test").C("problems")
    err = c.Update(bson.M{"id": problem.Id}, problem)
    if err != nil {
        return err
    }

    return nil
}

// Liefert alle Tags die in allen Problemen vorkommen.
func (self *MongoProblemRepository) AllTags() ([]string, error) {
    session, err := mgo.Dial(self.connectionString)
    if err != nil {
        return nil, err
    }
    defer session.Close()

    c := session.DB("jutraak_test").C("$cmd")
    result := &QueryValues{}
    err = c.Find(bson.M{"distinct": "problems", "key": "tags"}).One(result)
    if err != nil {
        return nil, err
    }

    return result.Values, nil
}

// Liefert ein Array mit allen Problemen.
func (self *MongoProblemRepository) All() ([]*Problem, error) {
    session, err := mgo.Dial(self.connectionString)
    if err != nil {
        return nil, err
    }
    defer session.Close()

    c := session.DB("jutraak_test").C("problems")
    var result []*Problem
    err = c.Find(nil).All(&result)

    return result, err
}

// Listert ein Array mit allen Problemen denen diese Tags zugeordnet sind.
func (self *MongoProblemRepository) GetProblemsByTag(tags []string) ([]*Problem, error) {
    session, err := mgo.Dial(self.connectionString)
    if err != nil {
        return nil, err
    }
    defer session.Close()

    c := session.DB("jutraak_test").C("problems")
    var result []*Problem
    err = c.Find(bson.M{"tags": bson.M{"$all": tags}}).All(&result)

    return result, err
}

// Problem gefiltert nach Suchstring und Tags
func (self *MongoProblemRepository) Filtered(tags []string, q string) ([]*Problem, error) {
    session, err := mgo.Dial(self.connectionString)
    if err != nil {
        return nil, err
    }
    defer session.Close()
    mgo.SetDebug(true)
    mgo.SetLogger(log.New(os.Stdout, "", 1))

    c := session.DB("jutraak_test").C("problems")
    var result []*Problem
    // err = c.Find(bson.M{"tags": bson.M{"$all": tags}}).All(&result)
    err = c.Find(bson.M{"$or": []bson.M{
        // bson.M{"tags": bson.M{"$all": tags}},
        // bson.M{"$or": []bson.M{
        bson.M{"summary": bson.RegEx{q, "i"}},
        bson.M{"description": bson.RegEx{q, "i"}},
        // }},
    }},
    ).All(&result)

    return result, err
}

func (self *MongoProblemRepository) GetById(id ProblemId) (*Problem, error) {
    session, err := mgo.Dial(self.connectionString)
    if err != nil {
        return nil, err
    }
    defer session.Close()

    c := session.DB("jutraak_test").C("problems")
    result := &Problem{}
    err = c.Find(bson.M{"id": id}).One(result)
    if err != nil {
        return nil, err
    }

    return result, nil
}

type QueryValues struct {
    Values []string
}
