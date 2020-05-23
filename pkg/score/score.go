package score

type Scorer interface {
	Score(int, int) bool
}

func New(line int) Scorer {
	return &scorerImpl{line: line}
}

type scorerImpl struct {
	line int
}

func (scorer *scorerImpl) Score(sc, absent int) bool {
	return sc-5*absent >= scorer.line
}
