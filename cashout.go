package uof

type Cashout struct {
	Product          int              `xml:"product,attr" json:"product"`
	EventID          string           `xml:"event_id,attr" json:"event_id"`
	Timestamp        int              `xml:"timestamp,attr" json:"timestamp"`
	SportEventStatus SportEventStatus `xml:"sport_event_status" json:"sport_event_status"`
	Odds             []Market         `xml:"odds>market" json:"odds"`
}
