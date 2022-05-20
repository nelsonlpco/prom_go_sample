package gene

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Gene(t *testing.T) {
	gene := "ATAATGAC"
	compressedGene, err := New(gene)

	assert.NoError(t, err)
	assert.Equal(t, gene, compressedGene.GetGene())
}
