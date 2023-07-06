package soduku

import (
	"math"

	"github.com/gofiber/fiber/v2"

	"soduku-madness-BE/internal/api"
	params "soduku-madness-BE/internal/param/soduku"
	"soduku-madness-BE/internal/soduku"
)

const (
	ErrUnsolvableBoard = "Unsolvable soduku puzzle."
	SuccesfulSolve     = "Soduku puzzle solved."
)

func HandleSolve(c *fiber.Ctx) error {
	var sodukuparams params.SodukuParams
	err := c.BodyParser(&sodukuparams)
	if err != nil {
		return err
	}

	err = sodukuparams.ValidateParams()
	if err != nil {
		return err
	}

	solver := soduku.SodokuSolver{
		Board:     sodukuparams.ToSodukuBoard(),
		Size:      len(sodukuparams.Board),
		ChunkSize: int(math.Sqrt(float64(len(sodukuparams.Board)))),
	}

	if !solver.Solve(0, 0) {
		return c.Status(fiber.StatusOK).JSON(api.APIResponse{
			Data:     solver.Board,
			Messages: []string{ErrUnsolvableBoard},
		})
	}

	return c.Status(fiber.StatusOK).JSON(api.APIResponse{
		Data:     solver.Board,
		Messages: []string{SuccesfulSolve},
	})
}
