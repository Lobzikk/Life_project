package life

import (
	"log"
	"math/rand"
	"slices"
	"strconv"
	"time"
)

// Getting all the cell's neighbours to decide its fate later.
func (c Cell) GetNeighbours() []Cell {
	Map, height, width := c.World.Map, c.World.Height, c.World.Width
	res := make([]Cell, 0)
	AddY := func(k int) {
		if k != 0 {
			res = append(res, Map[c.PosY][c.PosX+k])
		}
		if c.PosY > 0 {
			res = append(res, Map[c.PosY-1][c.PosX+k])
		}
		if c.PosY < height-1 {
			res = append(res, Map[c.PosY+1][c.PosX+k])
		}
	}
	if c.PosX > 0 {
		AddY(-1)
	}
	AddY(0)
	if c.PosX < width-1 {
		AddY(1)
	}
	return res
}

// Checking the cell's fate - will it be alive in the next generation or not.
func FateCheck(c Cell) bool {
	var AliveNum int
	for _, cell := range c.GetNeighbours() {
		if cell.State {
			AliveNum++
		}
	}
	if AliveNum >= 3 && !c.State || (AliveNum == 2 || AliveNum == 3) && c.State {
		return true
	}
	return false
}

// Kind of the World's structure objects' constructor. It is recommended to use only this function to avoid any errors.
func WorldGen(height, width int) World {
	newWorld := World{Width: width, Height: height}
	res := make([][]Cell, height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			res[i] = append(res[i], Cell{PosX: j, PosY: i, State: false, World: &newWorld})
		}
	}
	newWorld.Map = res
	return newWorld
}

// Returns randomly generated 20-digit seed needed to fill the world with alive seeds.
func SeedGen() string {
	var res string
	for i := 0; i < 20; i++ {
		res += strconv.Itoa(rand.Intn(5)) //0-4
	}
	return res
}

// Uses given seed to fill the world with life
func (w *World) FillWorld(seed string) {
	var steps int //how much steps left to making the cell alive
	var err error
	origSeed := seed
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			if seed == "" {
				seed = origSeed //"refilling" the seed, so we can use it multiple times
			}
			if steps == 0 {
				steps, err = strconv.Atoi(string(seed[len(seed)-1]))
				if err != nil { //checking if the seed is valid
					log.Fatal(err)
				}
				seed = seed[:len(seed)-1]
				w.Map[i][j].State = true
			} else {
				steps--
			}
		}
	}
}

// "Updating" the world to the next generation by deciding each of the cells' individual fate.
func (w *World) NextGen() {
	NextGenMap := WorldGen(w.Height, w.Width).Map
	for i := 0; i < w.Height; i++ {
		for j := 0; j < w.Width; j++ {
			NextGenMap[i][j].State = FateCheck(w.Map[i][j])
		}
	}
	w.Map = slices.Clone(NextGenMap)
	NextGenMap = nil // deleting the supporting matrix to avoid memory leaks
}

// Change the world's state every n ms/s/min/h (run as goroutine)
func (w *World) TickingWorldGen(duration time.Duration) {
	ticker := time.Tick(duration)
	for range ticker {
		w.NextGen()
	}
}
