package lisp

type Scope struct {
	Enverment map[string]interface{}
	name      string
}
type Statement struct {
	node *Node
	name string
}
