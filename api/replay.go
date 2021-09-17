package api

import (
	"github.com/minus5/go-uof-sdk"
)

// replay api paths
const (
	startScenario = "/v1/replay/scenario/play/{{.ScenarioID}}?speed={{.Speed}}&max_delay={{.MaxDelay}}&use_replay_timestamp={{.UseReplayTimestamp}}"
	replayStop    = "/v1/replay/stop"
	replayReset   = "/v1/replay/reset"
	replayAdd     = "/v1/replay/events/{{.EventURN}}"
	replayPlay    = "/v1/replay/play?speed={{.Speed}}&max_delay={{.MaxDelay}}&use_replay_timestamp={{.UseReplayTimestamp}}"
)

// StartScenario replay of the scenario from replay queue. Your current playlist will be
// wiped, and populated with events from specified scenario. Events are played
// in the order they were played in reality. Parameters 'speed' and 'max_delay'
// specify the speed of replay and what should be the maximum delay between
// messages. Default values for these are speed = 10 and max_delay = 10000. This
// means that messages will be sent 10x faster than in reality, and that if
// there was some delay between messages that was longer than 10 seconds it will
// be reduced to exactly 10 seconds/10 000 ms (this is helpful especially in
// pre-match odds where delay can be even a few hours or more). If player is
// already in play, nothing will happen.
func (a *API) StartScenario(scenarioID, speed, maxDelay int) error {
	return a.post(startScenario, &params{ScenarioID: scenarioID, Speed: speed, MaxDelay: maxDelay})
}

// StartEvent starts replay of a single event.
func (a *API) StartEvent(eventURN uof.URN, speed, maxDelay int) error {
	if err := a.Reset(); err != nil {
		return err
	}
	if err := a.Add(eventURN); err != nil {
		return err
	}
	return a.Play(speed, maxDelay)
}

// Add to the end of the replay queue.
func (a *API) Add(eventURN uof.URN) error {
	return a.put(replayAdd, &params{EventURN: eventURN})
}

// Play replay the events from replay queue. Events are played in the order
// they were played in reality. Parameters 'speed' and 'max_delay' specify the
// speed of replay and what should be the maximum delay between messages.
// Default values for these are speed = 10 and max_delay = 10000. This means
// that messages will be sent 10x faster than in reality, and that if there was
// some delay between messages that was longer than 10 seconds it will be
// reduced to exactly 10 seconds/10 000 ms (this is helpful especially in
// pre-match odds where delay can be even a few hours or more). If player is
// already in play, nothing will happen.
func (a *API) Play(speed, maxDelay int) error {
	return a.post(replayPlay, &params{Speed: speed, MaxDelay: maxDelay})
}

// Stop the player if it is currently playing. If player is already stopped,
// nothing will happen.
func (a *API) Stop() error {
	return a.post(replayStop, nil)
}

// Reset the player if it is currently playing and clear the replay queue. If
// player is already stopped, the queue is cleared.
func (a *API) Reset() error {
	return a.post(replayReset, nil)
}
