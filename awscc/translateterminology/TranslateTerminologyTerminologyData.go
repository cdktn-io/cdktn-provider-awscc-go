// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package translateterminology


type TranslateTerminologyTerminologyData struct {
	// The directionality of the terminology resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/translate_terminology#directionality TranslateTerminology#directionality}
	Directionality *string `field:"optional" json:"directionality" yaml:"directionality"`
	// The file containing the custom terminology data, base64-encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/translate_terminology#file TranslateTerminology#file}
	File *string `field:"optional" json:"file" yaml:"file"`
	// The data format of the custom terminology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/translate_terminology#format TranslateTerminology#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
}

