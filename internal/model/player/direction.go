package player

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func (d *Direction) IsValid() bool {
	switch *d {
	case Left, Right, Up, Down:
		return true
	default:
		return false
	}
}
