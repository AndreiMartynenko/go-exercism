package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	if f, ok := cb[file]; ok {
		count := 0
		for _, occupied := range f {
			if occupied {
				count++
			}
		}
		return count
	}
	return 0
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	// Chess ranks are numbered 1-8, but array indices are 0-7
	// Convert to 0-based index
	index := rank - 1

	// Check if rank is within valid range (1-8, which becomes 0-7 after conversion)
	if index < 0 || index >= 8 {
		return 0
	}

	count := 0
	for _, file := range cb {
		if index < len(file) && file[index] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		count += len(file)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for _, occupied := range file {
			if occupied {
				count++
			}
		}
	}
	return count
}
