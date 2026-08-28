// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationdataflowendpointgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GroundstationDataflowEndpointGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/groundstation_dataflow_endpoint_group#endpoint_details GroundstationDataflowEndpointGroup#endpoint_details}.
	EndpointDetails interface{} `field:"required" json:"endpointDetails" yaml:"endpointDetails"`
	// Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state.
	//
	// A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/groundstation_dataflow_endpoint_group#contact_post_pass_duration_seconds GroundstationDataflowEndpointGroup#contact_post_pass_duration_seconds}
	ContactPostPassDurationSeconds *float64 `field:"optional" json:"contactPostPassDurationSeconds" yaml:"contactPostPassDurationSeconds"`
	// Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state.
	//
	// A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/groundstation_dataflow_endpoint_group#contact_pre_pass_duration_seconds GroundstationDataflowEndpointGroup#contact_pre_pass_duration_seconds}
	ContactPrePassDurationSeconds *float64 `field:"optional" json:"contactPrePassDurationSeconds" yaml:"contactPrePassDurationSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/groundstation_dataflow_endpoint_group#tags GroundstationDataflowEndpointGroup#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

