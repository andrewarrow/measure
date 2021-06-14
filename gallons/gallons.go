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

	// start with larger of x and y
	larger := xJug
	smaller := yJug
	if yJug.MaxSize > xJug.MaxSize {
		larger = yJug
		smaller = xJug
	}

	fmt.Println(xJug.String())
	fmt.Println(yJug.String())

	largeSum := 0
	largeN := 0
	for {
		if largeSum+larger.MaxSize > zGallons {
			break
		}
		largeSum += larger.MaxSize
		largeN++
	}

	if largeSum == zGallons {
		fmt.Println("simple solution, use large bucket", largeN, "times.")
		return
	}

	smallSum := largeSum
	smallN := 0
	for {
		if smallSum+smaller.MaxSize > zGallons {
			break
		}
		smallSum += smaller.MaxSize
		smallN++
	}

	if smallSum == zGallons {
		fmt.Println("simple solution, use large bucket", largeN, "times and smaller bucket", smallN, "times.")
		return
	}

	larger.Fill()
	left := larger.Transfer(smaller)
	fmt.Println("left over", left)
	smaller.Empty()
	left = larger.Transfer(smaller)
	fmt.Println("left over", left)
	larger.Fill()
	left = larger.Transfer(smaller)
	fmt.Println("left over", left)
	fmt.Println(xJug.String())
	fmt.Println(yJug.String())

	if smaller.GallonsOfWater == zGallons {
		fmt.Println("smaller bucket now has", zGallons)
		return
	}
	if larger.GallonsOfWater == zGallons {
		fmt.Println("larger bucket now has", zGallons)
		return
	}
	fmt.Println("No Solution")
}
