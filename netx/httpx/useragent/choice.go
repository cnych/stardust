package useragent

import "github.com/cnych/stardust/mathx/randx"

func ChoiceUA(filters ...Filter) *UA {
	l := FindUA(filters...)
	if len(l) == 0 {
		return nil
	}
	return randx.ChoiceSlice(l).(*UA)
}

func Choice(filters ...Filter) string {
	ua := ChoiceUA(filters...)
	if ua == nil {
		return ""
	}
	return ua.UA
}
