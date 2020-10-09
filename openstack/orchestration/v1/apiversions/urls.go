package apiversions

import "gophercloud-lc"

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}
