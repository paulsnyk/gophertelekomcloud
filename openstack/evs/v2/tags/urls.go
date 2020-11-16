package tags

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

func createURL(c *golangsdk.ServiceClient, resourceType, resourceId string) string {
	return c.ServiceURL("os-vendor-tags", resourceType, resourceId)
}

func getURL(c *golangsdk.ServiceClient, resourceType, resourceId string) string {
	return c.ServiceURL("os-vendor-tags", resourceType, resourceId)
}
