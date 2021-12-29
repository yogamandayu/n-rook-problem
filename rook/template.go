package rook

var (
	template3x3 = [][]CellState{
		{CellStateEmpty, CellStateWall, CellStateEmpty},
		{CellStateWall, CellStateEmpty, CellStateEmpty},
		{CellStateEmpty, CellStateWall, CellStateEmpty},
	}

	template4x4 = [][]CellState{
		{CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty},
		{CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty},
		{CellStateWall, CellStateWall, CellStateEmpty, CellStateEmpty},
		{CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty},
	}

	template6x6 = [][]CellState{
		{CellStateWall, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty},
		{CellStateWall, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty},
		{CellStateEmpty, CellStateWall, CellStateWall, CellStateEmpty, CellStateEmpty, CellStateWall},
		{CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty},
		{CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty},
		{CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty},
	}

	template7x7 = [][]CellState{
		{CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty},
		{CellStateWall, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty},
		{CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateWall},
		{CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty},
		{CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateWall, CellStateEmpty, CellStateEmpty},
		{CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall},
		{CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty},
		{CellStateWall, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateWall},
	}

	template8x8 = [][]CellState{
		{CellStateEmpty, CellStateWall, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateWall},
		{CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty},
		{CellStateWall, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateWall},
		{CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty},
		{CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty},
		{CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateWall},
		{CellStateEmpty, CellStateWall, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty},
		{CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateEmpty, CellStateWall, CellStateEmpty, CellStateWall, CellStateEmpty},
	}
)

func Template3x3() [][]CellState {
	return template3x3
}

func Template4x4() [][]CellState {
	return template4x4
}

func Template6x6() [][]CellState {
	return template6x6
}

func Template7x7() [][]CellState {
	return template7x7
}

func Template8x8() [][]CellState {
	return template8x8
}
