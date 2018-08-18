package vareval

import "math"

type expr interface {
	Evaluate() (float64, error)
}

type number struct {
	value float64
}

func (n number) Evaluate() (float64, error) {
	return n.value, nil
}

type add struct {
	left, right expr
}

func (a add) Evaluate() (float64, error) {
	l, err := a.left.Evaluate()
	if err != nil {
		return 1, err
	}
	r, err := a.right.Evaluate()
	if err != nil {
		return 1, err
	}
	return l + r, nil
}

type subtract struct {
	left, right expr
}

func (s subtract) Evaluate() (float64, error) {
	l, err := s.left.Evaluate()
	if err != nil {
		return 1, err
	}
	r, err := s.right.Evaluate()
	if err != nil {
		return 1, err
	}
	return l - r, nil
}

type multiply struct {
	left, right expr
}

func (m multiply) Evaluate() (float64, error) {
	l, err := m.left.Evaluate()
	if err != nil {
		return 1, err
	}
	r, err := m.right.Evaluate()
	if err != nil {
		return 1, err
	}
	return l * r, nil
}

type divide struct {
	left, right expr
}

func (d divide) Evaluate() (float64, error) {
	l, err := d.left.Evaluate()
	if err != nil {
		return 1, err
	}
	r, err := d.right.Evaluate()
	if err != nil {
		return 1, err
	}
	if r == 0 {
		return 1, DivideByZeroError
	}
	return l / r, nil
}

type negate struct {
	operand expr
}

func (n negate) Evaluate() (float64, error) {
	o, err := n.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	return -o, nil
}

type power struct {
	base, exponent expr
}

func (p power) Evaluate() (float64, error) {
	l, err := p.base.Evaluate()
	if err != nil {
		return 1, err
	}
	r, err := p.exponent.Evaluate()
	if err != nil {
		return 1, err
	}
	result := math.Pow(l, r)
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type cos struct {
	operand expr
}

func (c cos) Evaluate() (float64, error) {
	o, err := c.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := math.Cos(toRadians(o))
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type sin struct {
	operand expr
}

func (s sin) Evaluate() (float64, error) {
	o, err := s.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := math.Sin(toRadians(o))
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type tan struct {
	operand expr
}

func (t tan) Evaluate() (float64, error) {
	o, err := t.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := math.Tan(toRadians(o))
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type acos struct {
	operand expr
}

func (a acos) Evaluate() (float64, error) {
	o, err := a.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := toDegrees(math.Acos(o))
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type asin struct {
	operand expr
}

func (a asin) Evaluate() (float64, error) {
	o, err := a.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := toDegrees(math.Asin(o))
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type atan struct {
	operand expr
}

func (a atan) Evaluate() (float64, error) {
	o, err := a.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := toDegrees(math.Atan(o))
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type round struct {
	operand expr
}

func (r round) Evaluate() (float64, error) {
	o, err := r.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := math.Round(o)
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type ceil struct {
	operand expr
}

func (c ceil) Evaluate() (float64, error) {
	o, err := c.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := math.Ceil(o)
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type floor struct {
	operand expr
}

func (f floor) Evaluate() (float64, error) {
	o, err := f.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := math.Floor(o)
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type sqrt struct {
	operand expr
}

func (s sqrt) Evaluate() (float64, error) {
	o, err := s.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := math.Sqrt(o)
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type log struct {
	operand expr
}

func (l log) Evaluate() (float64, error) {
	o, err := l.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := math.Log10(o)
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

type ln struct {
	operand expr
}

func (l ln) Evaluate() (float64, error) {
	o, err := l.operand.Evaluate()
	if err != nil {
		return 1, err
	}
	result := math.Log(o)
	if math.IsNaN(result) {
		return 1, DomainError
	}
	return result, nil
}

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func toDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
