package jsonresolve

import (
	"life_project/life"
	"os"
)

// Saves the world's current state to the JSON file in the 'save' directory. You can pass an empty string to this function, that way file will have a default name.
func WorldToFile(name string, world life.World) error {
	if len(name) == 0 {
		name = "quicksave.json"
	}
	data, err := MapToJSON(world)
	if err != nil {
		return err
	}
	err = os.WriteFile("../save/"+name, data, 0222)
	return err
}
