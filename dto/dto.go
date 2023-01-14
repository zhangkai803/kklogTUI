package dto

type State int
type Choices []string
type LogFinishMsg struct {Err error}
type PodType string

type Env struct {
	Name string
	Alias string
}

func (e Env) String() string {
	return e.Alias + " - " + e.Name
}

func (e Env) IsProd() bool {
	return e.Name == "prod"
}

type Namespace struct {
	Env *Env
	Name string
	Alias string
}

func (ns Namespace) String() string {
	return ns.Alias + " - " + ns.Name
}

type Deployment struct {
	ProdNamespace *Namespace
	Name string
	Alias string
}

func (d Deployment) String() string {
	return d.Alias + " - " + d.Name
}

type Pod struct {
	Type PodType
	Deployment *Deployment
	Name string
	Alias string
}

func (p Pod) String() string {
	return p.Alias + " - " + p.Name
}
