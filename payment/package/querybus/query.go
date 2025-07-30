package querybus

var _ IQuery = (*query)(nil)

type IQuery interface {
	QueryName() string
}

type query struct {
	name string
}

func NewQuery(name string) IQuery {
	return &query{
		name: name,
	}
}

func (c *query) QueryName() string {
	return c.name
}
