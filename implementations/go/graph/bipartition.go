package graph

func Bipartition(d Digraph) (a, b List) {
	a = List(make(map[string]bool))
	b = List(make(map[string]bool))
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
