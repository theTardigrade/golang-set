package set

type Config struct {
	EqualityTest      equalityTestFunc
	Filter            filterFunc
	MaximumValueCount *int
	MultiMode         bool
	Capacity          *int
}
