package tags

import "github.com/opentelekomcloud/gophertelekomcloud"

func createURL(c *golangsdk.ServiceClient, volumeId string) string {
	return c.ServiceURL(c.ProjectID, "os-vendor-tags", volumeId, "tags", "action")
}
