package score_test

import (
	"testing"

	"github.com/hsmtkk/c056_score/pkg/score"
	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	scorer := score.New(25)
	assert.True(t, scorer.Score(80, 11), "should be true")
	assert.False(t, scorer.Score(20, 1), "should be false")
	assert.True(t, scorer.Score(50, 2), "should be true")
	assert.True(t, scorer.Score(70, 0), "should be true")
	assert.False(t, scorer.Score(25, 1), "should be false")
}
