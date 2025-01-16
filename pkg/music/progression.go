package music

// e.g. 1-6-4-5
type Progression struct {
	Chords []*Chord
}

func (p *Progression) Apply(key *Key) {
	panic("TODO")
}
