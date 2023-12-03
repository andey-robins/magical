package memory

import (
	"testing"
)

func TestMemory(t *testing.T) {
	mem := NewMemory()

	// fill in two initial cells
	mem.ProcessNode(1, 1, []int{})
	mem.ProcessNode(2, 1, []int{})

	// process a node with one of the initials as a parent
	mem.ProcessNode(3, 1, []int{1})

	// confirm cell 1 is no longer valid
	if mem.cells[0].valid {
		t.Errorf("Cell 1 should be invalid")
	}

	// confirm cell 2 is still valid
	if !mem.cells[1].valid {
		t.Errorf("Cell 2 should be valid")
	}

	// process a node with two parents
	mem.ProcessNode(4, 1, []int{2, 3})

	// confirm cell 2 is no longer valid
	if mem.cells[1].valid || mem.cells[2].valid {
		t.Errorf("Cells 2 and 3 should be invalid")
	}
}
