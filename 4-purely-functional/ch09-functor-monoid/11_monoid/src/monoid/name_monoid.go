package monoid

type NameMonoid interface {
	Append(s string) NameMonoid
	Zero() string
}

func WrapName(s string) NameMonoid {
	return nameContainer{name: s}
}

type nameContainer struct {
	name string
}

func (s nameContainer) Append(name string) NameMonoid {
	s.name = s.name + name
	return s
}

func (nameContainer) Zero() string {
	return ""
}

func (s nameContainer) String() string {
	return s.name
}
