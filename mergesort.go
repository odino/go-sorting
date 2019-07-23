package main

// Mergesort is a very fascinating algorithm,
// which typically runs in O(n logn).
// We start with our list and split it into
// sub-lists cut in half (there's your logn).
//
// After we've recursively reduced it to lists of 1 elements
// we start merging the lists together while
// sorting them, so l[0] gets compared to r[0].
// Say that l[0] is lower, it gets pushed to the first
// position. Then we compare l[1] with r[0],
// and so on. Once the comparison is done we return the
// result, and upper layers proceed to merge the combination
// of lists.
//
// Say we start with [3, 1, 4, 2]. We split in [3, 1] and [4, 2],
// then split in [3] [1] [4] [2]. We can't split anymore, so we start
// merging these lists and pushing them up the recursive stack:
// [3] [1] [4] [2] become merge([3], [1]) = [1, 3] and merge([4, 2]) = [2, 4],
// then merge([1, 3], [2, 4]) = [1, 2, 3, 4].
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
