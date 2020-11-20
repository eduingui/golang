package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Square struct {
	Center Point
	Length int
}

// Moves
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

//Moves
func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

// Calculates area
func (s *Square) Area() int {
	return s.Length * s.Length
}

// New square
func NewSquare(x int, y int, length int) (*Square, error) {
	if x <= 0 {
		return nil, fmt.Errorf("X can not be lower than 0")
	}

	if y <= 0 {
		return nil, fmt.Errorf("Y can not be lower than 0")
	}

	if length <= 0 {
		return nil, fmt.Errorf("LENGTH can not be lower than 0")
	}

	point := Point{
		X: x,
		Y: y,
	}

	square := &Square{
		Center: point,
		Length: length,
	}

	return square, nil
}

func main() {
	square, err := NewSquare(10, 10, 20)

	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	square.Move(2, 3)
	fmt.Printf("%+v\n", square)
	fmt.Println(square.Area())

}
