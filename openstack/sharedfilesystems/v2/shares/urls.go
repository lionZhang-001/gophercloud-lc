package shares

import "gophercloud-lc"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("shares")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("shares", id)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("shares", id)
}
