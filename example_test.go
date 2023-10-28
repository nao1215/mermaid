//go:build linux || darwin

package markdown_test

import (
	"os"

	"github.com/go-spectest/markdown"
	"github.com/go-spectest/mermaid/sequence"
)

// Examle is example code. Skip this test on Windows.
// The newline codes in the comment section where
// the expected values are written are represented as '\n',
// causing failures when testing on Windows.
func Example() {
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
		String()

	markdown.NewMarkdown(os.Stdout).
		H2("Sequence Diagram").
		CodeBlocks(markdown.SyntaxHighlightMermaid, diagram).
		Build() // nolint:errcheck

	// Output:
	// ## Sequence Diagram
	// ```mermaid
	// sequenceDiagram
	//     participant Viktoriya
	//     participant Naohiro
	//     participant Naoyuki
	//
	//     Viktoriya->>Naohiro: Please wake up Naoyuki
	//     Naohiro-->>Viktoriya: OK
	//
	//     loop until Naoyuki wake up
	//     Naohiro->>Naoyuki: Wake up!
	//     Naoyuki-->>Naohiro: zzz
	//     Naohiro->>Naoyuki: Hey!!!
	//     break if Naoyuki wake up
	//     Naoyuki-->>Naohiro: ......
	//     end
	//     end
	//
	//     Naohiro-->>Viktoriya: wake up, wake up
	// ```
}
