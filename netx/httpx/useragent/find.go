package useragent

type Filter func(ua *UA) bool

func FindUA(filters ...Filter) []*UA {
	var r []*UA
	for _, ua := range All {
		ok := true
		for _, filter := range filters {
			if filter != nil && !filter(ua) {
				ok = false
			}
		}
		if ok {
			r = append(r, ua)
		}
	}
	return r
}

func Find(filters ...Filter) []string {
	l := FindUA(filters...)
	n := len(l)
	r := make([]string, n, n)
	for i := 0; i < n; i++ {
		r[i] = l[i].UA
	}
	return r
}
