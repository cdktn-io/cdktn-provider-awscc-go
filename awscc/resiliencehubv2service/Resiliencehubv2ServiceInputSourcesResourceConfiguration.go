// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceInputSourcesResourceConfiguration struct {
	// ARN of a CloudFormation stack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/resiliencehubv2_service#cfn_stack_arn Resiliencehubv2Service#cfn_stack_arn}
	CfnStackArn *string `field:"optional" json:"cfnStackArn" yaml:"cfnStackArn"`
	// S3 URL of a design file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/resiliencehubv2_service#design_file_s3_url Resiliencehubv2Service#design_file_s3_url}
	DesignFileS3Url *string `field:"optional" json:"designFileS3Url" yaml:"designFileS3Url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/resiliencehubv2_service#eks Resiliencehubv2Service#eks}.
	Eks *Resiliencehubv2ServiceInputSourcesResourceConfigurationEks `field:"optional" json:"eks" yaml:"eks"`
	// Resource tags to discover resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/resiliencehubv2_service#resource_tags Resiliencehubv2Service#resource_tags}
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// URL of a Terraform state file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/resiliencehubv2_service#tf_state_file_url Resiliencehubv2Service#tf_state_file_url}
	TfStateFileUrl *string `field:"optional" json:"tfStateFileUrl" yaml:"tfStateFileUrl"`
}

