// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccstoragegatewaysmbfileshare

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/dataawsccstoragegatewaysmbfileshare/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/data-sources/storagegateway_smb_file_share awscc_storagegateway_smb_file_share}.
type DataAwsccStoragegatewaySmbFileShare interface {
	cdktn.TerraformDataSource
	AccessBasedEnumeration() cdktn.IResolvable
	AdminUserList() *[]*string
	AuditDestinationArn() *string
	Authentication() *string
	BucketRegion() *string
	CacheAttributes() DataAwsccStoragegatewaySmbFileShareCacheAttributesOutputReference
	CaseSensitivity() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultStorageClass() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EncryptionType() *string
	FileShareArn() *string
	FileShareId() *string
	FileShareName() *string
	FileShareStatus() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayArn() *string
	GuessMimeTypeEnabled() cdktn.IResolvable
	Id() *string
	SetId(val *string)
	IdInput() *string
	InvalidUserList() *[]*string
	KmsEncrypted() cdktn.IResolvable
	KmsKey() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LocationArn() *string
	// The tree node.
	Node() constructs.Node
	NotificationPolicy() *string
	ObjectAcl() *string
	OplocksEnabled() cdktn.IResolvable
	Path() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ReadOnly() cdktn.IResolvable
	RequesterPays() cdktn.IResolvable
	Role() *string
	SmbaclEnabled() cdktn.IResolvable
	Tags() DataAwsccStoragegatewaySmbFileShareTagsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ValidUserList() *[]*string
	VpcEndpointDnsName() *string
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

// The jsii proxy struct for DataAwsccStoragegatewaySmbFileShare
type jsiiProxy_DataAwsccStoragegatewaySmbFileShare struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) AccessBasedEnumeration() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"accessBasedEnumeration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) AdminUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) AuditDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) Authentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) BucketRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) CacheAttributes() DataAwsccStoragegatewaySmbFileShareCacheAttributesOutputReference {
	var returns DataAwsccStoragegatewaySmbFileShareCacheAttributesOutputReference
	_jsii_.Get(
		j,
		"cacheAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) CaseSensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caseSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) DefaultStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) FileShareArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) FileShareId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) FileShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) FileShareStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) GuessMimeTypeEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"guessMimeTypeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) InvalidUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) KmsEncrypted() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) LocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) NotificationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) OplocksEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"oplocksEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) ReadOnly() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) RequesterPays() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"requesterPays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) SmbaclEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"smbaclEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) Tags() DataAwsccStoragegatewaySmbFileShareTagsList {
	var returns DataAwsccStoragegatewaySmbFileShareTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) ValidUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) VpcEndpointDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsName",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/data-sources/storagegateway_smb_file_share awscc_storagegateway_smb_file_share} Data Source.
func NewDataAwsccStoragegatewaySmbFileShare(scope constructs.Construct, id *string, config *DataAwsccStoragegatewaySmbFileShareConfig) DataAwsccStoragegatewaySmbFileShare {
	_init_.Initialize()

	if err := validateNewDataAwsccStoragegatewaySmbFileShareParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccStoragegatewaySmbFileShare{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccStoragegatewaySmbFileShare.DataAwsccStoragegatewaySmbFileShare",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/data-sources/storagegateway_smb_file_share awscc_storagegateway_smb_file_share} Data Source.
func NewDataAwsccStoragegatewaySmbFileShare_Override(d DataAwsccStoragegatewaySmbFileShare, scope constructs.Construct, id *string, config *DataAwsccStoragegatewaySmbFileShareConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccStoragegatewaySmbFileShare.DataAwsccStoragegatewaySmbFileShare",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccStoragegatewaySmbFileShare)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsccStoragegatewaySmbFileShare resource upon running "cdktn plan <stack-name>".
func DataAwsccStoragegatewaySmbFileShare_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccStoragegatewaySmbFileShare_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccStoragegatewaySmbFileShare.DataAwsccStoragegatewaySmbFileShare",
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
func DataAwsccStoragegatewaySmbFileShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccStoragegatewaySmbFileShare_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccStoragegatewaySmbFileShare.DataAwsccStoragegatewaySmbFileShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccStoragegatewaySmbFileShare_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccStoragegatewaySmbFileShare_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccStoragegatewaySmbFileShare.DataAwsccStoragegatewaySmbFileShare",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccStoragegatewaySmbFileShare_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccStoragegatewaySmbFileShare_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccStoragegatewaySmbFileShare.DataAwsccStoragegatewaySmbFileShare",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccStoragegatewaySmbFileShare_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.dataAwsccStoragegatewaySmbFileShare.DataAwsccStoragegatewaySmbFileShare",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccStoragegatewaySmbFileShare) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

