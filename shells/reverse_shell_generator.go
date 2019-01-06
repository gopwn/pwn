package shells

import (
	"fmt"
	"os"
)

// ShellError represents an error with the shell
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
func ReShellGen(lang string, rhost string, rport int) (string, error) {

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

func loadTemplate(lang string) error {
	file, err := os.Open("templates/shell_" + lang)
	if err != nil {
		return err
	}

	return nil
}
