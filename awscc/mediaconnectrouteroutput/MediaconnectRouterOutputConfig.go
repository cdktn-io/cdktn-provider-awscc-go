// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectRouterOutputConfig struct {
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
	// The configuration settings for a router output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#configuration MediaconnectRouterOutput#configuration}
	Configuration *MediaconnectRouterOutputConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The maximum bitrate for the router output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#maximum_bitrate MediaconnectRouterOutput#maximum_bitrate}
	MaximumBitrate *float64 `field:"required" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The name of the router output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#name MediaconnectRouterOutput#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#routing_scope MediaconnectRouterOutput#routing_scope}.
	RoutingScope *string `field:"required" json:"routingScope" yaml:"routingScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#tier MediaconnectRouterOutput#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// The Availability Zone where you want to create the router output.
	//
	// This must be a valid Availability Zone for the region specified by regionName, or the current region if no regionName is provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#availability_zone MediaconnectRouterOutput#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The fabric configuration settings for the router output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#fabric_configuration MediaconnectRouterOutput#fabric_configuration}
	FabricConfiguration *MediaconnectRouterOutputFabricConfiguration `field:"optional" json:"fabricConfiguration" yaml:"fabricConfiguration"`
	// The configuration settings for maintenance operations, including preferred maintenance windows and schedules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#maintenance_configuration MediaconnectRouterOutput#maintenance_configuration}
	MaintenanceConfiguration *MediaconnectRouterOutputMaintenanceConfiguration `field:"optional" json:"maintenanceConfiguration" yaml:"maintenanceConfiguration"`
	// The Amazon Web Services Region for the router output. Defaults to the current region if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#region_name MediaconnectRouterOutput#region_name}
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Key-value pairs that can be used to tag this router output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_router_output#tags MediaconnectRouterOutput#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

