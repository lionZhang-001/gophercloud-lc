package testing

import (
	"testing"

	"gophercloud-lc/openstack/orchestration/v1/buildinfo"
	th "gophercloud-lc/testhelper"
	fake "gophercloud-lc/testhelper/client"
)

func TestGetTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t, GetOutput)

	actual, err := buildinfo.Get(fake.ServiceClient()).Extract()
	th.AssertNoErr(t, err)

	expected := GetExpected
	th.AssertDeepEquals(t, expected, actual)
}
