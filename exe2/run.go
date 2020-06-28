package exe2

import (
	"os/exec"
	"strings"
	"time"

	"aduu.dev/utils/errors2"
	"k8s.io/klog/v2"
)

func runWithSettings(cmd *exec.Cmd, setting *ExecuteSetting) (out string, err error) {
	defer func() {
		err = cleanup(err, cmd, setting)
	}()

	var timer *time.Timer
	exited := false

	// Run it in a goroutine to start the timeout timer after.
	go func() {
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
	}()

	if setting.timeout != 0 {
		timer = time.AfterFunc(setting.timeout, func() {
			// Only call kill if we did not exit already.
			klog.InfoS("State",
				"isProcessNil", cmd.Process == nil,
				"isProcessStateNil", cmd.ProcessState == nil)

			// What if we did not start (cmd.ProcessState == nil)
			// but the timeout triggered already?
			//
			// If we don't wait then sending the kill signal fails
			// and we get a panic. Wait for 1s and kill it then when
			// we are sure the process must have been started.
			//
			// Starting the timer after the cmd.Run/cmd.Start is
			// superior though. Maybe run them in a Goroutine.
			if cmd.ProcessState == nil {
				_ = timer.Stop()

				// Best effort.
				timer = time.AfterFunc(time.Second, func() {
					if !exited &&
						cmd.Process != nil {
						if timeoutErr := cmd.Process.Kill(); timeoutErr != nil {
							klog.ErrorS(err, "Command timeout",
								"args", cmd.Args)
						}
					}
				})

				return
			}

			if !exited {
				if killErr := cmd.Process.Kill(); killErr != nil {
					klog.ErrorS(err, "Kill error",
						"args", cmd.Args)
				}
			}
		})
	}

	exited = true
	if timer != nil {
		timer.Stop()
	}
	return out, err
}

func cleanup(err error, cmd *exec.Cmd, setting *ExecuteSetting) error {
	klog.V(5).InfoS("Files to close", "count", len(setting.openFiles))

	// Closing file-pipes if used.
	for _, file := range setting.openFiles {
		klog.V(5).InfoS("Closing file",
			"name", file.Name())
		if closeErr := file.Close(); closeErr != nil {
			err = errors2.CombineErrors(err, closeErr)
		}
	}

	return err
}
