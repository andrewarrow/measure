package gallons

import (
	"errors"
	"fmt"
)

type Jug struct {
	GallonsOfWater int
	MaxSize        int
	Name           string
}

func NewJug(name string, max int) *Jug {
	j := Jug{}
	j.MaxSize = max
	j.Name = name
	j.GallonsOfWater = 0
	return &j
}

func (j *Jug) String() string {
	return fmt.Sprintf("JUG %s MAX %d CURRENT %d", j.Name, j.MaxSize, j.GallonsOfWater)
}

func (j *Jug) Fill() {
	j.GallonsOfWater = j.MaxSize
}
func (j *Jug) Empty() {
	j.GallonsOfWater = 0
}
func (j *Jug) Transfer(to *Jug) error {
	if to.MaxSize < to.GallonsOfWater+j.GallonsOfWater {
		to.GallonsOfWater += j.GallonsOfWater
		j.GallonsOfWater = 0
		return nil
	}
	return errors.New("to jug isn't big enough")
}
