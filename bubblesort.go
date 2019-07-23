package main

// The bubblesort algorithm is a fairly
// expensive one, with a complexity of N^2.
// Given a list, it will loop through it
// and do a sub-loop: for each of our elements,
// it will see if the next ones are lower and,
// if so, swap their position.
//
// Given [2, 1, 0] we will start the first loop
// with 2 and then start an inner loop of elements
// after 2 ([1, 0]). The first element after 2 is 1,
// which is lower, so we swap them. The list is now
// [1, 2, 0]. Then we compare our current element again
// (the element at index 0 is now 1), with the next element, 0,
// and swap them again. The list is now [0, 2, 1]:
// the smallest element is at the first position.
// Continuing the loop will make sure that the second smallest
// is at the second position, the third...and so on.
func bubblesort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]

			}
		}
	}

	return nums
}

// Other implementations:
// - https://tutorialedge.net/golang/implementing-bubble-sort-with-golang/
