package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spaceuptech/space-cli/cmd/model"
	"github.com/spaceuptech/space-cli/cmd/utils"
)

// GetServicesRoutes gets services routes
func GetServicesRoutes(project, commandName string, params map[string]string) ([]*model.SpecObject, error) {
	url := fmt.Sprintf("/v1/runner/%s/service-routes", project)

	// Get the spec from the server
	payload := new(model.Response)
	if err := utils.Get(http.MethodGet, url, params, payload); err != nil {
		return nil, err
	}

	var objs []*model.SpecObject
	for _, item := range payload.Result {
		spec := item.(map[string]interface{})
		meta := map[string]string{"project": project, "id": spec["id"].(string)}

		// Delete the unwanted keys from spec
		delete(spec, "id")
		delete(spec, "project")
		delete(spec, "version")

		// Printing the object on the screen
		s, err := utils.CreateSpecObject("/v1/runner/{project}/service-routes/{id}", commandName, meta, spec)
		if err != nil {
			return nil, err
		}
		objs = append(objs, s)
	}

	return objs, nil
}

// GetServicesSecrets gets services secrets
func GetServicesSecrets(project, commandName string, params map[string]string) ([]*model.SpecObject, error) {
	url := fmt.Sprintf("/v1/runner/%s/secrets", project)

	// Get the spec from the server
	result := new(model.Response)
	if err := utils.Get(http.MethodGet, url, params, result); err != nil {
		return nil, err
	}

	var services []*model.SpecObject
	for _, item := range result.Result {
		spec := item.(map[string]interface{})

		meta := map[string]string{"project": project, "id": spec["id"].(string)}

		// Delete the unwanted keys from spec
		delete(spec, "id")
		delete(spec, "name")

		// Printing the object on the screen
		s, err := utils.CreateSpecObject("//v1/runner/{project}/secrets/{id}", commandName, meta, spec)
		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}

	return services, nil
}

// GetServices gets services
func GetServices(project, commandName string, params map[string]string) ([]*model.SpecObject, error) {
	url := fmt.Sprintf("/v1/runner/%s/services", project)
	// Get the spec from the server
	payload := new(model.Response)
	if err := utils.Get(http.MethodGet, url, params, payload); err != nil {
		return nil, err
	}

	var objs []*model.SpecObject
	for _, item := range payload.Result {
		spec := item.(map[string]interface{})
		id, ok := spec["id"]
		if !ok {
			// array may have an empty object
			continue
		}
		str := strings.Split(id.(string), "-")
		meta := map[string]string{"project": project, "version": str[1], "serviceId": str[0]}

		// Delete the unwanted keys from spec
		delete(spec, "id")
		delete(spec, "name")
		delete(spec, "version")
		delete(spec, "projectId")

		// Generating the object
		s, err := utils.CreateSpecObject("/v1/runner/{project}/services/{serviceId}/{version}", commandName, meta, spec)
		if err != nil {
			return nil, err
		}
		objs = append(objs, s)
	}
	return objs, nil
}