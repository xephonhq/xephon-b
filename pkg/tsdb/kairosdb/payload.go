package kairosdb

type KairosDBPayload struct {
	// TODO: use byte writer
}

func (p *KairosDBPayload) AddPoint() {
	// this write the point as bytes into buffer
}

func (p *KairosDBPayload) AddPointToBuffer() {
	// this store the struct and merge into one series when get the string ([]byte) , actually it's a group by
}

func (p *KairosDBPayload) groupBySeries() {

}

func (p *KairosDBPayload) DataSize() int {
	// the real data size,
	// TODO: count series data several times?
	// TODO: the payload size, they are all different
	return 0
}
