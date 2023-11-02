package main

import (
	"encoding/json"
	"fmt"
)

type Block struct {
	Gym    bool
	School bool
	Store  bool
}

func calculateDistance(blocks []Block, reqs [][]int) (distance int) {
	minDistance := 1000000
	mapMinDistance := make(map[int]int)
	for idx, _ := range blocks {
		currentDistance := 0
		if blocks[idx].Gym {
			if reqs[idx][0] > 1 {
				currentDistance += reqs[idx][0]
			} else {
				currentDistance += 1000 // get max allocate
			}
		}
		if blocks[idx].School {
			if reqs[idx][1] > 1 {
				currentDistance += reqs[idx][1]
			} else {
				currentDistance += 1000 // get max allocate
			}
		}
		if blocks[idx].Store {
			if reqs[idx][2] > 1 {
				currentDistance += reqs[idx][2]
			} else {
				currentDistance += 1000 // get max allocate
			}
		}
		mapMinDistance[currentDistance] = idx
	}

	for key, _ := range mapMinDistance {
		if key < minDistance {
			minDistance = key
		}
	}

	return mapMinDistance[minDistance]
}

func main() {
	jsonBlocks := `[
    { 
        "gym":false,
        "school":true,
        "store":false
    },
     { 
        "gym":true,
        "school":false,
        "store":false
    },
     { 
        "gym":true,
        "school":true,
        "store":false
    },
     { 
        "gym":false,
        "school":true,
        "store":true
    },
    { 
        "gym":false,
        "school":true,
        "store":true
    }
]`

	reqs := [][]int{
		{1, 0, 4},
		{0, 1, 3},
		{0, 0, 2},
		{1, 0, 1},
		{2, 0, 0},
	}

	var blocks []Block
	err := json.Unmarshal([]byte(jsonBlocks), &blocks)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Block with the most minimum distance to all facilities: Block %d\n", calculateDistance(blocks, reqs))
}
