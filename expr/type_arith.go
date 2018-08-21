package expr

// arithmeticValue holds an integral or real value
type arithmeticValue interface {
	value
	floatValue() float64
	almostEquals(other value, epsilon float64) bool
	add(other arithmeticValue) arithmeticValue
	sub(other arithmeticValue) arithmeticValue
	mul(other arithmeticValue) arithmeticValue
	div(other arithmeticValue) (arithmeticValue, error)
	pow(other arithmeticValue) (arithmeticValue, error)
	negate() arithmeticValue
	equality(other arithmeticValue) bool
	lessThan(other arithmeticValue) bool
}
