package mdtracer

import (
	"log"
	"time"

	"github.com/short-d/app/fw"

	uuid "github.com/satori/go.uuid"
)

type Local struct{}

type LocalTrace struct {
	id    string
	name  string
	start time.Time
}

func (t LocalTrace) Next(name string) fw.Segment {
	start := time.Now()
	log.Printf("[Trace Start id=%s name=%s startAt=%v]", t.id, name, start)
	return LocalTrace{
		id:    t.id,
		name:  name,
		start: start,
	}
}

func (t LocalTrace) End() {
	end := time.Now()
	diff := end.Sub(t.start)
	log.Printf("[Trace End   id=%s name=%s endAt=%v duration=%v]", t.id, t.name, end, diff)
}

func (Local) BeginTrace(name string) fw.Segment {
	id := uuid.NewV4().String()
	start := time.Now()

	log.Printf("[Trace Start id=%s name=%s startAt=%v]", id, name, start)
	return LocalTrace{
		id:    id,
		name:  name,
		start: start,
	}
}

func NewLocal() fw.Tracer {
	return Local{}
}
