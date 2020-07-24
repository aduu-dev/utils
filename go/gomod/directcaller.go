package gomod

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/xerrors"
)

// DirectCallerPackageE is same as DirectCallerPackagePanic
// but returns an error instead of panicing once an error occurs.
func DirectCallerPackageE() (s string, err error) {
	return directCallerPackageE(3)
}

type callInfo struct {
	packageName string
	fileName    string
	funcName    string
	line        int
}

func retrieveCallInfo() *callInfo {
	pc, file, line, _ := runtime.Caller(3)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	fmt.Println("parts:", parts)
	// parts: [
	// 		go/gomod/directcallertest/go_default_test
	//   	TestDirectCallerPackage
	// ]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &callInfo{
		packageName: packageName,
		fileName:    fileName,
		funcName:    funcName,
		line:        line,
	}
}

func directCallerPackageE(depth int) (s string, err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("determining current package failed: %w", err)
		}
	}()

	call := retrieveCallInfo()

	return call.packageName, nil

	/*
		currentDir := packagePathOfCaller(depth)

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

		fmt.Println("currentDir:", currentDir, "gomodDir:", gomodDir)

		// Add module path in front, then add current file with root removed.
		return strings.ReplaceAll(
			ws.Module+"/"+strings.TrimPrefix(strings.TrimSuffix(strings.TrimPrefix(currentDir, gomodDir), string(filepath.Separator)), string(filepath.Separator)),
			string(filepath.Separator), "/"), nil
	*/
}

// DirectCallerPackagePanic returns the package name the working directory currently points to.
// The function panics if working directory is not inside a workspace.
func DirectCallerPackagePanic() string {
	pkg, err := directCallerPackageE(3)
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
