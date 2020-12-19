package main

// Austen McClernon
// 834063
// Assignment 1 COMP90077

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/kvoli/COMP90077_ASS/pkg/treap"

	"github.com/kvoli/COMP90077_ASS/pkg/darry"
	"github.com/kvoli/COMP90077_ASS/pkg/dgen"
	"github.com/kvoli/COMP90077_ASS/pkg/elm"
)

const testsize = 10e5

func experimentOne() ([]float64, []float64) {
	darry1, treap1 := handleOperations(genSequence(0, 1, 0, testsize*0.1))
	darry2, treap2 := handleOperations(genSequence(0, 1, 0, testsize*0.2))
	darry3, treap3 := handleOperations(genSequence(0, 1, 0, testsize*0.5))
	darry4, treap4 := handleOperations(genSequence(0, 1, 0, testsize*0.8))
	darry5, treap5 := handleOperations(genSequence(0, 1, 0, testsize))
	return []float64{darry1, darry2, darry3, darry4, darry5}, []float64{treap1, treap2, treap3, treap4, treap5}
}

func experimentTwo() ([]float64, []float64) {
	darry1, treap1 := handleOperations(genSequence(0.001, 0.999, 0, testsize))
	darry2, treap2 := handleOperations(genSequence(0.005, 0.995, 0, testsize))
	darry3, treap3 := handleOperations(genSequence(0.01, 0.99, 0, testsize))
	darry4, treap4 := handleOperations(genSequence(0.05, 0.95, 0, testsize))
	darry5, treap5 := handleOperations(genSequence(0.10, 0.90, 0, testsize))
	return []float64{darry1, darry2, darry3, darry4, darry5}, []float64{treap1, treap2, treap3, treap4, treap5}
}

func experimentThree() ([]float64, []float64) {
	darry1, treap1 := handleOperations(genSequence(0, 0.999, 0.001, testsize))
	darry2, treap2 := handleOperations(genSequence(0, 0.995, 0.005, testsize))
	darry3, treap3 := handleOperations(genSequence(0, 0.99, 0.01, testsize))
	darry4, treap4 := handleOperations(genSequence(0, 0.95, 0.05, testsize))
	darry5, treap5 := handleOperations(genSequence(0, 0.90, 0.10, testsize))
	return []float64{darry1, darry2, darry3, darry4, darry5}, []float64{treap1, treap2, treap3, treap4, treap5}
}

func experimentFour() ([]float64, []float64) {
	darry1, treap1 := handleOperations(genSequence(0.05, 0.90, 0.05, testsize*0.1))
	darry2, treap2 := handleOperations(genSequence(0.05, 0.90, 0.05, testsize*0.2))
	darry3, treap3 := handleOperations(genSequence(0.05, 0.90, 0.05, testsize*0.5))
	darry4, treap4 := handleOperations(genSequence(0.05, 0.90, 0.05, testsize*0.8))
	darry5, treap5 := handleOperations(genSequence(0.05, 0.90, 0.05, testsize))
	return []float64{darry1, darry2, darry3, darry4, darry5}, []float64{treap1, treap2, treap3, treap4, treap5}
}

func genSequence(del, ins, search float64, size int) []*elm.Operation {
	generator := dgen.New()
	operations := make([]*elm.Operation, 0)
	insertions := 0
	deletions := 0
	ser := 0
	for i := 0; i < size; i++ {
		prob := rand.Float64()
		if prob <= ins {
			operations = append(operations, generator.GenInsertion())
			insertions++
		} else if prob <= ins+del {
			operations = append(operations, generator.GenDeletion())
			deletions++
		} else {
			operations = append(operations, generator.GenSearch())
			ser++
		}
	}
	fmt.Printf("Insertions %d, Deletions %d, Searches %d, Total %d, Size %d\n", insertions, deletions, ser, insertions+ser+deletions, size)
	return operations
}

func handleOperations(operations []*elm.Operation) (float64, float64) {
	darryStart := time.Now()
	handleDarry(operations)
	darryFinish := (time.Now().Sub(darryStart)).Seconds() * 1000

	treapStart := time.Now()
	handleTreap(operations)
	treapFinish := (time.Now().Sub(treapStart)).Seconds() * 1000

	return darryFinish, treapFinish
}

func handleDarry(operations []*elm.Operation) {
	dynamicArray := darry.New()
	for _, v := range operations {
		dynamicArray.OperationEval(v)
	}
}

func handleTreap(operations []*elm.Operation) {
	var root *treap.Node
	for _, v := range operations {
		root = treap.OperationEval(root, v)
	}
	fmt.Printf("%d\n", treap.MaxDepth(root))
}

func outputResults(filename string, resultsDarry, resultsTreap []float64) {
	f, _ := os.Create(fmt.Sprintf("%s.csv", filename))
	rows := make([][]string, 0)
	rows = append(rows, []string{"Dynamic Array", "Treap"})
	for i := range resultsDarry {
		rows = append(rows, []string{fmt.Sprintf("%f", resultsDarry[i]), fmt.Sprintf("%f", resultsTreap[i])})
	}
	_ = csv.NewWriter(f).WriteAll(rows)
}

func main() {
	x, y := experimentOne()
	fmt.Printf("Experiment One: darry: %+v treap: %+v\n", x, y)
	outputResults("experiment_one", x, y)
	m, n := experimentTwo()
	fmt.Printf("Experiment Two: darry: %+v treap: %+v\n", m, n)
	outputResults("experiment_two", m, n)
	l, o := experimentThree()
	fmt.Printf("Experiment Three: darry: %+v treap: %+v\n", l, o)
	outputResults("experiment_three", l, o)
	h, k := experimentFour()
	fmt.Printf("Experiment Four: darry: %+v treap: %+v\n", h, k)
	outputResults("experiment_four", h, k)
}
