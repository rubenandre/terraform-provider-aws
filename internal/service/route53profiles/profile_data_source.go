// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53profiles

// **PLEASE DELETE THIS AND ALL TIP COMMENTS BEFORE SUBMITTING A PR FOR REVIEW!**
//
// TIP: ==== INTRODUCTION ====
// Thank you for trying the skaff tool!
//
// You have opted to include these helpful comments. They all include "TIP:"
// to help you find and remove them when you're done with them.
//
// While some aspects of this file are customized to your input, the
// scaffold tool does *not* look at the AWS API and ensure it has correct
// function, structure, and variable names. It makes guesses based on
// commonalities. You will need to make significant adjustments.
//
// In other words, as generated, this is a rough outline of the work you will
// need to do. If something doesn't make sense for your situation, get rid of
// it.

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/route53profiles"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	fwflex "github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
)

// TIP: ==== FILE STRUCTURE ====
// All data sources should follow this basic outline. Improve this data source's
// maintainability by sticking to it.
//
// 1. Package declaration
// 2. Imports
// 3. Main data source struct with schema method
// 4. Read method
// 5. Other functions (flatteners, expanders, waiters, finders, etc.)

// Function annotations are used for datasource registration to the Provider. DO NOT EDIT.
// @FrameworkDataSource(name="Profile")
func newprofileDataSource(context.Context) (datasource.DataSourceWithConfigure, error) {
	return &profileDataSource{}, nil
}

type profileDataSource struct {
	framework.DataSourceWithConfigure
}

func (d *profileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) { // nosemgrep:ci.meta-in-func-name
	resp.TypeName = "aws_route53profiles_profile"
}

func (d *profileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"arn": framework.ARNAttributeComputedOnly(),
			"client_token": schema.StringAttribute{
				Computed: true,
			},
			"creation_time": schema.NumberAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
			"modified_time": schema.NumberAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"owner_id": schema.StringAttribute{
				Computed: true,
			},
			"share_status": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"status_message": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *profileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data profileDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	conn := d.Meta().Route53ProfilesClient(ctx)

	input := &route53profiles.GetProfileInput{
		ProfileId: fwflex.StringFromFramework(ctx, data.ID),
	}

	output, err := conn.GetProfile(ctx, input)

	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("reading Route53 Profile (%s)", data.ID), err.Error())
		return
	}

	resp.Diagnostics.Append(fwflex.Flatten(ctx, output.Profile, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

type profileDataSourceModel struct {
	ARN           types.String `tfsdk:"arn"`
	ClientToken   types.String `tfsdk:"client_token"`
	CreationTime  types.Number `tfsdk:"creation_time"`
	ID            types.String `tfsdk:"id"`
	ModifiedTime  types.Number `tfsdk:"modified_time"`
	Name          types.String `tfsdk:"name"`
	OwnerID       types.String `tfsdk:"owner_id"`
	ShareStatus   types.String `tfsdk:"share_status"`
	Status        types.String `tfsdk:"status"`
	StatusMessage types.String `tfsdk:"status_message"`
}
