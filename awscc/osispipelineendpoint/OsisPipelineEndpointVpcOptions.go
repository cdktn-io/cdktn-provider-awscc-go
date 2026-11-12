// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package osispipelineendpoint


type OsisPipelineEndpointVpcOptions struct {
	// A list of security group IDs that control network access to the pipeline endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/osis_pipeline_endpoint#security_group_ids OsisPipelineEndpoint#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs where the pipeline endpoint network interfaces are created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/osis_pipeline_endpoint#subnet_ids OsisPipelineEndpoint#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

