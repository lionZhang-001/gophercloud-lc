// +build acceptance

package v3

import (
	"testing"

	"gophercloud-lc/acceptance/clients"
	"gophercloud-lc/acceptance/tools"
	"gophercloud-lc/openstack"
	"gophercloud-lc/openstack/identity/v3/tokens"
)

func TestGetToken(t *testing.T) {
	client, err := clients.NewIdentityV3UnauthenticatedClient()
	if err != nil {
		t.Fatalf("Unable to obtain an identity client: %v")
	}

	ao, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		t.Fatalf("Unable to obtain environment auth options: %v", err)
	}

	authOptions := tokens.AuthOptions{
		Username:   ao.Username,
		Password:   ao.Password,
		DomainName: "default",
	}

	token, err := tokens.Create(client, &authOptions).Extract()
	if err != nil {
		t.Fatalf("Unable to get token: %v", err)
	}

	tools.PrintResource(t, token)
}
