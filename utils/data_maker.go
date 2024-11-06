package utils

type ValueType uint8

const (
	ValueType_Int   ValueType = 0
	ValueType_Float ValueType = 1
	ValueType_Char  ValueType = 2
)

func (v ValueType) String() string {
	switch v {
	case ValueType_Int:
		return "int"
	case ValueType_Char:
		return "char"
	case ValueType_Float:
		return "float"
	default:
		return "unknown"
	}
}

type CommonTestcaseInput struct {
	Rows []CommonTestcaseRow
}

type CommonTestcaseRow struct {
	CommonValue []CommonTestcaseValue
}

type CommonTestcaseValue struct {
	Type ValueType
	Max  float64
	Min  float64
}

func (c *CommonTestcaseInput) AppendRow(row CommonTestcaseRow) {
	c.Rows = append(c.Rows, row)
}
func (c *CommonTestcaseInput) Size() uint64 {
	return uint64(len(c.Rows))
}

func (c *CommonTestcaseInput) GetRow(index uint64) CommonTestcaseRow {
	return c.Rows[index]
}

func (c *CommonTestcaseRow) AppendValue(value CommonTestcaseValue) {
	c.CommonValue = append(c.CommonValue, value)
}

func (c *CommonTestcaseRow) Size() uint64 {
	return uint64(len(c.CommonValue))
}

func (c *CommonTestcaseRow) GetValue(index uint64) CommonTestcaseValue {
	return c.CommonValue[index]
}
