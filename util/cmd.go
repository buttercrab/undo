package util

type Command struct {
	Raw  string
	Main string
	Args []string
	Pipe *Command
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
		cmd.Main = split[0]
		cmd.Args = split[1:]
	}
	return &cmd
}
