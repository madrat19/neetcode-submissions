type Pair struct {
	value     string
	timestamp int
}

type TimeMap struct {
	storage map[string][]Pair
}

func Constructor() TimeMap {
	return TimeMap{storage: map[string][]Pair{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.storage[key] = append(this.storage[key], Pair{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if len(this.storage[key]) == 0 {
		return ""
	}
	l, r := 0, len(this.storage[key])-1
	for l < r {
		m := (l + r + 1) / 2
		ts := this.storage[key][m].timestamp
		if ts > timestamp {
			r = m - 1
		} else {
			l = m
		}
	}
	if this.storage[key][l].timestamp <= timestamp {
		return this.storage[key][l].value
	} else {
		return ""
	}
}