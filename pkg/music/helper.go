package music

func getNotesFrom(start *Note, steps []*Step) []*Note {
	base, natural := start.BaseNote(), start.Natural
	var out []*Note
	for _, step := range steps {
		nextBase := base.Next(step.Base)
		nextNatural := natural.Next(step.Natural)
		out = append(out, nextBase.KeyNote(nextNatural))
	}
	return out
}
