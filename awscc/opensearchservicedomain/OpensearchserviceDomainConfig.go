// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpensearchserviceDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#access_policies OpensearchserviceDomain#access_policies}.
	AccessPolicies *string `field:"optional" json:"accessPolicies" yaml:"accessPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#advanced_options OpensearchserviceDomain#advanced_options}.
	AdvancedOptions *map[string]*string `field:"optional" json:"advancedOptions" yaml:"advancedOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#advanced_security_options OpensearchserviceDomain#advanced_security_options}.
	AdvancedSecurityOptions *OpensearchserviceDomainAdvancedSecurityOptions `field:"optional" json:"advancedSecurityOptions" yaml:"advancedSecurityOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#aiml_options OpensearchserviceDomain#aiml_options}.
	AimlOptions *OpensearchserviceDomainAimlOptions `field:"optional" json:"aimlOptions" yaml:"aimlOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#automated_snapshot_pause_options OpensearchserviceDomain#automated_snapshot_pause_options}.
	AutomatedSnapshotPauseOptions *OpensearchserviceDomainAutomatedSnapshotPauseOptions `field:"optional" json:"automatedSnapshotPauseOptions" yaml:"automatedSnapshotPauseOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#cluster_config OpensearchserviceDomain#cluster_config}.
	ClusterConfig *OpensearchserviceDomainClusterConfig `field:"optional" json:"clusterConfig" yaml:"clusterConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#cognito_options OpensearchserviceDomain#cognito_options}.
	CognitoOptions *OpensearchserviceDomainCognitoOptions `field:"optional" json:"cognitoOptions" yaml:"cognitoOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#deployment_strategy_options OpensearchserviceDomain#deployment_strategy_options}.
	DeploymentStrategyOptions *OpensearchserviceDomainDeploymentStrategyOptions `field:"optional" json:"deploymentStrategyOptions" yaml:"deploymentStrategyOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#domain_endpoint_options OpensearchserviceDomain#domain_endpoint_options}.
	DomainEndpointOptions *OpensearchserviceDomainDomainEndpointOptions `field:"optional" json:"domainEndpointOptions" yaml:"domainEndpointOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#domain_name OpensearchserviceDomain#domain_name}.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#ebs_options OpensearchserviceDomain#ebs_options}.
	EbsOptions *OpensearchserviceDomainEbsOptions `field:"optional" json:"ebsOptions" yaml:"ebsOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#encryption_at_rest_options OpensearchserviceDomain#encryption_at_rest_options}.
	EncryptionAtRestOptions *OpensearchserviceDomainEncryptionAtRestOptions `field:"optional" json:"encryptionAtRestOptions" yaml:"encryptionAtRestOptions"`
	// The engine mode of the domain.
	//
	// Determines whether the domain runs the standard (GENERAL) engine or the optimized multi-engine (OPTIMIZED) engine. This value cannot be changed after the domain is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#engine_mode OpensearchserviceDomain#engine_mode}
	EngineMode *string `field:"optional" json:"engineMode" yaml:"engineMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#engine_version OpensearchserviceDomain#engine_version}.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Options for configuring Identity Center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#identity_center_options OpensearchserviceDomain#identity_center_options}
	IdentityCenterOptions *OpensearchserviceDomainIdentityCenterOptions `field:"optional" json:"identityCenterOptions" yaml:"identityCenterOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#ip_address_type OpensearchserviceDomain#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#log_publishing_options OpensearchserviceDomain#log_publishing_options}.
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#node_to_node_encryption_options OpensearchserviceDomain#node_to_node_encryption_options}.
	NodeToNodeEncryptionOptions *OpensearchserviceDomainNodeToNodeEncryptionOptions `field:"optional" json:"nodeToNodeEncryptionOptions" yaml:"nodeToNodeEncryptionOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#off_peak_window_options OpensearchserviceDomain#off_peak_window_options}.
	OffPeakWindowOptions *OpensearchserviceDomainOffPeakWindowOptions `field:"optional" json:"offPeakWindowOptions" yaml:"offPeakWindowOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#skip_shard_migration_wait OpensearchserviceDomain#skip_shard_migration_wait}.
	SkipShardMigrationWait interface{} `field:"optional" json:"skipShardMigrationWait" yaml:"skipShardMigrationWait"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#snapshot_options OpensearchserviceDomain#snapshot_options}.
	SnapshotOptions *OpensearchserviceDomainSnapshotOptions `field:"optional" json:"snapshotOptions" yaml:"snapshotOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#software_update_options OpensearchserviceDomain#software_update_options}.
	SoftwareUpdateOptions *OpensearchserviceDomainSoftwareUpdateOptions `field:"optional" json:"softwareUpdateOptions" yaml:"softwareUpdateOptions"`
	// An arbitrary set of tags (key-value pairs) for this Domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#tags OpensearchserviceDomain#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The primary use case of the domain.
	//
	// Determines the default configuration tuned for the workload. For GENERAL engine-mode domains, this value can be changed after creation. For OPTIMIZED engine-mode domains, this value cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#use_case OpensearchserviceDomain#use_case}
	UseCase *string `field:"optional" json:"useCase" yaml:"useCase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#vpc_options OpensearchserviceDomain#vpc_options}.
	VpcOptions *OpensearchserviceDomainVpcOptions `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

