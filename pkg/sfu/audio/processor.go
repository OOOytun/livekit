package audio

import "github.com/livekit/protocol/livekit"

// Processor processes PCM audio frames. Implementations may modify the frame in place.
// Returns true if the frame was modified.
type Processor interface {
	ProcessPCM(frame []float32) bool
}

// ProcessorFactory creates per-track audio processors.
type ProcessorFactory interface {
	NewProcessor(trackID livekit.TrackID) Processor
}

// NoopProcessorFactory is the default passthrough factory.
var NoopProcessorFactory ProcessorFactory = noopProcessorFactory{}

type noopProcessorFactory struct{}

func (noopProcessorFactory) NewProcessor(livekit.TrackID) Processor {
	return noopProcessor{}
}

type noopProcessor struct{}

func (noopProcessor) ProcessPCM([]float32) bool {
	return false
}
