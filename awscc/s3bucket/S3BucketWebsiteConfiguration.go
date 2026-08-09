// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketWebsiteConfiguration struct {
	// The name of the error document for the website.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_bucket#error_document S3Bucket#error_document}
	ErrorDocument *string `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// The name of the index document for the website.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_bucket#index_document S3Bucket#index_document}
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// The redirect behavior for every request to this bucket's website endpoint.
	//
	// If you specify this property, you can't specify any other property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_bucket#redirect_all_requests_to S3Bucket#redirect_all_requests_to}
	RedirectAllRequestsTo *S3BucketWebsiteConfigurationRedirectAllRequestsTo `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// Rules that define when a redirect is applied and the redirect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_bucket#routing_rules S3Bucket#routing_rules}
	RoutingRules interface{} `field:"optional" json:"routingRules" yaml:"routingRules"`
}

