package model

import "encoding/xml"

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
