package kata

import "sort"

func QueueTime(customers []int, n int) int {
	result := 0
	if n == 1 {
		for _, v := range customers {
			result += v
		}
		return result
	}

	queues := make([]int, n)
	for _, v := range customers {
		queues[0] += v
		sort.Ints(queues)	
	}
	result = queues[n-1]
	return result
}
