package receiver

type receiver struct {
	Something uint8
}

func New() *receiver {
	return new(receiver)
}
