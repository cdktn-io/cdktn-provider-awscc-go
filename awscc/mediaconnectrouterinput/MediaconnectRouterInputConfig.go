// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectRouterInputConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The configuration settings for a router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#configuration MediaconnectRouterInput#configuration}
	Configuration *MediaconnectRouterInputConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The maximum bitrate for the router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#maximum_bitrate MediaconnectRouterInput#maximum_bitrate}
	MaximumBitrate *float64 `field:"required" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The name of the router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#name MediaconnectRouterInput#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#routing_scope MediaconnectRouterInput#routing_scope}.
	RoutingScope *string `field:"required" json:"routingScope" yaml:"routingScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#tier MediaconnectRouterInput#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// The Availability Zone where you want to create the router input.
	//
	// This must be a valid Availability Zone for the region specified by regionName, or the current region if no regionName is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#availability_zone MediaconnectRouterInput#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The content quality analysis configuration for the router input.
	//
	// The content quality analysis feature only monitors the first video stream and the first audio stream it encounters within the router input source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#content_quality_analysis_configuration MediaconnectRouterInput#content_quality_analysis_configuration}
	ContentQualityAnalysisConfiguration *MediaconnectRouterInputContentQualityAnalysisConfiguration `field:"optional" json:"contentQualityAnalysisConfiguration" yaml:"contentQualityAnalysisConfiguration"`
	// The configuration settings for maintenance operations, including preferred maintenance windows and schedules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#maintenance_configuration MediaconnectRouterInput#maintenance_configuration}
	MaintenanceConfiguration *MediaconnectRouterInputMaintenanceConfiguration `field:"optional" json:"maintenanceConfiguration" yaml:"maintenanceConfiguration"`
	// The Amazon Web Services Region for the router input. Defaults to the current region if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#region_name MediaconnectRouterInput#region_name}
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Key-value pairs that can be used to tag and organize this router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#tags MediaconnectRouterInput#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The transit encryption settings for a router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#transit_encryption MediaconnectRouterInput#transit_encryption}
	TransitEncryption *MediaconnectRouterInputTransitEncryption `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
}

