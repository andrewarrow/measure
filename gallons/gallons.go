package gallons

import (
	"errors"
	"fmt"
	"strconv"
)

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

func HandleInput(argMap map[string]string) {

	x := argMap["x"]
	y := argMap["y"]
	z := argMap["z"]

	if x == "" || y == "" || z == "" {
		fmt.Println("pass in values for --x and --y and --z")
		return
	}

	xGallons, _ := strconv.Atoi(x)
	yGallons, _ := strconv.Atoi(y)
	zGallons, _ := strconv.Atoi(z)
	if xGallons == 0 || yGallons == 0 || zGallons == 0 {
		fmt.Println("pass in int values for --x and --y and --z")
		return
	}

	xJug := NewJug(xGallons)
	yJug := NewJug(yGallons)
	zJug := NewJug(zGallons)

	fmt.Println("Input", xJug, yJug, zJug)
}
