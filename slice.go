package ku

// 去除切片中的空元素
func RemoveEmptyItems[T []any | string | map[any]any](slice []T) []T {
	ret := make([]T, 0, len(slice))
	for _, item := range slice {
		if len(item) != 0 {
			ret = append(ret, item)
		}
	}
	return ret
}

// 切片去重，保留顺序
func Unique[T comparable](slice []T) []T {
	ret := make([]T, 0, len(slice))
	hold := make(map[T]struct{}, len(slice))
	for _, item := range slice {
		if _, ok := hold[item]; ok {
			continue
		}
		hold[item] = struct{}{}
		ret = append(ret, item)
	}
	return ret
}

// 求笛卡尔积
//
// [[a,b],[c],[d,e]]->[[a,c,d],[a,c,e],[b,c,d],[b,c,e]]
func Product[T any](sli [][]T) [][]T {
	if len(sli) == 0 {
		return sli
	}
	ret := make([][]T, 0, len(sli[0]))
	for i := 0; i < len(sli[0]); i++ {
		ret = append(ret, []T{sli[0][i]})
	}
	for i := 1; i < len(sli); i++ {
		ret = product(ret, sli[i])
	}
	return ret
}

// [[a],[b]],[c]->[[a,c],[b,c]]
func product[T any](sli [][]T, curr []T) [][]T {
	if len(curr) == 0 {
		return sli
	}
	ret := make([][]T, 0, len(sli)*len(curr))
	for i := 0; i < len(sli); i++ {
		for j := 0; j < len(curr); j++ {
			tmp := make([]T, len(sli[i]))
			copy(tmp, sli[i])
			ret = append(ret, append(tmp, curr[j]))
		}
	}
	return ret
}
