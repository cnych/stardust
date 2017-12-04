package mysqlx

type Query struct {
	Spec        string
	Sort        []string
	Select      []string
	Skip, Limit int
}

func (q *Query) One() {

}
