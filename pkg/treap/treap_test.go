package treap

// Austen McClernon
// 834063
// Assignment 1 COMP90077

import (
	"reflect"
	"testing"

	"github.com/kvoli/COMP90077_ASS/pkg/elm"
)

func TestInsert(t *testing.T) {
	type args struct {
		n *Node
		e *elm.Element
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.n, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		n   *Node
		key int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(tt.args.n, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	type args struct {
		n   *Node
		key int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.n, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
