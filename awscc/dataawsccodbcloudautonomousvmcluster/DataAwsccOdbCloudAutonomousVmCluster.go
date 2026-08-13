// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccodbcloudautonomousvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccodbcloudautonomousvmcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/data-sources/odb_cloud_autonomous_vm_cluster awscc_odb_cloud_autonomous_vm_cluster}.
type DataAwsccOdbCloudAutonomousVmCluster interface {
	cdktn.TerraformDataSource
	AutonomousDataStoragePercentage() *float64
	AutonomousDataStorageSizeInTBs() *float64
	AvailableAutonomousDataStorageSizeInTBs() *float64
	AvailableContainerDatabases() *float64
	AvailableCpus() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CloudAutonomousVmClusterArn() *string
	CloudAutonomousVmClusterId() *string
	CloudExadataInfrastructureId() *string
	ComputeModel() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuCoreCount() *float64
	CpuCoreCountPerNode() *float64
	CpuPercentage() *float64
	DataStorageSizeInGBs() *float64
	DataStorageSizeInTBs() *float64
	DbNodeStorageSizeInGBs() *float64
	DbServers() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	DisplayName() *string
	Domain() *string
	ExadataStorageInTBsLowestScaledValue() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	IamRoles() DataAwsccOdbCloudAutonomousVmClusterIamRolesList
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsMtlsEnabledVmCluster() cdktn.IResolvable
	LicenseModel() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaintenanceWindow() DataAwsccOdbCloudAutonomousVmClusterMaintenanceWindowOutputReference
	MaxAcdsLowestScaledValue() *float64
	MemoryPerOracleComputeUnitInGBs() *float64
	MemorySizeInGBs() *float64
	// The tree node.
	Node() constructs.Node
	NodeCount() *float64
	NonProvisionableAutonomousContainerDatabases() *float64
	Ocid() *string
	OciResourceAnchorName() *string
	OciUrl() *string
	OdbNetworkId() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProvisionableAutonomousContainerDatabases() *float64
	ProvisionedAutonomousContainerDatabases() *float64
	ProvisionedCpus() *float64
	// Experimental.
	RawOverrides() interface{}
	ReclaimableCpus() *float64
	ReservedCpus() *float64
	ScanListenerPortNonTls() *float64
	ScanListenerPortTls() *float64
	Shape() *string
	Tags() DataAwsccOdbCloudAutonomousVmClusterTagsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeZone() *string
	TotalContainerDatabases() *float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAwsccOdbCloudAutonomousVmCluster
type jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) AutonomousDataStoragePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autonomousDataStoragePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) AutonomousDataStorageSizeInTBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autonomousDataStorageSizeInTBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) AvailableAutonomousDataStorageSizeInTBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableAutonomousDataStorageSizeInTBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) AvailableContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) AvailableCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) CloudAutonomousVmClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudAutonomousVmClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) CloudAutonomousVmClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudAutonomousVmClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) CloudExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) CpuCoreCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) CpuPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) DataStorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) DataStorageSizeInTBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) DbNodeStorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) DbServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ExadataStorageInTBsLowestScaledValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"exadataStorageInTBsLowestScaledValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) IamRoles() DataAwsccOdbCloudAutonomousVmClusterIamRolesList {
	var returns DataAwsccOdbCloudAutonomousVmClusterIamRolesList
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) IsMtlsEnabledVmCluster() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isMtlsEnabledVmCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) MaintenanceWindow() DataAwsccOdbCloudAutonomousVmClusterMaintenanceWindowOutputReference {
	var returns DataAwsccOdbCloudAutonomousVmClusterMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) MaxAcdsLowestScaledValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAcdsLowestScaledValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) MemoryPerOracleComputeUnitInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) MemorySizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) NonProvisionableAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nonProvisionableAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) OdbNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ProvisionableAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionableAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ProvisionedAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ProvisionedCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ReclaimableCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reclaimableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ReservedCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ScanListenerPortNonTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortNonTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ScanListenerPortTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) Tags() DataAwsccOdbCloudAutonomousVmClusterTagsList {
	var returns DataAwsccOdbCloudAutonomousVmClusterTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) TotalContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalContainerDatabases",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/data-sources/odb_cloud_autonomous_vm_cluster awscc_odb_cloud_autonomous_vm_cluster} Data Source.
func NewDataAwsccOdbCloudAutonomousVmCluster(scope constructs.Construct, id *string, config *DataAwsccOdbCloudAutonomousVmClusterConfig) DataAwsccOdbCloudAutonomousVmCluster {
	_init_.Initialize()

	if err := validateNewDataAwsccOdbCloudAutonomousVmClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccOdbCloudAutonomousVmCluster.DataAwsccOdbCloudAutonomousVmCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/data-sources/odb_cloud_autonomous_vm_cluster awscc_odb_cloud_autonomous_vm_cluster} Data Source.
func NewDataAwsccOdbCloudAutonomousVmCluster_Override(d DataAwsccOdbCloudAutonomousVmCluster, scope constructs.Construct, id *string, config *DataAwsccOdbCloudAutonomousVmClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccOdbCloudAutonomousVmCluster.DataAwsccOdbCloudAutonomousVmCluster",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsccOdbCloudAutonomousVmCluster resource upon running "cdktn plan <stack-name>".
func DataAwsccOdbCloudAutonomousVmCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccOdbCloudAutonomousVmCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOdbCloudAutonomousVmCluster.DataAwsccOdbCloudAutonomousVmCluster",
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
func DataAwsccOdbCloudAutonomousVmCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccOdbCloudAutonomousVmCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOdbCloudAutonomousVmCluster.DataAwsccOdbCloudAutonomousVmCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccOdbCloudAutonomousVmCluster_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccOdbCloudAutonomousVmCluster_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOdbCloudAutonomousVmCluster.DataAwsccOdbCloudAutonomousVmCluster",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccOdbCloudAutonomousVmCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccOdbCloudAutonomousVmCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOdbCloudAutonomousVmCluster.DataAwsccOdbCloudAutonomousVmCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccOdbCloudAutonomousVmCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.dataAwsccOdbCloudAutonomousVmCluster.DataAwsccOdbCloudAutonomousVmCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudAutonomousVmCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

