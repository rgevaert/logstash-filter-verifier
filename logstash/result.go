// Copyright (c) 2015 Magnus Bäck <magnus@noun.se>

package logstash

// Result contains the result of a Logstash execution.
type Result struct {
	// Success indicates whether the execution was successful,
	// i.e. whether the Logstash process terminated with a zero
	// exit status.
	Success bool

	// Events contains a slice of the events emitted from
	// Logstash.
	Events []Event

	// Log contains the contents of the Logstash log file.
	Log string
}

// Event represents a Logstash event, i.e. basically a JSON document.
type Event map[string]interface{}
