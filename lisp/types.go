package lisp

type Node struct {
	root   Token
	left   *Node
	right  *Node
	name   string
	parent *Node
}

type Scope struct {
	Enverment map[string]interface{}
	name      string
}
type Statement struct {
	node *Node
	name string
}
type Function struct {
	name  string
	scope Scope
	stmt  *Statement
}
