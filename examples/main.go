package main

import (
	"fmt"

	"github.com/scigno/gizmo"
)

func main() {

	graph := gizmo.Graph()
	g := graph.Traversal("g")
	compundQuery := g.New()
	getUser := g.New().Raw("t=").Append(g.New("g").V().Has("user", "username", "scigno"))
	newQuery := g.New().TernaryOp(
		g.New("t").HasNext(),
		g.New("t").Next(),
		g.New("g").AddV("user").Property("userId", "744be509-a1cc-466d-bb10-0bb9a376da2e").Property("username", "scigno").Next(),
	)
	compundQuery.Append(getUser).AddLine(newQuery)
	fmt.Println(compundQuery)

}
