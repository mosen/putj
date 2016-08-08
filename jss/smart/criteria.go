package smart

type Criterion struct {
	Name       string `xml:"name" json:"name"`
	Priority   int    `xml:"priority" json:"priority"`
	AndOr      string `xml:"and_or" json:"and_or"`
	SearchType string `xml:"not_like" json:"not_like"`
	Value      string `xml:"value" json:"value"`
}
