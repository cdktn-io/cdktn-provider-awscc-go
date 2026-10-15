// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationSlackBidirectional struct {
	// Whether bidirectional communication is enabled for this association.
	//
	// When set to true, you can mention the agent in a configured Slack channel and it responds in that channel. When omitted or set to false, the agent ignores mentions and only sends notifications
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/devopsagent_association#enabled DevopsagentAssociation#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// IAM role ARN that AWS DevOps Agent assumes to exchange messages with your Slack workspace on behalf of this association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/devopsagent_association#role_arn DevopsagentAssociation#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

