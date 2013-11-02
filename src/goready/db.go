package main

import (
	"labix.org/v2/mgo"
	_ "labix.org/v2/mgo/bson"
)

// Deployment stores all infomration of a remote deployment
type Deployment struct {
	Name           string
	Description    string
	TargetHost     string
	DeploymentKey  string
	Notes          string
	Options struct {
		AutoUpdate  bool
		ValidBranch string
	}
}

func (d *Deployment) Update() error {
	return nil
}

func (d *Deployment) QueryVersion() (string, error) {
	return "", nil
}

func (d *Deployment) Remove() (string, error) {
	return "", nil
}

func (d *Deployment) Start() (string, error) {
	return "", nil
}
func (d *Deployment) Stop(reason string) error {
	return nil
}

func (d *Deployment) Restart(reason string ) error {
	return nil
}

func LoadDeployments() ([]Deployment, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("goready").C("deployments")
	result := []Deployment{}
	err = c.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func InsertDeployments() error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("goready").C("deployments")
	err = c.Insert(
		&Deployment{"Testna", "Testna postavitev", "localhost", "127.0.0.1", "AAABBB", struct{AutoUpdate bool; ValidBranch string}{false, "master"}},
		&Deployment{"Druga", "Druga testna postavitev", "localhost", "127.0.0.1", "CCCDDD", struct{AutoUpdate bool; ValidBranch string}{true, "unstable"}},
		&Deployment{"Beta staging", "Beta staging server", "localhost", "127.0.0.1", "CCCDDD",struct{AutoUpdate bool; ValidBranch string}{false, "master"}})
	if err != nil {
		return err
	}
	return nil
}
