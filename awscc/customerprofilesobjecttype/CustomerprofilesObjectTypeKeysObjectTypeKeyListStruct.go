// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesobjecttype


type CustomerprofilesObjectTypeKeysObjectTypeKeyListStruct struct {
	// The reference for the key name of the fields map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/customerprofiles_object_type#field_names CustomerprofilesObjectType#field_names}
	FieldNames *[]*string `field:"optional" json:"fieldNames" yaml:"fieldNames"`
	// The types of keys that a ProfileObject can have.
	//
	// Each ProfileObject can have only 1 UNIQUE key but multiple PROFILE keys. PROFILE means that this key can be used to tie an object to a PROFILE. UNIQUE means that it can be used to uniquely identify an object. If a key a is marked as SECONDARY, it will be used to search for profiles after all other PROFILE keys have been searched. A LOOKUP_ONLY key is only used to match a profile but is not persisted to be used for searching of the profile. A NEW_ONLY key is only used if the profile does not already exist before the object is ingested, otherwise it is only used for matching objects to profiles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/customerprofiles_object_type#standard_identifiers CustomerprofilesObjectType#standard_identifiers}
	StandardIdentifiers *[]*string `field:"optional" json:"standardIdentifiers" yaml:"standardIdentifiers"`
}

