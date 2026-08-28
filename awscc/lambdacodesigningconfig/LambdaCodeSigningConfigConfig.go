// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdacodesigningconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaCodeSigningConfigConfig struct {
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
	// When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_code_signing_config#allowed_publishers LambdaCodeSigningConfig#allowed_publishers}
	AllowedPublishers *LambdaCodeSigningConfigAllowedPublishers `field:"required" json:"allowedPublishers" yaml:"allowedPublishers"`
	// Policies to control how to act if a signature is invalid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_code_signing_config#code_signing_policies LambdaCodeSigningConfig#code_signing_policies}
	CodeSigningPolicies *LambdaCodeSigningConfigCodeSigningPolicies `field:"optional" json:"codeSigningPolicies" yaml:"codeSigningPolicies"`
	// A description of the CodeSigningConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_code_signing_config#description LambdaCodeSigningConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of tags to apply to CodeSigningConfig resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lambda_code_signing_config#tags LambdaCodeSigningConfig#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

