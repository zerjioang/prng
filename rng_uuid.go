package prng

const (
	hextable = "0123456789abcdef"
)

type RandUUID [16]byte

/*
 * UUID generates a random uuid with Uniform distribution
 *
 * algorithm by
 * DP Mitchell and JA Reeds
 */
func (rng *rngSource) UUID() string {
	var uid RandUUID
	rng.UUID2(&uid)
	return string(uid[:])
}

func (rng *rngSource) UUID2(uid *RandUUID) {
	u := rng.Int63()
	_ = uid[15]
	uid[0] = byte(u)
	uid[2] = byte(u >> 8)
	uid[4] = byte(u >> 16)
	uid[6] = byte(u >> 24)
	uid[8] = byte(u >> 32)
	uid[10] = byte(u >> 40)
	uid[12] = byte(u >> 48)
	uid[14] = byte(u >> 56)
	// compressed loop
	/*
		for i := 0; i < size; i++ {
			idx := i * 2
			v := uid[idx]
			uid[idx] = hextable[v>>4]
			uid[idx+1] = hextable[v&0x0f]
		}
	*/

	// expanded for loop
	var v byte
	_ = uid[15]
	v = uid[0]
	uid[0] = hextable[v>>4]
	uid[1] = hextable[v&0x0f]

	v = uid[2]
	uid[2] = hextable[v>>4]
	uid[3] = hextable[v&0x0f]

	v = uid[4]
	uid[4] = hextable[v>>4]
	uid[5] = hextable[v&0x0f]

	v = uid[6]
	uid[6] = hextable[v>>4]
	uid[7] = hextable[v&0x0f]

	v = uid[8]
	uid[8] = hextable[v>>4]
	uid[9] = hextable[v&0x0f]

	v = uid[10]
	uid[10] = hextable[v>>4]
	uid[11] = hextable[v&0x0f]

	v = uid[12]
	uid[12] = hextable[v>>4]
	uid[13] = hextable[v&0x0f]

	v = uid[14]
	uid[14] = hextable[v>>4]
	uid[15] = hextable[v&0x0f]
}
