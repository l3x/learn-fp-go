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

type Fbs func(bool) string
type Fss func(string) string

func Compose(g Fss, f Fbs) Fbs {
	return func(x bool) string {
		return g(f(x))
	}
}

var Emphasize_Humanize = Compose(Emphasize, Humanize)
