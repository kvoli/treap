package elm

// Austen McClernon
// 834063
// Assignment 1 COMP90077

const (
	// SearchOperation represents the search for an elementQ
	SearchOperation = "search_operation"

	// DeleteOperation represents the deletion of an element
	DeleteOperation = "delete_operation"

	// InsertOperation represents the insertion of an element
	InsertOperation = "insert_operation"
)

// Element represent the generated element
type Element struct {
	ID  int
	Key int
}

// Operation represnts an operation on a data structure
type Operation struct {
	Element *Element
	Type    string
}
