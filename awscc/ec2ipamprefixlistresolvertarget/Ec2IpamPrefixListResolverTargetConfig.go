// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipamprefixlistresolvertarget

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2IpamPrefixListResolverTargetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Id of the IPAM Prefix List Resolver associated with this Target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ipam_prefix_list_resolver_target#ipam_prefix_list_resolver_id Ec2IpamPrefixListResolverTarget#ipam_prefix_list_resolver_id}
	IpamPrefixListResolverId *string `field:"required" json:"ipamPrefixListResolverId" yaml:"ipamPrefixListResolverId"`
	// The Id of the Managed Prefix List.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ipam_prefix_list_resolver_target#prefix_list_id Ec2IpamPrefixListResolverTarget#prefix_list_id}
	PrefixListId *string `field:"required" json:"prefixListId" yaml:"prefixListId"`
	// The region that the Managed Prefix List is located in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ipam_prefix_list_resolver_target#prefix_list_region Ec2IpamPrefixListResolverTarget#prefix_list_region}
	PrefixListRegion *string `field:"required" json:"prefixListRegion" yaml:"prefixListRegion"`
	// Indicates whether this Target automatically tracks the latest version of the Prefix List Resolver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ipam_prefix_list_resolver_target#track_latest_version Ec2IpamPrefixListResolverTarget#track_latest_version}
	TrackLatestVersion interface{} `field:"required" json:"trackLatestVersion" yaml:"trackLatestVersion"`
	// The desired version of the Prefix List Resolver that this Target should synchronize with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ipam_prefix_list_resolver_target#desired_version Ec2IpamPrefixListResolverTarget#desired_version}
	DesiredVersion *float64 `field:"optional" json:"desiredVersion" yaml:"desiredVersion"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ec2_ipam_prefix_list_resolver_target#tags Ec2IpamPrefixListResolverTarget#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

