package nodepools

import "github.com/opentelekomcloud/gophertelekomcloud"

const (
	rootPath     = "clusters"
	resourcePath = "nodepools"
)

func rootURL(c *golangsdk.ServiceClient, clusterid string) string {
	return c.ServiceURL(rootPath, clusterid, resourcePath)
}

func resourceURL(c *golangsdk.ServiceClient, clusterid, nodepoolid string) string {
	return c.ServiceURL(rootPath, clusterid, resourcePath, nodepoolid)
}
