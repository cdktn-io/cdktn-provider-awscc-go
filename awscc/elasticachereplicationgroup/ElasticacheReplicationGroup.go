// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticachereplicationgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticachereplicationgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticache_replication_group awscc_elasticache_replication_group}.
type ElasticacheReplicationGroup interface {
	cdktn.TerraformResource
	AtRestEncryptionEnabled() interface{}
	SetAtRestEncryptionEnabled(val interface{})
	AtRestEncryptionEnabledInput() interface{}
	AuthToken() *string
	SetAuthToken(val *string)
	AuthTokenInput() *string
	AutomaticFailoverEnabled() interface{}
	SetAutomaticFailoverEnabled(val interface{})
	AutomaticFailoverEnabledInput() interface{}
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	AutoMinorVersionUpgradeInput() interface{}
	CacheNodeType() *string
	SetCacheNodeType(val *string)
	CacheNodeTypeInput() *string
	CacheParameterGroupName() *string
	SetCacheParameterGroupName(val *string)
	CacheParameterGroupNameInput() *string
	CacheSecurityGroupNames() *[]*string
	SetCacheSecurityGroupNames(val *[]*string)
	CacheSecurityGroupNamesInput() *[]*string
	CacheSubnetGroupName() *string
	SetCacheSubnetGroupName(val *string)
	CacheSubnetGroupNameInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClusterMode() *string
	SetClusterMode(val *string)
	ClusterModeInput() *string
	ConfigurationEndPoint() ElasticacheReplicationGroupConfigurationEndPointOutputReference
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DataTieringEnabled() interface{}
	SetDataTieringEnabled(val interface{})
	DataTieringEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Durability() *string
	SetDurability(val *string)
	DurabilityInput() *string
	EffectiveDurability() *string
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlobalReplicationGroupId() *string
	SetGlobalReplicationGroupId(val *string)
	GlobalReplicationGroupIdInput() *string
	Id() *string
	IpDiscovery() *string
	SetIpDiscovery(val *string)
	IpDiscoveryInput() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogDeliveryConfigurations() ElasticacheReplicationGroupLogDeliveryConfigurationsList
	LogDeliveryConfigurationsInput() interface{}
	MultiAzEnabled() interface{}
	SetMultiAzEnabled(val interface{})
	MultiAzEnabledInput() interface{}
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
	// The tree node.
	Node() constructs.Node
	NodeGroupConfiguration() ElasticacheReplicationGroupNodeGroupConfigurationList
	NodeGroupConfigurationInput() interface{}
	NotificationTopicArn() *string
	SetNotificationTopicArn(val *string)
	NotificationTopicArnInput() *string
	NumCacheClusters() *float64
	SetNumCacheClusters(val *float64)
	NumCacheClustersInput() *float64
	NumNodeGroups() *float64
	SetNumNodeGroups(val *float64)
	NumNodeGroupsInput() *float64
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PreferredCacheClusterAZs() *[]*string
	SetPreferredCacheClusterAZs(val *[]*string)
	PreferredCacheClusterAZsInput() *[]*string
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	PreferredMaintenanceWindowInput() *string
	PrimaryClusterId() *string
	SetPrimaryClusterId(val *string)
	PrimaryClusterIdInput() *string
	PrimaryEndPoint() ElasticacheReplicationGroupPrimaryEndPointOutputReference
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
	ReadEndPoint() ElasticacheReplicationGroupReadEndPointOutputReference
	ReaderEndPoint() ElasticacheReplicationGroupReaderEndPointOutputReference
	ReplicasPerNodeGroup() *float64
	SetReplicasPerNodeGroup(val *float64)
	ReplicasPerNodeGroupInput() *float64
	ReplicationGroupDescription() *string
	SetReplicationGroupDescription(val *string)
	ReplicationGroupDescriptionInput() *string
	ReplicationGroupId() *string
	SetReplicationGroupId(val *string)
	ReplicationGroupIdInput() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SnapshotArns() *[]*string
	SetSnapshotArns(val *[]*string)
	SnapshotArnsInput() *[]*string
	SnapshotName() *string
	SetSnapshotName(val *string)
	SnapshotNameInput() *string
	SnapshotRetentionLimit() *float64
	SetSnapshotRetentionLimit(val *float64)
	SnapshotRetentionLimitInput() *float64
	SnapshottingClusterId() *string
	SetSnapshottingClusterId(val *string)
	SnapshottingClusterIdInput() *string
	SnapshotWindow() *string
	SetSnapshotWindow(val *string)
	SnapshotWindowInput() *string
	Tags() ElasticacheReplicationGroupTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TransitEncryptionEnabled() interface{}
	SetTransitEncryptionEnabled(val interface{})
	TransitEncryptionEnabledInput() interface{}
	TransitEncryptionMode() *string
	SetTransitEncryptionMode(val *string)
	TransitEncryptionModeInput() *string
	UserGroupIds() *[]*string
	SetUserGroupIds(val *[]*string)
	UserGroupIdsInput() *[]*string
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
	PutLogDeliveryConfigurations(value interface{})
	PutNodeGroupConfiguration(value interface{})
	PutTags(value interface{})
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
	ResetAtRestEncryptionEnabled()
	ResetAuthToken()
	ResetAutomaticFailoverEnabled()
	ResetAutoMinorVersionUpgrade()
	ResetCacheNodeType()
	ResetCacheParameterGroupName()
	ResetCacheSecurityGroupNames()
	ResetCacheSubnetGroupName()
	ResetClusterMode()
	ResetDataTieringEnabled()
	ResetDurability()
	ResetEngine()
	ResetEngineVersion()
	ResetGlobalReplicationGroupId()
	ResetIpDiscovery()
	ResetKmsKeyId()
	ResetLogDeliveryConfigurations()
	ResetMultiAzEnabled()
	ResetNetworkType()
	ResetNodeGroupConfiguration()
	ResetNotificationTopicArn()
	ResetNumCacheClusters()
	ResetNumNodeGroups()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetPreferredCacheClusterAZs()
	ResetPreferredMaintenanceWindow()
	ResetPrimaryClusterId()
	ResetReplicasPerNodeGroup()
	ResetReplicationGroupId()
	ResetSecurityGroupIds()
	ResetSnapshotArns()
	ResetSnapshotName()
	ResetSnapshotRetentionLimit()
	ResetSnapshottingClusterId()
	ResetSnapshotWindow()
	ResetTags()
	ResetTransitEncryptionEnabled()
	ResetTransitEncryptionMode()
	ResetUserGroupIds()
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

// The jsii proxy struct for ElasticacheReplicationGroup
type jsiiProxy_ElasticacheReplicationGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ElasticacheReplicationGroup) AtRestEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atRestEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) AtRestEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atRestEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) AuthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) AuthTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) AutomaticFailoverEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticFailoverEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) AutomaticFailoverEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticFailoverEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) CacheNodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) CacheNodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) CacheParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) CacheParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) CacheSecurityGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheSecurityGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) CacheSecurityGroupNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheSecurityGroupNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) CacheSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) CacheSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ClusterMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ClusterModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ConfigurationEndPoint() ElasticacheReplicationGroupConfigurationEndPointOutputReference {
	var returns ElasticacheReplicationGroupConfigurationEndPointOutputReference
	_jsii_.Get(
		j,
		"configurationEndPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) DataTieringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTieringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) DataTieringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTieringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Durability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) DurabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) EffectiveDurability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveDurability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) GlobalReplicationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalReplicationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) GlobalReplicationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalReplicationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) IpDiscovery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) IpDiscoveryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipDiscoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) LogDeliveryConfigurations() ElasticacheReplicationGroupLogDeliveryConfigurationsList {
	var returns ElasticacheReplicationGroupLogDeliveryConfigurationsList
	_jsii_.Get(
		j,
		"logDeliveryConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) LogDeliveryConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeliveryConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) MultiAzEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) MultiAzEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) NodeGroupConfiguration() ElasticacheReplicationGroupNodeGroupConfigurationList {
	var returns ElasticacheReplicationGroupNodeGroupConfigurationList
	_jsii_.Get(
		j,
		"nodeGroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) NodeGroupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeGroupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) NotificationTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) NotificationTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) NumCacheClusters() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCacheClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) NumCacheClustersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCacheClustersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) NumNodeGroups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numNodeGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) NumNodeGroupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numNodeGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) PreferredCacheClusterAZs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredCacheClusterAZs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) PreferredCacheClusterAZsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredCacheClusterAZsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) PrimaryClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) PrimaryClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) PrimaryEndPoint() ElasticacheReplicationGroupPrimaryEndPointOutputReference {
	var returns ElasticacheReplicationGroupPrimaryEndPointOutputReference
	_jsii_.Get(
		j,
		"primaryEndPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ReadEndPoint() ElasticacheReplicationGroupReadEndPointOutputReference {
	var returns ElasticacheReplicationGroupReadEndPointOutputReference
	_jsii_.Get(
		j,
		"readEndPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ReaderEndPoint() ElasticacheReplicationGroupReaderEndPointOutputReference {
	var returns ElasticacheReplicationGroupReaderEndPointOutputReference
	_jsii_.Get(
		j,
		"readerEndPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ReplicasPerNodeGroup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasPerNodeGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ReplicasPerNodeGroupInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasPerNodeGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ReplicationGroupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ReplicationGroupDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ReplicationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) ReplicationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SnapshotArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SnapshotArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SnapshotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SnapshotRetentionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SnapshotRetentionLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SnapshottingClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshottingClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SnapshottingClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshottingClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SnapshotWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) SnapshotWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) Tags() ElasticacheReplicationGroupTagsList {
	var returns ElasticacheReplicationGroupTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) TransitEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) TransitEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) TransitEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) TransitEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) UserGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheReplicationGroup) UserGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticache_replication_group awscc_elasticache_replication_group} Resource.
