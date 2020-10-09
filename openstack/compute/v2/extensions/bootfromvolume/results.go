package bootfromvolume

import (
	os "gophercloud-lc/openstack/compute/v2/servers"
)

// CreateResult temporarily contains the response from a Create call.
type CreateResult struct {
	os.CreateResult
}
