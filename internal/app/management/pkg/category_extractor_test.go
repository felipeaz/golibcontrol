package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractCategoryId(t *testing.T) {
	categoryString := "0, 3, -5,4,9,     15, 13, -7, A, bc  , testing  , 20,  , ,, 25"
	expectedSlice := []uint{3, 4, 9, 15, 13, 20, 25}
	returnedSlice := ExtractCategoryId(categoryString)

	assert.NotNil(t, returnedSlice)
	assert.Equal(t, expectedSlice, returnedSlice)
}

func TestExtractCategoryWithEmptyString(t *testing.T) {
	returnedSlice := ExtractCategoryId("")
	assert.Nil(t, returnedSlice)
}
