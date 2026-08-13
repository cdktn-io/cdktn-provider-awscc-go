// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsCustomExtensions struct {
	// Specifies the critical flag of the X.509 extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/acmpca_certificate#critical AcmpcaCertificate#critical}
	Critical interface{} `field:"optional" json:"critical" yaml:"critical"`
	// Specifies the object identifier (OID) of the X.509 extension. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/acmpca_certificate#object_identifier AcmpcaCertificate#object_identifier}
	ObjectIdentifier *string `field:"optional" json:"objectIdentifier" yaml:"objectIdentifier"`
	// Specifies the base64-encoded value of the X.509 extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/acmpca_certificate#value AcmpcaCertificate#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

