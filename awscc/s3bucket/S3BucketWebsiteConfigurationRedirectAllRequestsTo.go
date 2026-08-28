// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketWebsiteConfigurationRedirectAllRequestsTo struct {
	// Name of the host where requests are redirected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#host_name S3Bucket#host_name}
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#protocol S3Bucket#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

