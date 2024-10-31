package house

type Home struct {
	Address string
	Size    int
	Price   float64
	Type    Type
}

const (
	APT = iota
)

type Type struct {
	Name string
}
