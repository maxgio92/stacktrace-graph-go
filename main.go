package main

import (
	"fmt"
	"github.com/maxgio92/stacktrace-graph-go/internal/graph"
	"github.com/maxgio92/stacktrace-graph-go/internal/trace"
)

var (
	trace1 = trace.StackTrace{[]string{"main", "foo", "qux", "grault"}, 4}
	trace5 = trace.StackTrace{[]string{"main", "foo", "quux"}, 2}
	trace2 = trace.StackTrace{[]string{"main", "bar", "quux"}, 2}
	trace3 = trace.StackTrace{[]string{"main", "foo"}, 3}
	trace4 = trace.StackTrace{[]string{"main", "foo", "corge", "garply"}, 5}
	traces = []trace.StackTrace{trace1, trace2, trace3, trace4, trace5}
)

func main() {
	var sampleCountTotal int
	for _, v := range traces {
		sampleCountTotal += v.Samples
	}

	g := graph.NewGraph()
	for kt, _ := range traces {
		for ks, sym := range traces[kt].Syms {
			var parent string
			if ks > 0 {
				parent = traces[kt].Syms[ks-1]
			}

			// If it's the traced function, that is, the last symbol/IP in the stack trace,
			// update also its weight.
			if ks == len(traces[kt].Syms)-1 {
				g.UpsertNode(sym, parent, float32(traces[kt].Samples)/float32(sampleCountTotal))
			} else {
				g.UpsertNode(sym, parent)
			}
		}
	}
	fmt.Println(g)
}
