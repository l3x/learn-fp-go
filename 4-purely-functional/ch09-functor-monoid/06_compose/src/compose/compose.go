package compose

func Humanize(b bool) string {
	if b { return "yes" } else { return "no" }
}

func Emphasize(s string) string {
	return s + "!!"
}

func EmphasizeHumanize(b bool) string {
	return Emphasize(Humanize(b))
}
