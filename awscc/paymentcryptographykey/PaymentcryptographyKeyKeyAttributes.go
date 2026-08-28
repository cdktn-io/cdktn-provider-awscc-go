// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paymentcryptographykey


type PaymentcryptographyKeyKeyAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/paymentcryptography_key#key_algorithm PaymentcryptographyKey#key_algorithm}.
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/paymentcryptography_key#key_class PaymentcryptographyKey#key_class}.
	KeyClass *string `field:"required" json:"keyClass" yaml:"keyClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/paymentcryptography_key#key_modes_of_use PaymentcryptographyKey#key_modes_of_use}.
	KeyModesOfUse *PaymentcryptographyKeyKeyAttributesKeyModesOfUse `field:"required" json:"keyModesOfUse" yaml:"keyModesOfUse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/paymentcryptography_key#key_usage PaymentcryptographyKey#key_usage}.
	KeyUsage *string `field:"required" json:"keyUsage" yaml:"keyUsage"`
}

