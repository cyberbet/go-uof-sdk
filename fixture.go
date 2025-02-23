package uof

import (
	"encoding/xml"
	"fmt"
	"time"
)

type FixtureRsp struct {
	Fixture     Fixture   `xml:"fixture" json:"fixture"`
	GeneratedAt time.Time `xml:"generated_at,attr,omitempty" json:"generatedAt,omitempty"`
}

// Fixtures describe static or semi-static information about matches and races.
// Reference: https://docs.betradar.com/display/BD/UOF+-+Fixtures+in+the+API
type Fixture struct {
	ID                 int       `xml:"-" json:"id"`
	URN                URN       `xml:"id,attr,omitempty" json:"urn"`
	StartTime          time.Time `xml:"start_time,attr,omitempty" json:"startTime,omitempty"`
	StartTimeConfirmed bool      `xml:"start_time_confirmed,attr,omitempty" json:"startTimeConfirmed,omitempty"`
	StartTimeTbd       bool      `xml:"start_time_tbd,attr,omitempty" json:"startTimeTbd,omitempty"`
	NextLiveTime       time.Time `xml:"next_live_time,attr,omitempty" json:"nextLiveTime,omitempty"`
	Liveodds           string    `xml:"liveodds,attr,omitempty" json:"liveodds,omitempty"`
	Status             string    `xml:"status,attr,omitempty" json:"status,omitempty"`
	Name               string    `xml:"name,attr,omitempty" json:"name,omitempty"`
	Type               string    `xml:"type,attr,omitempty" json:"type,omitempty"`
	Scheduled          time.Time `xml:"scheduled,attr,omitempty" json:"scheduled,omitempty"`
	ScheduledEnd       time.Time `xml:"scheduled_end,attr,omitempty" json:"scheduledEnd,omitempty"`
	ReplacedBy         string    `xml:"replaced_by,attr,omitempty" json:"replacedBy,omitempty"`

	Sport      Sport      `xml:"sport" json:"sport"`
	Category   Category   `xml:"category" json:"category"`
	Tournament Tournament `xml:"tournament,omitempty" json:"tournament,omitempty"`

	Round  Round  `xml:"tournament_round,omitempty" json:"round,omitempty"`
	Season Season `xml:"season,omitempty" json:"season,omitempty"`
	Venue  Venue  `xml:"venue,omitempty" json:"venue,omitempty"`

	ProductInfo ProductInfo  `xml:"product_info,omitempty" json:"productInfo,omitempty"`
	Competitors []Competitor `xml:"competitors>competitor,omitempty" json:"competitors,omitempty"`
	TvChannels  []TvChannel  `xml:"tv_channels>tv_channel,omitempty" json:"tvChannels,omitempty"`

	Home Competitor `json:"home"`
	Away Competitor `json:"away"`

	ExtraInfo []ExtraInfo  `xml:"extra_info>info,omitempty" json:"extraInfo,omitempty"`
	Races     []SportEvent `xml:"races>sport_event,omitempty" json:"races,omitempty"`
	// this also exists but we are skiping for the time being
	//ReferenceIDs         ReferenceIDs         `xml:"reference_ids,omitempty" json:"referenceIDs,omitempty"`
	//SportEventConditions SportEventConditions `xml:"sport_event_conditions,omitempty" json:"sportEventConditions,omitempty"`
	//DelayedInfo DelayedInfo `xml:"delayed_info,omitempty" json:"delayedInfo,omitempty"`
	//CoverageInfo CoverageInfo `xml:"coverage_info,omitempty" json:"coverageInfo,omitempty"`
	//ScheduledStartTimeChanges []ScheduledStartTimeChange `xml:"scheduled_start_time_changes>scheduled_start_time_change,omitempty" json:"scheduledStartTimeChanges,omitempty"`
	//Parent *ParentStage `xml:"parent,omitempty" json:"parent,omitempty"`
}

