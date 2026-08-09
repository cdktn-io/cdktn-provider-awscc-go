// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3vectorsindex


type S3VectorsIndexTags struct {
	// Tag key must be between 1 to 128 characters in length.
	//
	// Tag key cannot start with 'aws:' and can only contain alphanumeric characters, spaces, _, ., /, =, +, -, and
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Tag value must be between 0 to 256 characters in length.
	//
	// Tag value can only contain alphanumeric characters, spaces, _, ., /, =, +, -, and
	Value *string `field:"optional" json:"value" yaml:"value"`
}

