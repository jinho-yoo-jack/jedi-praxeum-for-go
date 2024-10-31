package direction

const (
	None  Direction = iota
	North Direction = iota
	East  Direction = iota
	South Direction = iota
)

type Direction float64

func GetDirection(angle float64) Direction {
	switch {
	case angle >= 315, angle >= 0 && angle < 45:
		return North
	case angle >= 45 && angle < 135:
		return East
	case angle >= 135 && angle < 225:
		return South
	default:
		return None
	}
}
