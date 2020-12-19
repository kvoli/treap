package darry

// Austen McClernon
// 834063
// Assignment 1 COMP90077

import (
	"reflect"
	"testing"

	"github.com/kvoli/COMP90077_ASS/pkg/elm"
)

func TestDynamicArray_resize(t *testing.T) {
	type fields struct {
		Array    []elm.Element
		Size     int
		Capacity int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DynamicArray{
				Array:    tt.fields.Array,
				Size:     tt.fields.Size,
				Capacity: tt.fields.Capacity,
			}
			d.resize()
		})
	}
}

func TestDynamicArray_Search(t *testing.T) {
	type fields struct {
		Array    []elm.Element
		Size     int
		Capacity int
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *elm.Element
		want1  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DynamicArray{
				Array:    tt.fields.Array,
				Size:     tt.fields.Size,
				Capacity: tt.fields.Capacity,
			}
			got, got1 := d.Search(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DynamicArray.Search() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DynamicArray.Search() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDynamicArray_Delete(t *testing.T) {
	type fields struct {
		Array    []elm.Element
		Size     int
		Capacity int
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DynamicArray{
				Array:    tt.fields.Array,
				Size:     tt.fields.Size,
				Capacity: tt.fields.Capacity,
			}
			d.Delete(tt.args.key)
		})
	}
}

func TestDynamicArray_Insert(t *testing.T) {
	type fields struct {
		Array    []elm.Element
		Size     int
		Capacity int
	}
	type args struct {
		e *elm.Element
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DynamicArray{
				Array:    tt.fields.Array,
				Size:     tt.fields.Size,
				Capacity: tt.fields.Capacity,
			}
			d.Insert(tt.args.e)
		})
	}
}

func TestDynamicArray_desize(t *testing.T) {
	type fields struct {
		Array    []elm.Element
		Size     int
		Capacity int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DynamicArray{
				Array:    tt.fields.Array,
				Size:     tt.fields.Size,
				Capacity: tt.fields.Capacity,
			}
			d.desize()
		})
	}
}

func TestDynamicArray_swap(t *testing.T) {
	type fields struct {
		Array    []elm.Element
		Size     int
		Capacity int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DynamicArray{
				Array:    tt.fields.Array,
				Size:     tt.fields.Size,
				Capacity: tt.fields.Capacity,
			}
			d.swap(tt.args.i, tt.args.j)
		})
	}
}
