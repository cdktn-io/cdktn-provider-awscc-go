// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueidentitycenterconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueIdentityCenterConfigurationConfig struct {
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
	// The IAM identity center instance arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_identity_center_configuration#instance_arn GlueIdentityCenterConfiguration#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The downstream scopes that Glue identity center configuration can access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_identity_center_configuration#scopes GlueIdentityCenterConfiguration#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Enable or disable user background sessions for Glue Identity Center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_identity_center_configuration#user_background_sessions_enabled GlueIdentityCenterConfiguration#user_background_sessions_enabled}
	UserBackgroundSessionsEnabled interface{} `field:"optional" json:"userBackgroundSessionsEnabled" yaml:"userBackgroundSessionsEnabled"`
}

