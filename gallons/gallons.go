package gallons

import (
	"fmt"
	"strconv"
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

	// start with larger of x and y
	larger := xJug
	smaller := yJug
	if yJug.MaxSize > xJug.MaxSize {
		larger = yJug
		smaller = xJug
	}

	fmt.Println(xJug.String())
	fmt.Println(yJug.String())
	fmt.Println(zJug.String())

	fmt.Println("filling using the larger bucket...")
	for {
		larger.Fill()
		fmt.Println(larger.String())
		e := larger.Transfer(zJug)
		if e != nil {
			larger.Empty()
			break
		}
		fmt.Println(larger.String())
		fmt.Println(zJug.String())
	}

	fmt.Println("filling using the smaller bucket...")
	for {
		smaller.Fill()
		fmt.Println(smaller.String())
		e := smaller.Transfer(zJug)
		if e != nil {
			smaller.Empty()
			break
		}
		fmt.Println(smaller.String())
		fmt.Println(zJug.String())
	}
	fmt.Println("")

	fmt.Println(xJug.String())
	fmt.Println(yJug.String())
	fmt.Println(zJug.String())

	if zJug.GallonsOfWater == zJug.MaxSize {
		return
	}

	fmt.Println("No Solution")
}
