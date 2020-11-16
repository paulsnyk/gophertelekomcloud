package tags

import "github.com/opentelekomcloud/gophertelekomcloud"

type commonResult struct {
	golangsdk.Result
}

// Tags model
type Tags struct {
	// Tags is a list of any tags. Tags are arbitrarily defined strings
	// attached to a resource.
	Tags map[string]string `json:"tags"`
}

// CreateResult represents the result of a Create operation
type CreateResult struct {
	commonResult
}

// Extract interprets any commonResult as a Tags.
func (r commonResult) Extract() (*Tags, error) {
	var s *Tags
	err := r.ExtractInto(&s)
	return s, err
}
