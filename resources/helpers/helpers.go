package helpers

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func CheckArgsAndExec(callee func(args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if err := cobra.OnlyValidArgs(cmd, args); err != nil {
			return ShowValidArgument(cmd.ValidArgs)
		}
		return callee(args)
	}
}

func ShowValidArgument(validArgs []string) error {
	return fmt.Errorf("valid arguments are: %v", validArgs)
}

func ReadResourcesFromInputFile(absPath string) ([]Resource, error) {
	resources := make([]Resource, 0)

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		filesInfo, err := listFilesInDir(absPath)
		if err != nil {
			return nil, err
		}

		for _, fileInfo := range filesInfo {
			if err := appendResource(&resources, path.Join(absPath, fileInfo.Name())); err != nil {
				return nil, err
			}
		}
	} else {
		if err := appendResource(&resources, absPath); err != nil {
			return nil, err
		}
	}

	return resources, nil
}

func listFilesInDir(dirPath string, result ...os.FileInfo) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return result, err
	}

	for _, file := range files {
		if file.IsDir() {
			listFilesInDir(path.Join(dirPath, file.Name()), result...)
		} else {
			result = append(result, file)
		}
	}

	return result, nil
}

func appendResource(resources *[]Resource, fileAbsPath string) error {
	data, err := ioutil.ReadFile(fileAbsPath)
	if err != nil {
		return err
	}

	kind, err := determineResourceKind(fileAbsPath)
	if err != nil {
		return err
	}

	*resources = append(*resources, Resource{
		FileName: path.Base(fileAbsPath),
		Kind:     kind,
		Data:     data,
	})

	return nil
}

func determineResourceKind(fileAbsPath string) (string, error) {
	name := path.Base(fileAbsPath)
	parts := strings.Split(name, ".")
	if len(parts) < 3 {
		return "", fmt.Errorf("invalid file name: %s", name)
	}

	kind := parts[len(parts)-2]
	kindIsValid := false
	for _, validKind := range ValidResourceKinds {
		if kind == validKind {
			kindIsValid = true
			break
		}
	}

	if !kindIsValid {
		return "", fmt.Errorf("invalid resource kind: %s", kind)
	}

	return kind, nil
}

func SendCoreRequest(req *http.Request) ([]byte, error) {
	cli := http.Client{Timeout: 10 * time.Second}

	response, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("%s %s response: %d", req.Method, req.URL, response.StatusCode)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
