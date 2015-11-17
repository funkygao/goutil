package gofmt

import (
	"fmt"
	"time"
)

// Seconds-based time units
const (
	Minute = 60
	Hour   = 60 * Minute
	Day    = 24 * Hour
	Week   = 7 * Day
	Month  = 30 * Day
	Year   = 12 * Month
)

func PrettySince(t time.Time) string {
	sec := int(time.Since(t).Seconds())
	r := ""
	y := sec / Year
	if y > 0 {
		r += fmt.Sprintf("%dy", y)
		sec = sec % Year
	}
	m := sec / Month
	if m > 0 {
		r += fmt.Sprintf("%dm", m)
		sec = sec % Month
	}
	d := sec / Day
	if d > 0 {
		r += fmt.Sprintf("%dd", d)
		sec = sec % Day
	}
	h := sec / Hour
	if h > 0 {
		r += fmt.Sprintf("%dh", h)
		sec = sec % Hour
	}
	mi := sec / Minute
	if mi > 0 {
		r += fmt.Sprintf("%dm", mi)
		sec = sec % Minute
	}
	r += fmt.Sprintf("%ds", sec)
	return r

}

// Humanize the time.
func TimePretty(then time.Time) string {
	now := time.Now()

	lbl := "ago"
	diff := now.Unix() - then.Unix()
	if then.After(now) {
		lbl = "from now"
		diff = then.Unix() - now.Unix()
	}

	switch {

	case diff <= 0:
		return "now"
	case diff <= 2:
		return fmt.Sprintf("1 second %s", lbl)
	case diff < 1*Minute:
		return fmt.Sprintf("%d seconds %s", diff, lbl)

	case diff < 2*Minute:
		return fmt.Sprintf("1 minute %s", lbl)
	case diff < 1*Hour:
		return fmt.Sprintf("%d minutes %s", diff/Minute, lbl)

	case diff < 2*Hour:
		return fmt.Sprintf("1 hour %s", lbl)
	case diff < 1*Day:
		return fmt.Sprintf("%d hours %s", diff/Hour, lbl)

	case diff < 2*Day:
		return fmt.Sprintf("1 day %s", lbl)
	case diff < 1*Week:
		return fmt.Sprintf("%d days %s", diff/Day, lbl)

	case diff < 2*Week:
		return fmt.Sprintf("1 week %s", lbl)
	case diff < 1*Month:
		return fmt.Sprintf("%d weeks %s", diff/Week, lbl)

	case diff < 2*Month:
		return fmt.Sprintf("1 month %s", lbl)
	case diff < 1*Year:
		return fmt.Sprintf("%d months %s", diff/Month, lbl)

	case diff < 18*Month:
		return fmt.Sprintf("1 year %s", lbl)
	}
	return then.String()
}
