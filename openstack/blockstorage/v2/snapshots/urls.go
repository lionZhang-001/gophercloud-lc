package snapshots

import gophercloud "gophercloud-lc"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("snapshots")
}
