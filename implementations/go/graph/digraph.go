package graph

import "fmt"

type Digraph map[string]List

func (d Digraph) String() (s string) {
	s += fmt.Sprintf("%d", len(d))
	for name, nameList := range d {
		s += fmt.Sprintf("\n%s %d", name, len(nameList))
		for name, _ := range nameList {
			s += "\n" + name
		}
	}
	return s
}

func (d Digraph) Add(name string, list List) {
	d[name] = list
}
func (d Digraph) Del(name string) {
	d[name] = nil, false
}
