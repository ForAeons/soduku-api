package sodukuparams

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	ErrInvalidBoardSize      = "Invalid board size. Must be a square number."
	ErrInconsistentBoardSize = "Inconsistent board size. Board is not a square."
	ErrInvalidBoardEntries   = "Invalid board entries. Must be a number between 0 and the max board length."
	ErrPuzzleTooComplex      = "Puzzle too complex."
)

type SodukuParams struct {
	Board [][]string `json:"puzzle"`
}

func (s *SodukuParams) ValidateParams() error {
	length := len(s.Board)

	if length > 9 {
		return fiber.NewError(fiber.ErrBadRequest.Code, ErrPuzzleTooComplex)
	}

	if length != int(math.Pow(math.Sqrt(float64(length)), 2.0)) {
		return fiber.NewError(fiber.ErrBadRequest.Code, ErrInvalidBoardSize)
	}

	for _, row := range s.Board {
		if len(row) != length {
			return fiber.NewError(fiber.ErrBadRequest.Code, ErrInconsistentBoardSize)
		}
		for _, str := range row {
			if str == "." {
				continue
			}
			num, err := strconv.Atoi(str)
			if err != nil || num <= 0 || num > length {
				return fiber.NewError(fiber.ErrBadRequest.Code, ErrInvalidBoardEntries)
			}
		}
	}

	return nil
}

func (s *SodukuParams) ToSodukuBoard() [][]int {
	board := [][]int{}
	for _, row := range s.Board {
		var intRow []int
		for _, str := range row {
			if str == "." {
				intRow = append(intRow, 0)
				continue
			}
			num, _ := strconv.Atoi(str)
			intRow = append(intRow, num)
		}
		board = append(board, intRow)
	}
	return board
}
