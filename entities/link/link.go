package link

import "sort"

type Link struct {
	Formatted string
	FromFile  string
}

func Unique(list []Link) []Link {
	if len(list) == 0 {
		return nil
	}
	tmp := make([]Link, len(list))
	copy(tmp, list)
	list = tmp
	sort.Slice(list, func(i, j int) bool {
		return list[i].Formatted < list[j].Formatted
	})
	r := make([]Link, 0, len(list))
	uniqueIDX := 0
	for i := 0; i < len(list); i++ {
		if list[i] != list[uniqueIDX] {
			uniqueIDX++
			list[i], list[uniqueIDX] = list[uniqueIDX], list[i]
		}
	}
	for i := 0; i < uniqueIDX+1; i++ {
		r = append(r, list[i])
	}
	return r
}
