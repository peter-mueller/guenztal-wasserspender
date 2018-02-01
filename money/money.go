package money

type Money struct {
	Cents uint
}

const (
	Cent = 1
	Euro = Cent * 100
)
