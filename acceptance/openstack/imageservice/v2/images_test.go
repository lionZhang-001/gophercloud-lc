// +build acceptance imageservice images

package v2

import (
	"testing"

	"gophercloud-lc/acceptance/clients"
	"gophercloud-lc/acceptance/tools"
)

func TestImagesCreateDestroyEmptyImage(t *testing.T) {
	client, err := clients.NewImageServiceV2Client()
	if err != nil {
		t.Fatalf("Unable to create an image service client: %v", err)
	}

	image, err := CreateEmptyImage(t, client)
	if err != nil {
		t.Fatalf("Unable to create empty image: %v", err)
	}

	defer DeleteImage(t, client, image)

	tools.PrintResource(t, image)
}
