package uof

type MatchStatusDescriptions []MatchStatus

type MatchStatus struct {
	ID           string            `xml:"id,attr"`
	Description  string            `xml:"description,attr"`
	PeriodNumber string            `xml:"period_number,attr"`
	Sports       MatchStatusSports `xml:"sports"`
}

type MatchStatusSports struct {
	All   bool               `xml:"all,attr"`
	Sport []MatchStatusSport `xml:"sport"`
}

type MatchStatusSport struct {
	ID string `xml:"id,attr"`
}
