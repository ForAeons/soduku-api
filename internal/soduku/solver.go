package soduku

type SodokuSolver struct {
	Board     [][]int
	Size      int
	ChunkSize int
}

// returns the valid numbers that can be placed at the given row and col
func (s *SodokuSolver) getValidNums(row, col int) []int {
	validNums := make([]int, s.Size+1)

	// 1-s.Size are valid
	for i := 1; i < s.Size+1; i++ {
		validNums[i] = i
	}

	for i := 0; i < s.Size; i++ {
		if s.Board[row][i] != 0 {
			validNums[s.Board[row][i]] = 0
		}
		if s.Board[i][col] != 0 {
			validNums[s.Board[i][col]] = 0
		}
	}

	// Check chunk
	chunkRow := row / s.ChunkSize
	chunkCol := col / s.ChunkSize
	for i := 0; i < s.ChunkSize; i++ {
		for j := 0; j < s.ChunkSize; j++ {
			rowIndex := chunkRow*s.ChunkSize + i
			colIndex := chunkCol*s.ChunkSize + j
			if s.Board[rowIndex][colIndex] != 0 {
				validNums[s.Board[rowIndex][colIndex]] = 0
			}
		}
	}

	return validNums[1:]
}

func (s *SodokuSolver) isValidRow(row int) bool {
	seen := make([]bool, s.Size+1)
	for _, i := range s.Board[row] {
		if i == 0 {
			continue
		}
		if seen[i] {
			return false
		}
		seen[i] = true
	}
	return true
}

func (s *SodokuSolver) isValidCol(col int) bool {
	seen := make([]bool, s.Size+1)
	for _, row := range s.Board {
		if row[col] == 0 {
			continue
		}
		if seen[row[col]] {
			return false
		}
		seen[row[col]] = true
	}
	return true
}

func (s *SodokuSolver) isValidBox(box int) bool {
	seen := make([]bool, s.Size+1)
	row := box / s.ChunkSize
	col := box % s.ChunkSize
	for i := 0; i < s.ChunkSize; i++ {
		for j := 0; j < s.ChunkSize; j++ {
			if s.Board[i+row*s.ChunkSize][j+col*s.ChunkSize] == 0 {
				continue
			}
			if seen[s.Board[i+row*s.ChunkSize][j+col*s.ChunkSize]] {
				return false
			}
			seen[s.Board[i+row*s.ChunkSize][j+col*s.ChunkSize]] = true
		}
	}
	return true
}

func (s *SodokuSolver) isValidMove(row, col int, entry int) bool {
	s.Board[row][col] = entry
	valid := s.isValidRow(row) && s.isValidCol(col) && s.isValidBox(row/s.ChunkSize*s.ChunkSize+col/s.ChunkSize)
	s.Board[row][col] = 0
	return valid
}

func (s *SodokuSolver) Solve(row, col int) bool {
	// Gradually fill in the board from left to right, top to bottom
	if col == s.Size {
		col = 0
		row++
	}

	// If we've reached the end of the board, we're done
	if row == s.Size {
		return true
	}

	// If the current cell is already filled, move on to the next one
	if s.Board[row][col] != 0 {
		return s.Solve(row, col+1)
	}

	// Try all valid numbers for the current cell
	for _, num := range s.getValidNums(row, col) {
		if num == 0 {
			continue
		}
		s.Board[row][col] = num
		// If we've solved the s.Board, we're done
		if s.Solve(row, col+1) {
			return true
		}
	}
	// If we've tried all valid numbers and none of them worked, backtrack
	s.Board[row][col] = 0
	return false
}
