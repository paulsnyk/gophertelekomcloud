package v1

import (
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/eips"
)

func TestEipList(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	if err != nil {
		t.Fatalf("Unable to create NetworkV1 client : %s", err)
	}

	eipsList, err := eips.List(client).Extract()
	if err != nil {
		t.Fatalf("Error querying eips: %s", err)
	}
	tools.PrintResource(t, eipsList)
}

func TestEipCRUD(t *testing.T) {

}