func NewElasticacheReplicationGroup(scope constructs.Construct, id *string, config *ElasticacheReplicationGroupConfig) ElasticacheReplicationGroup {
	_init_.Initialize()

	if err := validateNewElasticacheReplicationGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticacheReplicationGroup{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticacheReplicationGroup.ElasticacheReplicationGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticache_replication_group awscc_elasticache_replication_group} Resource.
func NewElasticacheReplicationGroup_Override(e ElasticacheReplicationGroup, scope constructs.Construct, id *string, config *ElasticacheReplicationGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticacheReplicationGroup.ElasticacheReplicationGroup",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetAtRestEncryptionEnabled(val interface{}) {
	if err := j.validateSetAtRestEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atRestEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetAuthToken(val *string) {
	if err := j.validateSetAuthTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authToken",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetAutomaticFailoverEnabled(val interface{}) {
	if err := j.validateSetAutomaticFailoverEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticFailoverEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetCacheNodeType(val *string) {
	if err := j.validateSetCacheNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheNodeType",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetCacheParameterGroupName(val *string) {
	if err := j.validateSetCacheParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetCacheSecurityGroupNames(val *[]*string) {
	if err := j.validateSetCacheSecurityGroupNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheSecurityGroupNames",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetCacheSubnetGroupName(val *string) {
	if err := j.validateSetCacheSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetClusterMode(val *string) {
	if err := j.validateSetClusterModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterMode",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetDataTieringEnabled(val interface{}) {
	if err := j.validateSetDataTieringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTieringEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetDurability(val *string) {
	if err := j.validateSetDurabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durability",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetGlobalReplicationGroupId(val *string) {
	if err := j.validateSetGlobalReplicationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalReplicationGroupId",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetIpDiscovery(val *string) {
	if err := j.validateSetIpDiscoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipDiscovery",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetMultiAzEnabled(val interface{}) {
	if err := j.validateSetMultiAzEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAzEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetNotificationTopicArn(val *string) {
	if err := j.validateSetNotificationTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationTopicArn",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetNumCacheClusters(val *float64) {
	if err := j.validateSetNumCacheClustersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numCacheClusters",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetNumNodeGroups(val *float64) {
	if err := j.validateSetNumNodeGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numNodeGroups",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetPreferredCacheClusterAZs(val *[]*string) {
	if err := j.validateSetPreferredCacheClusterAZsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredCacheClusterAZs",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetPrimaryClusterId(val *string) {
	if err := j.validateSetPrimaryClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryClusterId",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetReplicasPerNodeGroup(val *float64) {
	if err := j.validateSetReplicasPerNodeGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicasPerNodeGroup",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetReplicationGroupDescription(val *string) {
	if err := j.validateSetReplicationGroupDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationGroupDescription",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetReplicationGroupId(val *string) {
	if err := j.validateSetReplicationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationGroupId",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetSnapshotArns(val *[]*string) {
	if err := j.validateSetSnapshotArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotArns",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetSnapshotName(val *string) {
	if err := j.validateSetSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotName",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetSnapshotRetentionLimit(val *float64) {
	if err := j.validateSetSnapshotRetentionLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotRetentionLimit",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetSnapshottingClusterId(val *string) {
	if err := j.validateSetSnapshottingClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshottingClusterId",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetSnapshotWindow(val *string) {
	if err := j.validateSetSnapshotWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotWindow",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetTransitEncryptionEnabled(val interface{}) {
	if err := j.validateSetTransitEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetTransitEncryptionMode(val *string) {
	if err := j.validateSetTransitEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_ElasticacheReplicationGroup)SetUserGroupIds(val *[]*string) {
	if err := j.validateSetUserGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a ElasticacheReplicationGroup resource upon running "cdktn plan <stack-name>".
func ElasticacheReplicationGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateElasticacheReplicationGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticacheReplicationGroup.ElasticacheReplicationGroup",
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
func ElasticacheReplicationGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheReplicationGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticacheReplicationGroup.ElasticacheReplicationGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticacheReplicationGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheReplicationGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticacheReplicationGroup.ElasticacheReplicationGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticacheReplicationGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheReplicationGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticacheReplicationGroup.ElasticacheReplicationGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElasticacheReplicationGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.elasticacheReplicationGroup.ElasticacheReplicationGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElasticacheReplicationGroup) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElasticacheReplicationGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_ElasticacheReplicationGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElasticacheReplicationGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElasticacheReplicationGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElasticacheReplicationGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElasticacheReplicationGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElasticacheReplicationGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElasticacheReplicationGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElasticacheReplicationGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_ElasticacheReplicationGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (e *jsiiProxy_ElasticacheReplicationGroup) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) PutLogDeliveryConfigurations(value interface{}) {
	if err := e.validatePutLogDeliveryConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogDeliveryConfigurations",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) PutNodeGroupConfiguration(value interface{}) {
	if err := e.validatePutNodeGroupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNodeGroupConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetAtRestEncryptionEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAtRestEncryptionEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetAuthToken() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthToken",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetAutomaticFailoverEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAutomaticFailoverEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetCacheNodeType() {
	_jsii_.InvokeVoid(
		e,
		"resetCacheNodeType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetCacheParameterGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetCacheParameterGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetCacheSecurityGroupNames() {
	_jsii_.InvokeVoid(
		e,
		"resetCacheSecurityGroupNames",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetCacheSubnetGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetCacheSubnetGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetClusterMode() {
	_jsii_.InvokeVoid(
		e,
		"resetClusterMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetDataTieringEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetDataTieringEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetDurability() {
	_jsii_.InvokeVoid(
		e,
		"resetDurability",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetEngine() {
	_jsii_.InvokeVoid(
		e,
		"resetEngine",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetGlobalReplicationGroupId() {
	_jsii_.InvokeVoid(
		e,
		"resetGlobalReplicationGroupId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetIpDiscovery() {
	_jsii_.InvokeVoid(
		e,
		"resetIpDiscovery",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		e,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetLogDeliveryConfigurations() {
	_jsii_.InvokeVoid(
		e,
		"resetLogDeliveryConfigurations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetMultiAzEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetMultiAzEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetNetworkType() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetNodeGroupConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetNodeGroupConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetNotificationTopicArn() {
	_jsii_.InvokeVoid(
		e,
		"resetNotificationTopicArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetNumCacheClusters() {
	_jsii_.InvokeVoid(
		e,
		"resetNumCacheClusters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetNumNodeGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetNumNodeGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetPort() {
	_jsii_.InvokeVoid(
		e,
		"resetPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetPreferredCacheClusterAZs() {
	_jsii_.InvokeVoid(
		e,
		"resetPreferredCacheClusterAZs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		e,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetPrimaryClusterId() {
	_jsii_.InvokeVoid(
		e,
		"resetPrimaryClusterId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetReplicasPerNodeGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetReplicasPerNodeGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetReplicationGroupId() {
	_jsii_.InvokeVoid(
		e,
		"resetReplicationGroupId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetSnapshotArns() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotArns",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetSnapshotName() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetSnapshotRetentionLimit() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotRetentionLimit",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetSnapshottingClusterId() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshottingClusterId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetSnapshotWindow() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotWindow",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetTransitEncryptionEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetTransitEncryptionEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetTransitEncryptionMode() {
	_jsii_.InvokeVoid(
		e,
		"resetTransitEncryptionMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ResetUserGroupIds() {
	_jsii_.InvokeVoid(
		e,
		"resetUserGroupIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheReplicationGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheReplicationGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheReplicationGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheReplicationGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

