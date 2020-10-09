package hypervisors

import (
	gophercloud "gophercloud-lc"
	//"gophercloud-lc"
	//"gophercloud-lc/pagination"
	"gophercloud-lc/pagination"
)

// List makes a request against the API to list hypervisors.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, hypervisorsListDetailURL(client), func(r pagination.PageResult) pagination.Page {
		return HypervisorPage{pagination.SinglePageBase(r)}
	})
}

func AggregateList(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, aggregatesListURL(client), func(r pagination.PageResult) pagination.Page {
		return AggregatePage{pagination.SinglePageBase(r)}
	})
}
