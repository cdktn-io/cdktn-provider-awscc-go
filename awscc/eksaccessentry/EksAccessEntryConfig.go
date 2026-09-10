// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eksaccessentry

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EksAccessEntryConfig struct {
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
	// The cluster that the access entry is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_access_entry#cluster_name EksAccessEntry#cluster_name}
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The principal ARN that the access entry is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_access_entry#principal_arn EksAccessEntry#principal_arn}
	PrincipalArn *string `field:"required" json:"principalArn" yaml:"principalArn"`
	// An array of access policies that are associated with the access entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_access_entry#access_policies EksAccessEntry#access_policies}
	AccessPolicies interface{} `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// The Kubernetes groups that the access entry is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_access_entry#kubernetes_groups EksAccessEntry#kubernetes_groups}
	KubernetesGroups *[]*string `field:"optional" json:"kubernetesGroups" yaml:"kubernetesGroups"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_access_entry#tags EksAccessEntry#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The node type to associate with the access entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_access_entry#type EksAccessEntry#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The Kubernetes user that the access entry is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_access_entry#username EksAccessEntry#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

