// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccnetworkmanagervpcattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccnetworkmanagervpcattachment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/data-sources/networkmanager_vpc_attachment awscc_networkmanager_vpc_attachment}.
type DataAwsccNetworkmanagerVpcAttachment interface {
	cdktn.TerraformDataSource
	AttachmentId() *string
	AttachmentPolicyRuleNumber() *float64
	AttachmentType() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CoreNetworkArn() *string
	CoreNetworkId() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EdgeLocation() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastModificationErrors() *[]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	NetworkFunctionGroupName() *string
	// The tree node.
	Node() constructs.Node
	Options() DataAwsccNetworkmanagerVpcAttachmentOptionsOutputReference
	OwnerAccountId() *string
	ProposedNetworkFunctionGroupChange() DataAwsccNetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference
	ProposedSegmentChange() DataAwsccNetworkmanagerVpcAttachmentProposedSegmentChangeOutputReference
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ResourceArn() *string
	RoutingPolicyLabel() *string
	SegmentName() *string
	State() *string
	SubnetArns() *[]*string
	Tags() DataAwsccNetworkmanagerVpcAttachmentTagsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedAt() *string
	VpcArn() *string
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

// The jsii proxy struct for DataAwsccNetworkmanagerVpcAttachment
type jsiiProxy_DataAwsccNetworkmanagerVpcAttachment struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) AttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) AttachmentPolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attachmentPolicyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) AttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) CoreNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) CoreNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) EdgeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) LastModificationErrors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lastModificationErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) NetworkFunctionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkFunctionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) Options() DataAwsccNetworkmanagerVpcAttachmentOptionsOutputReference {
	var returns DataAwsccNetworkmanagerVpcAttachmentOptionsOutputReference
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) OwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) ProposedNetworkFunctionGroupChange() DataAwsccNetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference {
	var returns DataAwsccNetworkmanagerVpcAttachmentProposedNetworkFunctionGroupChangeOutputReference
	_jsii_.Get(
		j,
		"proposedNetworkFunctionGroupChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) ProposedSegmentChange() DataAwsccNetworkmanagerVpcAttachmentProposedSegmentChangeOutputReference {
	var returns DataAwsccNetworkmanagerVpcAttachmentProposedSegmentChangeOutputReference
	_jsii_.Get(
		j,
		"proposedSegmentChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) RoutingPolicyLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingPolicyLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) SegmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) SubnetArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) Tags() DataAwsccNetworkmanagerVpcAttachmentTagsList {
	var returns DataAwsccNetworkmanagerVpcAttachmentTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) VpcArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcArn",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/data-sources/networkmanager_vpc_attachment awscc_networkmanager_vpc_attachment} Data Source.
func NewDataAwsccNetworkmanagerVpcAttachment(scope constructs.Construct, id *string, config *DataAwsccNetworkmanagerVpcAttachmentConfig) DataAwsccNetworkmanagerVpcAttachment {
	_init_.Initialize()

	if err := validateNewDataAwsccNetworkmanagerVpcAttachmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccNetworkmanagerVpcAttachment{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccNetworkmanagerVpcAttachment.DataAwsccNetworkmanagerVpcAttachment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/data-sources/networkmanager_vpc_attachment awscc_networkmanager_vpc_attachment} Data Source.
func NewDataAwsccNetworkmanagerVpcAttachment_Override(d DataAwsccNetworkmanagerVpcAttachment, scope constructs.Construct, id *string, config *DataAwsccNetworkmanagerVpcAttachmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccNetworkmanagerVpcAttachment.DataAwsccNetworkmanagerVpcAttachment",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsccNetworkmanagerVpcAttachment resource upon running "cdktn plan <stack-name>".
func DataAwsccNetworkmanagerVpcAttachment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccNetworkmanagerVpcAttachment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccNetworkmanagerVpcAttachment.DataAwsccNetworkmanagerVpcAttachment",
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
func DataAwsccNetworkmanagerVpcAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccNetworkmanagerVpcAttachment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccNetworkmanagerVpcAttachment.DataAwsccNetworkmanagerVpcAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccNetworkmanagerVpcAttachment_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccNetworkmanagerVpcAttachment_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccNetworkmanagerVpcAttachment.DataAwsccNetworkmanagerVpcAttachment",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccNetworkmanagerVpcAttachment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccNetworkmanagerVpcAttachment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccNetworkmanagerVpcAttachment.DataAwsccNetworkmanagerVpcAttachment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccNetworkmanagerVpcAttachment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.dataAwsccNetworkmanagerVpcAttachment.DataAwsccNetworkmanagerVpcAttachment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccNetworkmanagerVpcAttachment) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

