package structers

type AlGraph struct {
	Graph map[int]map[int]int
}

type AMGraph struct {
	Graph [][]int
}

type Graph interface {
	AMGraph
	AlGraph
}

func (g AMGraph) AMtoAL() AlGraph {
	gal := AlGraph{}
	for i, e := range g.Graph {
		for j, q := range e {
			if q != 0 {
				gal.Graph[i][j] = q
			}
		}
	}
	return gal
}

/*
func (g AlGraph) BFS() map[int]int {

}

func (g AlGraph) DFS() bool {

}*/
