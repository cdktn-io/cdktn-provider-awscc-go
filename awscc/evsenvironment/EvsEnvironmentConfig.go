// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evsenvironment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EvsEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#service_access_subnet_id EvsEnvironment#service_access_subnet_id}.
	ServiceAccessSubnetId *string `field:"required" json:"serviceAccessSubnetId" yaml:"serviceAccessSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#terms_accepted EvsEnvironment#terms_accepted}.
	TermsAccepted interface{} `field:"required" json:"termsAccepted" yaml:"termsAccepted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#vcf_version EvsEnvironment#vcf_version}.
	VcfVersion *string `field:"required" json:"vcfVersion" yaml:"vcfVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#vpc_id EvsEnvironment#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#connectivity_info EvsEnvironment#connectivity_info}.
	ConnectivityInfo *EvsEnvironmentConnectivityInfo `field:"optional" json:"connectivityInfo" yaml:"connectivityInfo"`
	// The name of an EVS environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#environment_name EvsEnvironment#environment_name}
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// The initial hosts for environment only required upon creation. Modification after creation will have no effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#hosts EvsEnvironment#hosts}
	Hosts interface{} `field:"optional" json:"hosts" yaml:"hosts"`
	// The initial Vlan configuration only required upon creation. Modification after creation will have no effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#initial_vlans EvsEnvironment#initial_vlans}
	InitialVlans *EvsEnvironmentInitialVlans `field:"optional" json:"initialVlans" yaml:"initialVlans"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#kms_key_id EvsEnvironment#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The license information for an EVS environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#license_info EvsEnvironment#license_info}
	LicenseInfo *EvsEnvironmentLicenseInfo `field:"optional" json:"licenseInfo" yaml:"licenseInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#service_access_security_groups EvsEnvironment#service_access_security_groups}.
	ServiceAccessSecurityGroups *EvsEnvironmentServiceAccessSecurityGroups `field:"optional" json:"serviceAccessSecurityGroups" yaml:"serviceAccessSecurityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#site_id EvsEnvironment#site_id}.
	SiteId *string `field:"optional" json:"siteId" yaml:"siteId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#tags EvsEnvironment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evs_environment#vcf_hostnames EvsEnvironment#vcf_hostnames}.
	VcfHostnames *EvsEnvironmentVcfHostnames `field:"optional" json:"vcfHostnames" yaml:"vcfHostnames"`
}

