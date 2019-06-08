package main

func mergesort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	middle := len(nums) / 2
	l := mergesort(nums[0:middle])
	r := mergesort(nums[middle:])

	return merge(l, r)
}

func merge(l []int, r []int) []int {
	result := []int{}

	for len(l) != 0 && len(r) != 0 {
		if l[0] < r[0] {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
	}

	for _, v := range l {
		result = append(result, v)
	}

	for _, v := range r {
		result = append(result, v)
	}

	return result
}

// Other implementations:
// - https://austingwalters.com/merge-sort-in-go-golang/
