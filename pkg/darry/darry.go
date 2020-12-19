package darry

// Austen McClernon
// 834063
// Assignment 1 COMP90077

import (
	"github.com/kvoli/COMP90077_ASS/pkg/elm"
)

// DynamicArray represents a self managed Array
type DynamicArray struct {
	Array    []elm.Element
	Size     int
	Capacity int
}

// New returns a new Dynamic
func New() *DynamicArray {
	return &DynamicArray{
		Array:    make([]elm.Element, 0, 1),
		Size:     0,
		Capacity: 1,
	}
}

// OperationEval takes an operation and performs the corresponding operation
func (d *DynamicArray) OperationEval(o *elm.Operation) {
	switch o.Type {
	case elm.InsertOperation:
		d.Insert(o.Element)
	case elm.DeleteOperation:
		d.Delete(o.Element.Key)
	case elm.SearchOperation:
		d.Search(o.Element.Key)
	}
}

// Insert puts a new element into the Dynamic Array
func (d *DynamicArray) Insert(e *elm.Element) {
	if d.Capacity <= (d.Size) {
		d.resize()
	}
	d.Array = append(d.Array, *e)
	d.Size++
}

// Delete searches for a key in an array, if it exists it deletes it
func (d *DynamicArray) Delete(key int) {
	elm, index := d.Search(key)
	if elm != nil {
		d.swap(index, d.Size-1)
	}

	d.Size--
	d.Array = d.Array[:len(d.Array)-1]

	if d.Size <= d.Capacity/4 {
		d.desize()
	}
}

// Search implements a linear scan of the array for an element with the key
func (d *DynamicArray) Search(key int) (*elm.Element, int) {
	for index, v := range d.Array {
		if v.Key == key {
			return &v, index
		}
	}
	return nil, -1
}

func (d *DynamicArray) desize() {
	newArray := make([]elm.Element, d.Size, d.Capacity/2)
	copy(newArray, d.Array)
	d.Array = newArray
	d.Capacity /= 2
}

// resizes the array (slice) to be double the size
func (d *DynamicArray) resize() {
	newArray := make([]elm.Element, d.Size, d.Capacity*2)
	copy(newArray, d.Array)
	d.Array = newArray
	d.Capacity *= 2
}

func (d *DynamicArray) swap(i, j int) {
	if i > 0 && j > 0 && i <= d.Size && j <= d.Size {
		tmp := d.Array[i]
		d.Array[i] = d.Array[j]
		d.Array[j] = tmp
	}
}
