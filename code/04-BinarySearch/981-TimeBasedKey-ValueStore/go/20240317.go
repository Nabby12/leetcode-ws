package main

// https://leetcode.com/problems/time-based-key-value-store/description/
type TimeMap2 struct {
	hash map[string][]Value
}

type Value struct {
	value     string
	timestamp int
}

func Constructor2() TimeMap2 {
	return TimeMap2{
		hash: make(map[string][]Value),
	}
}

func (this *TimeMap2) Set(key string, value string, timestamp int) {
	this.hash[key] = append(this.hash[key], Value{
		value:     value,
		timestamp: timestamp,
	})
}

func (this *TimeMap2) Get(key string, timestamp int) string {
	values, ok := this.hash[key]
	if !ok {
		return ""
	}

	ans := ""
	l, r := 0, len(values)-1
	for l <= r {
		m := (l + r) / 2
		mid := values[m]
		if mid.timestamp <= timestamp {
			ans = mid.value
			l = m + 1
			if timestamp == mid.timestamp {
				break
			}
		} else {
			r = m - 1
		}
	}
	return ans
}
