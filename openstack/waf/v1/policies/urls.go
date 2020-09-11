package policies

import "github.com/opentelekomcloud/gophertelekomcloud"

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("policy")
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("policy", id)
}

func hostsURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("policy", id, "hosts")
}
