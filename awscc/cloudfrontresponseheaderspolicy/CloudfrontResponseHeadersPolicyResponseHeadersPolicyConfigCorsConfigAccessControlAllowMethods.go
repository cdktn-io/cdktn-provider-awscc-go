// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowMethods struct {
	// The list of HTTP methods.
	//
	// Valid values are:
	//   +   ``GET``
	//   +   ``DELETE``
	//   +   ``HEAD``
	//   +   ``OPTIONS``
	//   +   ``PATCH``
	//   +   ``POST``
	//   +   ``PUT``
	//   +   ``ALL``
	//
	//  ``ALL`` is a special value that includes all of the listed HTTP methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cloudfront_response_headers_policy#items CloudfrontResponseHeadersPolicy#items}
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
}

