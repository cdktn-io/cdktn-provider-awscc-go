// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ekscluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_cluster awscc_eks_cluster}.
type EksCluster interface {
	cdktn.TerraformResource
	AccessConfig() EksClusterAccessConfigOutputReference
	AccessConfigInput() interface{}
	Arn() *string
	BootstrapSelfManagedAddons() interface{}
	SetBootstrapSelfManagedAddons(val interface{})
	BootstrapSelfManagedAddonsInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CertificateAuthorityData() *string
	ClusterId() *string
	ClusterSecurityGroupId() *string
	ComputeConfig() EksClusterComputeConfigOutputReference
	ComputeConfigInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlaneScalingConfig() EksClusterControlPlaneScalingConfigOutputReference
	ControlPlaneScalingConfigInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EncryptionConfig() EksClusterEncryptionConfigList
	EncryptionConfigInput() interface{}
	EncryptionConfigKeyArn() *string
	Endpoint() *string
	Force() interface{}
	SetForce(val interface{})
	ForceInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KubeApiServerConfig() EksClusterKubeApiServerConfigOutputReference
	KubeApiServerConfigInput() interface{}
	KubeControllerManagerConfig() EksClusterKubeControllerManagerConfigOutputReference
	KubeControllerManagerConfigInput() interface{}
	KubernetesNetworkConfig() EksClusterKubernetesNetworkConfigOutputReference
	KubernetesNetworkConfigInput() interface{}
	KubeSchedulerConfig() EksClusterKubeSchedulerConfigOutputReference
	KubeSchedulerConfigInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Logging() EksClusterLoggingOutputReference
	LoggingInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OpenIdConnectIssuerUrl() *string
	OutpostConfig() EksClusterOutpostConfigOutputReference
	OutpostConfigInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RemoteNetworkConfig() EksClusterRemoteNetworkConfigOutputReference
	RemoteNetworkConfigInput() interface{}
	ResourcesVpcConfig() EksClusterResourcesVpcConfigOutputReference
	ResourcesVpcConfigInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	RollbackConfig() EksClusterRollbackConfigOutputReference
	RollbackConfigInput() interface{}
	StorageConfig() EksClusterStorageConfigOutputReference
	StorageConfigInput() interface{}
	Tags() EksClusterTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpgradePolicy() EksClusterUpgradePolicyOutputReference
	UpgradePolicyInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	ZonalShiftConfig() EksClusterZonalShiftConfigOutputReference
	ZonalShiftConfigInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAccessConfig(value *EksClusterAccessConfig)
	PutComputeConfig(value *EksClusterComputeConfig)
	PutControlPlaneScalingConfig(value *EksClusterControlPlaneScalingConfig)
	PutEncryptionConfig(value interface{})
	PutKubeApiServerConfig(value *EksClusterKubeApiServerConfig)
	PutKubeControllerManagerConfig(value *EksClusterKubeControllerManagerConfig)
	PutKubernetesNetworkConfig(value *EksClusterKubernetesNetworkConfig)
	PutKubeSchedulerConfig(value *EksClusterKubeSchedulerConfig)
	PutLogging(value *EksClusterLogging)
	PutOutpostConfig(value *EksClusterOutpostConfig)
	PutRemoteNetworkConfig(value *EksClusterRemoteNetworkConfig)
	PutResourcesVpcConfig(value *EksClusterResourcesVpcConfig)
	PutRollbackConfig(value *EksClusterRollbackConfig)
	PutStorageConfig(value *EksClusterStorageConfig)
	PutTags(value interface{})
	PutUpgradePolicy(value *EksClusterUpgradePolicy)
	PutZonalShiftConfig(value *EksClusterZonalShiftConfig)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAccessConfig()
	ResetBootstrapSelfManagedAddons()
	ResetComputeConfig()
	ResetControlPlaneScalingConfig()
	ResetDeletionProtection()
	ResetEncryptionConfig()
	ResetForce()
	ResetKubeApiServerConfig()
	ResetKubeControllerManagerConfig()
	ResetKubernetesNetworkConfig()
	ResetKubeSchedulerConfig()
	ResetLogging()
	ResetName()
	ResetOutpostConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRemoteNetworkConfig()
	ResetRollbackConfig()
	ResetStorageConfig()
	ResetTags()
	ResetUpgradePolicy()
	ResetVersion()
	ResetZonalShiftConfig()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for EksCluster
type jsiiProxy_EksCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_EksCluster) AccessConfig() EksClusterAccessConfigOutputReference {
	var returns EksClusterAccessConfigOutputReference
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) AccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) BootstrapSelfManagedAddons() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapSelfManagedAddons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) BootstrapSelfManagedAddonsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapSelfManagedAddonsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) CertificateAuthorityData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ClusterSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ComputeConfig() EksClusterComputeConfigOutputReference {
	var returns EksClusterComputeConfigOutputReference
	_jsii_.Get(
		j,
		"computeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ComputeConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"computeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ControlPlaneScalingConfig() EksClusterControlPlaneScalingConfigOutputReference {
	var returns EksClusterControlPlaneScalingConfigOutputReference
	_jsii_.Get(
		j,
		"controlPlaneScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ControlPlaneScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"controlPlaneScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) EncryptionConfig() EksClusterEncryptionConfigList {
	var returns EksClusterEncryptionConfigList
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) EncryptionConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) EncryptionConfigKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionConfigKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Force() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"force",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ForceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) KubeApiServerConfig() EksClusterKubeApiServerConfigOutputReference {
	var returns EksClusterKubeApiServerConfigOutputReference
	_jsii_.Get(
		j,
		"kubeApiServerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) KubeApiServerConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kubeApiServerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) KubeControllerManagerConfig() EksClusterKubeControllerManagerConfigOutputReference {
	var returns EksClusterKubeControllerManagerConfigOutputReference
	_jsii_.Get(
		j,
		"kubeControllerManagerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) KubeControllerManagerConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kubeControllerManagerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) KubernetesNetworkConfig() EksClusterKubernetesNetworkConfigOutputReference {
	var returns EksClusterKubernetesNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"kubernetesNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) KubernetesNetworkConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kubernetesNetworkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) KubeSchedulerConfig() EksClusterKubeSchedulerConfigOutputReference {
	var returns EksClusterKubeSchedulerConfigOutputReference
	_jsii_.Get(
		j,
		"kubeSchedulerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) KubeSchedulerConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kubeSchedulerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Logging() EksClusterLoggingOutputReference {
	var returns EksClusterLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) LoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) OpenIdConnectIssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) OutpostConfig() EksClusterOutpostConfigOutputReference {
	var returns EksClusterOutpostConfigOutputReference
	_jsii_.Get(
		j,
		"outpostConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) OutpostConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outpostConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RemoteNetworkConfig() EksClusterRemoteNetworkConfigOutputReference {
	var returns EksClusterRemoteNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"remoteNetworkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RemoteNetworkConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteNetworkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ResourcesVpcConfig() EksClusterResourcesVpcConfigOutputReference {
	var returns EksClusterResourcesVpcConfigOutputReference
	_jsii_.Get(
		j,
		"resourcesVpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ResourcesVpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesVpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RollbackConfig() EksClusterRollbackConfigOutputReference {
	var returns EksClusterRollbackConfigOutputReference
	_jsii_.Get(
		j,
		"rollbackConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) RollbackConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollbackConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) StorageConfig() EksClusterStorageConfigOutputReference {
	var returns EksClusterStorageConfigOutputReference
	_jsii_.Get(
		j,
		"storageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) StorageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Tags() EksClusterTagsList {
	var returns EksClusterTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) UpgradePolicy() EksClusterUpgradePolicyOutputReference {
	var returns EksClusterUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"upgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) UpgradePolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ZonalShiftConfig() EksClusterZonalShiftConfigOutputReference {
	var returns EksClusterZonalShiftConfigOutputReference
	_jsii_.Get(
		j,
		"zonalShiftConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCluster) ZonalShiftConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zonalShiftConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_cluster awscc_eks_cluster} Resource.
func NewEksCluster(scope constructs.Construct, id *string, config *EksClusterConfig) EksCluster {
	_init_.Initialize()

	if err := validateNewEksClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EksCluster{}

	_jsii_.Create(
		"@cdktn/provider-awscc.eksCluster.EksCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/eks_cluster awscc_eks_cluster} Resource.
func NewEksCluster_Override(e EksCluster, scope constructs.Construct, id *string, config *EksClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.eksCluster.EksCluster",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EksCluster)SetBootstrapSelfManagedAddons(val interface{}) {
	if err := j.validateSetBootstrapSelfManagedAddonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrapSelfManagedAddons",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetForce(val interface{}) {
	if err := j.validateSetForceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"force",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_EksCluster)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a EksCluster resource upon running "cdktn plan <stack-name>".
func EksCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateEksCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.eksCluster.EksCluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func EksCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.eksCluster.EksCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EksCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.eksCluster.EksCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EksCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.eksCluster.EksCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EksCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.eksCluster.EksCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EksCluster) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EksCluster) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EksCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EksCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := e.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EksCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EksCluster) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EksCluster) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EksCluster) PutAccessConfig(value *EksClusterAccessConfig) {
	if err := e.validatePutAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAccessConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutComputeConfig(value *EksClusterComputeConfig) {
	if err := e.validatePutComputeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putComputeConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutControlPlaneScalingConfig(value *EksClusterControlPlaneScalingConfig) {
	if err := e.validatePutControlPlaneScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putControlPlaneScalingConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutEncryptionConfig(value interface{}) {
	if err := e.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutKubeApiServerConfig(value *EksClusterKubeApiServerConfig) {
	if err := e.validatePutKubeApiServerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putKubeApiServerConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutKubeControllerManagerConfig(value *EksClusterKubeControllerManagerConfig) {
	if err := e.validatePutKubeControllerManagerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putKubeControllerManagerConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutKubernetesNetworkConfig(value *EksClusterKubernetesNetworkConfig) {
	if err := e.validatePutKubernetesNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putKubernetesNetworkConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutKubeSchedulerConfig(value *EksClusterKubeSchedulerConfig) {
	if err := e.validatePutKubeSchedulerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putKubeSchedulerConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutLogging(value *EksClusterLogging) {
	if err := e.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogging",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutOutpostConfig(value *EksClusterOutpostConfig) {
	if err := e.validatePutOutpostConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOutpostConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutRemoteNetworkConfig(value *EksClusterRemoteNetworkConfig) {
	if err := e.validatePutRemoteNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRemoteNetworkConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutResourcesVpcConfig(value *EksClusterResourcesVpcConfig) {
	if err := e.validatePutResourcesVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putResourcesVpcConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutRollbackConfig(value *EksClusterRollbackConfig) {
	if err := e.validatePutRollbackConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRollbackConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutStorageConfig(value *EksClusterStorageConfig) {
	if err := e.validatePutStorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStorageConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutUpgradePolicy(value *EksClusterUpgradePolicy) {
	if err := e.validatePutUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putUpgradePolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) PutZonalShiftConfig(value *EksClusterZonalShiftConfig) {
	if err := e.validatePutZonalShiftConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putZonalShiftConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EksCluster) ResetAccessConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAccessConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetBootstrapSelfManagedAddons() {
	_jsii_.InvokeVoid(
		e,
		"resetBootstrapSelfManagedAddons",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetComputeConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetComputeConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetControlPlaneScalingConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetControlPlaneScalingConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		e,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetForce() {
	_jsii_.InvokeVoid(
		e,
		"resetForce",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetKubeApiServerConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetKubeApiServerConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetKubeControllerManagerConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetKubeControllerManagerConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetKubernetesNetworkConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetKubernetesNetworkConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetKubeSchedulerConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetKubeSchedulerConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetLogging() {
	_jsii_.InvokeVoid(
		e,
		"resetLogging",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetOutpostConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetOutpostConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetRemoteNetworkConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetRemoteNetworkConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetRollbackConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetRollbackConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetStorageConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetUpgradePolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetUpgradePolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) ResetZonalShiftConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetZonalShiftConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

