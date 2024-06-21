package entity

import (
	"errors"
)

type Group struct {
	Name string `json:"name"`
	ID   uint
}

func NewGroup(name string) *Group {
	return &Group{
		Name: name,
	}
}

func (g *Group) Validate() error {
	if g.Name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}
