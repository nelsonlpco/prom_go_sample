package gene

import (
	"fmt"
	"io"
	"log"
	"strings"
	"unsafe"
)

type CompressedGene struct {
	gene []byte
}

func New(gene string) (*CompressedGene, error) {
	return new(CompressedGene).withGene(gene)
}

func (g *CompressedGene) withGene(gene string) (*CompressedGene, error) {
	compressedGene, err := compress(gene)

	if err != nil {
		log.Printf("invalid gene %v", gene)
		return nil, err
	}

	g.gene = compressedGene
	return g, nil
}

func (g *CompressedGene) GetGene() string {
	var gene string
	for _, b := range g.gene {
		bits := fmt.Sprintf("%08b", b)

		for n := 0; n < 8; n += 2 {
			nucleotid := bits[n : n+2]

			if nucleotid == "00" {
				gene += "A"
			} else if nucleotid == "01" {
				gene += "C"
			} else if nucleotid == "10" {
				gene += "G"
			} else if nucleotid == "11" {
				gene += "T"
			}
		}
	}

	size := unsafe.Sizeof(gene)
	t := len(gene)

	fmt.Printf("a: %T, %d, %d\n", gene, size, t)
	return gene
}

func compress(gene string) ([]byte, error) {
	var g []byte
	reader := strings.NewReader(gene)
	buf := make([]byte, 4)
	for {
		_, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		var b byte = 0x1

		for _, n := range buf {
			b = b << 2

			if n == 'A' {
				b = b | 0x0
			} else if n == 'C' {
				b = b | 0x1
			} else if n == 'G' {
				b = b | 0x2
			} else if n == 'T' {
				b = b | 0x3
			} else {
				err = fmt.Errorf("invalid nucletidio %v", b)
				break
			}
		}

		if err != nil {
			return nil, err
		}

		g = append(g, b)
	}
	return g, nil
}
