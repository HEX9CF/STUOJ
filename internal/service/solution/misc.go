package solution

import (
	"STUOJ/external/neko"
	"STUOJ/internal/model"
)

func Generate(si model.NekoSolutionInstruction) (model.NekoSolution, error) {
	return neko.GenerateSolution(si)
}
