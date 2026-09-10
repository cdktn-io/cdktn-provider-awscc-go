// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekspodidentityassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EksPodIdentityAssociationConfig struct {
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
	// The cluster that the pod identity association is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_pod_identity_association#cluster_name EksPodIdentityAssociation#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The Kubernetes namespace that the pod identity association is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_pod_identity_association#namespace EksPodIdentityAssociation#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The IAM role ARN that the pod identity association is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_pod_identity_association#role_arn EksPodIdentityAssociation#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Kubernetes service account that the pod identity association is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_pod_identity_association#service_account EksPodIdentityAssociation#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
	// The Disable Session Tags of the pod identity association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_pod_identity_association#disable_session_tags EksPodIdentityAssociation#disable_session_tags}
	DisableSessionTags interface{} `field:"optional" json:"disableSessionTags" yaml:"disableSessionTags"`
	// The policy of the pod identity association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_pod_identity_association#policy EksPodIdentityAssociation#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_pod_identity_association#tags EksPodIdentityAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The Target Role Arn of the pod identity association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_pod_identity_association#target_role_arn EksPodIdentityAssociation#target_role_arn}
	TargetRoleArn *string `field:"optional" json:"targetRoleArn" yaml:"targetRoleArn"`
}

