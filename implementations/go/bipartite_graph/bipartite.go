// Yup, its just a bipartite graph
// http://www.facebook.com/careers/puzzles.php?puzzle_id=20
// Alex Ray (2011) <ajray@ncsu.edu>

// makes a basic digraph where each node is a unique string (name)
// then divides it up into two parts

package bipartite

import "fmt"

type List map[string]bool
type Digraph map[string] List

func (d Digraph) String() (s string) {
	s += fmt.Sprintf("%d",len(d))
	for name, nameList := range d {
		s += fmt.Sprintf("\n%s %d",name,len(nameList))
		for name, _ := range nameList {
			s += "\n" + name
		}
	}
	return s
}
func (d Digraph) Del(name string) {
	d[name] = nil, false
}
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

func Bipartition(d Digraph) (a, b List) {
	//fmt.Println(d)
	a = List(make(map[string] bool))
	b = List(make(map[string] bool))
	for len(d) > 0 {
		// loop until d is exhausted
		for name, nameList := range d {
			if a.Has(name) { // name in A
				b.AddList(nameList)
				d.Del(name)
			} else if b.Has(name) { // name in B
				a.AddList(nameList)
				d.Del(name)
			} else if a.HasList(nameList) { // accused in A
				b.Add(name)
				a.AddList(nameList)
				d.Del(name)
			} else if b.HasList(nameList) { // accused in B
				a.Add(name)
				b.AddList(nameList)
				d.Del(name)
			} else if len(a) == 0 && len(b) == 0 { // init
				a.Add(name)
				b.AddList(nameList)
				d.Del(name)
			}
		}
	}
	return a, b
}
