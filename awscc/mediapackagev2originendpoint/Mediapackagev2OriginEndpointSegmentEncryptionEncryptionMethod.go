// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2originendpoint


type Mediapackagev2OriginEndpointSegmentEncryptionEncryptionMethod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackagev2_origin_endpoint#cmaf_encryption_method Mediapackagev2OriginEndpoint#cmaf_encryption_method}.
	CmafEncryptionMethod *string `field:"optional" json:"cmafEncryptionMethod" yaml:"cmafEncryptionMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackagev2_origin_endpoint#ism_encryption_method Mediapackagev2OriginEndpoint#ism_encryption_method}.
	IsmEncryptionMethod *string `field:"optional" json:"ismEncryptionMethod" yaml:"ismEncryptionMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediapackagev2_origin_endpoint#ts_encryption_method Mediapackagev2OriginEndpoint#ts_encryption_method}.
	TsEncryptionMethod *string `field:"optional" json:"tsEncryptionMethod" yaml:"tsEncryptionMethod"`
}

