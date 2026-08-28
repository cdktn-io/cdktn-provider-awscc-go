// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightTopicConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#aws_account_id QuicksightTopic#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#config_options QuicksightTopic#config_options}.
	ConfigOptions *QuicksightTopicConfigOptions `field:"optional" json:"configOptions" yaml:"configOptions"`
	// <p>Instructions that provide additional guidance and context for response generation.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#custom_instructions QuicksightTopic#custom_instructions}
	CustomInstructions *QuicksightTopicCustomInstructions `field:"optional" json:"customInstructions" yaml:"customInstructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#data_sets QuicksightTopic#data_sets}.
	DataSets interface{} `field:"optional" json:"dataSets" yaml:"dataSets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#description QuicksightTopic#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#folder_arns QuicksightTopic#folder_arns}.
	FolderArns *[]*string `field:"optional" json:"folderArns" yaml:"folderArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#name QuicksightTopic#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#tags QuicksightTopic#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#topic_id QuicksightTopic#topic_id}.
	TopicId *string `field:"optional" json:"topicId" yaml:"topicId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_topic#user_experience_version QuicksightTopic#user_experience_version}.
	UserExperienceVersion *string `field:"optional" json:"userExperienceVersion" yaml:"userExperienceVersion"`
}

