package sdk

const (
	triggerEveryEvent           = "EveryEvent"
	triggerFirstEvent           = "FirstEvent"
	triggerSubsequentEvents     = "SubsequentEvents"
	triggerEveryThreshold       = "EveryThreshold"
	triggerFirstThreshold       = "FirstThreshold"
	triggerSubsequentThresholds = "SubsequentThreshold"
	triggerAllThresholdsReached = "AllThresholdsReached"
	triggerTimeout              = "Timeout"
)

type triggerCallback func(trigger string, payload *triggerPayload)

type triggerPayload struct {
	eventSpecs []*eventSpecification
	eventSpec  *eventSpecification
	events     []*Event
	counter    *eventCounter
	cellKey    string
	cellReset  chan string
}
