package main

import "fmt"

type Human struct {
	Name     string `royal:"true"`
	DickSize int
	Gender   uint8
	BoobSize int
}

func (h Human) Bio() string {
	return fmt.Sprintf("%s %d %d %d\n", h.Name, h.Gender, h.DickSize, h.BoobSize)
}
