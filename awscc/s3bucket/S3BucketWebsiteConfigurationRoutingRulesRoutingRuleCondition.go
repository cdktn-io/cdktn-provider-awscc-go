// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketWebsiteConfigurationRoutingRulesRoutingRuleCondition struct {
	// The HTTP error code when the redirect is applied.
	//
	// In the event of an error, if the error code equals this value, then the specified redirect is applied.
	//  Required when parent element ``Condition`` is specified and sibling ``KeyPrefixEquals`` is not specified. If both are specified, then both must be true for the redirect to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#http_error_code_returned_equals S3Bucket#http_error_code_returned_equals}
	HttpErrorCodeReturnedEquals *string `field:"optional" json:"httpErrorCodeReturnedEquals" yaml:"httpErrorCodeReturnedEquals"`
	// The object key name prefix when the redirect is applied.
	//
	// For example, to redirect requests for ``ExamplePage.html``, the key prefix will be ``ExamplePage.html``. To redirect request for all pages with the prefix ``docs/``, the key prefix will be ``/docs``, which identifies all objects in the docs/ folder.
	//  Required when the parent element ``Condition`` is specified and sibling ``HttpErrorCodeReturnedEquals`` is not specified. If both conditions are specified, both must be true for the redirect to be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#key_prefix_equals S3Bucket#key_prefix_equals}
	KeyPrefixEquals *string `field:"optional" json:"keyPrefixEquals" yaml:"keyPrefixEquals"`
}

