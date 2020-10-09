package os_services

import gophercloud "gophercloud-lc"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-services")
}
