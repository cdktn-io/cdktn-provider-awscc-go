// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontkeyvaluestore


type CloudfrontKeyValueStoreTags struct {
	// A string that contains ``Tag`` key.
	//
	// The string length should be between 1 and 128 characters. Valid characters include ``a-z``, ``A-Z``, ``0-9``, space, and the special characters ``_ - . : / = +
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A string that contains an optional ``Tag`` value.
	//
	// The string length should be between 0 and 256 characters. Valid characters include ``a-z``, ``A-Z``, ``0-9``, space, and the special characters ``_ - . : / = +
	Value *string `field:"optional" json:"value" yaml:"value"`
}

