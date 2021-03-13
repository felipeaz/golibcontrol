package pkg

import (
	"regexp"
	"strconv"
)

// ExtractCategoryId get categories ID from request string and return as uint slice.
func ExtractCategoryId(categoriesString string) (categoriesId []uint) {
	if len(categoriesString) <= 0 {
		return
	}

	rgx := regexp.MustCompile(`[,\s+]`)
	tmpSlice := rgx.Split(categoriesString, -1)

	categoriesId = make([]uint, 0)
	for _, val := range tmpSlice {
		tmpId, err := strconv.Atoi(val)
		if err != nil || tmpId < 0 {
			continue
		}

		categoriesId = append(categoriesId, uint(tmpId))
	}

	return
}
