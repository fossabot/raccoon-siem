package resource

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

type resource struct {
	fileName string
	kind     string
	data     []byte
}

func readResourcesFromInputFile(absPath string) ([]resource, error) {
	resources := make([]resource, 0)

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

func appendResource(resources *[]resource, fileAbsPath string) error {
	data, err := ioutil.ReadFile(fileAbsPath)
	if err != nil {
		return err
	}

	kind, err := determineResourceKind(fileAbsPath)
	if err != nil {
		return err
	}

	*resources = append(*resources, resource{
		fileName: path.Base(fileAbsPath),
		kind:     kind,
		data:     data,
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
	for _, validKind := range validResourceKinds {
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

func sendRequest(req *http.Request) ([]byte, error) {
	response, err := httpClient.Do(req)
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
