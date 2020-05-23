package exe

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

// Runner combines all run functions into one interface.
type Runner interface {
	Run(cmd *exec.Cmd, name string, settings ...SettingsFunc)
	RunE(cmd *exec.Cmd, name string, settings ...SettingsFunc) (err error)
	RunWithOutput(cmd *exec.Cmd, name string, settings ...SettingsFunc) string
	RunWithOutputE(cmd *exec.Cmd, name string, settings ...SettingsFunc) (out string, err error)
	SplitCommand(s string) []string
	Execute(s string, obj interface{}, settings ...SettingsFunc) *exec.Cmd
	ExecuteE(s string, obj interface{}, settings ...SettingsFunc) (*exec.Cmd, error)
	ExecuteWithContext(ctx context.Context, s string, obj interface{}, settings ...SettingsFunc) *exec.Cmd
	ExecuteEWithContext(ctx context.Context, s string, obj interface{}, settings ...SettingsFunc) (*exec.Cmd, error)
	ExecuteWithOutput(s string, obj interface{}, settings ...SettingsFunc) string
	ExecuteWithOutputE(s string, obj interface{}, settings ...SettingsFunc) (out string, err error)
}

// DefaultRunner returns the default Runner implementation.
func DefaultRunner() Runner {
	return &runner{}
}

type runner struct {
}

func (r *runner) Run(cmd *exec.Cmd, name string, settings ...SettingsFunc) {
	if err := r.RunE(cmd, name, settings...); err != nil {
		os.Exit(1)
	}
}

func (r *runner) RunE(cmd *exec.Cmd, name string, settings ...SettingsFunc) (err error) {
	setting := getSetting(settings)
	applySettings(cmd, setting)

	_, err = runWithSettings(cmd, setting)
	if err != nil {
		fmt.Printf("%s failed: %v\n", name, err)
		return err
	}
	return nil
}

func (r *runner) RunWithOutput(cmd *exec.Cmd, name string, settings ...SettingsFunc) string {
	out, err := r.RunWithOutputE(cmd, name, settings...)
	if err != nil {
		os.Exit(1)
	}
	return out
}

func (r *runner) RunWithOutputE(cmd *exec.Cmd, name string, settings ...SettingsFunc) (out string, err error) {
	setting := getSetting(append(settings, withOutput))
	applySettings(cmd, setting)

	out, err = runWithSettings(cmd, setting)
	if err != nil {
		fmt.Printf("%s - %v failed with %v\n", name, cmd, err)
		return out, err
	}
	return
}

func (r *runner) SplitCommand(s string) []string {
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

func (r *runner) Execute(s string, obj interface{}, settings ...SettingsFunc) *exec.Cmd {
	return r.ExecuteWithContext(context.Background(), s, obj, settings...)
}

func (r *runner) ExecuteE(s string, obj interface{}, settings ...SettingsFunc) (*exec.Cmd, error) {
	return r.ExecuteEWithContext(context.Background(), s, obj, settings...)
}

func (r *runner) ExecuteWithContext(ctx context.Context, s string, obj interface{}, settings ...SettingsFunc) *exec.Cmd {
	cmd, _ := r.ExecuteEWithContext(ctx, s, obj, settings...)
	return cmd
}

func (r *runner) ExecuteEWithContext(ctx context.Context, s string, obj interface{}, settings ...SettingsFunc) (*exec.Cmd, error) {
	t := template.Must(template.New("").Parse(s))

	var tpl bytes.Buffer

	if err := t.Execute(&tpl, obj); err != nil {
		panic(err)
	}

	fmt.Println(tpl.String())

	args := r.SplitCommand(os.ExpandEnv(tpl.String()))
	//args := strings.Split(tpl.String(), " ")

	if len(args) == 0 {
		panic("template length was zero.")
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	return cmd, r.RunE(cmd, args[0], settings...)
}

func (r *runner) ExecuteWithOutput(s string, obj interface{}, settings ...SettingsFunc) string {
	out, err := r.ExecuteWithOutputE(s, obj, settings...)
	if err != nil {
		fmt.Printf("Command failed with %v\n", err)
		os.Exit(1)
	}
	return out
}

func (r *runner) ExecuteWithOutputE(s string, obj interface{}, settings ...SettingsFunc) (out string, err error) {
	t := template.Must(template.New("").Parse(s))

	var tpl bytes.Buffer

	if err := t.Execute(&tpl, obj); err != nil {
		panic(err)
	}

	fmt.Println(tpl.String())

	args := r.SplitCommand(os.ExpandEnv(tpl.String()))

	if len(args) == 0 {
		panic("template length was zero.")
	}

	cmd := exec.Command(args[0], args[1:]...)
	return r.RunWithOutputE(cmd, args[0], settings...)
}

func getSetting(settings []SettingsFunc) ExecuteSetting {
	setting := defaultSettings()

	for _, settingsFunc := range settings {
		settingsFunc(&setting)
	}

	return setting
}

func applySettings(cmd *exec.Cmd, setting ExecuteSetting) {
	if setting.dir != "" {
		cmd.Dir = os.ExpandEnv(setting.dir)
	}

	if setting.output {
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
}

func defaultSettings() ExecuteSetting {
	return ExecuteSetting{}
}

func runWithSettings(cmd *exec.Cmd, setting ExecuteSetting) (out string, err error) {
	var timer *time.Timer
	exited := false
	var errFromTimeout error

	if setting.timeout != 0 {
		timer = time.AfterFunc(setting.timeout, func() {
			if !exited {
				errFromTimeout = cmd.Process.Kill()
			}
		})
	}

	switch {
	case setting.start && setting.output:
		panic("can't do start with output")
	case setting.start && !setting.output:
		err = cmd.Start()
	case !setting.start && setting.output:
		var byteOut []byte
		byteOut, err = cmd.Output()
		if byteOut != nil {
			out = strings.TrimSpace(string(byteOut))
		}
	default:
		err = cmd.Run()
	}

	if errFromTimeout != nil {
		if err != nil {
			err = fmt.Errorf("error from timeout %v wraps err: %v", errFromTimeout, err)
		} else {
			err = fmt.Errorf("error from timeout: %v", errFromTimeout)
		}
	}

	exited = true
	if timer != nil {
		timer.Stop()
	}
	return out, err
}
