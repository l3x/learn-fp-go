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


func Compose(g Fbs, f Fss) Fbs {
	return func(x bool) string {
		return f(g(x))
	}
}

var EmphasizeHumanizeFoG = Compose(Emphasize, Humanize)

