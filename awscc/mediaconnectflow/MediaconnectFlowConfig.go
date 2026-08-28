// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectFlowConfig struct {
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
	// The name of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#name MediaconnectFlow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#source MediaconnectFlow#source}
	Source *MediaconnectFlowSource `field:"required" json:"source" yaml:"source"`
	// The Availability Zone that you want to create the flow in.
	//
	// These options are limited to the Availability Zones within the current AWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#availability_zone MediaconnectFlow#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The maintenance settings you want to use for the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#maintenance MediaconnectFlow#maintenance}
	Maintenance *MediaconnectFlowMaintenance `field:"optional" json:"maintenance" yaml:"maintenance"`
	// The media streams associated with the flow.
	//
	// You can associate any of these media streams with sources and outputs on the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#media_streams MediaconnectFlow#media_streams}
	MediaStreams interface{} `field:"optional" json:"mediaStreams" yaml:"mediaStreams"`
	// The source failover config of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#source_failover_config MediaconnectFlow#source_failover_config}
	SourceFailoverConfig *MediaconnectFlowSourceFailoverConfig `field:"optional" json:"sourceFailoverConfig" yaml:"sourceFailoverConfig"`
	// The source monitoring config of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#source_monitoring_config MediaconnectFlow#source_monitoring_config}
	SourceMonitoringConfig *MediaconnectFlowSourceMonitoringConfig `field:"optional" json:"sourceMonitoringConfig" yaml:"sourceMonitoringConfig"`
	// The VPC interfaces that you added to this flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_flow#vpc_interfaces MediaconnectFlow#vpc_interfaces}
	VpcInterfaces interface{} `field:"optional" json:"vpcInterfaces" yaml:"vpcInterfaces"`
}

