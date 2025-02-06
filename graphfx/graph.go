package graphfx

import (
	"log"
	"os"

	"go.uber.org/fx"
)

type GraphPlotter struct {
	dotGraph string
}

func NewGraphPlotter(dotGraph fx.DotGraph) *GraphPlotter {
	return &GraphPlotter{dotGraph: string(dotGraph)}
}

func (g *GraphPlotter) Plot() {
	// Write the dotGraph to a file
	file, err := os.Create("dotGraph.dot")
	if err != nil {
		log.Printf("Failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(g.dotGraph)
	if err != nil {
		log.Printf("Failed to write dotGraph to file: %v", err)
	}

}
