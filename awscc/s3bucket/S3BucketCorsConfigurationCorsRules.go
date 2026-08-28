// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketCorsConfigurationCorsRules struct {
	// Headers that are specified in the ``Access-Control-Request-Headers`` header.
	//
	// These headers are allowed in a preflight OPTIONS request. In response to any preflight OPTIONS request, Amazon S3 returns any requested headers that are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#allowed_headers S3Bucket#allowed_headers}
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// An HTTP method that you allow the origin to run.
	//
	// *Allowed values*: ``GET`` | ``PUT`` | ``HEAD`` | ``POST`` | ``DELETE``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#allowed_methods S3Bucket#allowed_methods}
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// One or more origins you want customers to be able to access the bucket from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#allowed_origins S3Bucket#allowed_origins}
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript ``XMLHttpRequest`` object).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#exposed_headers S3Bucket#exposed_headers}
	ExposedHeaders *[]*string `field:"optional" json:"exposedHeaders" yaml:"exposedHeaders"`
	// A unique identifier for this rule. The value must be no more than 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#id S3Bucket#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_bucket#max_age S3Bucket#max_age}
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

