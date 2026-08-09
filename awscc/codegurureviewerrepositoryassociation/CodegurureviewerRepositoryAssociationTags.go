// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codegurureviewerrepositoryassociation


type CodegurureviewerRepositoryAssociationTags struct {
	// The key name of the tag.
	//
	// You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : /
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : /
	Value *string `field:"optional" json:"value" yaml:"value"`
}

