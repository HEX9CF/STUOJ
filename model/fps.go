package model

import (
	"encoding/xml"
	"strconv"
)

type FPS struct {
	XMLName   xml.Name  `xml:"fps"`
	Version   string    `xml:"version,attr"`
	URL       string    `xml:"url,attr,omitempty"`
	Generator Generator `xml:"generator,omitempty"`
	Items     []Item    `xml:"item"`
}

type Generator struct {
	Name string `xml:"name,attr"`
	URL  string `xml:"url,attr,omitempty"`
}

type Item struct {
	Title        string      `xml:"title"`
	TimeLimit    TimeLimit   `xml:"time_limit,omitempty"`
	MemoryLimit  MemoryLimit `xml:"memory_limit,omitempty"`
	Description  string      `xml:"description"`
	Input        string      `xml:"input"`
	Output       string      `xml:"output"`
	SampleInput  string      `xml:"sample_input"`
	SampleOutput string      `xml:"sample_output"`
	TestInput    []string    `xml:"test_input,omitempty"`
	TestOutput   []string    `xml:"test_output,omitempty"`
	Hint         string      `xml:"hint,omitempty"`
	Source       string      `xml:"source,omitempty"`
	Solution     []Solution  `xml:"solution,omitempty"`
}

type TimeLimit struct {
	Unit string `xml:"unit,attr"`
	Data string `xml:",chardata"`
}

type MemoryLimit struct {
	Unit string `xml:"unit,attr"`
	Data string `xml:",chardata"`
}

type Solution struct {
	Language string `xml:"language,attr"`
	Code     string `xml:",chardata"`
}

func (i *Item) ToProblem() Problem {
	var timeLimit float64
	var memoryLimit uint64
	timetmp, _ := strconv.ParseFloat(i.TimeLimit.Data, 64)
	memorytmp, _ := strconv.ParseInt(i.MemoryLimit.Data, 10, 64)
	if i.TimeLimit.Unit == "ms" {
		timeLimit = 1000 * timetmp
	} else if i.TimeLimit.Unit == "s" {
		timeLimit = timetmp
	}
	if i.MemoryLimit.Unit == "kb" {
		memoryLimit = uint64(memorytmp)
	} else if i.MemoryLimit.Unit == "mb" {
		memoryLimit = 1024 * uint64(memorytmp)
	}

	return Problem{
		Title:        i.Title,
		Description:  i.Description,
		Input:        i.Input,
		Output:       i.Output,
		SampleInput:  i.SampleInput,
		SampleOutput: i.SampleOutput,
		Hint:         i.Hint,
		TimeLimit:    timeLimit,
		MemoryLimit:  memoryLimit,
	}
}
