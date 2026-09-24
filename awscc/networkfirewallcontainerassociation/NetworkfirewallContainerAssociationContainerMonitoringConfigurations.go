// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallcontainerassociation


type NetworkfirewallContainerAssociationContainerMonitoringConfigurations struct {
	// The ARN of the Amazon ECS or Amazon EKS cluster to monitor.
	//
	// The cluster must be in the same Region and account as the container association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/networkfirewall_container_association#cluster_arn NetworkfirewallContainerAssociation#cluster_arn}
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// Key-value pairs that filter which containers are tracked.
	//
	// For Amazon EKS, you can filter by namespace and Kubernetes labels. For Amazon ECS, you can filter by container instance attributes (EC2 launch type only).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/networkfirewall_container_association#attribute_filters NetworkfirewallContainerAssociation#attribute_filters}
	AttributeFilters interface{} `field:"optional" json:"attributeFilters" yaml:"attributeFilters"`
}

