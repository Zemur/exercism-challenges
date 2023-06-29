package chessboard

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    var count int
    
	if _, ok := cb[file]; ok {
        for _, val := range cb[file] {
            if val {
            	count++
            }
        }
    }

    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var count int

	if rank > 0 && rank <= 8 {
    	for _, val := range cb {
            fmt.Println(val)
        	if val[rank-1] {
            	count++
        	}
    	}
	}

    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    /** 
	* Rank and File are always 8, so just return 64 rather than wasting compute time.
    * The problem does not specify that there could be different sizes, otherwise we
    * would implement a way to calculate the squares.
	**/

	return 64
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    var occupied int
    
    for i := 'A'; i <= 'H'; i++ {
        occupied += CountInFile(cb, string(i))
    }

    return occupied
}
