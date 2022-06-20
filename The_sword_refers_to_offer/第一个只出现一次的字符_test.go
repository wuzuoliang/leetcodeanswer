package The_sword_refers_to_offer

// https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/solution/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-by-3zqv5/
func firstUniqChar(s string) byte {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	for i, ch := range s {
		ch -= 'a'
		if pos[ch] == n {
			pos[ch] = i
		} else {
			pos[ch] = n + 1
		}
	}
	ans := n
	for _, p := range pos[:] {
		if p < ans {
			ans = p
		}
	}
	if ans < n {
		return s[ans]
	}
	return ' '
}

type pair struct {
	ch  byte
	pos int
}

func firstUniqChar2(s string) byte {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	q := []pair{}
	for i := range s {
		ch := s[i] - 'a'
		if pos[ch] == n {
			pos[ch] = i
			q = append(q, pair{ch, i})
		} else {
			pos[ch] = n + 1
			for len(q) > 0 && pos[q[0].ch] == n+1 {
				q = q[1:]
			}
		}
	}
	if len(q) > 0 {
		return s[q[0].pos]
	}
	return ' '
}
