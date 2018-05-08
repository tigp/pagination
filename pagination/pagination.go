package pagination

import (
	"errors"
	"fmt"
)

// separator is put in places of skipped links
const separator = "..."

var (
	ErrTotalPagesShouldNotBeNil           = errors.New("Total pages shouldn't be nil")
	ErrCurrentPageShouldNotBeNil          = errors.New("Current page shouldn't be nil")
	ErrCurrentPageShouldNotBeGTTotalPages = errors.New("Current page shouldn't be greater than total pages")
)

// pagination represents data passed from the user
type pagination struct {
	currentPage, // actual page
	totalPages, // total available pages
	boundaries, // how many pages we want to link in the beginning and in the end
	around int // how many pages we want to link before and after the actual page
}

// Paginate validates the input and returns an array of links with separator
func Paginate(currentPage, totalPages, boundaries, around int) ([]string, error) {
	p := pagination{currentPage, totalPages, boundaries, around}

	var result []string

	// validate
	err := p.validate()
	if err != nil {
		return result, err
	}

	// paginate
	result = p.paginate()

	return result, nil
}

// validate validates data passed by client
func (p pagination) validate() error {
	// validate total pages
	if p.totalPages == 0 {
		return ErrTotalPagesShouldNotBeNil
	}

	// validate current page
	if p.currentPage == 0 {
		return ErrCurrentPageShouldNotBeNil
	}
	if p.currentPage > p.totalPages {
		return ErrCurrentPageShouldNotBeGTTotalPages
	}

	return nil
}

// paginate returns an array of links with separator
func (p pagination) paginate() []string {
	// initialize map
	var pageMap = make(map[int]bool)

	// add current page to map
	pageMap[p.currentPage] = true

	// add boundaries to map
	for i := 0; i < p.boundaries; i++ {
		pageMap[1+i] = true
		pageMap[p.totalPages-i] = true
	}

	// add around on each side of current page
	for i := 1; i <= p.around; i++ {
		pageMap[p.currentPage+i] = true
		pageMap[p.currentPage-i] = true
	}

	var result []string

	// iterate over total pages and append links from the map with separator
	var addSeparator = false
	for i := 1; i <= p.totalPages; i++ {
		switch {
		case pageMap[i]:
			result = append(result, fmt.Sprint(i))
			addSeparator = true
		case !pageMap[i] && addSeparator:
			result = append(result, separator)
			addSeparator = false
		}
	}

	// remove trailing separator
	if result[len(result)-1] == separator {
		result = result[:len(result)-1]
	}

	return result
}
