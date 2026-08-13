// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectpredefinedattribute

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectPredefinedAttributeConfig struct {
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
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_predefined_attribute#instance_arn ConnectPredefinedAttribute#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the predefined attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_predefined_attribute#name ConnectPredefinedAttribute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Custom metadata associated to a Predefined attribute that controls how the attribute behaves when used by upstream services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_predefined_attribute#attribute_configuration ConnectPredefinedAttribute#attribute_configuration}
	AttributeConfiguration *ConnectPredefinedAttributeAttributeConfiguration `field:"optional" json:"attributeConfiguration" yaml:"attributeConfiguration"`
	// The assigned purposes of the predefined attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_predefined_attribute#purposes ConnectPredefinedAttribute#purposes}
	Purposes *[]*string `field:"optional" json:"purposes" yaml:"purposes"`
	// The values of a predefined attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/connect_predefined_attribute#values ConnectPredefinedAttribute#values}
	Values *ConnectPredefinedAttributeValues `field:"optional" json:"values" yaml:"values"`
}

