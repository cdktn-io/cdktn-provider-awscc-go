// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package protonenvironmentaccountconnection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ProtonEnvironmentAccountConnectionConfig struct {
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
	// The Amazon Resource Name (ARN) of an IAM service role in the environment account.
	//
	// AWS Proton uses this role to provision infrastructure resources using CodeBuild-based provisioning in the associated environment account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/proton_environment_account_connection#codebuild_role_arn ProtonEnvironmentAccountConnection#codebuild_role_arn}
	CodebuildRoleArn *string `field:"optional" json:"codebuildRoleArn" yaml:"codebuildRoleArn"`
	// The Amazon Resource Name (ARN) of the IAM service role that AWS Proton uses when provisioning directly defined components in the associated environment account.
	//
	// It determines the scope of infrastructure that a component can provision in the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/proton_environment_account_connection#component_role_arn ProtonEnvironmentAccountConnection#component_role_arn}
	ComponentRoleArn *string `field:"optional" json:"componentRoleArn" yaml:"componentRoleArn"`
	// The environment account that's connected to the environment account connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/proton_environment_account_connection#environment_account_id ProtonEnvironmentAccountConnection#environment_account_id}
	EnvironmentAccountId *string `field:"optional" json:"environmentAccountId" yaml:"environmentAccountId"`
	// The name of the AWS Proton environment that's created in the associated management account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/proton_environment_account_connection#environment_name ProtonEnvironmentAccountConnection#environment_name}
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// The ID of the management account that accepts or rejects the environment account connection.
	//
	// You create an manage the AWS Proton environment in this account. If the management account accepts the environment account connection, AWS Proton can use the associated IAM role to provision environment infrastructure resources in the associated environment account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/proton_environment_account_connection#management_account_id ProtonEnvironmentAccountConnection#management_account_id}
	ManagementAccountId *string `field:"optional" json:"managementAccountId" yaml:"managementAccountId"`
	// The Amazon Resource Name (ARN) of the IAM service role that's created in the environment account.
	//
	// AWS Proton uses this role to provision infrastructure resources in the associated environment account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/proton_environment_account_connection#role_arn ProtonEnvironmentAccountConnection#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// <p>An optional list of metadata items that you can associate with the Proton environment account connection.
	//
	// A tag is a key-value pair.</p>
	//          <p>For more information, see <a href="https://docs.aws.amazon.com/proton/latest/userguide/resources.html">Proton resources and tagging</a> in the
	//         <i>Proton User Guide</i>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/proton_environment_account_connection#tags ProtonEnvironmentAccountConnection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

