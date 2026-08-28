// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsconfiguredtableassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsConfiguredTableAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table_association#configured_table_identifier CleanroomsConfiguredTableAssociation#configured_table_identifier}.
	ConfiguredTableIdentifier *string `field:"required" json:"configuredTableIdentifier" yaml:"configuredTableIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table_association#membership_identifier CleanroomsConfiguredTableAssociation#membership_identifier}.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table_association#name CleanroomsConfiguredTableAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table_association#role_arn CleanroomsConfiguredTableAssociation#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table_association#configured_table_association_analysis_rules CleanroomsConfiguredTableAssociation#configured_table_association_analysis_rules}.
	ConfiguredTableAssociationAnalysisRules interface{} `field:"optional" json:"configuredTableAssociationAnalysisRules" yaml:"configuredTableAssociationAnalysisRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table_association#description CleanroomsConfiguredTableAssociation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_configured_table_association#tags CleanroomsConfiguredTableAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

