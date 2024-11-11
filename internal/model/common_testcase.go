package model

import (
	"STUOJ/utils"
	"golang.org/x/exp/rand"
	"time"
)

type CommonTestcaseInput struct {
	Rows []CommonTestcaseRow `json:"rows,omitempty"`
}

type CommonTestcaseRow struct {
	RowSizeId uint64                `json:"row_size_id,omitempty"`
	Values    []CommonTestcaseValue `json:"values,omitempty"`
}

type CommonTestcaseValue struct {
	ValueSizeId uint64  `json:"value_size_id,omitempty"`
	Type        string  `json:"type,omitempty"`
	Max         float64 `json:"max,omitempty"`
	Min         float64 `json:"min,omitempty"`
	MaxId       uint64  `json:"max_id,omitempty"`
	MinId       uint64  `json:"min_id,omitempty"`
}

func (c *CommonTestcaseInput) Unfold() utils.CommonTestcaseInput {
	rand.Seed(uint64(time.Now().UnixNano()))
	var input utils.CommonTestcaseInput
	var hsh []float64
	hsh = append(hsh, 0)
	for _, row := range c.Rows {
		if row.RowSizeId > 0 && row.RowSizeId < uint64(len(hsh)) {
			for i := 0; i < int(hsh[row.RowSizeId]); i++ {
				input.AppendRow(row.Unfold(&hsh))
			}
		} else {
			input.AppendRow(row.Unfold(&hsh))
		}
	}
	return input
}

func (c *CommonTestcaseRow) Unfold(hsh *[]float64) utils.CommonTestcaseRow {
	var row utils.CommonTestcaseRow
	for _, v := range c.Values {
		if v.ValueSizeId > 0 && v.ValueSizeId < uint64(len(*hsh)) {
			for i := 0; i < int((*hsh)[v.ValueSizeId]); i++ {
				row.AppendValue(v.Unfold(hsh))
			}
		} else {
			row.AppendValue(v.Unfold(hsh))
		}
	}
	return row
}

func (c *CommonTestcaseValue) Unfold(hsh *[]float64) utils.CommonTestcaseValue {
	if c.MaxId > 0 && c.MaxId < uint64(len(*hsh)) {
		c.Max = (*hsh)[c.MaxId]
	}
	if c.MinId > 0 && c.MinId < uint64(len(*hsh)) {
		c.Min = (*hsh)[c.MinId]
	}
	t := utils.GetValueType(c.Type)
	v := rand.Float64()*(c.Max-c.Min) + c.Min
	*hsh = append(*hsh, v)
	return utils.CommonTestcaseValue{
		Type:  t,
		Value: v,
	}
}