type FixtureTournament struct {
	ID         int        `xml:"-" json:"id"`
	URN        URN        `xml:"id,attr,omitempty" json:"urn"`
	Name       string     `xml:"name,attr,omitempty" json:"name,omitempty"`
	Sport      Sport      `xml:"sport" json:"sport"`
	Category   Category   `xml:"category" json:"category"`
	Tournament Tournament `xml:"tournament,omitempty" json:"tournament,omitempty"`
	Season     Season     `xml:"season,omitempty" json:"season,omitempty"`
	Groups     []Group    `xml:"groups>group,omitempty" json:"groups,omitempty"`
}

type Group struct {
	Name        string       `xml:"name,attr,omitempty" json:"name"`
	Competitors []Competitor `xml:"competitor,omitempty" json:"competitors,omitempty"`
}

type Tournament struct {
	ID   int    `json:"id"`
	URN  URN    `json:"urn"`
	Name string `xml:"name,attr" json:"name"`
}

type Sport struct {
	ID   int    `json:"id"`
	URN  URN    `json:"urn"`
	Name string `xml:"name,attr" json:"name"`
}

type Category struct {
	ID          int    `json:"id"`
	URN         URN    `json:"urn"`
	Name        string `xml:"name,attr" json:"name"`
	CountryCode string `xml:"country_code,attr,omitempty" json:"countryCode,omitempty"`
}

type Competitor struct {
	ID           int                `json:"id"`
	URN          URN                `json:"urn"`
	Qualifier    string             `xml:"qualifier,attr,omitempty" json:"qualifier,omitempty"`
	Name         string             `xml:"name,attr" json:"name"`
	Abbreviation string             `xml:"abbreviation,attr" json:"abbreviation"`
	Country      string             `xml:"country,attr,omitempty" json:"country,omitempty"`
	CountryCode  string             `xml:"country_code,attr,omitempty" json:"countryCode,omitempty"`
	Virtual      bool               `xml:"virtual,attr,omitempty" json:"virtual,omitempty"`
	Players      []CompetitorPlayer `xml:"players>player,omitempty" json:"players,omitempty"`
	//ReferenceIDs CompetitorReferenceIDs `xml:"reference_ids,omitempty" json:"referenceIDs,omitempty"`
}

type CompetitorPlayer struct {
	ID           int    `json:"id"`
	URN          URN    `json:"urn"`
	Name         string `xml:"name,attr" json:"name"`
	Abbreviation string `xml:"abbreviation,attr" json:"abbreviation"`
	Nationality  string `xml:"nationality,attr,omitempty" json:"nationality,omitempty"`
}

type Venue struct {
	ID             int    `json:"id"`
	Name           string `xml:"name,attr" json:"name"`
	Capacity       int    `xml:"capacity,attr,omitempty" json:"capacity,omitempty"`
	CityName       string `xml:"city_name,attr,omitempty" json:"cityName,omitempty"`
	CountryName    string `xml:"country_name,attr,omitempty" json:"countryName,omitempty"`
	CountryCode    string `xml:"country_code,attr,omitempty" json:"countryCode,omitempty"`
	MapCoordinates string `xml:"map_coordinates,attr,omitempty" json:"mapCoordinates,omitempty"`
}

type TvChannel struct {
	Name string `xml:"name,attr" json:"name"`
	// seams to be always zero
	// StartTime time.Time `xml:"start_time,attr,omitempty" json:"startTime,omitempty"`
}

type StreamingChannel struct {
	ID   int    `xml:"id,attr" json:"id"`
	Name string `xml:"name,attr" json:"name"`
}
type ProductInfoLink struct {
	Name string `xml:"name,attr" json:"name"`
	Ref  string `xml:"ref,attr" json:"ref"`
}

