package chessboard

//File ...
// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

//Chessboard ...
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	value, _ := cb[file]
	var count int

	for _, v := range value {
		if v {
			count++
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	var count int
	for _, place := range cb {
		if place[rank-1] {
			count++
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var squares int

	for _, v := range cb {
		squares += len(v)
	}

	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var occupied int

	for _, row := range cb {
		for _, v := range row {
			if v {
				occupied++
			}
		}
	}

	return occupied
}
