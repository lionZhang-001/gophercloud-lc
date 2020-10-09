package roles

import "gophercloud-lc"

func listAssignmentsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("roles")
}
