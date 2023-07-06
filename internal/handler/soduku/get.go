package soduku

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"soduku-madness-BE/internal/api"
	"soduku-madness-BE/internal/soduku"
)

const (
	ErrInvalidSize    = "Invalid size"
	ErrNoSizeProvided = "Puzzle size is required"
	GenerationSuccess = "%s difficulty puzzle generated"
)

func HandleGeneratePuzzle(c *fiber.Ctx) error {
	size := c.QueryInt("size")
	if size == 0 {
		return fiber.NewError(fiber.StatusBadRequest, ErrInvalidSize)
	}

	difficulty := c.Query("difficulty")
	puzzle, err := soduku.GeneratePuzzle(size, difficulty)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(api.APIResponse{
		Data:     puzzle,
		Messages: []string{fmt.Sprintf(GenerationSuccess, difficulty)},
	})
}
