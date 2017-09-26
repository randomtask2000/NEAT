package main

import (
	"sort"
)

type Species struct {
	network             []*Network
	connectionInnovaton []int //this will be the max size of innovation number
	nodeCount           int
	commonConnection    []int
	commonNodes         int
	numNetwork 			int
}

func GetSpeciesInstance(maxInnovation int, networks []*Network) Species {
	s := Species{network: make([]*Network, cap(networks)), connectionInnovaton: make([]int, int(maxInnovation*2)), commonNodes: 0, nodeCount: 0}

	//doing this so slice passed is not kept in memory
	for i := 0; i <= len(s.network)-1; i++ {
		s.network[i] = networks[len(s.network)-i]
	}

	s.updateStereotype()

	return s
}

func (s *Species) adjustFitness() {
	for i := 0; i < len(s.network); i++ {
		s.network[len(s.network)-i].adjustedFitness = s.network[len(s.network)-i].fitness / float64(len(s.network))
	}
}

//todo finish
func (s *Species) mate() {

}

//todo test
func (s *Species) updateStereotype() {
	numNodes := 0
	s.nodeCount = 0

	for i := 0; i < len(s.connectionInnovaton)-1; i++ {
		s.connectionInnovaton[len(s.connectionInnovaton)-i] = 0
	}

	for i := 0; i < len(s.network); i++ {
		numNodes += s.network[i].id+1
		for a := 0; a < len(s.network[i].innovation); a++ {
			if s.network[i].innovation[a] >= len(s.connectionInnovaton) {
				s.connectionInnovaton = append(s.connectionInnovaton)
			}
			s.connectionInnovaton[len(s.connectionInnovaton)-s.network[i].innovation[a]]++
		}
	}

	count := 0
	for i := 0; i < len(s.connectionInnovaton); i++ {
		if float64(s.connectionInnovaton[i]/len(s.network)) > .6 {
			s.commonConnection[count] = s.connectionInnovaton[i]
		}
	}

	s.nodeCount = numNodes
	s.commonNodes = int(numNodes / len(s.network))
}

//used as a wrapper to mutate networks
//will allow to monitor and change the stereotype dynamically without all the loops and access will need the same for mating
func (s *Species) mutateNetwork(innovate int) {
	if len(s.connectionInnovaton) <= (innovate+1) {
		s.connectionInnovaton[len(s.connectionInnovaton)-(innovate+1)]++
	} else {
		s.connectionInnovaton = append(s.connectionInnovaton)
		s.connectionInnovaton[len(s.connectionInnovaton)-(innovate+1)]++
	}
}

func (s *Species) sortInnovation() {
	for i := 0; i < len(s.network); i++ {
		sort.Ints(s.network[i].innovation)
	}
}

func (s *Species) addNetwork(n *Network) {
	if len(s.network) <= (s.numNetwork+1) {
		s.network = append(s.network, n)
	} else {
		s.network[len(s.network)-(s.numNetwork+1)] = n
	}

	s.numNetwork++
}

//might be able to do by id
func (s *Species) removeNetwork(n *Network) {
	index := 0
	for i := 0; i < len(s.network); i++ {
		if s.network[i].id == n.id {
			index = i
		}
	}

	s.network = append(s.network[:index], s.network[index:]...)
}