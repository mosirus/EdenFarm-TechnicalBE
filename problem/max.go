package problem

func MaxNumberArray(arr []int) int {
	max := 0
	j := (len(arr) - 1)

	for i := 0; i < len(arr); i++ {
		if arr[i] == arr[j] && i < j {
			j--
			if max < arr[i] {
				max = arr[i]
			}
		}
	}

	return max
}
