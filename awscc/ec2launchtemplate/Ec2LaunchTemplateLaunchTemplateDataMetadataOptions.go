// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataMetadataOptions struct {
	// Enables or disables the HTTP metadata endpoint on your instances.
	//
	// If the parameter is not specified, the default state is ``enabled``.
	//   If you specify a value of ``disabled``, you will not be able to access your instance metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_launch_template#http_endpoint Ec2LaunchTemplate#http_endpoint}
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Enables or disables the IPv6 endpoint for the instance metadata service.  Default: ``disabled``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_launch_template#http_protocol_ipv_6 Ec2LaunchTemplate#http_protocol_ipv_6}
	HttpProtocolIpv6 *string `field:"optional" json:"httpProtocolIpv6" yaml:"httpProtocolIpv6"`
	// The desired HTTP PUT response hop limit for instance metadata requests.
	//
	// The larger the number, the further instance metadata requests can travel.
	//  Default: ``1``
	//  Possible values: Integers from 1 to 64
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_launch_template#http_put_response_hop_limit Ec2LaunchTemplate#http_put_response_hop_limit}
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Indicates whether IMDSv2 is required.
	//
	// +   ``optional`` - IMDSv2 is optional. You can choose whether to send a session token in your instance metadata retrieval requests. If you retrieve IAM role credentials without a session token, you receive the IMDSv1 role credentials. If you retrieve IAM role credentials using a valid session token, you receive the IMDSv2 role credentials.
	//   +   ``required`` - IMDSv2 is required. You must send a session token in your instance metadata retrieval requests. With this option, retrieving the IAM role credentials always returns IMDSv2 credentials; IMDSv1 credentials are not available.
	//
	//  Default: If the value of ``ImdsSupport`` for the Amazon Machine Image (AMI) for your instance is ``v2.0``, the default is ``required``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_launch_template#http_tokens Ec2LaunchTemplate#http_tokens}
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
	// Set to ``enabled`` to allow access to instance tags from the instance metadata.
	//
	// Set to ``disabled`` to turn off access to instance tags from the instance metadata. For more information, see [Work with instance tags using the instance metadata](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#work-with-tags-in-IMDS).
	//  Default: ``disabled``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_launch_template#instance_metadata_tags Ec2LaunchTemplate#instance_metadata_tags}
	InstanceMetadataTags *string `field:"optional" json:"instanceMetadataTags" yaml:"instanceMetadataTags"`
}

