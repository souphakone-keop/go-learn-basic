package functionstuct

type Student struct {
	Firstname string
	Lastname  string
}

func (s Student) FullName() string {
	return s.Firstname + " " + s.Lastname
}
