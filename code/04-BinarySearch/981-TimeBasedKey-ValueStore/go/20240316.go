package main

// // https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
type TimeMap struct {
	hash map[string][]Pair
}

type Pair struct {
	value     string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{
		hash: make(map[string][]Pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	pair := Pair{
		value:     value,
		timestamp: timestamp,
	}
	this.hash[key] = append(this.hash[key], pair)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	ans := ""
	pairs := this.hash[key]
	l, r := 0, len(pairs)-1
	for l <= r {
		m := (l + r) / 2
		mid := pairs[m]
		v, ts := mid.value, mid.timestamp
		if ts <= timestamp {
			ans = v
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return ans
}
