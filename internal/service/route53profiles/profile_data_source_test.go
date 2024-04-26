// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53profiles_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"

	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccRoute53ProfilesProfileDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	datasourceName := "data.aws_route53profiles_profile.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.Route53ProfilesServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRoute53ProfilesProfileDataSource_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttrSet(datasourceName, "arn"),
					resource.TestCheckResourceAttrSet(datasourceName, "name"),
					resource.TestCheckResourceAttrSet(datasourceName, "client_token"),
					resource.TestCheckResourceAttrSet(datasourceName, "creation_time"),
					resource.TestCheckResourceAttrSet(datasourceName, "modified_time"),
					resource.TestCheckResourceAttrSet(datasourceName, "status"),
					resource.TestCheckResourceAttrSet(datasourceName, "owner_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "share_status"),
					resource.TestCheckResourceAttrSet(datasourceName, "status_message"),
				),
			},
		},
	})
}

func testAccRoute53ProfilesProfileDataSource_basic() string {
	return `
data "aws_route53profiles_profile" "test" {
	id = "rp-f080c59961fa4516"
}
`
}
