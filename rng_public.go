package prng

import "time"

var (
	g = rngSource{}
)

func init() {
	g.Seed(time.Now().Unix())
}

// New returns a new UUID
func New() string {
	var dst RandUUID
	g.UUID2(&dst)
	return string(dst[:])
}
