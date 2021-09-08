package foldright

func FoldRight(numsOfNums ...[]int) (totals []int) {
	for _, nums := range numsOfNums {
		total := 0
		for _, n := range nums {
			total += n
		}
		totals = append(totals, total)
	}
	return totals
}