type Round struct {
	ID                  int    `xml:"betradar_id,attr,omitempty" json:"id,omitempty"`
	Type                string `xml:"type,attr,omitempty" json:"type,omitempty"`
	Number              int    `xml:"number,attr,omitempty" json:"number,omitempty"`
	Name                string `xml:"name,attr,omitempty" json:"name,omitempty"`
	GroupLongName       string `xml:"group_long_name,attr,omitempty" json:"groupLongName,omitempty"`
	Group               string `xml:"group,attr,omitempty" json:"group,omitempty"`
	GroupID             string `xml:"group_id,attr,omitempty" json:"groupID,omitempty"`
	CupRoundMatches     int    `xml:"cup_round_matches,attr,omitempty" json:"cupRoundMatches,omitempty"`
	CupRoundMatchNumber int    `xml:"cup_round_match_number,attr,omitempty" json:"cupRoundMatchNumber,omitempty"`
	OtherMatchID        string `xml:"other_match_id,attr,omitempty" json:"otherMatchID,omitempty"`
}

type Season struct {
	ID        int    `json:"id"`
	URN       URN    `json:"urn"`
	StartDate string `xml:"start_date,attr" json:"startDate"`
	EndDate   string `xml:"end_date,attr" json:"endDate"`
	StartTime string `xml:"start_time,attr,omitempty" json:"startTime,omitempty"`
	EndTime   string `xml:"end_time,attr,omitempty" json:"endTime,omitempty"`
	Year      string `xml:"year,attr,omitempty" json:"year,omitempty"`
	Name      string `xml:"name,attr" json:"name"`
	//TournamentID string    `xml:"tournament_id,attr,omitempty" json:"tournamentID,omitempty"`
}

// type ParentStage struct {
// 	URN          URN       `xml:"id,attr,omitempty" json:"urn,omitempty"`
// 	Name         string    `xml:"name,attr,omitempty" json:"name,omitempty"`
// 	Type         string    `xml:"type,attr,omitempty" json:"type,omitempty"`
// 	Scheduled    time.Time `xml:"scheduled,attr,omitempty" json:"scheduled,omitempty"`
// 	StartTimeTbd bool      `xml:"start_time_tbd,attr,omitempty" json:"startTimeTbd,omitempty"`
// 	ScheduledEnd time.Time `xml:"scheduled_end,attr,omitempty" json:"scheduledEnd,omitempty"`
// 	ReplacedBy   string    `xml:"replaced_by,attr,omitempty" json:"replacedBy,omitempty"`
// }

// type ScheduledStartTimeChange struct {
// 	OldTime   time.Time `xml:"old_time,attr" json:"oldTime"`
// 	NewTime   time.Time `xml:"new_time,attr" json:"newTime"`
// 	ChangedAt time.Time `xml:"changed_at,attr" json:"changedAt"`
// }

type ProductInfo struct {
	Streaming            []StreamingChannel `xml:"streaming>channel,omitempty" json:"streaming,omitempty"`
	IsInLiveScore        string             `xml:"is_in_live_score,omitempty" json:"isInLiveScore,omitempty"`
	IsInHostedStatistics string             `xml:"is_in_hosted_statistics,omitempty" json:"isInHostedStatistics,omitempty"`
	IsInLiveCenterSoccer string             `xml:"is_in_live_center_soccer,omitempty" json:"isInLiveCenterSoccer,omitempty"`
	IsAutoTraded         string             `xml:"is_auto_traded,omitempty" json:"isAutoTraded,omitempty"`
	Links                []ProductInfoLink  `xml:"links>link,omitempty" json:"links,omitempty"`
}

