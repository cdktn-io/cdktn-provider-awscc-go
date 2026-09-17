// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelqualityjobdefinition


type SagemakerModelQualityJobDefinitionNetworkConfigVpcConfig struct {
	// The VPC security group IDs, in the form sg-xxxxxxxx.
	//
	// Specify the security groups for the VPC that is specified in the Subnets field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_quality_job_definition#security_group_ids SagemakerModelQualityJobDefinition#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnets in the VPC to which you want to connect to your monitoring jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_quality_job_definition#subnets SagemakerModelQualityJobDefinition#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

