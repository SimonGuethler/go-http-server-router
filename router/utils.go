package router

import (
	"errors"
	"regexp"
	"strings"
)

const (
	pathElementPattern      = `^[a-zA-Z0-9_-]+$`
	dynamicPathAllowedChars = `^:*[a-zA-Z0-9_-]+$`
	dynamicPathPattern      = `^:[a-zA-Z0-9_-]+$`
)

func SanitizePath(path string, allowDynamic bool) (string, error) {
	var pathPattern string
	if allowDynamic {
		pathPattern = dynamicPathAllowedChars
	} else {
		pathPattern = pathElementPattern
	}

	path = strings.TrimSpace(path)
	path = strings.Trim(path, "/")
	pathElements := SplitPath(path)

	if len(pathElements) == 0 {
		return "/", nil
	}
	if len(pathElements) == 1 && pathElements[0] == "" {
		return "/", nil
	}

	var sanitizedElements []string
	for _, pathElement := range pathElements {
		sanitizedElement, err := SanitizePathElement(pathElement, pathPattern)
		if err != nil {
			return "", errors.New("invalid path element")
		}
		sanitizedElements = append(sanitizedElements, sanitizedElement)
	}

	return "/" + strings.Join(sanitizedElements, "/"), nil
}

func SanitizePathElement(pathElement string, allowPattern string) (string, error) {
	re := regexp.MustCompile(allowPattern)
	if !re.MatchString(pathElement) {
		return "", errors.New("invalid path element")
	}
	return strings.Trim(strings.ToLower(pathElement), "/"), nil
}

func SplitPath(path string) []string {
	path = strings.Trim(path, "/")
	pathParts := strings.Split(path, "/")
	return pathParts
}

func IsPathParam(pathSegment string) bool {
	re := regexp.MustCompile(dynamicPathPattern)
	return re.MatchString(pathSegment)
}

func ExtractPathVariables(routePath string, requestPath string) map[string]string {
	routeParts := SplitPath(routePath)
	requestParts := SplitPath(requestPath)

	var pathVariables = make(map[string]string)
	for i, part := range routeParts {
		if IsPathParam(part) {
			pathVariables[part[1:]] = requestParts[i]
		}
	}
	return pathVariables
}
