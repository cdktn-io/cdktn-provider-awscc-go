// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eksaddon

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EksAddonConfig struct {
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
	// Name of Addon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_addon#addon_name EksAddon#addon_name}
	AddonName *string `field:"required" json:"addonName" yaml:"addonName"`
	// Name of Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_addon#cluster_name EksAddon#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Version of Addon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_addon#addon_version EksAddon#addon_version}
	AddonVersion *string `field:"optional" json:"addonVersion" yaml:"addonVersion"`
	// The configuration values to use with the add-on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_addon#configuration_values EksAddon#configuration_values}
	ConfigurationValues *string `field:"optional" json:"configurationValues" yaml:"configurationValues"`
	// The custom namespace configuration to use with the add-on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_addon#namespace_config EksAddon#namespace_config}
	NamespaceConfig *EksAddonNamespaceConfig `field:"optional" json:"namespaceConfig" yaml:"namespaceConfig"`
	// An array of pod identities to apply to this add-on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_addon#pod_identity_associations EksAddon#pod_identity_associations}
	PodIdentityAssociations interface{} `field:"optional" json:"podIdentityAssociations" yaml:"podIdentityAssociations"`
	// PreserveOnDelete parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_addon#preserve_on_delete EksAddon#preserve_on_delete}
	PreserveOnDelete interface{} `field:"optional" json:"preserveOnDelete" yaml:"preserveOnDelete"`
	// Resolve parameter value conflicts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_addon#resolve_conflicts EksAddon#resolve_conflicts}
	ResolveConflicts *string `field:"optional" json:"resolveConflicts" yaml:"resolveConflicts"`
	// IAM role to bind to the add-on's service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_addon#service_account_role_arn EksAddon#service_account_role_arn}
	ServiceAccountRoleArn *string `field:"optional" json:"serviceAccountRoleArn" yaml:"serviceAccountRoleArn"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_addon#tags EksAddon#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

