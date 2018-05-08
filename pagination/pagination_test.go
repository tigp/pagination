package pagination

import (
	"reflect"
	"testing"
)

func TestPagination(t *testing.T) {
	// create test data
	result := []struct {
		pagination pagination
		expected   []string
	}{
		{pagination{currentPage: 4, totalPages: 5, boundaries: 1, around: 0},
			[]string{"1", "...", "4", "5"},
		},
		{pagination{currentPage: 4, totalPages: 10, boundaries: 2, around: 2},
			[]string{"1", "2", "3", "4", "5", "6", "...", "9", "10"},
		},
	}

	// run tests
	for i, r := range result {
		t.Log("Test case: %v", i)

		res, err := Paginate(
			r.pagination.currentPage,
			r.pagination.totalPages,
			r.pagination.boundaries,
			r.pagination.around,
		)

		if err != nil {
			t.Fatalf("in test number: %v, error happened: %v", i, err)
		}

		// check if the result equal to expected
		if !reflect.DeepEqual(r.expected, res) {
			t.Fatalf("in test number: %v, expected: %s, but got: %s", i, r.expected, res)
		}
	}
}
