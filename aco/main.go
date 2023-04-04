package main

import (
	"fmt"
)

type City struct {
	Name string // name of the city
}

type Distances struct {
	DistancesTable [][]float64  // 2D array of distances
	Cities [] City // array of cities
}

func initData() Distances {
	// create a few City objects
	sanAntonio := City{"San Antonio"};
	newYork := City{"New York"};
	boston := City{"Boston"};
	washington := City{"Washington"};

	// create a slice called cities with the above City objects
	cities := []City{sanAntonio, newYork, boston, washington};

	// create Distances object with the above cities slice and an empty 4x4 2D array of distances

	distances := Distances{Cities: cities, DistancesTable: [][]float64{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}};

	distances.DistancesTable[0][0] = 0;
	distances.DistancesTable[0][1] = 2000;
	distances.DistancesTable[0][2] = 1700;
	distances.DistancesTable[0][3] = 800;
	distances.DistancesTable[1][0] = 2000;
	distances.DistancesTable[1][1] = 0;
	distances.DistancesTable[1][2] = 800;
	distances.DistancesTable[1][3] = 1500;
	distances.DistancesTable[2][0] = 1700;
	distances.DistancesTable[2][1] = 800;
	distances.DistancesTable[2][2] = 0;
	distances.DistancesTable[2][3] = 1200;
	distances.DistancesTable[3][0] = 800;
	distances.DistancesTable[3][1] = 1500;
	distances.DistancesTable[3][2] = 1200;
	distances.DistancesTable[3][3] = 0;

	return distances;
}

func main () {
	fmt.Println("Hello, World!");
	var distances = initData();
	fmt.Println(distances);

}