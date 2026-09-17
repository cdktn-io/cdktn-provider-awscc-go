// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdalayerversionpermission

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaLayerVersionPermissionConfig struct {
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
	// The API action that grants access to the layer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_layer_version_permission#action LambdaLayerVersionPermission#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_layer_version_permission#layer_version_arn LambdaLayerVersionPermission#layer_version_arn}
	LayerVersionArn *string `field:"required" json:"layerVersionArn" yaml:"layerVersionArn"`
	// An account ID, or * to grant layer usage permission to all accounts in an organization, or all AWS accounts (if organizationId is not specified).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_layer_version_permission#principal LambdaLayerVersionPermission#principal}
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// With the principal set to *, grant permission to all accounts in the specified organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_layer_version_permission#organization_id LambdaLayerVersionPermission#organization_id}
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
}

