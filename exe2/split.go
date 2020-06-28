package exe2

import "strings"

// SplitResult stores the result of a split method and is passed
// to Runner.RunE or Runner.RunWithOutputE.
//
// If SplitResult.Err != nil the Runner methods will return this error
// without executing the command.
type SplitResult struct {
	Name string
	Args []string

	Err error
}

// Split just splits the argument and returns the result.
func Split(arg string) SplitResult {
	return splitCommand(arg)
}

func (result *SplitResult) command() []string {
	return append([]string{result.Name}, result.Args...)
}

func splitCommand(arg string) SplitResult {
	split := splitCommandOld(arg)

	name := ""
	args := []string{}

	switch {
	case len(split) == 1:
		name = split[0]
	case len(split) > 1:
		name = split[0]
		args = split[1:]
	default:
		break
	}

	return SplitResult{
		Name: name,
		Args: args,
	}
}

func splitCommandOld(s string) []string {
	runes := []rune(s)
	var values []string

	for len(runes) != 0 {
		// Workaround for bug which triggers if "" is detected,
		// then space is the first character instead of a possible " afterwards.
		runes = []rune(strings.TrimSpace(string(runes)))
		if len(runes) == 0 {
			break
		}

		currentRunes := string(runes)
		_ = currentRunes
		first := string(runes[0])
		_ = first

		// TODO: ' and " can be in the middle of a token, so checking the start of a string is not enough.
		if runes[0] == '"' {
			next := search(runes[1:], "\"")
			if next == -1 {
				panic("no matching \": " + s)
			}

			values = append(values, string(runes[1:next+1]))
			runes = runes[next+2:]
		} else {
			// TODO: If ' before next space, do skip this space and go instead to the next ' to continue from there.

			next := search(runes, " ")
			if next == -1 {
				values = append(values, string(runes))
				break
			}

			nextSpecial := search(runes, "'")

			// Early return if nextSpecial is not so.
			if nextSpecial == -1 || nextSpecial > next {
				values = append(values, string(runes[:next]))
				runes = runes[next+1:]
				continue
			}

			var current []rune

			for next != -1 && nextSpecial != -1 && nextSpecial < next {
				// Skip newlines enclosed by specials '  '.

				// Add to current.
				current = append(current, runes[:nextSpecial]...)
				// Move ahead.
				runes = runes[nextSpecial+1:]

				// Find the closing special character.
				nextSpecial = search(runes, "'")
				if nextSpecial == -1 {
					panic("no matching ': " + s)
				}
				// Add up to the closing parameter to current.
				current = append(current, runes[:nextSpecial]...)
				// Mve ahead.
				runes = runes[nextSpecial+1:]

				nextSpecial = search(runes, "'")
				next = search(runes, " ")
			}

			next = search(runes, " ")
			if next == -1 {
				// Add current to last part.
				values = append(values, string(append(current, runes...)))
				// Move ahead.
				break
			}

			// Add current to closing part.
			values = append(values, string(append(current, runes[:next]...)))
			// Move ahead.
			runes = runes[next+1:]
		}
	}

	out := make([]string, len(values))
	copy(out, values)
	return out
}

func search(text []rune, what string) int {
	whatRunes := []rune(what)

	for i := range text {
		found := true
		for j := range whatRunes {
			if text[i+j] != whatRunes[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}
