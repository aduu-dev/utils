package exe2

import (
	"fmt"
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
