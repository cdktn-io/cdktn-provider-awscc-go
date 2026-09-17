// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketWebsiteConfigurationRoutingRulesRedirectRule struct {
	// The host name to use in the redirect request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#host_name S3Bucket#host_name}
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// The HTTP redirect code to use on the response. Not required if one of the siblings is present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#http_redirect_code S3Bucket#http_redirect_code}
	HttpRedirectCode *string `field:"optional" json:"httpRedirectCode" yaml:"httpRedirectCode"`
	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#protocol S3Bucket#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The object key prefix to use in the redirect request.
	//
	// For example, to redirect requests for all pages with prefix ``docs/`` (objects in the ``docs/`` folder) to ``documents/``, you can set a condition block with ``KeyPrefixEquals`` set to ``docs/`` and in the Redirect set ``ReplaceKeyPrefixWith`` to ``/documents``. Not required if one of the siblings is present. Can be present only if ``ReplaceKeyWith`` is not provided.
	//   Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#replace_key_prefix_with S3Bucket#replace_key_prefix_with}
	ReplaceKeyPrefixWith *string `field:"optional" json:"replaceKeyPrefixWith" yaml:"replaceKeyPrefixWith"`
	// The specific object key to use in the redirect request.
	//
	// For example, redirect request to ``error.html``. Not required if one of the siblings is present. Can be present only if ``ReplaceKeyPrefixWith`` is not provided.
	//   Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3_bucket#replace_key_with S3Bucket#replace_key_with}
	ReplaceKeyWith *string `field:"optional" json:"replaceKeyWith" yaml:"replaceKeyWith"`
}

