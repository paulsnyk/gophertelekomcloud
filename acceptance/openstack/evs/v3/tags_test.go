package v3

import (
	"testing"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/evs/v3/volumes"
)

func TestEvsTagsCRUD(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	if err != nil {
		t.Fatalf("Unable to create a BlockStorageV3 client: %s", err)
	}

	createTagsOpts := tags.CreateOpts{}
}

func createEVSv3Volume(t *testing.T, client *golangsdk.ServiceClient) (*volumes.Volume, error) {
	createOpts := &volumes.CreateOpts{
		AvailabilityZone: clients.OS_AVAILABILITY_ZONE,
		VolumeType:       "",
	}
}

func deleteEVSv3Volume(t *testing.T, client *golangsdk.ServiceClient, rdsId string) {

}

func updateEVSTags(t *testing.T, client *golangsdk.ServiceClient) error {

}
