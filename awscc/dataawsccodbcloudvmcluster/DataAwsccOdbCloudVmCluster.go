// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccodbcloudvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccodbcloudvmcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/data-sources/odb_cloud_vm_cluster awscc_odb_cloud_vm_cluster}.
type DataAwsccOdbCloudVmCluster interface {
	cdktn.TerraformDataSource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CloudExadataInfrastructureId() *string
	CloudVmClusterArn() *string
	CloudVmClusterId() *string
	ClusterName() *string
	ComputeModel() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuCoreCount() *float64
	DataCollectionOptions() DataAwsccOdbCloudVmClusterDataCollectionOptionsOutputReference
	DataStorageSizeInTBs() *float64
	DbNodes() DataAwsccOdbCloudVmClusterDbNodesList
	DbNodeStorageSizeInGBs() *float64
	DbServers() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskRedundancy() *string
	DisplayName() *string
	Domain() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GiVersion() *string
	Hostname() *string
	IamRoles() DataAwsccOdbCloudVmClusterIamRolesList
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsLocalBackupEnabled() cdktn.IResolvable
	IsSparseDiskgroupEnabled() cdktn.IResolvable
	LicenseModel() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ListenerPort() *float64
	MemorySizeInGBs() *float64
	// The tree node.
	Node() constructs.Node
	NodeCount() *float64
	Ocid() *string
	OciResourceAnchorName() *string
	OciUrl() *string
	OdbNetworkId() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ScanDnsName() *string
	ScanIpIds() *[]*string
	ScanListenerPortTcp() *float64
	Shape() *string
	SshPublicKeys() *[]*string
	StorageSizeInGBs() *float64
	SystemVersion() *string
	Tags() DataAwsccOdbCloudVmClusterTagsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeZone() *string
	VipIds() *[]*string
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

// The jsii proxy struct for DataAwsccOdbCloudVmCluster
type jsiiProxy_DataAwsccOdbCloudVmCluster struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) CloudExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) CloudVmClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudVmClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) CloudVmClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudVmClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) DataCollectionOptions() DataAwsccOdbCloudVmClusterDataCollectionOptionsOutputReference {
	var returns DataAwsccOdbCloudVmClusterDataCollectionOptionsOutputReference
	_jsii_.Get(
		j,
		"dataCollectionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) DataStorageSizeInTBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) DbNodes() DataAwsccOdbCloudVmClusterDbNodesList {
	var returns DataAwsccOdbCloudVmClusterDbNodesList
	_jsii_.Get(
		j,
		"dbNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) DbNodeStorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) DbServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) DiskRedundancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskRedundancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) GiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) IamRoles() DataAwsccOdbCloudVmClusterIamRolesList {
	var returns DataAwsccOdbCloudVmClusterIamRolesList
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) IsLocalBackupEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isLocalBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) IsSparseDiskgroupEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isSparseDiskgroupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) ListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"listenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) MemorySizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) OdbNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) ScanDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) ScanIpIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scanIpIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) ScanListenerPortTcp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) StorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) SystemVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) Tags() DataAwsccOdbCloudVmClusterTagsList {
	var returns DataAwsccOdbCloudVmClusterTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster) VipIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vipIds",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/data-sources/odb_cloud_vm_cluster awscc_odb_cloud_vm_cluster} Data Source.
func NewDataAwsccOdbCloudVmCluster(scope constructs.Construct, id *string, config *DataAwsccOdbCloudVmClusterConfig) DataAwsccOdbCloudVmCluster {
	_init_.Initialize()

	if err := validateNewDataAwsccOdbCloudVmClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccOdbCloudVmCluster{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccOdbCloudVmCluster.DataAwsccOdbCloudVmCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/data-sources/odb_cloud_vm_cluster awscc_odb_cloud_vm_cluster} Data Source.
func NewDataAwsccOdbCloudVmCluster_Override(d DataAwsccOdbCloudVmCluster, scope constructs.Construct, id *string, config *DataAwsccOdbCloudVmClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccOdbCloudVmCluster.DataAwsccOdbCloudVmCluster",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOdbCloudVmCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsccOdbCloudVmCluster resource upon running "cdktn plan <stack-name>".
func DataAwsccOdbCloudVmCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccOdbCloudVmCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOdbCloudVmCluster.DataAwsccOdbCloudVmCluster",
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
func DataAwsccOdbCloudVmCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccOdbCloudVmCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOdbCloudVmCluster.DataAwsccOdbCloudVmCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccOdbCloudVmCluster_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccOdbCloudVmCluster_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOdbCloudVmCluster.DataAwsccOdbCloudVmCluster",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccOdbCloudVmCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccOdbCloudVmCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOdbCloudVmCluster.DataAwsccOdbCloudVmCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccOdbCloudVmCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.dataAwsccOdbCloudVmCluster.DataAwsccOdbCloudVmCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOdbCloudVmCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

