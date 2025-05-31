package parser

import (
	"fmt"
	"strings"
)

func GenerateDOT(node *ASTNode) string {
	var builder strings.Builder
	builder.WriteString("digraph AST {\n")
	builder.WriteString("  node [shape=box, fontname=\"Courier\", fontsize=10];\n")
	builder.WriteString("  edge [fontname=\"Courier\", fontsize=10];\n\n")

	var nodeCounter int
	generateDOTNode(&builder, node, &nodeCounter)

	builder.WriteString("}\n")
	return builder.String()
}

func generateDOTNode(builder *strings.Builder, node *ASTNode, counter *int) int {
	if node == nil {
		return -1
	}

	currentID := *counter
	*counter++

	label := node.Type
	if node.Value != "" {
		label += fmt.Sprintf("\\n%s", node.Value)
	}

	builder.WriteString(fmt.Sprintf("  node%d [label=\"%s\"];\n", currentID, label))

	for _, child := range node.Children {
		childID := generateDOTNode(builder, child, counter)
		if childID >= 0 {
			builder.WriteString(fmt.Sprintf("  node%d -> node%d;\n", currentID, childID))
		}
	}

	return currentID
}
