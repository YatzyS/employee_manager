package employee

import (
	"fmt"
	"github.com/employee_manager/internal/common/constants"
	"github.com/employee_manager/internal/dao"
	"sort"
)

func getStarAndEndIdx(page int, pageSize int, maxLen int) (int, int, error) {
	start := (page - 1) * pageSize
	end := page * pageSize
	if start >= maxLen {
		return 0, 0, fmt.Errorf(constants.NotEnoughDataError)
	}
	if end > maxLen {
		end = maxLen
	}
	return start, end, nil
}

func getKeys(inMemoryStore map[int]*dao.Employee) []int {
	keys := make([]int, 0)
	for key := range inMemoryStore {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}
