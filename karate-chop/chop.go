package karate_chop

func Chop(i int, arr []int, start int, end int) int {
	if len(arr) == 0 || start > end {
		return -1
	}

	mid := (start + end) / 2

	if arr[mid] == i {
		return mid
	}
	if arr[mid] < i {
		return Chop(i, arr, mid+1, end)
	}

	return Chop(i, arr, start, mid-1)
}
