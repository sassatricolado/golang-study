package main

func Sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

func SumAllOfRest(numsToSum ...[]int) (sums []int) {
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			final := nums[1:]
			sums = append(sums, Sum(final))
		}
	}
	return
}
