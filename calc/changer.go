package calc

type Changer struct {
	state int
}

type IChanger interface {
	Mul(val int)
	Add(val int)
	Divide(val int)
	Substract(val int)
	GetState() int
}

func NewChanger(state int) *Changer {
	return &Changer{state: state}
}

func (c *Changer) Mul(val int) {
	c.state *= val
}

func (c *Changer) Add(val int) {
	c.state += val
}

func (c *Changer) Divide(val int, r string) {
	c.state /= val
}

func (c *Changer) Substract(val int) {
	c.state -= val
}

func (c *Changer) GetState() int {
	return c.state
}
