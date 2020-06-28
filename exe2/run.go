package exe2

import (
	"os/exec"
	"strings"

	"aduu.dev/utils/errors2"
	"k8s.io/klog/v2"
)

func runWithSettings(cmd *exec.Cmd, setting *ExecuteSetting) (out string, err error) {
	defer func() {
		err = cleanup(err, cmd, setting)
	}()

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
