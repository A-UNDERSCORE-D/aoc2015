package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"awesome-dragon.science/go/adventofcode2015/util"
)

type graphNode struct {
	name             string
	neighboursString []struct {
		distance int
		name     string
	}
	neighbours map[*graphNode]int
	edges      []*edge
}

func (g *graphNode) hasEdgeTo(other *graphNode) bool {
	for _, e := range g.edges {
		if e.left == other || e.right == other {
			return true
		}
	}
	return false
}

func (g *graphNode) edgesByDistance() []*edge {
	edges := make([]*edge, len(g.edges))
	copy(edges, g.edges)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})
	return edges
}

type edge struct {
	distance int
	left     *graphNode
	right    *graphNode
}

func (e *edge) String() string {
	return fmt.Sprint(e.left.name, " - ", e.distance, " > ", e.right.name)
}

func main() {
	input := util.ReadLines("input.txt")
	startTime := time.Now()
	res := part1(input)
	fmt.Println("Part 1:", res, "Took:", time.Since(startTime))
	startTime = time.Now()
	res = part2(input)
	fmt.Println("Part 2:", res, "Took:", time.Since(startTime))
}

func parseGraph(input []string) map[string]*graphNode {
	out := map[string]*graphNode{}
	for _, line := range input {
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		start := split[0]
		end := split[2]
		distance, _ := strconv.Atoi(split[4])

		startNode, exists := out[start]
		if !exists {
			startNode = &graphNode{name: start}
			out[start] = startNode
		}

		endNode, exists := out[end]
		if !exists {
			endNode = &graphNode{name: end}

			out[end] = endNode
		}

		edg := &edge{left: startNode, right: endNode, distance: distance}
		if !startNode.hasEdgeTo(endNode) {
			startNode.edges = append(startNode.edges, edg)
		}

		if !endNode.hasEdgeTo(startNode) {
			endNode.edges = append(endNode.edges, edg)
		}
	}

	return out
}

func part1(input []string) string {
	// 	input = strings.Split(
	// 		`London to Dublin = 464
	// London to Belfast = 518
	// Dublin to Belfast = 141`, "\n")
	graph := parseGraph(input)
	dst, _ := bestNearestNeighbour(graph)
	return fmt.Sprint(dst)
}

func bestNearestNeighbour(graph map[string]*graphNode) (int, []*edge) {
	bestDst := math.MaxInt64
	bestRoute := ([]*edge)(nil)
	keys := make([]string, 0, len(graph))
	for k := range graph {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		route := []*edge{}
		visited := []string{k}
		current := k
		currentDst := 0
		for len(visited) != len(keys) {
			neighboursByDst := graph[current].edgesByDistance()
			for _, e := range neighboursByDst {
				if util.StringSliceContains(visited, e.left.name) && util.StringSliceContains(visited, e.right.name) {
					// we've visited both sides, skip!
					continue
				}

				route = append(route, e)
				if e.left.name == current {
					current = e.right.name
				} else {
					current = e.left.name
				}
				currentDst += e.distance
				visited = append(visited, current)
				break
			}
		}

		if currentDst < bestDst {
			bestDst = currentDst
			bestRoute = route
		}

	}

	return bestDst, bestRoute
}

func bestFurthestNeighbour(graph map[string]*graphNode) (int, []*edge) {
	bestDst := math.MinInt64
	bestRoute := ([]*edge)(nil)
	keys := make([]string, 0, len(graph))
	for k := range graph {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		route := []*edge{}
		visited := []string{k}
		current := k
		currentDst := 0
		for len(visited) != len(keys) {
			neighboursByDst := graph[current].edgesByDistance()
			sort.Slice(neighboursByDst, func(i, j int) bool {
				return neighboursByDst[i].distance > neighboursByDst[j].distance
			})
			for _, e := range neighboursByDst {
				if util.StringSliceContains(visited, e.left.name) && util.StringSliceContains(visited, e.right.name) {
					// we've visited both sides, skip!
					continue
				}

				route = append(route, e)
				if e.left.name == current {
					current = e.right.name
				} else {
					current = e.left.name
				}
				currentDst += e.distance
				visited = append(visited, current)
				break
			}
		}

		if currentDst > bestDst {
			bestDst = currentDst
			bestRoute = route
		}

	}

	return bestDst, bestRoute
}

func part2(input []string) string {
	graph := parseGraph(input)
	dst, _ := bestFurthestNeighbour(graph)
	return fmt.Sprint(dst)
}
