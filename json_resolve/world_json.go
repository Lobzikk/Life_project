package jsonresolve

import (
	"encoding/json"
	"life_project/life"
)

// Converts world's map to JSON boolean matrix.
func MapToJSON(w life.World) ([]byte, error) {
	res := make([][]bool, 0)
	for _, row := range w.Map {
		resrow := make([]bool, w.Width)
		for ind, cell := range row {
			resrow[ind] = cell.State
		}
		res = append(res, resrow)
	}
	return json.Marshal(res)
}
