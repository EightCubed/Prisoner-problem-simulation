package init

import "math/rand/v2"

func InitializeGame(MAX_SIZE int) ([]int, map[int]int) {
	prisonerArr := createPrisonerArray(MAX_SIZE)
	boxMap := createBoxMap(MAX_SIZE)
	return prisonerArr, boxMap
}

func createPrisonerArray(maxSize int) []int {
	if maxSize <= 0 {
		return []int{}
	}
	prisoners := make([]int, maxSize)
	for idx := range prisoners {
		prisoners[idx] = idx + 1
	}
	return prisoners
}

func createBoxMap(maxSize int) map[int]int {
	if maxSize <= 0 {
		return map[int]int{}
	}
	nums := rand.Perm(maxSize)
	for i := range nums {
		nums[i] += 1
	}
	boxes := make(map[int]int)
	for idx, val := range nums {
		boxes[idx+1] = val
	}
	return boxes
}
