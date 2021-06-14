package gallons

import "errors"

type Jug struct {
	GallonsOfWater int
	MaxSize        int
}

func NewJug(max int) *Jug {
	j := Jug{}
	j.MaxSize = max
	j.GallonsOfWater = 0
	return &j
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
