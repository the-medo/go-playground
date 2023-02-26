//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Coordinate struct {
	x int
	y int
}

type Rectangle struct {
	point1 Coordinate
	point2 Coordinate
}

func getAreaOfRectangle(rect Rectangle) int {
	length := math.Abs(float64(rect.point1.y - rect.point2.y))
	width := math.Abs(float64(rect.point1.x - rect.point2.x))

	return int(length * width)
}

func generateRandomCoordinate() Coordinate {
	return Coordinate{
		x: rand.Intn(10) + 1,
		y: rand.Intn(10) + 1,
	}
}

func createRandomRectangle() Rectangle {
	rec := Rectangle{
		generateRandomCoordinate(),
		generateRandomCoordinate(),
	}
	fmt.Println("Generated rectangle: ", rec)
	return rec
}

func main() {

	var (
		rec1 = createRandomRectangle()
		rec2 = createRandomRectangle()
		rec3 = createRandomRectangle()
	)

	fmt.Println("Rec1 area:", getAreaOfRectangle(rec1))
	fmt.Println("Rec2 area:", getAreaOfRectangle(rec2))
	fmt.Println("Rec3 area:", getAreaOfRectangle(rec3))

}
