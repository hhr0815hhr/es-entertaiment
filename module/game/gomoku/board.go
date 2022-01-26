package gomoku

type Location struct {
	X int
	Y int
}

type BoardStruct struct {
	Board [][]int
	Steps []Location
}

var BoardIns *BoardStruct

type IBoard interface {
	Put(who, x, y int) bool
}

func NewBoard() IBoard {
	b := make([][]int, 15)
	for i := range b {
		b[i] = make([]int, 15)
	}
	BoardIns = &BoardStruct{
		Board: b,
		Steps: make([]Location, 225),
	}
	return BoardIns
}

func (b *BoardStruct) Put(who, x, y int) bool {
	if x < 0 || x > 14 || y < 0 || y > 14 {
		return false
	}
	if b.Board[x][y] != 0 {
		return false
	}
	b.Board[x][y] = who
	return true
}