// ExtraInfo covers additional fixture information about the match,
// such as coverage information, extended markets offer, additional rules etc.
type ExtraInfo struct {
	Key   string `xml:"key,attr,omitempty" json:"key,omitempty"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

// SportEvent covers information about scheduled races in a stage
// For VHC and VDR information is in vdr/vhc:stage:<int> fixture with type="parent"
type SportEvent struct {
	ID           string    `xml:"id,attr,omitempty" json:"id,omitempty"`
	Name         string    `xml:"name,attr,omitempty" json:"name,omitempty"`
	Type         string    `xml:"type,attr,omitempty" json:"type,omitempty"`
	Scheduled    time.Time `xml:"scheduled,attr,omitempty" json:"scheduled,omitempty"`
	ScheduledEnd time.Time `xml:"scheduled_end,attr,omitempty" json:"scheduled_end,omitempty"`
}

type Summary struct {
	SportEvent         Fixture             `xml:"sport_event" json:"sportEvent"`
	GeneratedAt        time.Time           `xml:"generated_at,attr,omitempty" json:"generatedAt,omitempty"`
	SummaryEventStatus *SummaryEventStatus `xml:"sport_event_status,omitempty" json:"summaryEventStatus,omitempty"`
}

type MatchTimeline struct {
	SportEvent  Fixture   `xml:"sport_event" json:"sportEvent"`
	GeneratedAt time.Time `xml:"generated_at,attr,omitempty" json:"generatedAt,omitempty"`
	//SportEventStatus SportEventStatus `xml:"sport_event_status" json:"sportEventStatus,omitempty"`
	Events []TimelineEvent `xml:"timeline>event" json:"events,omitempty"`
	//SportEventConditions struct {
	//	Text  string `xml:",chardata"`
	//	Venue struct {
	//		Text           string `xml:",chardata"`
	//		ID             string `xml:"id,attr"`
	//		Name           string `xml:"name,attr"`
	//		Capacity       string `xml:"capacity,attr"`
	//		CityName       string `xml:"city_name,attr"`
	//		CountryName    string `xml:"country_name,attr"`
	//		MapCoordinates string `xml:"map_coordinates,attr"`
	//		CountryCode    string `xml:"country_code,attr"`
	//	} `xml:"venue"`
	//} `xml:"sport_event_conditions"`
}

type TimelineEvent struct {
	ID              string     `xml:"id,attr"  json:"id,omitempty"`
	Type            string     `xml:"type,attr" json:"type,omitempty"`
	Time            time.Time  `xml:"time,attr" json:"time,omitempty"`
	Period          string     `xml:"period,attr" json:"period,omitempty"`
	PeriodName      string     `xml:"period_name,attr" json:"periodName,omitempty"`
	MatchStatusCode string     `xml:"match_status_code,attr" json:"matchStatusCode,omitempty"`
	MatchTime       string     `xml:"match_time,attr" json:"matchTime,omitempty"`
	MatchClock      string     `xml:"match_clock,attr" json:"matchClock,omitempty"`
	Team            string     `xml:"team,attr" json:"team,omitempty"`
	X               string     `xml:"x,attr" json:"x,omitempty"`
	Y               string     `xml:"y,attr" json:"y,omitempty"`
	HomeScore       string     `xml:"home_score,attr" json:"homeScore,omitempty"`
	AwayScore       string     `xml:"away_score,attr" json:"awayScore,omitempty"`
	GoalScorer      GoalScorer `xml:"goal_scorer" json:"goalScorer,omitempty"`
	Assist          Assist     `xml:"assist" json:"assist,omitempty"`
}

type GoalScorer struct {
	ID   string `xml:"id,attr" json:"id,omitempty"`
	Name string `xml:"name,attr" json:"name,omitempty"`
}

type Assist struct {
	ID   string `xml:"id,attr" json:"id,omitempty"`
	Name string `xml:"name,attr" json:"name,omitempty"`
	Type string `xml:"type,attr" json:"type,omitempty"`
}

//type Statistics struct {
//	Text   string `xml:",chardata"`
//	Totals struct {
//		Text  string `xml:",chardata"`
//		Teams struct {
//			Text string `xml:",chardata"`
//			Team []struct {
//				Text       string `xml:",chardata"`
//				ID         string `xml:"id,attr"`
//				Name       string `xml:"name,attr"`
//				Statistics struct {
//					Text        string `xml:",chardata"`
//					Cards       string `xml:"cards,attr"`
//					CornerKicks string `xml:"corner_kicks,attr"`
//					YellowCards string `xml:"yellow_cards,attr"`
//				} `xml:"statistics"`
//			} `xml:"team"`
//		} `xml:"teams"`
//	} `xml:"totals"`
//	Periods struct {
//		Text   string `xml:",chardata"`
//		Period []struct {
//			Text  string `xml:",chardata"`
//			Name  string `xml:"name,attr"`
//			Teams struct {
//				Text string `xml:",chardata"`
//				Team []struct {
//					Text       string `xml:",chardata"`
//					ID         string `xml:"id,attr"`
//					Name       string `xml:"name,attr"`
//					Statistics struct {
//						Text        string `xml:",chardata"`
//						Cards       string `xml:"cards,attr"`
//						CornerKicks string `xml:"corner_kicks,attr"`
//						YellowCards string `xml:"yellow_cards,attr"`
//					} `xml:"statistics"`
//				} `xml:"team"`
//			} `xml:"teams"`
//		} `xml:"period"`
//	} `xml:"periods"`
//}

// slici na sport_event_status ali statusi nisu int nego string
type SummaryEventStatus struct {
	Status       string        `xml:"status,attr" json:"status"`
	MatchStatus  string        `xml:"match_status,attr" json:"matchStatus"`
	HomeScore    *int          `xml:"home_score,attr,omitempty" json:"homeScore,omitempty"`
	AwayScore    *int          `xml:"away_score,attr,omitempty" json:"awayScore,omitempty"`
	PeriodScores []PeriodScore `xml:"period_scores>period_score,omitempty" json:"periodScores,omitempty"`
}

func (f *Fixture) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Fixture
	var overlay struct {
		*T
		Tournament *struct {
			URN      URN      `xml:"id,attr"`
			Name     string   `xml:"name,attr"`
			Sport    Sport    `xml:"sport"`
			Category Category `xml:"category"`
		} `xml:"tournament,omitempty"`
	}
	overlay.T = (*T)(f)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	f.ID = overlay.URN.EventID()
	f.Sport = overlay.Tournament.Sport
	f.Category = overlay.Tournament.Category
	f.Tournament.ID = overlay.Tournament.URN.ID()
	f.Tournament.URN = overlay.Tournament.URN
	f.Tournament.Name = overlay.Tournament.Name

	for _, c := range f.Competitors {
		if c.Qualifier == "home" {
			f.Home = c
		}
		if c.Qualifier == "away" {
			f.Away = c
		}
	}
	return nil
}

func (t *Sport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Sport
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	t.URN = overlay.URN
	return nil
}

func (t *Category) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Category
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	t.URN = overlay.URN
	return nil
}

func (t *Season) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Season
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	t.URN = overlay.URN
	return nil
}

func (t *Venue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Venue
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	return nil
}

func (t *Competitor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Competitor
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	t.URN = overlay.URN
	return nil
}

func (t *CompetitorPlayer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T CompetitorPlayer
	var overlay struct {
		*T
		URN URN `xml:"id,attr"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	t.ID = overlay.URN.ID()
	t.URN = overlay.URN
	return nil
}

func (t *FixtureTournament) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T FixtureTournament
	var overlay struct {
		*T
		Tournament *struct {
			URN      URN      `xml:"id,attr"`
			Name     string   `xml:"name,attr"`
			Sport    Sport    `xml:"sport"`
			Category Category `xml:"category"`
		} `xml:"tournament,omitempty"`
	}
	overlay.T = (*T)(t)
	if err := d.DecodeElement(&overlay, &start); err != nil {
		return err
	}
	if overlay.Tournament != nil {
		t.Sport = overlay.Tournament.Sport
		t.Category = overlay.Tournament.Category
		t.Tournament.ID = overlay.Tournament.URN.ID()
		t.Tournament.URN = overlay.Tournament.URN
		t.Tournament.Name = overlay.Tournament.Name
		t.ID = overlay.Tournament.URN.ID()
		t.URN = overlay.Tournament.URN
	}
	return nil
}

// PP pretty prints fixure row
func (f *Fixture) PP() string {
	name := fmt.Sprintf("%s - %s", f.Home.Name, f.Away.Name)
	return fmt.Sprintf("%-90s %12s %15s", name, f.Scheduled.Format("02.01. 15:04"), f.Status)
}
