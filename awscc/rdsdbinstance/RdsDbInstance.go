// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rdsdbinstance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_instance awscc_rds_db_instance}.
type RdsDbInstance interface {
	cdktn.TerraformResource
	AdditionalStorageVolumes() RdsDbInstanceAdditionalStorageVolumesList
	AdditionalStorageVolumesInput() interface{}
	AllocatedStorage() *string
	SetAllocatedStorage(val *string)
	AllocatedStorageInput() *string
	AllowMajorVersionUpgrade() interface{}
	SetAllowMajorVersionUpgrade(val interface{})
	AllowMajorVersionUpgradeInput() interface{}
	ApplyImmediately() interface{}
	SetApplyImmediately(val interface{})
	ApplyImmediatelyInput() interface{}
	AssociatedRoles() RdsDbInstanceAssociatedRolesList
	AssociatedRolesInput() interface{}
	AutomaticBackupReplicationKmsKeyId() *string
	SetAutomaticBackupReplicationKmsKeyId(val *string)
	AutomaticBackupReplicationKmsKeyIdInput() *string
	AutomaticBackupReplicationRegion() *string
	SetAutomaticBackupReplicationRegion(val *string)
	AutomaticBackupReplicationRegionInput() *string
	AutomaticBackupReplicationRetentionPeriod() *float64
	SetAutomaticBackupReplicationRetentionPeriod(val *float64)
	AutomaticBackupReplicationRetentionPeriodInput() *float64
	AutomaticRestartTime() *string
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	AutoMinorVersionUpgradeInput() interface{}
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	BackupRetentionPeriodInput() *float64
	BackupTarget() *string
	SetBackupTarget(val *string)
	BackupTargetInput() *string
	CaCertificateIdentifier() *string
	SetCaCertificateIdentifier(val *string)
	CaCertificateIdentifierInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CertificateDetails() RdsDbInstanceCertificateDetailsOutputReference
	CertificateRotationRestart() interface{}
	SetCertificateRotationRestart(val interface{})
	CertificateRotationRestartInput() interface{}
	CharacterSetName() *string
	SetCharacterSetName(val *string)
	CharacterSetNameInput() *string
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
	CustomIamInstanceProfile() *string
	SetCustomIamInstanceProfile(val *string)
	CustomIamInstanceProfileInput() *string
	DatabaseInsightsMode() *string
	SetDatabaseInsightsMode(val *string)
	DatabaseInsightsModeInput() *string
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	DbClusterIdentifierInput() *string
	DbClusterSnapshotIdentifier() *string
	SetDbClusterSnapshotIdentifier(val *string)
	DbClusterSnapshotIdentifierInput() *string
	DbInstanceArn() *string
	DbInstanceClass() *string
	SetDbInstanceClass(val *string)
	DbInstanceClassInput() *string
	DbInstanceIdentifier() *string
	SetDbInstanceIdentifier(val *string)
	DbInstanceIdentifierInput() *string
	DbInstanceStatus() *string
	DbiResourceId() *string
	DbName() *string
	SetDbName(val *string)
	DbNameInput() *string
	DbParameterGroupName() *string
	SetDbParameterGroupName(val *string)
	DbParameterGroupNameInput() *string
	DbSecurityGroups() *[]*string
	SetDbSecurityGroups(val *[]*string)
	DbSecurityGroupsInput() *[]*string
	DbSnapshotIdentifier() *string
	SetDbSnapshotIdentifier(val *string)
	DbSnapshotIdentifierInput() *string
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	DbSubnetGroupNameInput() *string
	DbSystemId() *string
	SetDbSystemId(val *string)
	DbSystemIdInput() *string
	DedicatedLogVolume() interface{}
	SetDedicatedLogVolume(val interface{})
	DedicatedLogVolumeInput() interface{}
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
	DomainAuthSecretArn() *string
	SetDomainAuthSecretArn(val *string)
	DomainAuthSecretArnInput() *string
	DomainDnsIps() *[]*string
	SetDomainDnsIps(val *[]*string)
	DomainDnsIpsInput() *[]*string
	DomainFqdn() *string
	SetDomainFqdn(val *string)
	DomainFqdnInput() *string
	DomainIamRoleName() *string
	SetDomainIamRoleName(val *string)
	DomainIamRoleNameInput() *string
	DomainInput() *string
	DomainOu() *string
	SetDomainOu(val *string)
	DomainOuInput() *string
	EnableCloudwatchLogsExports() *[]*string
	SetEnableCloudwatchLogsExports(val *[]*string)
	EnableCloudwatchLogsExportsInput() *[]*string
	EnableIamDatabaseAuthentication() interface{}
	SetEnableIamDatabaseAuthentication(val interface{})
	EnableIamDatabaseAuthenticationInput() interface{}
	EnablePerformanceInsights() interface{}
	SetEnablePerformanceInsights(val interface{})
	EnablePerformanceInsightsInput() interface{}
	Endpoint() RdsDbInstanceEndpointOutputReference
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
	EngineLifecycleSupport() *string
	SetEngineLifecycleSupport(val *string)
	EngineLifecycleSupportInput() *string
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
	Id() *string
	InstanceCreateTime() *string
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	IsStorageConfigUpgradeAvailable() cdktn.IResolvable
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	LatestRestorableTime() *string
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ListenerEndpoint() RdsDbInstanceListenerEndpointOutputReference
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
	MasterUserSecret() RdsDbInstanceMasterUserSecretOutputReference
	MasterUserSecretInput() interface{}
	MaxAllocatedStorage() *float64
	SetMaxAllocatedStorage(val *float64)
	MaxAllocatedStorageInput() *float64
	MonitoringInterval() *float64
	SetMonitoringInterval(val *float64)
	MonitoringIntervalInput() *float64
	MonitoringRoleArn() *string
	SetMonitoringRoleArn(val *string)
	MonitoringRoleArnInput() *string
	MultiAz() interface{}
	SetMultiAz(val interface{})
	MultiAzInput() interface{}
	NcharCharacterSetName() *string
	SetNcharCharacterSetName(val *string)
	NcharCharacterSetNameInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
	// The tree node.
	Node() constructs.Node
	OptionGroupName() *string
	SetOptionGroupName(val *string)
	OptionGroupNameInput() *string
	PercentProgress() *string
	PerformanceInsightsKmsKeyId() *string
	SetPerformanceInsightsKmsKeyId(val *string)
	PerformanceInsightsKmsKeyIdInput() *string
	PerformanceInsightsRetentionPeriod() *float64
	SetPerformanceInsightsRetentionPeriod(val *float64)
	PerformanceInsightsRetentionPeriodInput() *float64
	Port() *string
	SetPort(val *string)
	PortInput() *string
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	PreferredBackupWindowInput() *string
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	PreferredMaintenanceWindowInput() *string
	ProcessorFeatures() RdsDbInstanceProcessorFeaturesList
	ProcessorFeaturesInput() interface{}
	PromotionTier() *float64
	SetPromotionTier(val *float64)
	PromotionTierInput() *float64
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
	ReadReplicaDbClusterIdentifiers() *[]*string
	ReadReplicaDbInstanceIdentifiers() *[]*string
	ReplicaMode() *string
	SetReplicaMode(val *string)
	ReplicaModeInput() *string
	RestoreTime() *string
	SetRestoreTime(val *string)
	RestoreTimeInput() *string
	ResumeFullAutomationModeTime() *string
	SecondaryAvailabilityZone() *string
	SourceDbClusterIdentifier() *string
	SetSourceDbClusterIdentifier(val *string)
	SourceDbClusterIdentifierInput() *string
	SourceDbInstanceAutomatedBackupsArn() *string
	SetSourceDbInstanceAutomatedBackupsArn(val *string)
	SourceDbInstanceAutomatedBackupsArnInput() *string
	SourceDbInstanceIdentifier() *string
	SetSourceDbInstanceIdentifier(val *string)
	SourceDbInstanceIdentifierInput() *string
	SourceDbiResourceId() *string
	SetSourceDbiResourceId(val *string)
	SourceDbiResourceIdInput() *string
	SourceRegion() *string
	SetSourceRegion(val *string)
	SourceRegionInput() *string
	StatusInfos() RdsDbInstanceStatusInfosList
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	StorageEncryptedInput() interface{}
	StorageThroughput() *float64
	SetStorageThroughput(val *float64)
	StorageThroughputInput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	Tags() RdsDbInstanceTagsList
	TagsInput() interface{}
	TdeCredentialArn() *string
	SetTdeCredentialArn(val *string)
	TdeCredentialArnInput() *string
	TdeCredentialPassword() *string
	SetTdeCredentialPassword(val *string)
	TdeCredentialPasswordInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	UseDefaultProcessorFeatures() interface{}
	SetUseDefaultProcessorFeatures(val interface{})
	UseDefaultProcessorFeaturesInput() interface{}
	UseLatestRestorableTime() interface{}
	SetUseLatestRestorableTime(val interface{})
	UseLatestRestorableTimeInput() interface{}
	VpcSecurityGroups() *[]*string
	SetVpcSecurityGroups(val *[]*string)
	VpcSecurityGroupsInput() *[]*string
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
	PutAdditionalStorageVolumes(value interface{})
	PutAssociatedRoles(value interface{})
	PutMasterUserSecret(value *RdsDbInstanceMasterUserSecret)
	PutProcessorFeatures(value interface{})
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
	ResetAdditionalStorageVolumes()
	ResetAllocatedStorage()
	ResetAllowMajorVersionUpgrade()
	ResetApplyImmediately()
	ResetAssociatedRoles()
	ResetAutomaticBackupReplicationKmsKeyId()
	ResetAutomaticBackupReplicationRegion()
	ResetAutomaticBackupReplicationRetentionPeriod()
	ResetAutoMinorVersionUpgrade()
	ResetAvailabilityZone()
	ResetBackupRetentionPeriod()
	ResetBackupTarget()
	ResetCaCertificateIdentifier()
	ResetCertificateRotationRestart()
	ResetCharacterSetName()
	ResetCopyTagsToSnapshot()
	ResetCustomIamInstanceProfile()
	ResetDatabaseInsightsMode()
	ResetDbClusterIdentifier()
	ResetDbClusterSnapshotIdentifier()
	ResetDbInstanceClass()
	ResetDbInstanceIdentifier()
	ResetDbName()
	ResetDbParameterGroupName()
	ResetDbSecurityGroups()
	ResetDbSnapshotIdentifier()
	ResetDbSubnetGroupName()
	ResetDbSystemId()
	ResetDedicatedLogVolume()
	ResetDeleteAutomatedBackups()
	ResetDeletionProtection()
	ResetDomain()
	ResetDomainAuthSecretArn()
	ResetDomainDnsIps()
	ResetDomainFqdn()
	ResetDomainIamRoleName()
	ResetDomainOu()
	ResetEnableCloudwatchLogsExports()
	ResetEnableIamDatabaseAuthentication()
	ResetEnablePerformanceInsights()
	ResetEngine()
	ResetEngineLifecycleSupport()
	ResetEngineVersion()
	ResetIops()
	ResetKmsKeyId()
	ResetLicenseModel()
	ResetManageMasterUserPassword()
	ResetMasterUserAuthenticationType()
	ResetMasterUsername()
	ResetMasterUserPassword()
	ResetMasterUserSecret()
	ResetMaxAllocatedStorage()
	ResetMonitoringInterval()
	ResetMonitoringRoleArn()
	ResetMultiAz()
	ResetNcharCharacterSetName()
	ResetNetworkType()
	ResetOptionGroupName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerformanceInsightsKmsKeyId()
	ResetPerformanceInsightsRetentionPeriod()
	ResetPort()
	ResetPreferredBackupWindow()
	ResetPreferredMaintenanceWindow()
	ResetProcessorFeatures()
	ResetPromotionTier()
	ResetPubliclyAccessible()
	ResetReplicaMode()
	ResetRestoreTime()
	ResetSourceDbClusterIdentifier()
	ResetSourceDbInstanceAutomatedBackupsArn()
	ResetSourceDbInstanceIdentifier()
	ResetSourceDbiResourceId()
	ResetSourceRegion()
	ResetStorageEncrypted()
	ResetStorageThroughput()
	ResetStorageType()
	ResetTags()
	ResetTdeCredentialArn()
	ResetTdeCredentialPassword()
	ResetTimezone()
	ResetUseDefaultProcessorFeatures()
	ResetUseLatestRestorableTime()
	ResetVpcSecurityGroups()
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

// The jsii proxy struct for RdsDbInstance
type jsiiProxy_RdsDbInstance struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_RdsDbInstance) AdditionalStorageVolumes() RdsDbInstanceAdditionalStorageVolumesList {
	var returns RdsDbInstanceAdditionalStorageVolumesList
	_jsii_.Get(
		j,
		"additionalStorageVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AdditionalStorageVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalStorageVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AllocatedStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AllocatedStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AllowMajorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AssociatedRoles() RdsDbInstanceAssociatedRolesList {
	var returns RdsDbInstanceAssociatedRolesList
	_jsii_.Get(
		j,
		"associatedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AssociatedRolesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutomaticBackupReplicationKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticBackupReplicationKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutomaticBackupReplicationKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticBackupReplicationKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutomaticBackupReplicationRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticBackupReplicationRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutomaticBackupReplicationRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticBackupReplicationRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutomaticBackupReplicationRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupReplicationRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutomaticBackupReplicationRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupReplicationRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutomaticRestartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticRestartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) BackupRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) BackupTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) BackupTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CaCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CaCertificateIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CertificateDetails() RdsDbInstanceCertificateDetailsOutputReference {
	var returns RdsDbInstanceCertificateDetailsOutputReference
	_jsii_.Get(
		j,
		"certificateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CertificateRotationRestart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateRotationRestart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CertificateRotationRestartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateRotationRestartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CharacterSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CopyTagsToSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CustomIamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CustomIamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DatabaseInsightsMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DatabaseInsightsModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbClusterSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbClusterSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbInstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbInstanceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbInstanceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DedicatedLogVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedLogVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DedicatedLogVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedLogVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DeleteAutomatedBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DeleteAutomatedBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainAuthSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainAuthSecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainDnsIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainDnsIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainFqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainIamRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainOuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnableCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnableCloudwatchLogsExportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnableIamDatabaseAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIamDatabaseAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnableIamDatabaseAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIamDatabaseAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnablePerformanceInsights() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceInsights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnablePerformanceInsightsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceInsightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Endpoint() RdsDbInstanceEndpointOutputReference {
	var returns RdsDbInstanceEndpointOutputReference
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EngineLifecycleSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EngineLifecycleSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) InstanceCreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) IsStorageConfigUpgradeAvailable() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isStorageConfigUpgradeAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) LatestRestorableTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ListenerEndpoint() RdsDbInstanceListenerEndpointOutputReference {
	var returns RdsDbInstanceListenerEndpointOutputReference
	_jsii_.Get(
		j,
		"listenerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ManageMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ManageMasterUserPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUserAuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserAuthenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUserAuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserAuthenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUserSecret() RdsDbInstanceMasterUserSecretOutputReference {
	var returns RdsDbInstanceMasterUserSecretOutputReference
	_jsii_.Get(
		j,
		"masterUserSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUserSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterUserSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MaxAllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MaxAllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MonitoringIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MonitoringRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) NcharCharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) NcharCharacterSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) OptionGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PercentProgress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"percentProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PerformanceInsightsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PerformanceInsightsRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PreferredBackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ProcessorFeatures() RdsDbInstanceProcessorFeaturesList {
	var returns RdsDbInstanceProcessorFeaturesList
	_jsii_.Get(
		j,
		"processorFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ProcessorFeaturesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processorFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PromotionTier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"promotionTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PromotionTierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"promotionTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ReadReplicaDbClusterIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readReplicaDbClusterIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ReadReplicaDbInstanceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readReplicaDbInstanceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ReplicaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ReplicaModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) RestoreTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) RestoreTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ResumeFullAutomationModeTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resumeFullAutomationModeTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SecondaryAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbInstanceAutomatedBackupsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceAutomatedBackupsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbInstanceAutomatedBackupsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceAutomatedBackupsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbInstanceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbiResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbiResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StatusInfos() RdsDbInstanceStatusInfosList {
	var returns RdsDbInstanceStatusInfosList
	_jsii_.Get(
		j,
		"statusInfos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Tags() RdsDbInstanceTagsList {
	var returns RdsDbInstanceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TdeCredentialArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TdeCredentialArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TdeCredentialPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TdeCredentialPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) UseDefaultProcessorFeatures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultProcessorFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) UseDefaultProcessorFeaturesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultProcessorFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) UseLatestRestorableTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) UseLatestRestorableTimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) VpcSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) VpcSecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_instance awscc_rds_db_instance} Resource.
