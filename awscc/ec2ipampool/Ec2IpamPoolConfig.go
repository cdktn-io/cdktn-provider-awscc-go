// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipampool

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2IpamPoolConfig struct {
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
	// The address family of the address space in this pool. Either IPv4 or IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#address_family Ec2IpamPool#address_family}
	AddressFamily *string `field:"required" json:"addressFamily" yaml:"addressFamily"`
	// The Id of the scope this pool is a part of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#ipam_scope_id Ec2IpamPool#ipam_scope_id}
	IpamScopeId *string `field:"required" json:"ipamScopeId" yaml:"ipamScopeId"`
	// The default netmask length for allocations made from this pool.
	//
	// This value is used when the netmask length of an allocation isn't specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#allocation_default_netmask_length Ec2IpamPool#allocation_default_netmask_length}
	AllocationDefaultNetmaskLength *float64 `field:"optional" json:"allocationDefaultNetmaskLength" yaml:"allocationDefaultNetmaskLength"`
	// The maximum allowed netmask length for allocations made from this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#allocation_max_netmask_length Ec2IpamPool#allocation_max_netmask_length}
	AllocationMaxNetmaskLength *float64 `field:"optional" json:"allocationMaxNetmaskLength" yaml:"allocationMaxNetmaskLength"`
	// The minimum allowed netmask length for allocations made from this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#allocation_min_netmask_length Ec2IpamPool#allocation_min_netmask_length}
	AllocationMinNetmaskLength *float64 `field:"optional" json:"allocationMinNetmaskLength" yaml:"allocationMinNetmaskLength"`
	// When specified, an allocation will not be allowed unless a resource has a matching set of tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#allocation_resource_tags Ec2IpamPool#allocation_resource_tags}
	AllocationResourceTags interface{} `field:"optional" json:"allocationResourceTags" yaml:"allocationResourceTags"`
	// Determines what to do if IPAM discovers resources that haven't been assigned an allocation.
	//
	// If set to true, an allocation will be made automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#auto_import Ec2IpamPool#auto_import}
	AutoImport interface{} `field:"optional" json:"autoImport" yaml:"autoImport"`
	// Limits which service in Amazon Web Services that the pool can be used in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#aws_service Ec2IpamPool#aws_service}
	AwsService *string `field:"optional" json:"awsService" yaml:"awsService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#description Ec2IpamPool#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The region of this pool.
	//
	// If not set, this will default to "None" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#locale Ec2IpamPool#locale}
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// A list of cidrs representing the address space available for allocation in this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#provisioned_cidrs Ec2IpamPool#provisioned_cidrs}
	ProvisionedCidrs interface{} `field:"optional" json:"provisionedCidrs" yaml:"provisionedCidrs"`
	// The IP address source for pools in the public scope.
	//
	// Only used for provisioning IP address CIDRs to pools in the public scope. Default is `byoip`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#public_ip_source Ec2IpamPool#public_ip_source}
	PublicIpSource *string `field:"optional" json:"publicIpSource" yaml:"publicIpSource"`
	// Determines whether or not address space from this pool is publicly advertised.
	//
	// Must be set if and only if the pool is IPv6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#publicly_advertisable Ec2IpamPool#publicly_advertisable}
	PubliclyAdvertisable interface{} `field:"optional" json:"publiclyAdvertisable" yaml:"publiclyAdvertisable"`
	// The Id of this pool's source.
	//
	// If set, all space provisioned in this pool must be free space provisioned in the parent pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#source_ipam_pool_id Ec2IpamPool#source_ipam_pool_id}
	SourceIpamPoolId *string `field:"optional" json:"sourceIpamPoolId" yaml:"sourceIpamPoolId"`
	// The resource associated with this pool's space.
	//
	// Depending on the ResourceType, setting a SourceResource changes which space can be provisioned in this pool and which types of resources can receive allocations
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#source_resource Ec2IpamPool#source_resource}
	SourceResource *Ec2IpamPoolSourceResource `field:"optional" json:"sourceResource" yaml:"sourceResource"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_pool#tags Ec2IpamPool#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

