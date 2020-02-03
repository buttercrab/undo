package util

type Command struct {
	Raw      string
	command  string
	argument []string
	pipe     *Command
}

func parseCommand(command string) []string {
	res := []string{""}
	for _, c := range command {
		if c == ' ' || c == '\n' || c == '\t' {
			// TODO
		}
	}
	return res
}

// TODO
func NewCommand(command string) *Command {
	cmd := Command{Raw: command}
	split := parseCommand(command)
	if len(split) > 0 {
		cmd.command = split[0]
		cmd.argument = split[1:]
	}
	return &cmd
}