func NewRdsDbInstance(scope constructs.Construct, id *string, config *RdsDbInstanceConfig) RdsDbInstance {
	_init_.Initialize()

	if err := validateNewRdsDbInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsDbInstance{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDbInstance.RdsDbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_instance awscc_rds_db_instance} Resource.
func NewRdsDbInstance_Override(r RdsDbInstance, scope constructs.Construct, id *string, config *RdsDbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rdsDbInstance.RdsDbInstance",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAllocatedStorage(val *string) {
	if err := j.validateSetAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAllowMajorVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowMajorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAutomaticBackupReplicationKmsKeyId(val *string) {
	if err := j.validateSetAutomaticBackupReplicationKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticBackupReplicationKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAutomaticBackupReplicationRegion(val *string) {
	if err := j.validateSetAutomaticBackupReplicationRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticBackupReplicationRegion",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAutomaticBackupReplicationRetentionPeriod(val *float64) {
	if err := j.validateSetAutomaticBackupReplicationRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticBackupReplicationRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetBackupRetentionPeriod(val *float64) {
	if err := j.validateSetBackupRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetBackupTarget(val *string) {
	if err := j.validateSetBackupTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupTarget",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCaCertificateIdentifier(val *string) {
	if err := j.validateSetCaCertificateIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCertificateRotationRestart(val interface{}) {
	if err := j.validateSetCertificateRotationRestartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateRotationRestart",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCharacterSetName(val *string) {
	if err := j.validateSetCharacterSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSetName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCustomIamInstanceProfile(val *string) {
	if err := j.validateSetCustomIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customIamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDatabaseInsightsMode(val *string) {
	if err := j.validateSetDatabaseInsightsModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseInsightsMode",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbClusterIdentifier(val *string) {
	if err := j.validateSetDbClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbClusterSnapshotIdentifier(val *string) {
	if err := j.validateSetDbClusterSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbInstanceClass(val *string) {
	if err := j.validateSetDbInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbInstanceClass",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbInstanceIdentifier(val *string) {
	if err := j.validateSetDbInstanceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbParameterGroupName(val *string) {
	if err := j.validateSetDbParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbSecurityGroups(val *[]*string) {
	if err := j.validateSetDbSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbSnapshotIdentifier(val *string) {
	if err := j.validateSetDbSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbSubnetGroupName(val *string) {
	if err := j.validateSetDbSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbSystemId(val *string) {
	if err := j.validateSetDbSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSystemId",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDedicatedLogVolume(val interface{}) {
	if err := j.validateSetDedicatedLogVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedLogVolume",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDeleteAutomatedBackups(val interface{}) {
	if err := j.validateSetDeleteAutomatedBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAutomatedBackups",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomainAuthSecretArn(val *string) {
	if err := j.validateSetDomainAuthSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAuthSecretArn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomainDnsIps(val *[]*string) {
	if err := j.validateSetDomainDnsIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainDnsIps",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomainFqdn(val *string) {
	if err := j.validateSetDomainFqdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainFqdn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomainIamRoleName(val *string) {
	if err := j.validateSetDomainIamRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomainOu(val *string) {
	if err := j.validateSetDomainOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainOu",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetEnableCloudwatchLogsExports(val *[]*string) {
	if err := j.validateSetEnableCloudwatchLogsExportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetEnableIamDatabaseAuthentication(val interface{}) {
	if err := j.validateSetEnableIamDatabaseAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIamDatabaseAuthentication",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetEnablePerformanceInsights(val interface{}) {
	if err := j.validateSetEnablePerformanceInsightsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePerformanceInsights",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetEngineLifecycleSupport(val *string) {
	if err := j.validateSetEngineLifecycleSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineLifecycleSupport",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetManageMasterUserPassword(val interface{}) {
	if err := j.validateSetManageMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterUserPassword",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMasterUserAuthenticationType(val *string) {
	if err := j.validateSetMasterUserAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserAuthenticationType",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMasterUsername(val *string) {
	if err := j.validateSetMasterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMasterUserPassword(val *string) {
	if err := j.validateSetMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMaxAllocatedStorage(val *float64) {
	if err := j.validateSetMaxAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllocatedStorage",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMonitoringInterval(val *float64) {
	if err := j.validateSetMonitoringIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInterval",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMonitoringRoleArn(val *string) {
	if err := j.validateSetMonitoringRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringRoleArn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMultiAz(val interface{}) {
	if err := j.validateSetMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetNcharCharacterSetName(val *string) {
	if err := j.validateSetNcharCharacterSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ncharCharacterSetName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetOptionGroupName(val *string) {
	if err := j.validateSetOptionGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPerformanceInsightsKmsKeyId(val *string) {
	if err := j.validateSetPerformanceInsightsKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPerformanceInsightsRetentionPeriod(val *float64) {
	if err := j.validateSetPerformanceInsightsRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPort(val *string) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPreferredBackupWindow(val *string) {
	if err := j.validateSetPreferredBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPromotionTier(val *float64) {
	if err := j.validateSetPromotionTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promotionTier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetReplicaMode(val *string) {
	if err := j.validateSetReplicaModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaMode",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetRestoreTime(val *string) {
	if err := j.validateSetRestoreTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreTime",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetSourceDbClusterIdentifier(val *string) {
	if err := j.validateSetSourceDbClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetSourceDbInstanceAutomatedBackupsArn(val *string) {
	if err := j.validateSetSourceDbInstanceAutomatedBackupsArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbInstanceAutomatedBackupsArn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetSourceDbInstanceIdentifier(val *string) {
	if err := j.validateSetSourceDbInstanceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetSourceDbiResourceId(val *string) {
	if err := j.validateSetSourceDbiResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbiResourceId",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetSourceRegion(val *string) {
	if err := j.validateSetSourceRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRegion",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetStorageThroughput(val *float64) {
	if err := j.validateSetStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageThroughput",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetTdeCredentialArn(val *string) {
	if err := j.validateSetTdeCredentialArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tdeCredentialArn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetTdeCredentialPassword(val *string) {
	if err := j.validateSetTdeCredentialPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tdeCredentialPassword",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetUseDefaultProcessorFeatures(val interface{}) {
	if err := j.validateSetUseDefaultProcessorFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDefaultProcessorFeatures",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetUseLatestRestorableTime(val interface{}) {
	if err := j.validateSetUseLatestRestorableTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLatestRestorableTime",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetVpcSecurityGroups(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroups",
		val,
	)
}

// Generates CDKTN code for importing a RdsDbInstance resource upon running "cdktn plan <stack-name>".
func RdsDbInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateRdsDbInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbInstance.RdsDbInstance",
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
func RdsDbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbInstance.RdsDbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsDbInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbInstance.RdsDbInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsDbInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.rdsDbInstance.RdsDbInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RdsDbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.rdsDbInstance.RdsDbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RdsDbInstance) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RdsDbInstance) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RdsDbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RdsDbInstance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RdsDbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RdsDbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RdsDbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RdsDbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RdsDbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RdsDbInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RdsDbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RdsDbInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RdsDbInstance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RdsDbInstance) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (r *jsiiProxy_RdsDbInstance) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsDbInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RdsDbInstance) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsDbInstance) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RdsDbInstance) PutAdditionalStorageVolumes(value interface{}) {
	if err := r.validatePutAdditionalStorageVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAdditionalStorageVolumes",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbInstance) PutAssociatedRoles(value interface{}) {
	if err := r.validatePutAssociatedRolesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAssociatedRoles",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbInstance) PutMasterUserSecret(value *RdsDbInstanceMasterUserSecret) {
	if err := r.validatePutMasterUserSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putMasterUserSecret",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbInstance) PutProcessorFeatures(value interface{}) {
	if err := r.validatePutProcessorFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putProcessorFeatures",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbInstance) PutTags(value interface{}) {
	if err := r.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTags",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbInstance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := r.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAdditionalStorageVolumes() {
	_jsii_.InvokeVoid(
		r,
		"resetAdditionalStorageVolumes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAllocatedStorage() {
	_jsii_.InvokeVoid(
		r,
		"resetAllocatedStorage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAllowMajorVersionUpgrade() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowMajorVersionUpgrade",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		r,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAssociatedRoles() {
	_jsii_.InvokeVoid(
		r,
		"resetAssociatedRoles",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAutomaticBackupReplicationKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetAutomaticBackupReplicationKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAutomaticBackupReplicationRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetAutomaticBackupReplicationRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAutomaticBackupReplicationRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetAutomaticBackupReplicationRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		r,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		r,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetBackupRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetBackupRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetBackupTarget() {
	_jsii_.InvokeVoid(
		r,
		"resetBackupTarget",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetCaCertificateIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetCaCertificateIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetCertificateRotationRestart() {
	_jsii_.InvokeVoid(
		r,
		"resetCertificateRotationRestart",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetCharacterSetName() {
	_jsii_.InvokeVoid(
		r,
		"resetCharacterSetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetCopyTagsToSnapshot() {
	_jsii_.InvokeVoid(
		r,
		"resetCopyTagsToSnapshot",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetCustomIamInstanceProfile() {
	_jsii_.InvokeVoid(
		r,
		"resetCustomIamInstanceProfile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDatabaseInsightsMode() {
	_jsii_.InvokeVoid(
		r,
		"resetDatabaseInsightsMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetDbClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbClusterSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetDbClusterSnapshotIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbInstanceClass() {
	_jsii_.InvokeVoid(
		r,
		"resetDbInstanceClass",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbInstanceIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetDbInstanceIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbParameterGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbParameterGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbSecurityGroups() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSecurityGroups",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSnapshotIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbSubnetGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSubnetGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbSystemId() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSystemId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDedicatedLogVolume() {
	_jsii_.InvokeVoid(
		r,
		"resetDedicatedLogVolume",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDeleteAutomatedBackups() {
	_jsii_.InvokeVoid(
		r,
		"resetDeleteAutomatedBackups",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		r,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomain() {
	_jsii_.InvokeVoid(
		r,
		"resetDomain",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomainAuthSecretArn() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainAuthSecretArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomainDnsIps() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainDnsIps",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomainFqdn() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainFqdn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomainIamRoleName() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainIamRoleName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomainOu() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainOu",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEnableCloudwatchLogsExports() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableCloudwatchLogsExports",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEnableIamDatabaseAuthentication() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableIamDatabaseAuthentication",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEnablePerformanceInsights() {
	_jsii_.InvokeVoid(
		r,
		"resetEnablePerformanceInsights",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEngine() {
	_jsii_.InvokeVoid(
		r,
		"resetEngine",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEngineLifecycleSupport() {
	_jsii_.InvokeVoid(
		r,
		"resetEngineLifecycleSupport",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetIops() {
	_jsii_.InvokeVoid(
		r,
		"resetIops",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		r,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetManageMasterUserPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetManageMasterUserPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMasterUserAuthenticationType() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUserAuthenticationType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMasterUsername() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUsername",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMasterUserPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUserPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMasterUserSecret() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUserSecret",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMaxAllocatedStorage() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxAllocatedStorage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMonitoringInterval() {
	_jsii_.InvokeVoid(
		r,
		"resetMonitoringInterval",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMonitoringRoleArn() {
	_jsii_.InvokeVoid(
		r,
		"resetMonitoringRoleArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMultiAz() {
	_jsii_.InvokeVoid(
		r,
		"resetMultiAz",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetNcharCharacterSetName() {
	_jsii_.InvokeVoid(
		r,
		"resetNcharCharacterSetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetNetworkType() {
	_jsii_.InvokeVoid(
		r,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetOptionGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetOptionGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPerformanceInsightsKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetPerformanceInsightsKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPerformanceInsightsRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetPerformanceInsightsRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPort() {
	_jsii_.InvokeVoid(
		r,
		"resetPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPreferredBackupWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetPreferredBackupWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetProcessorFeatures() {
	_jsii_.InvokeVoid(
		r,
		"resetProcessorFeatures",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPromotionTier() {
	_jsii_.InvokeVoid(
		r,
		"resetPromotionTier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		r,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetReplicaMode() {
	_jsii_.InvokeVoid(
		r,
		"resetReplicaMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetRestoreTime() {
	_jsii_.InvokeVoid(
		r,
		"resetRestoreTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetSourceDbClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceDbClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetSourceDbInstanceAutomatedBackupsArn() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceDbInstanceAutomatedBackupsArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetSourceDbInstanceIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceDbInstanceIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetSourceDbiResourceId() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceDbiResourceId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetSourceRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetStorageEncrypted() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageEncrypted",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetStorageThroughput() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageThroughput",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetStorageType() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetTdeCredentialArn() {
	_jsii_.InvokeVoid(
		r,
		"resetTdeCredentialArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetTdeCredentialPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetTdeCredentialPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetTimezone() {
	_jsii_.InvokeVoid(
		r,
		"resetTimezone",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetUseDefaultProcessorFeatures() {
	_jsii_.InvokeVoid(
		r,
		"resetUseDefaultProcessorFeatures",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetUseLatestRestorableTime() {
	_jsii_.InvokeVoid(
		r,
		"resetUseLatestRestorableTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetVpcSecurityGroups() {
	_jsii_.InvokeVoid(
		r,
		"resetVpcSecurityGroups",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

