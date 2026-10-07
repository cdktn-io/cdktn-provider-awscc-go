// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehendentityrecognizer


type ComprehendEntityRecognizerVpcConfig struct {
	// The ID number for a security group on an instance of your private VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer#security_group_ids ComprehendEntityRecognizer#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID for each subnet being used in your private VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer#subnets ComprehendEntityRecognizer#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

