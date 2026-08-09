// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscapability

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EksCapabilityConfig struct {
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
	// A unique name for the capability.
	//
	// The name must be unique within your cluster and can contain alphanumeric characters, hyphens, and underscores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/eks_capability#capability_name EksCapability#capability_name}
	CapabilityName *string `field:"required" json:"capabilityName" yaml:"capabilityName"`
	// The name of the EKS cluster where you want to create the capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/eks_capability#cluster_name EksCapability#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Specifies how Kubernetes resources managed by the capability should be handled when the capability is deleted.
	//
	// Currently, the only supported value is RETAIN which retains all Kubernetes resources managed by the capability when the capability is deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/eks_capability#delete_propagation_policy EksCapability#delete_propagation_policy}
	DeletePropagationPolicy *string `field:"required" json:"deletePropagationPolicy" yaml:"deletePropagationPolicy"`
	// The Amazon Resource Name (ARN) of the IAM role that the capability uses to interact with AWS services.
	//
	// This role must have a trust policy that allows the EKS service principal to assume it, and it must have the necessary permissions for the capability type you're creating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/eks_capability#role_arn EksCapability#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The type of capability to create.
	//
	// Valid values are: ACK (AWS Controllers for Kubernetes, which lets you manage AWS resources directly from Kubernetes), ARGOCD (Argo CD for GitOps-based continuous delivery), or KRO (Kube Resource Orchestrator for composing and managing custom Kubernetes resources).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/eks_capability#type EksCapability#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration settings for the capability.
	//
	// The structure of this object varies depending on the capability type. For Argo CD capabilities, you can configure IAM Identity Center integration, RBAC role mappings, and network access settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/eks_capability#configuration EksCapability#configuration}
	Configuration *EksCapabilityConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/eks_capability#tags EksCapability#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

