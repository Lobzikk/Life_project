// Currently used only for testing isolated features of the project.
package main

import (
	"life_project/json_resolve"
	"life_project/life"
)

func main() {
	world := life.WorldGen(10, 10)
	jsonresolve.WorldToFile("", world)
}
