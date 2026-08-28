// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmlconfiguredmodelalgorithmassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsmlConfiguredModelAlgorithmAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm_association#configured_model_algorithm_arn CleanroomsmlConfiguredModelAlgorithmAssociation#configured_model_algorithm_arn}.
	ConfiguredModelAlgorithmArn *string `field:"required" json:"configuredModelAlgorithmArn" yaml:"configuredModelAlgorithmArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm_association#membership_identifier CleanroomsmlConfiguredModelAlgorithmAssociation#membership_identifier}.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm_association#name CleanroomsmlConfiguredModelAlgorithmAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm_association#description CleanroomsmlConfiguredModelAlgorithmAssociation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm_association#privacy_configuration CleanroomsmlConfiguredModelAlgorithmAssociation#privacy_configuration}.
	PrivacyConfiguration *CleanroomsmlConfiguredModelAlgorithmAssociationPrivacyConfiguration `field:"optional" json:"privacyConfiguration" yaml:"privacyConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms-ml configured model algorithm association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanroomsml_configured_model_algorithm_association#tags CleanroomsmlConfiguredModelAlgorithmAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

