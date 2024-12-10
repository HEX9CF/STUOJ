package problem

import (
	"STUOJ/external/neko"
	"STUOJ/internal/model"
)

func Generate(pi model.NekoProblemInstruction) (model.NekoProblem, error) {
	return neko.GenerateProblem(pi)
}
