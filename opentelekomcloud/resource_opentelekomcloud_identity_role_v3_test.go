// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package opentelekomcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/huaweicloud/golangsdk"
)

func TestAccIdentityRoleV3_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIdentityRoleV3Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityRoleV3_basic(acctest.RandString(10)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIdentityRoleV3Exists(),
				),
			},
		},
	})
}

func testAccIdentityRoleV3_basic(val string) string {
	return fmt.Sprintf(`
resource "opentelekomcloud_identity_role_v3" "role" {
  description = "role"
  display_name = "custom_role%s"
  display_layer = "domain"
  statement = [
    {
      effect = "Allow"
      action = ["ecs:*:list*"]
    }
  ]
}
	`, val)
}

func testAccCheckIdentityRoleV3Destroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.sdkClient("", "identity", serviceDomainLevel)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}
	client.Endpoint = "https://iam.eu-de.otc.t-systems.com/v3.0/"

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opentelekomcloud_identity_role_v3" {
			continue
		}

		url, err := replaceVarsForTest(rs, "OS-ROLE/roles/{id}")
		if err != nil {
			return err
		}
		url = client.ServiceURL(url)

		_, err = client.Get(
			url, nil,
			&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
		if err == nil {
			return fmt.Errorf("opentelekomcloud_identity_role_v3 still exists at %s", url)
		}
	}

	return nil
}

func testAccCheckIdentityRoleV3Exists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)
		client, err := config.sdkClient("", "identity", serviceDomainLevel)
		if err != nil {
			return fmt.Errorf("Error creating sdk client, err=%s", err)
		}
		client.Endpoint = "https://iam.eu-de.otc.t-systems.com/v3.0/"

		rs, ok := s.RootModule().Resources["opentelekomcloud_identity_role_v3.role"]
		if !ok {
			return fmt.Errorf("Error checking opentelekomcloud_identity_role_v3.role exist, err=not found opentelekomcloud_identity_role_v3.role")
		}

		url, err := replaceVarsForTest(rs, "OS-ROLE/roles/{id}")
		if err != nil {
			return fmt.Errorf("Error checking opentelekomcloud_identity_role_v3.role exist, err=building url failed: %s", err)
		}
		url = client.ServiceURL(url)

		_, err = client.Get(
			url, nil,
			&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				return fmt.Errorf("opentelekomcloud_identity_role_v3.role is not exist")
			}
			return fmt.Errorf("Error checking opentelekomcloud_identity_role_v3.role exist, err=send request failed: %s", err)
		}
		return nil
	}
}
