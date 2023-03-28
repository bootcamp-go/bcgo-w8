package main

import "fmt"

type Position struct {
	X, Y int
}
func (p *Position) Move(x, y int) {
	p.X += x
	p.Y += y
}

type Ship struct {
	Model 		string
	Speed 		int		// units per second
	Position
}
func (s *Ship) Move(seconds int) {
	x := s.Speed * seconds
	y := s.Speed * seconds
	s.Position.Move(x, y)
}

func main() {
	s1 := Ship{
		Model: "Falcon",
		Speed: 2,
		Position: Position{
			X: 10,
			Y: 20,
		},
	}

	s1.Move(10)

	fmt.Println(s1.X, s1.Y)
}