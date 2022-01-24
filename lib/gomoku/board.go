package gomoku

var Board [][]int

func init() {
	//init board
	Board = make([][]int, 15)
	for i := range Board {
		Board[i] = make([]int, 15)
	}
}
