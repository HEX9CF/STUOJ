package testcase

import (
	"STUOJ/external/neko"
	"STUOJ/internal/model"
)

func Generate(ti model.NekoTestcaseInstruction) (model.NekoTestcase, error) {
	return neko.GenerateTestcase(ti)
}
