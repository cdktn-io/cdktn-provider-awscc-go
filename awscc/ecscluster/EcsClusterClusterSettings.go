// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecscluster


type EcsClusterClusterSettings struct {
	// The name of the cluster setting. The value is ``containerInsights``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_cluster#name EcsCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value to set for the cluster setting.
	//
	// The supported values are ``enhanced``, ``enabled``, and ``disabled``.
	//  To use Container Insights with enhanced observability, set the ``containerInsights`` account setting to ``enhanced``.
	//  To use Container Insights, set the ``containerInsights`` account setting to ``enabled``.
	//  If a cluster value is specified, it will override the ``containerInsights`` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ecs_cluster#value EcsCluster#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

