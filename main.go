package main

import (
	"encoding/hex"
	"fmt"
	"github.com/decred/dcrwallet/walletseed"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Ensure seed hex was specified.
	if len(os.Args) < 2 {
		appName := filepath.Base(os.Args[0])
		appName = strings.TrimSuffix(appName, filepath.Ext(appName))
		fmt.Fprintf(os.Stderr, "Seed hex argument not specified\n\n")
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintf(os.Stderr, "  %s <seedhex>\n", appName)
		os.Exit(1)
	}

	// Ensure the hex is the appropriate length for a 256-bit seed.
	seedStr := strings.TrimSpace(os.Args[1])
	if len(seedStr) != 64 {
		fmt.Println("The specified seed must have a length of 64 characters")
		os.Exit(1)
	}

	// Decode hex to bytes.
	seedBytes, err := hex.DecodeString(seedStr)
	if err != nil {
		fmt.Println("Specified seed hex is not valid.  It must only " +
			"contain numbers and letters in the range [a-f]")
		os.Exit(1)
	}

	// Convert bytes to seed words with checksum.
	fmt.Println(walletseed.EncodeMnemonic(seedBytes))
}
