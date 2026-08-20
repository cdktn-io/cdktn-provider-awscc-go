// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscapability


type EksCapabilityConfigurationArgoCd struct {
	// Configuration for integrating Argo CD with IAM Identity Center.
	//
	// This allows you to use your organization's identity provider for authentication to Argo CD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_capability#aws_idc EksCapability#aws_idc}
	AwsIdc *EksCapabilityConfigurationArgoCdAwsIdc `field:"optional" json:"awsIdc" yaml:"awsIdc"`
	// The Kubernetes namespace where Argo CD resources will be created. If not specified, the default namespace is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_capability#namespace EksCapability#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Configuration for network access to the Argo CD capability's managed API server endpoint.
	//
	// By default, the Argo CD server is accessible via a public endpoint. You can optionally specify one or more VPC endpoint IDs to enable private connectivity from your VPCs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_capability#network_access EksCapability#network_access}
	NetworkAccess *EksCapabilityConfigurationArgoCdNetworkAccess `field:"optional" json:"networkAccess" yaml:"networkAccess"`
	// A list of role mappings that define which IAM Identity Center users or groups have which Argo CD roles.
	//
	// Each mapping associates an Argo CD role (ADMIN, EDITOR, or VIEWER) with one or more IAM Identity Center identities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_capability#rbac_role_mappings EksCapability#rbac_role_mappings}
	RbacRoleMappings interface{} `field:"optional" json:"rbacRoleMappings" yaml:"rbacRoleMappings"`
}

