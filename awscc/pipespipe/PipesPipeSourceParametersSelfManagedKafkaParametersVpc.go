// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeSourceParametersSelfManagedKafkaParametersVpc struct {
	// List of SecurityGroupId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pipes_pipe#security_group PipesPipe#security_group}
	SecurityGroup *[]*string `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// List of SubnetId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pipes_pipe#subnets PipesPipe#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

