package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(p func(int) bool) Ints {
	result := Ints{}
	for _, v := range i {
		if p(v) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (i Ints) Discard(p func(int) bool) Ints {
	result := Ints{}
	for _, v := range i {
		if !p(v) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (l Lists) Keep(p func([]int) bool) Lists {
	result := Lists{}
	for _, v := range l {
		if p(v) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func (s Strings) Keep(p func(string) bool) Strings {
	result := Strings{}
	for _, v := range s {
		if p(v) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
