package pwn

import (
	"./shells"
	// "pwn/logger" // make a logger similar to pwntools?

	"fmt"
)

func main() {
	fmt.Println("Reverse Shell Generator - Testing:")
	
	available_options_lang := [4]string{'python2', 'go', 'php', 'fakeoption'}
	rhost := "192.168.0.196" // testing
	rport := 4456

	for i := 0; i < len(available_options_lang); i++ { // for option in available_options_lang
		shells.ReShellGen(available_options_lang[i], rhost, rport)
	}
}