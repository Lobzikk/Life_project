package life

type Cell struct {
	State bool
	PosX  int
	PosY  int
	World *World
}

type World struct {
	Map    [][]Cell
	Height int
	Width  int
}
