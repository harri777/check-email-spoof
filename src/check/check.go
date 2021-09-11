package check

import (
	"fmt"
	"net"
	"strings"

	"github.com/fatih/color"
)

func CheckVulnerability(host string) {
	fmt.Println("Find DNS register...")
	txtrecords, _ := net.LookupTXT(host)

	for _, txt := range txtrecords {
		if strings.Contains(txt, "~all") || strings.Contains(txt, "?all") {
			color.Yellow("\nThe host IS VULNERABLE !\n")
			return
		}
	}
	color.Green("\nThe host be safe.\n")
}
