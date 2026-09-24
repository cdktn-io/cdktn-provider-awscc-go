// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueuserdefinedfunction


type GlueUserDefinedFunctionResourceUris struct {
	// The type of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_user_defined_function#resource_type GlueUserDefinedFunction#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The URI for accessing the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_user_defined_function#uri GlueUserDefinedFunction#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

