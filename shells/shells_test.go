package shells

// "pwn/logger" // make a logger similar to pwntools?

func main() {
	available_options_lang := [4]string{"python2", "go", "php", "fakeoption"}
	rhost := "192.168.0.196" // testing
	rport := 4456

	// for option in available_options_lang
	for i := 0; i < len(available_options_lang); i++ {
		ReShellGen(available_options_lang[i], rhost, rport)
	}
}
