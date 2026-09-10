// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmparameter

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmParameterConfig struct {
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
	// The type of parameter.   Parameters of type ``SecureString`` are not supported by CFNlong.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_parameter#type SsmParameter#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The parameter value.
	//
	// If type is ``StringList``, the system returns a comma-separated string with no spaces between commas in the ``Value`` field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_parameter#value SsmParameter#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// A regular expression used to validate the parameter value.
	//
	// For example, for ``String`` types with values restricted to numbers, you can specify the following: ``AllowedPattern=^\d+$``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_parameter#allowed_pattern SsmParameter#allowed_pattern}
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// The data type of the parameter, such as ``text`` or ``aws:ec2:image``. The default is ``text``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_parameter#data_type SsmParameter#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Information about the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_parameter#description SsmParameter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	//
	// The reported maximum length of 2048 characters for a parameter name includes 1037 characters that are reserved for internal use by SYS. The maximum length for a parameter name that you specify is 1011 characters.
	//  This count of 1011 characters includes the characters in the ARN that precede the name you specify. This ARN length will vary depending on your partition and Region. For example, the following 45 characters count toward the 1011 character maximum for a parameter created in the US East (Ohio) Region: ``arn:aws:ssm:us-east-2:111122223333:parameter/``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_parameter#name SsmParameter#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Information about the policies assigned to a parameter.  [Assigning parameter policies](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-policies.html) in the *User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_parameter#policies SsmParameter#policies}
	Policies *string `field:"optional" json:"policies" yaml:"policies"`
	// Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs).
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a SYS parameter to identify the type of resource to which it applies, the environment, or the type of configuration data referenced by the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_parameter#tags SsmParameter#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The parameter tier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssm_parameter#tier SsmParameter#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

