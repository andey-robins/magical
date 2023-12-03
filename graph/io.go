package graph

// Helper functions for performing IO to create graphs
// examples include loading graphs from files, strings, etc.

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// we'll panic if we hit errors at this point because failing
// to parse the files means something outside the program's
// control has gone wrong (i.e. the wrong file was provided)
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// LoadGraphFromFile will load a graph from the given file path
// and open that file, parsing it into a graph. Panics if it
// can't open the file or decode it properly into a graph.
// This is a convenience function for LoadGraph.
func LoadGraphFromFile(path string) *Graph {
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	return LoadGraph(f)
}

// LoadGraphFromString will load a graph from the given string
// and parse it into a graph. Panics if it can't decode it.
// This is a convenience function for LoadGraph.
func LoadGraphFromString(graphString string) *Graph {
	return LoadGraph(strings.NewReader(graphString))
}

// LoadGraph will load a graph from the given io.Reader
// and parse it into a graph. Panics if it can't decode it.
func LoadGraph(encoding io.Reader) *Graph {
	nodes := make([]*Node, 0)

	scanner := bufio.NewScanner(encoding)
	scanner.Split(bufio.ScanLines)
	log.Println("Setup scanner for reading graph from file")

	// skip over input numbering line because we don't need it (they're listed in order)
	scanner.Scan()
	scanner.Scan()
	inputLabelsLine := scanner.Text()
	for _, label := range strings.Split(inputLabelsLine, " ") {
		id, err := strconv.Atoi(label)
		check(err)
		nodes = append(nodes, NewNode(id))
	}

	log.Println("Finished reading input labels")

	// skip over output numbering line for same reason as input numbering
	scanner.Scan()
	outputCountLine := scanner.Text()
	outputCount, err := strconv.Atoi(strings.Split(outputCountLine, " ")[1])
	check(err)
	scanner.Scan()
	outputLabelsLine := scanner.Text()
	for _, label := range strings.Split(outputLabelsLine, " ") {
		id, err := strconv.Atoi(label)
		check(err)
		nodes = append(nodes, NewNode(id))
	}

	log.Println("Finished reading output labels")

	scanner.Scan()
	nodesLine := scanner.Text()
	nodeCount, err := strconv.Atoi(strings.Split(nodesLine, " ")[1])
	check(err)
	log.Printf("Found %d intermediate nodes\n", nodeCount)
	totalNodes := len(nodes) + nodeCount - outputCount
	for i := len(nodes) + 1; i <= totalNodes; i++ {
		nodes = append(nodes, NewNode(i))
	}

	log.Println("Finished initializing all nodes")

	// we store the edge count just to make writing easier than having to traverse
	// and count them (if they're given, why spend the time when we can store an int)
	scanner.Scan()
	edgeCountLine := scanner.Text()
	edgeCount, err := strconv.Atoi(strings.Split(edgeCountLine, " ")[1])
	check(err)

	g := NewGraph(nodes, edgeCount)

	// build the edges into the graph
	for scanner.Scan() {
		edgeLine := scanner.Text()
		edgePair := strings.Split(edgeLine, " ")

		src, err := strconv.Atoi(edgePair[0])
		check(err)

		dest, err := strconv.Atoi(edgePair[1])
		check(err)

		srcNode, err := g.GetNodeById(src)
		check(err)

		destNode, err := g.GetNodeById(dest)
		check(err)

		srcNode.AddChild(destNode)
		destNode.AddParent(srcNode)
	}

	return g
}
