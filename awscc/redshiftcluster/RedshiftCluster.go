// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/redshiftcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_cluster awscc_redshift_cluster}.
type RedshiftCluster interface {
	cdktn.TerraformResource
	AllowVersionUpgrade() interface{}
	SetAllowVersionUpgrade(val interface{})
	AllowVersionUpgradeInput() interface{}
	AquaConfigurationStatus() *string
	SetAquaConfigurationStatus(val *string)
	AquaConfigurationStatusInput() *string
	AutomatedSnapshotRetentionPeriod() *float64
	SetAutomatedSnapshotRetentionPeriod(val *float64)
	AutomatedSnapshotRetentionPeriodInput() *float64
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	AvailabilityZoneRelocation() interface{}
	SetAvailabilityZoneRelocation(val interface{})
	AvailabilityZoneRelocationInput() interface{}
	AvailabilityZoneRelocationStatus() *string
	SetAvailabilityZoneRelocationStatus(val *string)
	AvailabilityZoneRelocationStatusInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Classic() interface{}
	SetClassic(val interface{})
	ClassicInput() interface{}
	ClusterIdentifier() *string
	SetClusterIdentifier(val *string)
	ClusterIdentifierInput() *string
	ClusterNamespaceArn() *string
	ClusterParameterGroupName() *string
	SetClusterParameterGroupName(val *string)
	ClusterParameterGroupNameInput() *string
	ClusterSecurityGroups() *[]*string
	SetClusterSecurityGroups(val *[]*string)
	ClusterSecurityGroupsInput() *[]*string
	ClusterSubnetGroupName() *string
	SetClusterSubnetGroupName(val *string)
	ClusterSubnetGroupNameInput() *string
	ClusterType() *string
	SetClusterType(val *string)
	ClusterTypeInput() *string
	ClusterVersion() *string
	SetClusterVersion(val *string)
	ClusterVersionInput() *string
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
	DbName() *string
	SetDbName(val *string)
	DbNameInput() *string
	DeferMaintenance() interface{}
	SetDeferMaintenance(val interface{})
	DeferMaintenanceDuration() *float64
	SetDeferMaintenanceDuration(val *float64)
	DeferMaintenanceDurationInput() *float64
	DeferMaintenanceEndTime() *string
	SetDeferMaintenanceEndTime(val *string)
	DeferMaintenanceEndTimeInput() *string
	DeferMaintenanceIdentifier() *string
	DeferMaintenanceInput() interface{}
	DeferMaintenanceStartTime() *string
	SetDeferMaintenanceStartTime(val *string)
	DeferMaintenanceStartTimeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestinationRegion() *string
	SetDestinationRegion(val *string)
	DestinationRegionInput() *string
	ElasticIp() *string
	SetElasticIp(val *string)
	ElasticIpInput() *string
	Encrypted() interface{}
	SetEncrypted(val interface{})
	EncryptedInput() interface{}
	Endpoint() RedshiftClusterEndpointOutputReference
	EndpointInput() interface{}
	EnhancedVpcRouting() interface{}
	SetEnhancedVpcRouting(val interface{})
	EnhancedVpcRoutingInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HsmClientCertificateIdentifier() *string
	SetHsmClientCertificateIdentifier(val *string)
	HsmClientCertificateIdentifierInput() *string
	HsmConfigurationIdentifier() *string
	SetHsmConfigurationIdentifier(val *string)
	HsmConfigurationIdentifierInput() *string
	IamRoles() *[]*string
	SetIamRoles(val *[]*string)
	IamRolesInput() *[]*string
	Id() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LoggingProperties() RedshiftClusterLoggingPropertiesOutputReference
	LoggingPropertiesInput() interface{}
	MaintenanceTrackName() *string
	SetMaintenanceTrackName(val *string)
	MaintenanceTrackNameInput() *string
	ManageMasterPassword() interface{}
	SetManageMasterPassword(val interface{})
	ManageMasterPasswordInput() interface{}
	ManualSnapshotRetentionPeriod() *float64
	SetManualSnapshotRetentionPeriod(val *float64)
	ManualSnapshotRetentionPeriodInput() *float64
	MasterPasswordSecretArn() *string
	MasterPasswordSecretKmsKeyId() *string
	SetMasterPasswordSecretKmsKeyId(val *string)
	MasterPasswordSecretKmsKeyIdInput() *string
	MasterUsername() *string
	SetMasterUsername(val *string)
	MasterUsernameInput() *string
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	MasterUserPasswordInput() *string
	MultiAz() interface{}
	SetMultiAz(val interface{})
	MultiAzInput() interface{}
	NamespaceResourcePolicy() *string
	SetNamespaceResourcePolicy(val *string)
	NamespaceResourcePolicyInput() *string
	// The tree node.
	Node() constructs.Node
	NodeType() *string
	SetNodeType(val *string)
	NodeTypeInput() *string
	NumberOfNodes() *float64
	SetNumberOfNodes(val *float64)
	NumberOfNodesInput() *float64
	OwnerAccount() *string
	SetOwnerAccount(val *string)
	OwnerAccountInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	ResourceAction() *string
	SetResourceAction(val *string)
	ResourceActionInput() *string
	RevisionTarget() *string
	SetRevisionTarget(val *string)
	RevisionTargetInput() *string
	RotateEncryptionKey() interface{}
	SetRotateEncryptionKey(val interface{})
	RotateEncryptionKeyInput() interface{}
	SnapshotClusterIdentifier() *string
	SetSnapshotClusterIdentifier(val *string)
	SnapshotClusterIdentifierInput() *string
	SnapshotCopyGrantName() *string
	SetSnapshotCopyGrantName(val *string)
	SnapshotCopyGrantNameInput() *string
	SnapshotCopyManual() interface{}
	SetSnapshotCopyManual(val interface{})
	SnapshotCopyManualInput() interface{}
	SnapshotCopyRetentionPeriod() *float64
	SetSnapshotCopyRetentionPeriod(val *float64)
	SnapshotCopyRetentionPeriodInput() *float64
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	SnapshotIdentifierInput() *string
	Tags() RedshiftClusterTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutEndpoint(value *RedshiftClusterEndpoint)
	PutLoggingProperties(value *RedshiftClusterLoggingProperties)
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
	ResetAllowVersionUpgrade()
	ResetAquaConfigurationStatus()
	ResetAutomatedSnapshotRetentionPeriod()
	ResetAvailabilityZone()
	ResetAvailabilityZoneRelocation()
	ResetAvailabilityZoneRelocationStatus()
	ResetClassic()
	ResetClusterIdentifier()
	ResetClusterParameterGroupName()
	ResetClusterSecurityGroups()
	ResetClusterSubnetGroupName()
	ResetClusterVersion()
	ResetDeferMaintenance()
	ResetDeferMaintenanceDuration()
	ResetDeferMaintenanceEndTime()
	ResetDeferMaintenanceStartTime()
	ResetDestinationRegion()
	ResetElasticIp()
	ResetEncrypted()
	ResetEndpoint()
	ResetEnhancedVpcRouting()
	ResetHsmClientCertificateIdentifier()
	ResetHsmConfigurationIdentifier()
	ResetIamRoles()
	ResetKmsKeyId()
	ResetLoggingProperties()
	ResetMaintenanceTrackName()
	ResetManageMasterPassword()
	ResetManualSnapshotRetentionPeriod()
	ResetMasterPasswordSecretKmsKeyId()
	ResetMasterUserPassword()
	ResetMultiAz()
	ResetNamespaceResourcePolicy()
	ResetNumberOfNodes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwnerAccount()
	ResetPort()
	ResetPreferredMaintenanceWindow()
	ResetPubliclyAccessible()
	ResetResourceAction()
	ResetRevisionTarget()
	ResetRotateEncryptionKey()
	ResetSnapshotClusterIdentifier()
	ResetSnapshotCopyGrantName()
	ResetSnapshotCopyManual()
	ResetSnapshotCopyRetentionPeriod()
	ResetSnapshotIdentifier()
	ResetTags()
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

// The jsii proxy struct for RedshiftCluster
type jsiiProxy_RedshiftCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_RedshiftCluster) AllowVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AllowVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AquaConfigurationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aquaConfigurationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AquaConfigurationStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aquaConfigurationStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AutomatedSnapshotRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automatedSnapshotRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AutomatedSnapshotRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automatedSnapshotRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AvailabilityZoneRelocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneRelocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AvailabilityZoneRelocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneRelocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AvailabilityZoneRelocationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneRelocationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AvailabilityZoneRelocationStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneRelocationStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Classic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"classic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClassicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"classicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterSecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DeferMaintenance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deferMaintenance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DeferMaintenanceDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deferMaintenanceDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DeferMaintenanceDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deferMaintenanceDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DeferMaintenanceEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deferMaintenanceEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DeferMaintenanceEndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deferMaintenanceEndTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DeferMaintenanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deferMaintenanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DeferMaintenanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deferMaintenanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DeferMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deferMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DeferMaintenanceStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deferMaintenanceStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DestinationRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DestinationRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ElasticIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ElasticIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) EncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Endpoint() RedshiftClusterEndpointOutputReference {
	var returns RedshiftClusterEndpointOutputReference
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) EndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) EnhancedVpcRouting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedVpcRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) EnhancedVpcRoutingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedVpcRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) HsmClientCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hsmClientCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) HsmClientCertificateIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hsmClientCertificateIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) HsmConfigurationIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hsmConfigurationIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) HsmConfigurationIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hsmConfigurationIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) IamRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) IamRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) LoggingProperties() RedshiftClusterLoggingPropertiesOutputReference {
	var returns RedshiftClusterLoggingPropertiesOutputReference
	_jsii_.Get(
		j,
		"loggingProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) LoggingPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MaintenanceTrackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceTrackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MaintenanceTrackNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceTrackNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ManageMasterPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ManageMasterPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ManualSnapshotRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manualSnapshotRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ManualSnapshotRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manualSnapshotRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MasterPasswordSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MasterPasswordSecretKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordSecretKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MasterPasswordSecretKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordSecretKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MasterUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MasterUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) NamespaceResourcePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceResourcePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) NamespaceResourcePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceResourcePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) NumberOfNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) NumberOfNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) OwnerAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) OwnerAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ResourceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ResourceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) RevisionTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) RevisionTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) RotateEncryptionKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotateEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) RotateEncryptionKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotateEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotCopyGrantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotCopyGrantName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotCopyGrantNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotCopyGrantNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotCopyManual() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotCopyManual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotCopyManualInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotCopyManualInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotCopyRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotCopyRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotCopyRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotCopyRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Tags() RedshiftClusterTagsList {
	var returns RedshiftClusterTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_cluster awscc_redshift_cluster} Resource.
func NewRedshiftCluster(scope constructs.Construct, id *string, config *RedshiftClusterConfig) RedshiftCluster {
	_init_.Initialize()

	if err := validateNewRedshiftClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftCluster{}

	_jsii_.Create(
		"@cdktn/provider-awscc.redshiftCluster.RedshiftCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/redshift_cluster awscc_redshift_cluster} Resource.
func NewRedshiftCluster_Override(r RedshiftCluster, scope constructs.Construct, id *string, config *RedshiftClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.redshiftCluster.RedshiftCluster",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetAllowVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetAquaConfigurationStatus(val *string) {
	if err := j.validateSetAquaConfigurationStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aquaConfigurationStatus",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetAutomatedSnapshotRetentionPeriod(val *float64) {
	if err := j.validateSetAutomatedSnapshotRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automatedSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetAvailabilityZoneRelocation(val interface{}) {
	if err := j.validateSetAvailabilityZoneRelocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneRelocation",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetAvailabilityZoneRelocationStatus(val *string) {
	if err := j.validateSetAvailabilityZoneRelocationStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneRelocationStatus",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClassic(val interface{}) {
	if err := j.validateSetClassicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classic",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterIdentifier(val *string) {
	if err := j.validateSetClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterParameterGroupName(val *string) {
	if err := j.validateSetClusterParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterSecurityGroups(val *[]*string) {
	if err := j.validateSetClusterSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterSubnetGroupName(val *string) {
	if err := j.validateSetClusterSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterType(val *string) {
	if err := j.validateSetClusterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterType",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterVersion(val *string) {
	if err := j.validateSetClusterVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetDeferMaintenance(val interface{}) {
	if err := j.validateSetDeferMaintenanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deferMaintenance",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetDeferMaintenanceDuration(val *float64) {
	if err := j.validateSetDeferMaintenanceDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deferMaintenanceDuration",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetDeferMaintenanceEndTime(val *string) {
	if err := j.validateSetDeferMaintenanceEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deferMaintenanceEndTime",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetDeferMaintenanceStartTime(val *string) {
	if err := j.validateSetDeferMaintenanceStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deferMaintenanceStartTime",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetDestinationRegion(val *string) {
	if err := j.validateSetDestinationRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationRegion",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetElasticIp(val *string) {
	if err := j.validateSetElasticIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticIp",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetEncrypted(val interface{}) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetEnhancedVpcRouting(val interface{}) {
	if err := j.validateSetEnhancedVpcRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedVpcRouting",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetHsmClientCertificateIdentifier(val *string) {
	if err := j.validateSetHsmClientCertificateIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hsmClientCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetHsmConfigurationIdentifier(val *string) {
	if err := j.validateSetHsmConfigurationIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hsmConfigurationIdentifier",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetIamRoles(val *[]*string) {
	if err := j.validateSetIamRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoles",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetMaintenanceTrackName(val *string) {
	if err := j.validateSetMaintenanceTrackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceTrackName",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetManageMasterPassword(val interface{}) {
	if err := j.validateSetManageMasterPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterPassword",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetManualSnapshotRetentionPeriod(val *float64) {
	if err := j.validateSetManualSnapshotRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manualSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetMasterPasswordSecretKmsKeyId(val *string) {
	if err := j.validateSetMasterPasswordSecretKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPasswordSecretKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetMasterUsername(val *string) {
	if err := j.validateSetMasterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetMasterUserPassword(val *string) {
	if err := j.validateSetMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetMultiAz(val interface{}) {
	if err := j.validateSetMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetNamespaceResourcePolicy(val *string) {
	if err := j.validateSetNamespaceResourcePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceResourcePolicy",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetNumberOfNodes(val *float64) {
	if err := j.validateSetNumberOfNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfNodes",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetOwnerAccount(val *string) {
	if err := j.validateSetOwnerAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownerAccount",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetResourceAction(val *string) {
	if err := j.validateSetResourceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceAction",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetRevisionTarget(val *string) {
	if err := j.validateSetRevisionTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revisionTarget",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetRotateEncryptionKey(val interface{}) {
	if err := j.validateSetRotateEncryptionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotateEncryptionKey",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetSnapshotClusterIdentifier(val *string) {
	if err := j.validateSetSnapshotClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetSnapshotCopyGrantName(val *string) {
	if err := j.validateSetSnapshotCopyGrantNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotCopyGrantName",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetSnapshotCopyManual(val interface{}) {
	if err := j.validateSetSnapshotCopyManualParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotCopyManual",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetSnapshotCopyRetentionPeriod(val *float64) {
	if err := j.validateSetSnapshotCopyRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotCopyRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetSnapshotIdentifier(val *string) {
	if err := j.validateSetSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a RedshiftCluster resource upon running "cdktn plan <stack-name>".
func RedshiftCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateRedshiftCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.redshiftCluster.RedshiftCluster",
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
func RedshiftCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.redshiftCluster.RedshiftCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedshiftCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.redshiftCluster.RedshiftCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedshiftCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.redshiftCluster.RedshiftCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RedshiftCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.redshiftCluster.RedshiftCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RedshiftCluster) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RedshiftCluster) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RedshiftCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RedshiftCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RedshiftCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RedshiftCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RedshiftCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RedshiftCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RedshiftCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RedshiftCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RedshiftCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RedshiftCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RedshiftCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (r *jsiiProxy_RedshiftCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (r *jsiiProxy_RedshiftCluster) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RedshiftCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RedshiftCluster) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RedshiftCluster) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RedshiftCluster) PutEndpoint(value *RedshiftClusterEndpoint) {
	if err := r.validatePutEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEndpoint",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftCluster) PutLoggingProperties(value *RedshiftClusterLoggingProperties) {
	if err := r.validatePutLoggingPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putLoggingProperties",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftCluster) PutTags(value interface{}) {
	if err := r.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTags",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := r.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetAllowVersionUpgrade() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowVersionUpgrade",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetAquaConfigurationStatus() {
	_jsii_.InvokeVoid(
		r,
		"resetAquaConfigurationStatus",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetAutomatedSnapshotRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetAutomatedSnapshotRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		r,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetAvailabilityZoneRelocation() {
	_jsii_.InvokeVoid(
		r,
		"resetAvailabilityZoneRelocation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetAvailabilityZoneRelocationStatus() {
	_jsii_.InvokeVoid(
		r,
		"resetAvailabilityZoneRelocationStatus",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClassic() {
	_jsii_.InvokeVoid(
		r,
		"resetClassic",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClusterParameterGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterParameterGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClusterSecurityGroups() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterSecurityGroups",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClusterSubnetGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterSubnetGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClusterVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetDeferMaintenance() {
	_jsii_.InvokeVoid(
		r,
		"resetDeferMaintenance",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetDeferMaintenanceDuration() {
	_jsii_.InvokeVoid(
		r,
		"resetDeferMaintenanceDuration",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetDeferMaintenanceEndTime() {
	_jsii_.InvokeVoid(
		r,
		"resetDeferMaintenanceEndTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetDeferMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		r,
		"resetDeferMaintenanceStartTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetDestinationRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetDestinationRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetElasticIp() {
	_jsii_.InvokeVoid(
		r,
		"resetElasticIp",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetEncrypted() {
	_jsii_.InvokeVoid(
		r,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetEndpoint() {
	_jsii_.InvokeVoid(
		r,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetEnhancedVpcRouting() {
	_jsii_.InvokeVoid(
		r,
		"resetEnhancedVpcRouting",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetHsmClientCertificateIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetHsmClientCertificateIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetHsmConfigurationIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetHsmConfigurationIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetIamRoles() {
	_jsii_.InvokeVoid(
		r,
		"resetIamRoles",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetLoggingProperties() {
	_jsii_.InvokeVoid(
		r,
		"resetLoggingProperties",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetMaintenanceTrackName() {
	_jsii_.InvokeVoid(
		r,
		"resetMaintenanceTrackName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetManageMasterPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetManageMasterPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetManualSnapshotRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetManualSnapshotRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetMasterPasswordSecretKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterPasswordSecretKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetMasterUserPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUserPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetMultiAz() {
	_jsii_.InvokeVoid(
		r,
		"resetMultiAz",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetNamespaceResourcePolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetNamespaceResourcePolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetNumberOfNodes() {
	_jsii_.InvokeVoid(
		r,
		"resetNumberOfNodes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetOwnerAccount() {
	_jsii_.InvokeVoid(
		r,
		"resetOwnerAccount",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetPort() {
	_jsii_.InvokeVoid(
		r,
		"resetPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		r,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetResourceAction() {
	_jsii_.InvokeVoid(
		r,
		"resetResourceAction",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetRevisionTarget() {
	_jsii_.InvokeVoid(
		r,
		"resetRevisionTarget",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetRotateEncryptionKey() {
	_jsii_.InvokeVoid(
		r,
		"resetRotateEncryptionKey",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetSnapshotClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSnapshotClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetSnapshotCopyGrantName() {
	_jsii_.InvokeVoid(
		r,
		"resetSnapshotCopyGrantName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetSnapshotCopyManual() {
	_jsii_.InvokeVoid(
		r,
		"resetSnapshotCopyManual",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetSnapshotCopyRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetSnapshotCopyRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSnapshotIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		r,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

