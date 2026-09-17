// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueuserdefinedfunction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueUserDefinedFunctionConfig struct {
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
	// The name of the catalog database in which the function is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_user_defined_function#database_name GlueUserDefinedFunction#database_name}
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The name of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_user_defined_function#function_name GlueUserDefinedFunction#function_name}
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The Java class that contains the function code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_user_defined_function#class_name GlueUserDefinedFunction#class_name}
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// The type of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_user_defined_function#function_type GlueUserDefinedFunction#function_type}
	FunctionType *string `field:"optional" json:"functionType" yaml:"functionType"`
	// The owner of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_user_defined_function#owner_name GlueUserDefinedFunction#owner_name}
	OwnerName *string `field:"optional" json:"ownerName" yaml:"ownerName"`
	// The owner type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_user_defined_function#owner_type GlueUserDefinedFunction#owner_type}
	OwnerType *string `field:"optional" json:"ownerType" yaml:"ownerType"`
	// The resource URIs for the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_user_defined_function#resource_uris GlueUserDefinedFunction#resource_uris}
	ResourceUris interface{} `field:"optional" json:"resourceUris" yaml:"resourceUris"`
}

