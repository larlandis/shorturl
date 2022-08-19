package metrics

import (
	"context"
	"sync"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var m metric
var once sync.Once

type (
	metric struct {
		writeApi api.WriteAPI
	}
	tag struct {
		key   string
		value string
	}
	segment struct {
		p   *write.Point
		ctx context.Context
	}
)

func Tag(k, v string) tag {
	return tag{k, v}
}

func (s *segment) End(tags ...tag) {
	duration := time.Since(s.p.Time()).Nanoseconds()
	p := s.p.AddField("duration", duration)
	for _, tag := range tags {
		p = p.AddTag(tag.key, tag.value)
	}
	m.writeApi.WritePoint(p)
}

func Init(metricServer string) {
	once.Do(func() {
		client := influxdb2.NewClient(metricServer, "token")
		m.writeApi = client.WriteAPI("", "shorturl")
	})
}

func Count(_ context.Context, event string, tags ...tag) {
	ts := make(map[string]string, len(tags))
	for _, tag := range tags {
		ts[tag.key] = tag.value
	}
	m.writeApi.WritePoint(influxdb2.NewPoint(
		event, ts,
		map[string]interface{}{"count": 1},
		time.Now(),
	))
}

func Segment(ctx context.Context, event string, tags ...tag) *segment {
	ts := make(map[string]string, len(tags))
	for _, tag := range tags {
		ts[tag.key] = tag.value
	}
	return &segment{
		ctx: ctx,
		p:   influxdb2.NewPoint(event, ts, nil, time.Now()),
	}
}
