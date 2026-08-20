// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2routeserver

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2RouteServerConfig struct {
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
	// The Amazon-side ASN of the Route Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_route_server#amazon_side_asn Ec2RouteServer#amazon_side_asn}
	AmazonSideAsn *float64 `field:"required" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// Whether to enable persistent routes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_route_server#persist_routes Ec2RouteServer#persist_routes}
	PersistRoutes *string `field:"optional" json:"persistRoutes" yaml:"persistRoutes"`
	// The duration of persistent routes in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_route_server#persist_routes_duration Ec2RouteServer#persist_routes_duration}
	PersistRoutesDuration *float64 `field:"optional" json:"persistRoutesDuration" yaml:"persistRoutesDuration"`
	// Whether to enable SNS notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_route_server#sns_notifications_enabled Ec2RouteServer#sns_notifications_enabled}
	SnsNotificationsEnabled interface{} `field:"optional" json:"snsNotificationsEnabled" yaml:"snsNotificationsEnabled"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_route_server#tags Ec2RouteServer#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

