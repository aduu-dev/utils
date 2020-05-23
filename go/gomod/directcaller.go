package gomod

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/xerrors"
)

// DirectCallerPackageE is same as DirectCallerPackage but returns an error instead of panicing once an error occurs.
func DirectCallerPackageE() (s string, err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("determining current package failed: %w", err)
		}
	}()

	currentDir := packagePathOfCaller(2)

	ws, err := GetWorkspaceFromWD()
	if err != nil {
		return
	}

	gomodDir := filepath.Dir(ws.GomodPath)

	if !strings.HasPrefix(currentDir, gomodDir) {
		return "", fmt.Errorf("the directory the caller function resides in (%s)"+
			"is not under the determined workspace (%s) with Module prefix (%s)"+
			"and so package name cannot be determined", currentDir, ws.GomodPath, ws.Module)
	}

	// Add module path in front, then add current file with root removed.
	return ws.Module + "/" + strings.TrimPrefix(currentDir, gomodDir+"/"), nil
}

// DirectCallerPackage returns the package name the working directory currently points to.
// The function panics if working directory is not inside a workspace.
func DirectCallerPackage() string {
	pkg, err := DirectCallerPackageE()
	if err != nil {
		panic(err)
	}
	return pkg
}

// DirectCallerPackagePath determines the path to the directory the caller function resides in and returns it.
func DirectCallerPackagePath() string {
	return packagePathOfCaller(2)
}

// packagePathOfCaller returns the file callerDepth calls removed from this call.
func packagePathOfCaller(callerDepth int) string {
	_, currentFile, _, _ := runtime.Caller(callerDepth)
	return filepath.Dir(currentFile)
}

// Source: https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
/*
func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}
*/

// DirectCallerFunctionName determines the function name of the caller and returns it.
func DirectCallerFunctionName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return strings.TrimPrefix(filepath.Ext(f.Name()), ".")
}

// DirectCallerFunctionNameWithPackageName returns the caller pkg.function name.
func DirectCallerFunctionNameWithPackageName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return filepath.Base(f.Name())
}
