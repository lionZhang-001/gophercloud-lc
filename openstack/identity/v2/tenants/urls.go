package tenants

import "gophercloud-lc"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("tenants")
}
