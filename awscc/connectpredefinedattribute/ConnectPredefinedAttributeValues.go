// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectpredefinedattribute


type ConnectPredefinedAttributeValues struct {
	// Predefined attribute values of type string list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connect_predefined_attribute#string_list ConnectPredefinedAttribute#string_list}
	StringList *[]*string `field:"optional" json:"stringList" yaml:"stringList"`
}

