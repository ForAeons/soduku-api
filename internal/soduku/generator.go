package soduku

import (
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	ErrInvalidDifficulty = "Invalid difficulty"
)

func GeneratePuzzle(size int, difficulty string) ([][]string, error) {
	if !isValidDifficulty(difficulty) {
		return nil, fiber.NewError(fiber.StatusBadRequest, ErrInvalidDifficulty)
	}

	// Generate a pre-filled board
	puzzle := populateBoard(size)
	var threshold float64
	switch difficulty {
	case "easy":
		threshold = 0.25
	case "medium":
		threshold = 0.4
	case "hard":
		threshold = 0.55
	case "expert":
		threshold = 0.7
	default:
		threshold = 0.5
	}
	for r, row := range puzzle {
		for c := range row {
			if rand.Float64() < threshold {
				puzzle[r][c] = ""
			}
		}
	}

	return puzzle, nil
}

func isValidDifficulty(difficulty string) bool {
	switch difficulty {
	case "easy", "medium", "hard", "expert":
		return true
	default:
		return false
	}
}

func populateBoard(size int) [][]string {
	puzzle := make([][]string, size)
	for i := range puzzle {
		puzzle[i] = make([]string, size)
	}

	// fills the first row with random numbers
	// then solves the board
	// then removes numbers from the board
	nums := make([]string, size)
	for i := range nums {
		nums[i] = strconv.Itoa(i + 1)

	}
	// shuffle nums
	for i := range nums {
		if i == 0 {
			continue
		}
		j := rand.Intn(i)
		nums[i], nums[j] = nums[j], nums[i]
	}

	for i := range puzzle[0] {
		puzzle[0][i] = nums[i]
	}

	for r, row := range puzzle {
		if r == 0 {
			continue
		}

		var step int
		if r%3 == 0 {
			step = 1
		} else {
			step = 3
		}

		for i := 0; i < size; i++ {
			row[i] = puzzle[r-1][(i+step)%size]
		}
	}

	return puzzle
}
