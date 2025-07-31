package eventbus

var _ IEvent = (*event)(nil)

type IEvent interface {
	EventName() string
}

type event struct {
	name string
}

func NewEvent(name string) IEvent {
	return &event{
		name: name,
	}
}

func (c *event) EventName() string {
	return c.name
}
