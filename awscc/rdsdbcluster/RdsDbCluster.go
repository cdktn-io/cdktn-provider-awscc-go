// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rdsdbcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rds_db_cluster awscc_rds_db_cluster}.
type RdsDbCluster interface {
	cdktn.TerraformResource
	AllocatedStorage() *float64
	SetAllocatedStorage(val *float64)
	AllocatedStorageInput() *float64
	AssociatedRoles() RdsDbClusterAssociatedRolesList
	AssociatedRolesInput() interface{}
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	AutoMinorVersionUpgradeInput() interface{}
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	BacktrackWindow() *float64
	SetBacktrackWindow(val *float64)
	BacktrackWindowInput() *float64
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	BackupRetentionPeriodInput() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClusterScalabilityType() *string
	SetClusterScalabilityType(val *string)
	ClusterScalabilityTypeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToSnapshot() interface{}
	SetCopyTagsToSnapshot(val interface{})
	CopyTagsToSnapshotInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseInsightsMode() *string
	SetDatabaseInsightsMode(val *string)
	DatabaseInsightsModeInput() *string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DbClusterArn() *string
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	DbClusterIdentifierInput() *string
	DbClusterInstanceClass() *string
	SetDbClusterInstanceClass(val *string)
	DbClusterInstanceClassInput() *string
	DbClusterParameterGroupName() *string
	SetDbClusterParameterGroupName(val *string)
	DbClusterParameterGroupNameInput() *string
	DbClusterResourceId() *string
	DbInstanceParameterGroupName() *string
	SetDbInstanceParameterGroupName(val *string)
	DbInstanceParameterGroupNameInput() *string
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	DbSubnetGroupNameInput() *string
	DbSystemId() *string
	SetDbSystemId(val *string)
	DbSystemIdInput() *string
	DeleteAutomatedBackups() interface{}
	SetDeleteAutomatedBackups(val interface{})
	DeleteAutomatedBackupsInput() interface{}
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Domain() *string
	SetDomain(val *string)
	DomainIamRoleName() *string
	SetDomainIamRoleName(val *string)
	DomainIamRoleNameInput() *string
	DomainInput() *string
	EnableCloudwatchLogsExports() *[]*string
	SetEnableCloudwatchLogsExports(val *[]*string)
	EnableCloudwatchLogsExportsInput() *[]*string
	EnableGlobalWriteForwarding() interface{}
	SetEnableGlobalWriteForwarding(val interface{})
	EnableGlobalWriteForwardingInput() interface{}
	EnableHttpEndpoint() interface{}
	SetEnableHttpEndpoint(val interface{})
	EnableHttpEndpointInput() interface{}
	EnableIamDatabaseAuthentication() interface{}
	SetEnableIamDatabaseAuthentication(val interface{})
	EnableIamDatabaseAuthenticationInput() interface{}
	EnableLocalWriteForwarding() interface{}
	SetEnableLocalWriteForwarding(val interface{})
	EnableLocalWriteForwardingInput() interface{}
	Endpoint() RdsDbClusterEndpointOutputReference
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
	EngineLifecycleSupport() *string
	SetEngineLifecycleSupport(val *string)
	EngineLifecycleSupportInput() *string
	EngineMode() *string
	SetEngineMode(val *string)
	EngineModeInput() *string
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
	GlobalClusterIdentifier() *string
	SetGlobalClusterIdentifier(val *string)
	GlobalClusterIdentifierInput() *string
	Id() *string
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ManageMasterUserPassword() interface{}
	SetManageMasterUserPassword(val interface{})
	ManageMasterUserPasswordInput() interface{}
	MasterUserAuthenticationType() *string
	SetMasterUserAuthenticationType(val *string)
	MasterUserAuthenticationTypeInput() *string
	MasterUsername() *string
	SetMasterUsername(val *string)
	MasterUsernameInput() *string
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	MasterUserPasswordInput() *string
	MasterUserSecret() RdsDbClusterMasterUserSecretOutputReference
	MasterUserSecretInput() interface{}
	MonitoringInterval() *float64
	SetMonitoringInterval(val *float64)
	MonitoringIntervalInput() *float64
	MonitoringRoleArn() *string
	SetMonitoringRoleArn(val *string)
	MonitoringRoleArnInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
	// The tree node.
	Node() constructs.Node
	PerformanceInsightsEnabled() interface{}
	SetPerformanceInsightsEnabled(val interface{})
	PerformanceInsightsEnabledInput() interface{}
	PerformanceInsightsKmsKeyId() *string
	SetPerformanceInsightsKmsKeyId(val *string)
	PerformanceInsightsKmsKeyIdInput() *string
	PerformanceInsightsRetentionPeriod() *float64
	SetPerformanceInsightsRetentionPeriod(val *float64)
	PerformanceInsightsRetentionPeriodInput() *float64
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	PreferredBackupWindowInput() *string
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	PreferredMaintenanceWindowInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	PubliclyAccessibleInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReadEndpoint() RdsDbClusterReadEndpointOutputReference
	ReplicationSourceIdentifier() *string
	SetReplicationSourceIdentifier(val *string)
	ReplicationSourceIdentifierInput() *string
	RestoreToTime() *string
	SetRestoreToTime(val *string)
	RestoreToTimeInput() *string
	RestoreType() *string
	SetRestoreType(val *string)
	RestoreTypeInput() *string
	ScalingConfiguration() RdsDbClusterScalingConfigurationOutputReference
	ScalingConfigurationInput() interface{}
	ServerlessV2ScalingConfiguration() RdsDbClusterServerlessV2ScalingConfigurationOutputReference
	ServerlessV2ScalingConfigurationInput() interface{}
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	SnapshotIdentifierInput() *string
	SourceDbClusterIdentifier() *string
	SetSourceDbClusterIdentifier(val *string)
	SourceDbClusterIdentifierInput() *string
	SourceDbClusterResourceId() *string
	SetSourceDbClusterResourceId(val *string)
	SourceDbClusterResourceIdInput() *string
	SourceRegion() *string
	SetSourceRegion(val *string)
	SourceRegionInput() *string
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	StorageEncryptedInput() interface{}
	StorageEncryptionType() *string
	StorageThroughput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	Tags() RdsDbClusterTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UseLatestRestorableTime() interface{}
	SetUseLatestRestorableTime(val interface{})
	UseLatestRestorableTimeInput() interface{}
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSecurityGroupIdsInput() *[]*string
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
	PutAssociatedRoles(value interface{})
	PutMasterUserSecret(value *RdsDbClusterMasterUserSecret)
	PutScalingConfiguration(value *RdsDbClusterScalingConfiguration)
	PutServerlessV2ScalingConfiguration(value *RdsDbClusterServerlessV2ScalingConfiguration)
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
	ResetAllocatedStorage()
	ResetAssociatedRoles()
	ResetAutoMinorVersionUpgrade()
	ResetAvailabilityZones()
	ResetBacktrackWindow()
	ResetBackupRetentionPeriod()
	ResetClusterScalabilityType()
	ResetCopyTagsToSnapshot()
	ResetDatabaseInsightsMode()
	ResetDatabaseName()
	ResetDbClusterIdentifier()
	ResetDbClusterInstanceClass()
	ResetDbClusterParameterGroupName()
	ResetDbInstanceParameterGroupName()
	ResetDbSubnetGroupName()
	ResetDbSystemId()
	ResetDeleteAutomatedBackups()
	ResetDeletionProtection()
	ResetDomain()
	ResetDomainIamRoleName()
	ResetEnableCloudwatchLogsExports()
	ResetEnableGlobalWriteForwarding()
	ResetEnableHttpEndpoint()
	ResetEnableIamDatabaseAuthentication()
	ResetEnableLocalWriteForwarding()
	ResetEngine()
	ResetEngineLifecycleSupport()
	ResetEngineMode()
	ResetEngineVersion()
	ResetGlobalClusterIdentifier()
	ResetIops()
	ResetKmsKeyId()
	ResetManageMasterUserPassword()
	ResetMasterUserAuthenticationType()
	ResetMasterUsername()
	ResetMasterUserPassword()
	ResetMasterUserSecret()
	ResetMonitoringInterval()
	ResetMonitoringRoleArn()
	ResetNetworkType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerformanceInsightsEnabled()
	ResetPerformanceInsightsKmsKeyId()
	ResetPerformanceInsightsRetentionPeriod()
	ResetPort()
	ResetPreferredBackupWindow()
	ResetPreferredMaintenanceWindow()
	ResetPubliclyAccessible()
	ResetReplicationSourceIdentifier()
	ResetRestoreToTime()
	ResetRestoreType()
	ResetScalingConfiguration()
	ResetServerlessV2ScalingConfiguration()
	ResetSnapshotIdentifier()
	ResetSourceDbClusterIdentifier()
	ResetSourceDbClusterResourceId()
	ResetSourceRegion()
	ResetStorageEncrypted()
	ResetStorageType()
	ResetTags()
	ResetUseLatestRestorableTime()
	ResetVpcSecurityGroupIds()
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

// The jsii proxy struct for RdsDbCluster
type jsiiProxy_RdsDbCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_RdsDbCluster) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) AllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) AssociatedRoles() RdsDbClusterAssociatedRolesList {
	var returns RdsDbClusterAssociatedRolesList
	_jsii_.Get(
		j,
		"associatedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) AssociatedRolesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) BacktrackWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backtrackWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) BacktrackWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backtrackWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) BackupRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ClusterScalabilityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterScalabilityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ClusterScalabilityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterScalabilityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) CopyTagsToSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DatabaseInsightsMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DatabaseInsightsModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbClusterInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbClusterInstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterInstanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbClusterParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbClusterResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbInstanceParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbInstanceParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DbSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DeleteAutomatedBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DeleteAutomatedBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DomainIamRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EnableCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EnableCloudwatchLogsExportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EnableGlobalWriteForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGlobalWriteForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EnableGlobalWriteForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGlobalWriteForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EnableHttpEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EnableHttpEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EnableIamDatabaseAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIamDatabaseAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EnableIamDatabaseAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIamDatabaseAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EnableLocalWriteForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalWriteForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EnableLocalWriteForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalWriteForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Endpoint() RdsDbClusterEndpointOutputReference {
	var returns RdsDbClusterEndpointOutputReference
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EngineLifecycleSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EngineLifecycleSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EngineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EngineModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) GlobalClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) GlobalClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ManageMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ManageMasterUserPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MasterUserAuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserAuthenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MasterUserAuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserAuthenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MasterUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MasterUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MasterUserSecret() RdsDbClusterMasterUserSecretOutputReference {
	var returns RdsDbClusterMasterUserSecretOutputReference
	_jsii_.Get(
		j,
		"masterUserSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MasterUserSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterUserSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MonitoringIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) MonitoringRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PerformanceInsightsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PerformanceInsightsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PerformanceInsightsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PerformanceInsightsRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PreferredBackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ReadEndpoint() RdsDbClusterReadEndpointOutputReference {
	var returns RdsDbClusterReadEndpointOutputReference
	_jsii_.Get(
		j,
		"readEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ReplicationSourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ReplicationSourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSourceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) RestoreToTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreToTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) RestoreToTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreToTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) RestoreType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) RestoreTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ScalingConfiguration() RdsDbClusterScalingConfigurationOutputReference {
	var returns RdsDbClusterScalingConfigurationOutputReference
	_jsii_.Get(
		j,
		"scalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ScalingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ServerlessV2ScalingConfiguration() RdsDbClusterServerlessV2ScalingConfigurationOutputReference {
	var returns RdsDbClusterServerlessV2ScalingConfigurationOutputReference
	_jsii_.Get(
		j,
		"serverlessV2ScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) ServerlessV2ScalingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessV2ScalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) SnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) SourceDbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) SourceDbClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) SourceDbClusterResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) SourceDbClusterResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) SourceRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) StorageEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) StorageEncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageEncryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) StorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) Tags() RdsDbClusterTagsList {
	var returns RdsDbClusterTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) UseLatestRestorableTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) UseLatestRestorableTimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbCluster) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rds_db_cluster awscc_rds_db_cluster} Resource.
func NewRdsDbCluster(scope constructs.Construct, id *string, config *RdsDbClusterConfig) RdsDbCluster {
	_init_.Initialize()

	if err := validateNewRdsDbClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsDbCluster{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDbCluster.RdsDbCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/rds_db_cluster awscc_rds_db_cluster} Resource.
func NewRdsDbCluster_Override(r RdsDbCluster, scope constructs.Construct, id *string, config *RdsDbClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDbCluster.RdsDbCluster",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetAllocatedStorage(val *float64) {
	if err := j.validateSetAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetBacktrackWindow(val *float64) {
	if err := j.validateSetBacktrackWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backtrackWindow",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetBackupRetentionPeriod(val *float64) {
	if err := j.validateSetBackupRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetClusterScalabilityType(val *string) {
	if err := j.validateSetClusterScalabilityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterScalabilityType",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDatabaseInsightsMode(val *string) {
	if err := j.validateSetDatabaseInsightsModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseInsightsMode",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDbClusterIdentifier(val *string) {
	if err := j.validateSetDbClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDbClusterInstanceClass(val *string) {
	if err := j.validateSetDbClusterInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterInstanceClass",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDbClusterParameterGroupName(val *string) {
	if err := j.validateSetDbClusterParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDbInstanceParameterGroupName(val *string) {
	if err := j.validateSetDbInstanceParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbInstanceParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDbSubnetGroupName(val *string) {
	if err := j.validateSetDbSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDbSystemId(val *string) {
	if err := j.validateSetDbSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSystemId",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDeleteAutomatedBackups(val interface{}) {
	if err := j.validateSetDeleteAutomatedBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAutomatedBackups",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetDomainIamRoleName(val *string) {
	if err := j.validateSetDomainIamRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetEnableCloudwatchLogsExports(val *[]*string) {
	if err := j.validateSetEnableCloudwatchLogsExportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetEnableGlobalWriteForwarding(val interface{}) {
	if err := j.validateSetEnableGlobalWriteForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableGlobalWriteForwarding",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetEnableHttpEndpoint(val interface{}) {
	if err := j.validateSetEnableHttpEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttpEndpoint",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetEnableIamDatabaseAuthentication(val interface{}) {
	if err := j.validateSetEnableIamDatabaseAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIamDatabaseAuthentication",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetEnableLocalWriteForwarding(val interface{}) {
	if err := j.validateSetEnableLocalWriteForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalWriteForwarding",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetEngineLifecycleSupport(val *string) {
	if err := j.validateSetEngineLifecycleSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineLifecycleSupport",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetEngineMode(val *string) {
	if err := j.validateSetEngineModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineMode",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetGlobalClusterIdentifier(val *string) {
	if err := j.validateSetGlobalClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetManageMasterUserPassword(val interface{}) {
	if err := j.validateSetManageMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterUserPassword",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetMasterUserAuthenticationType(val *string) {
	if err := j.validateSetMasterUserAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserAuthenticationType",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetMasterUsername(val *string) {
	if err := j.validateSetMasterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetMasterUserPassword(val *string) {
	if err := j.validateSetMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetMonitoringInterval(val *float64) {
	if err := j.validateSetMonitoringIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInterval",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetMonitoringRoleArn(val *string) {
	if err := j.validateSetMonitoringRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringRoleArn",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetPerformanceInsightsEnabled(val interface{}) {
	if err := j.validateSetPerformanceInsightsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsEnabled",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetPerformanceInsightsKmsKeyId(val *string) {
	if err := j.validateSetPerformanceInsightsKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetPerformanceInsightsRetentionPeriod(val *float64) {
	if err := j.validateSetPerformanceInsightsRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetPreferredBackupWindow(val *string) {
	if err := j.validateSetPreferredBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetReplicationSourceIdentifier(val *string) {
	if err := j.validateSetReplicationSourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationSourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetRestoreToTime(val *string) {
	if err := j.validateSetRestoreToTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreToTime",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetRestoreType(val *string) {
	if err := j.validateSetRestoreTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreType",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetSnapshotIdentifier(val *string) {
	if err := j.validateSetSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetSourceDbClusterIdentifier(val *string) {
	if err := j.validateSetSourceDbClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetSourceDbClusterResourceId(val *string) {
	if err := j.validateSetSourceDbClusterResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbClusterResourceId",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetSourceRegion(val *string) {
	if err := j.validateSetSourceRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRegion",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetUseLatestRestorableTime(val interface{}) {
	if err := j.validateSetUseLatestRestorableTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLatestRestorableTime",
		val,
	)
}

func (j *jsiiProxy_RdsDbCluster)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a RdsDbCluster resource upon running "cdktn plan <stack-name>".
func RdsDbCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateRdsDbCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbCluster.RdsDbCluster",
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
func RdsDbCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbCluster.RdsDbCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsDbCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbCluster.RdsDbCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsDbCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbCluster.RdsDbCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RdsDbCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.rdsDbCluster.RdsDbCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RdsDbCluster) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RdsDbCluster) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RdsDbCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RdsDbCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := r.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsDbCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RdsDbCluster) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsDbCluster) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RdsDbCluster) PutAssociatedRoles(value interface{}) {
	if err := r.validatePutAssociatedRolesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAssociatedRoles",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbCluster) PutMasterUserSecret(value *RdsDbClusterMasterUserSecret) {
	if err := r.validatePutMasterUserSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putMasterUserSecret",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbCluster) PutScalingConfiguration(value *RdsDbClusterScalingConfiguration) {
	if err := r.validatePutScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putScalingConfiguration",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbCluster) PutServerlessV2ScalingConfiguration(value *RdsDbClusterServerlessV2ScalingConfiguration) {
	if err := r.validatePutServerlessV2ScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putServerlessV2ScalingConfiguration",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbCluster) PutTags(value interface{}) {
	if err := r.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTags",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := r.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetAllocatedStorage() {
	_jsii_.InvokeVoid(
		r,
		"resetAllocatedStorage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetAssociatedRoles() {
	_jsii_.InvokeVoid(
		r,
		"resetAssociatedRoles",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		r,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		r,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetBacktrackWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetBacktrackWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetBackupRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetBackupRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetClusterScalabilityType() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterScalabilityType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetCopyTagsToSnapshot() {
	_jsii_.InvokeVoid(
		r,
		"resetCopyTagsToSnapshot",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDatabaseInsightsMode() {
	_jsii_.InvokeVoid(
		r,
		"resetDatabaseInsightsMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		r,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDbClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetDbClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDbClusterInstanceClass() {
	_jsii_.InvokeVoid(
		r,
		"resetDbClusterInstanceClass",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDbClusterParameterGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbClusterParameterGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDbInstanceParameterGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbInstanceParameterGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDbSubnetGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSubnetGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDbSystemId() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSystemId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDeleteAutomatedBackups() {
	_jsii_.InvokeVoid(
		r,
		"resetDeleteAutomatedBackups",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		r,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDomain() {
	_jsii_.InvokeVoid(
		r,
		"resetDomain",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetDomainIamRoleName() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainIamRoleName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetEnableCloudwatchLogsExports() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableCloudwatchLogsExports",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetEnableGlobalWriteForwarding() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableGlobalWriteForwarding",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetEnableHttpEndpoint() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableHttpEndpoint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetEnableIamDatabaseAuthentication() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableIamDatabaseAuthentication",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetEnableLocalWriteForwarding() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableLocalWriteForwarding",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetEngine() {
	_jsii_.InvokeVoid(
		r,
		"resetEngine",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetEngineLifecycleSupport() {
	_jsii_.InvokeVoid(
		r,
		"resetEngineLifecycleSupport",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetEngineMode() {
	_jsii_.InvokeVoid(
		r,
		"resetEngineMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetGlobalClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetGlobalClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetIops() {
	_jsii_.InvokeVoid(
		r,
		"resetIops",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetManageMasterUserPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetManageMasterUserPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetMasterUserAuthenticationType() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUserAuthenticationType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetMasterUsername() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUsername",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetMasterUserPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUserPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetMasterUserSecret() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUserSecret",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetMonitoringInterval() {
	_jsii_.InvokeVoid(
		r,
		"resetMonitoringInterval",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetMonitoringRoleArn() {
	_jsii_.InvokeVoid(
		r,
		"resetMonitoringRoleArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetNetworkType() {
	_jsii_.InvokeVoid(
		r,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetPerformanceInsightsEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetPerformanceInsightsEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetPerformanceInsightsKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetPerformanceInsightsKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetPerformanceInsightsRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetPerformanceInsightsRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetPort() {
	_jsii_.InvokeVoid(
		r,
		"resetPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetPreferredBackupWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetPreferredBackupWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		r,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetReplicationSourceIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetReplicationSourceIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetRestoreToTime() {
	_jsii_.InvokeVoid(
		r,
		"resetRestoreToTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetRestoreType() {
	_jsii_.InvokeVoid(
		r,
		"resetRestoreType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetScalingConfiguration() {
	_jsii_.InvokeVoid(
		r,
		"resetScalingConfiguration",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetServerlessV2ScalingConfiguration() {
	_jsii_.InvokeVoid(
		r,
		"resetServerlessV2ScalingConfiguration",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSnapshotIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetSourceDbClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceDbClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetSourceDbClusterResourceId() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceDbClusterResourceId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetSourceRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetStorageEncrypted() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageEncrypted",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetStorageType() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetUseLatestRestorableTime() {
	_jsii_.InvokeVoid(
		r,
		"resetUseLatestRestorableTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		r,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		r,
		"with",
		args,
		&returns,
	)

	return returns
}

