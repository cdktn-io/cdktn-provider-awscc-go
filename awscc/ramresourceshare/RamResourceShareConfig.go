// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ramresourceshare

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RamResourceShareConfig struct {
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
	// Specifies the name of the resource share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ram_resource_share#name RamResourceShare#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies whether principals outside your organization in AWS Organizations can be associated with a resource share.
	//
	// A value of `true` lets you share with individual AWS accounts that are not in your organization. A value of `false` only has meaning if your account is a member of an AWS Organization. The default value is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ram_resource_share#allow_external_principals RamResourceShare#allow_external_principals}
	AllowExternalPrincipals interface{} `field:"optional" json:"allowExternalPrincipals" yaml:"allowExternalPrincipals"`
	// Specifies the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the AWS RAM permission to associate with the resource share. If you do not specify an ARN for the permission, AWS RAM automatically attaches the default version of the permission for each resource type. You can associate only one permission with each resource type included in the resource share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ram_resource_share#permission_arns RamResourceShare#permission_arns}
	PermissionArns *[]*string `field:"optional" json:"permissionArns" yaml:"permissionArns"`
	// Specifies the principals to associate with the resource share. The possible values are:.
	//
	// - An AWS account ID
	//
	// - An Amazon Resource Name (ARN) of an organization in AWS Organizations
	//
	// - An ARN of an organizational unit (OU) in AWS Organizations
	//
	// - An ARN of an IAM role
	//
	// - An ARN of an IAM user
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ram_resource_share#principals RamResourceShare#principals}
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// Specifies a list of one or more ARNs of the resources to associate with the resource share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ram_resource_share#resource_arns RamResourceShare#resource_arns}
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// Specifies the configuration for the resource share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ram_resource_share#resource_share_configuration RamResourceShare#resource_share_configuration}
	ResourceShareConfiguration *RamResourceShareResourceShareConfiguration `field:"optional" json:"resourceShareConfiguration" yaml:"resourceShareConfiguration"`
	// Specifies from which source accounts the service principal has access to the resources in this resource share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ram_resource_share#sources RamResourceShare#sources}
	Sources *[]*string `field:"optional" json:"sources" yaml:"sources"`
	// Specifies one or more tags to attach to the resource share itself.
	//
	// It doesn't attach the tags to the resources associated with the resource share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ram_resource_share#tags RamResourceShare#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

