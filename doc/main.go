//go:build linux || darwin

// Package main is generating mermaid sequence diagram.
package main

import (
	"os"

	"github.com/go-spectest/markdown"
	"github.com/go-spectest/mermaid/sequence"
)

//go:generate go run main.go

func main() {
	f, err := os.Create("generated.md")
	if err != nil {
		panic(err)
	}

	diagram := sequence.NewDiagram(os.Stdout).
		Participant("Viktoriya").
		Participant("Naohiro").
		Participant("Naoyuki").
		LF().
		SyncRequest("Viktoriya", "Naohiro", "Please wake up Naoyuki").
		SyncResponse("Naohiro", "Viktoriya", "OK").
		LF().
		LoopStart("until Naoyuki wake up").
		SyncRequest("Naohiro", "Naoyuki", "Wake up!").
		SyncResponse("Naoyuki", "Naohiro", "zzz").
		SyncRequest("Naohiro", "Naoyuki", "Hey!!!").
		BreakStart("if Naoyuki wake up").
		SyncResponse("Naoyuki", "Naohiro", "......").
		BreakEnd().
		LoopEnd().
		LF().
		SyncResponse("Naohiro", "Viktoriya", "wake up, wake up").
		String() //nolint

	err = markdown.NewMarkdown(f).
		H2("Sequence Diagram").
		CodeBlocks(markdown.SyntaxHighlightMermaid, diagram).
		Build()

	if err != nil {
		panic(err)
	}
}
