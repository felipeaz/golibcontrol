package pkg

import (
	"regexp"
	"strconv"
)

// ExtractCategoryId get categories ID from request string and return as uint slice.
func ExtractCategoryId(categoriesString string) (categoriesId []uint) {
	if len(categoriesString) <= 0 {
		return nil
	}

	rgx := regexp.MustCompile(`[,\s+]`)
	tmpSlice := rgx.Split(categoriesString, -1)

	categoriesId = make([]uint, 0)
	for i := 0; i < len(tmpSlice); i++ {
		tmpId, err := strconv.Atoi(tmpSlice[i])
		if err != nil || tmpId <= 0 {
			continue
		}

		categoriesId = append(categoriesId, uint(tmpId))
	}

	return
}
