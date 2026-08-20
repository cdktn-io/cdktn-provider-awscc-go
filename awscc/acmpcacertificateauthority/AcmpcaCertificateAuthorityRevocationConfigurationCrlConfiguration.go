// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package acmpcacertificateauthority


type AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration struct {
	// Configures the default behavior of the CRL Distribution Point extension for certificates issued by your certificate authority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#crl_distribution_point_extension_configuration AcmpcaCertificateAuthority#crl_distribution_point_extension_configuration}
	CrlDistributionPointExtensionConfiguration *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationCrlDistributionPointExtensionConfiguration `field:"optional" json:"crlDistributionPointExtensionConfiguration" yaml:"crlDistributionPointExtensionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#crl_type AcmpcaCertificateAuthority#crl_type}.
	CrlType *string `field:"optional" json:"crlType" yaml:"crlType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#custom_cname AcmpcaCertificateAuthority#custom_cname}.
	CustomCname *string `field:"optional" json:"customCname" yaml:"customCname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#custom_path AcmpcaCertificateAuthority#custom_path}.
	CustomPath *string `field:"optional" json:"customPath" yaml:"customPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#enabled AcmpcaCertificateAuthority#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#expiration_in_days AcmpcaCertificateAuthority#expiration_in_days}.
	ExpirationInDays *float64 `field:"optional" json:"expirationInDays" yaml:"expirationInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#s3_bucket_name AcmpcaCertificateAuthority#s3_bucket_name}.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/acmpca_certificate_authority#s3_object_acl AcmpcaCertificateAuthority#s3_object_acl}.
	S3ObjectAcl *string `field:"optional" json:"s3ObjectAcl" yaml:"s3ObjectAcl"`
}

