package pointer_ex

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	return &Actor{name, hp, speed}
}

func newActor(name string) *Actor {
	actor := new(Actor)
	actor.Name = name
	actor.HP = 100
	actor.Speed = 100
	return actor
}
