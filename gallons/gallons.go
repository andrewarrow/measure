package gallons

import (
	"fmt"
	"strconv"
	"time"
)

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

	xJug := NewJug("x", xGallons)
	yJug := NewJug("y", yGallons)
	zJug := NewJug("z", zGallons)

	for {

		xJug.Fill()
		yJug.Fill()

		fmt.Println(xJug.String())
		fmt.Println(yJug.String())
		fmt.Println(zJug.String())

		xJug.Transfer(zJug)
		yJug.Transfer(zJug)

		fmt.Println(xJug.String())
		fmt.Println(yJug.String())
		fmt.Println(zJug.String())

		if zJug.GallonsOfWater == zJug.MaxSize {
			break
		}

		time.Sleep(time.Second * 1)
	}
}
