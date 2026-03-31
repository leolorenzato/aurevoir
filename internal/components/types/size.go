package types

type Size struct {
	Width  int
	Height int
}

func SubtractSize(s0, s1 Size) Size {
	return Size{
		Width:  s0.Width - s1.Width,
		Height: s0.Height - s1.Height,
	}
}
