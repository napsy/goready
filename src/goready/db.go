package main

// Deployment stores all infomration of a remote deployment
type Deployment struct {
	Name           string
	Description    string
	TargetHost     string
	DeploymentKey  string
	Notes          string
	Project struct {
		Name     string
		Version  string
		VCSHost  string
		VCSName  string
	}
	Options struct {
		AutoUpdate bool
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
	return "", nil
}

func (d *Deployment) Restart(reason string ) error {
	return "", nil
}

