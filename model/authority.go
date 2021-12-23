package model

import (
	"errors"
	"fmt"
	"strings"
)

const resourceSepWord = "#^#"

// NewResourceFromString load resource name.
func NewResourceFromString(resource string) (*Resource, error) {
	splitedResource := strings.SplitN(resource, resourceSepWord, 2)
	if len(splitedResource) != 2 {
		return nil, errors.New(fmt.Sprint("wrong resource name", resource))
	}

	res := &Resource{
		Type: splitedResource[0],
		Name: splitedResource[1],
	}

	return res, nil
}

// StringifyResource make a resource to a string type var.
func StringifyResource(resource *Resource) (string, error) {
	if strings.Contains(resource.Type, resourceSepWord) || strings.Contains(resource.Name, resourceSepWord) {
		return "", errors.New(fmt.Sprint(resourceSepWord, " can't appear in resource.Name or resource.Type"))
	}

	return fmt.Sprint(resource.Type, resourceSepWord, resource.Name), nil
}
