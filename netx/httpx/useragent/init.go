package useragent

var (
	All = []*UA{}
)

func init() {
	for _, raw := range rawUserAgents {
		ua, err := Parse(raw, nil)
		if err != nil {
			continue
		}
		All = append(All, ua)
	}
}
