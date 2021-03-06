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

func (j *Jug) Fill() error {
	if j.GallonsOfWater == j.MaxSize {
		return errors.New("jug is already full.")
	}
	j.GallonsOfWater = j.MaxSize
	return nil
}
func (j *Jug) Empty() {
	j.GallonsOfWater = 0
}
func (j *Jug) Transfer(to *Jug) int {
	for {
		if to.GallonsOfWater == to.MaxSize {
			break
		}
		if j.GallonsOfWater == 0 {
			break
		}
		to.GallonsOfWater++
		j.GallonsOfWater--
	}

	return j.GallonsOfWater
}
