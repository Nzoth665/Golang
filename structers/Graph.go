package structers

// Представление графа в виде списка смежности (Список смежности)(Adjacency List)
type ALGraph map[int]map[int]int

// Табличное представление графа (Матрица смежности)(Adjacency Matrix)
type AMGraph [][]int

type Graph interface {
	AMGraph
	ALGraph
}

func (g AMGraph) AMtoAL() ALGraph {
	gal := ALGraph{}
	for i, e := range g {
		for j, q := range e {
			if q != 0 {
				gal[i][j] = q
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
