package dgen

// Austen McClernon
// 834063
// Assignment 1 COMP90077

import (
	"math/rand"

	"github.com/kvoli/COMP90077_ASS/pkg/elm"
)

// Generator represents a generator for sequential random values
type Generator struct {
	IDNext    int
	Generated []*elm.Element
	Deleted   map[int]bool
}

// New returns a new generator
func New() *Generator {
	return &Generator{
		IDNext:    0,
		Generated: make([]*elm.Element, 0),
	}
}

// GenSizeList Generates a list of elements of size n
func (g *Generator) GenSizeList(n int) []*elm.Element {
	elmList := make([]*elm.Element, n)

	for i := 0; i < n; i++ {
		elmList = append(elmList, g.GenElement())
	}
	return elmList
}

// GenElement generates a new GenElement
func (g *Generator) GenElement() *elm.Element {

	newGenElement := &elm.Element{
		ID:  g.IDNext + 1,
		Key: int(rand.Int31n(MaximumKey)),
	}

	// update Generator
	g.Generated = append(g.Generated, newGenElement)
	g.IDNext++
	return newGenElement
}

// GenInsertion returns an insertion operation
func (g *Generator) GenInsertion() *elm.Operation {
	return &elm.Operation{
		Element: g.GenElement(),
		Type:    elm.InsertOperation,
	}
}

// GenDeletion returns a deletion operation
func (g *Generator) GenDeletion() *elm.Operation {
	randomID := int(rand.Int31n(int32(g.IDNext)))
	var key int
	_, ok := g.Deleted[randomID]
	if !ok {
		key = int(rand.Int31n(MaximumKey))
	} else {
		key = g.Generated[randomID].Key
		g.Deleted[randomID] = true
	}
	return &elm.Operation{
		Element: &elm.Element{
			ID:  -1,
			Key: key,
		},
		Type: elm.DeleteOperation,
	}
}

// GenSearch returns a search operation
func (g *Generator) GenSearch() *elm.Operation {
	randomSearchKey := int(rand.Int31n(MaximumKey))
	return &elm.Operation{
		Element: &elm.Element{
			ID:  -1,
			Key: randomSearchKey,
		},
		Type: elm.SearchOperation,
	}
}
