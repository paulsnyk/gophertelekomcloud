package addons

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

var RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json"},
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToAddonCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new addon
type CreateOpts struct {
	// API type, fixed value Addon
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Metadata required to create an addon
	Metadata CreateMetadata `json:"metadata" required:"true"`
	// specifications to create an addon
	Spec RequestSpec `json:"spec" required:"true"`
}

type CreateMetadata struct {
	Annotations Annotations `json:"annotations" required:"true"`
}

type Annotations struct {
	AddonInstallType string `json:"addon.install/type" required:"true"`
}

// Specifications to create an addon
type RequestSpec struct {
	// For the addon version.
	Version string `json:"version" required:"true"`
	// Cluster ID.
	ClusterID string `json:"clusterID" required:"true"`
	// Addon Template Name.
	AddonTemplateName string `json:"addonTemplateName" required:"true"`
	// Addon Parameters
	Values Values `json:"values" required:"true"`
}

type Values struct {
	Basic    map[string]interface{} `json:"basic" required:"true"`
	Advanced map[string]interface{} `json:"custom,omitempty"`
}

// ToAddonCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToAddonCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// addon.
func Create(c *golangsdk.ServiceClient, opts CreateOptsBuilder, clusterId string) (r CreateResult) {
	b, err := opts.ToAddonCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &golangsdk.RequestOpts{OkCodes: []int{201}}
	_, r.Err = c.Post(rootURL(c, clusterId), b, &r.Body, reqOpt)
	return
}

// Get retrieves a particular addon based on its unique ID.
func Get(c *golangsdk.ServiceClient, id, clusterId string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id, clusterId), &r.Body, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}

// Delete will permanently delete a particular addon based on its unique ID.
func Delete(c *golangsdk.ServiceClient, id, clusterId string) (r DeleteResult) {
	_, r.Err = c.Delete(resourceURL(c, id, clusterId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
