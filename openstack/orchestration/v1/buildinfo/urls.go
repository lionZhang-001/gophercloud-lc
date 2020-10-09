package buildinfo

import "gophercloud-lc"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
