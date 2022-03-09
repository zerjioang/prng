package prng

import "time"

var (
	g = rngSource{}
)

func init() {
	g.Seed(time.Now().Unix())
}

// New returns a new UUID as string
func New() string {
	var dst RandUUID
	g.UUID2(&dst)
	return string(dst[:])
}

// New returns a new UUID
func NewRandUUID() RandUUID {
	var dst RandUUID
	g.UUID2(&dst)
	return dst
}
