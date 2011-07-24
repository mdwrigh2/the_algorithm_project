package graph

type List map[string]bool

func (l List) Has(name string) bool {
	_, ok := l[name]
	return ok
}

func (l List) HasList(list List) bool {
	for name, _ := range list {
		if l.Has(name) {
			return true
		}
	}
	return false
}

func (l List) Add(name string) {
	l[name] = true
}

func (l List) AddList(list List) {
	for name, _ := range list {
		l[name] = true
	}
}
