package gtools

type Gslice struct {
	S1 []string
	S2 []string
}

func (g *Gslice) Delete(s string) []string {
	var ret []string
	for _, s1 := range g.S1 {
		if s1 != s {
			ret = append(ret, s1)
		}
	}
	return ret
}

func (g *Gslice) DeleteSlice() []string {
	m := g.ToGMap()
	for _, s := range g.S2 {
		if m.ContainsKey(s) {
			delete(m.M, s)
		}
	}
	var ret []string
	for s := range m.M {
		ret = append(ret, s)
	}
	return ret
}

func (g *Gslice) ToMap() map[string]interface{} {
	ret := make(map[string]interface{})
	for _, s := range g.S1 {
		ret[s] = true
	}
	return ret
}

func (g *Gslice) ToGMap() *Gmap {
	gm := &Gmap{M: make(map[string]interface{})}
	for _, s := range g.S1 {
		gm.M[s] = true
	}
	return gm
}
