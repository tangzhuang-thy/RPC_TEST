package shared

type Args struct {
	A, B int
}

type Calculator interface {
	Add(args Args, result *int) error
	Multiply(args Args, result *int) error
}
