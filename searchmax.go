package search

type MaxComparator func(i, max int) bool

func Max(arr []interface{}, comparator MaxComparator) interface{} {

	if len(arr) == 0 {
		return nil
	}

	max := 0

	for i := 0; i < len(arr); i++ {
		if comparator(i, max) {
			max = i
		}
	}

	return arr[max]
}
