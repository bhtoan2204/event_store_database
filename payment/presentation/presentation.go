package presentation

import (
	"context"
)

type Presentation struct {
}

func NewPresentation() *Presentation {
	return &Presentation{}
}

func (p *Presentation) Handle(ctx context.Context) error {
	return nil
}
