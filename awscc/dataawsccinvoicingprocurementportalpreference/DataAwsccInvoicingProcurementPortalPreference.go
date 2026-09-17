// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccinvoicingprocurementportalpreference

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccinvoicingprocurementportalpreference/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/data-sources/invoicing_procurement_portal_preference awscc_invoicing_procurement_portal_preference}.
type DataAwsccInvoicingProcurementPortalPreference interface {
	cdktn.TerraformDataSource
	AwsAccountId() *string
	BuyerDomain() *string
	BuyerIdentifier() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Contacts() DataAwsccInvoicingProcurementPortalPreferenceContactsList
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateDate() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EinvoiceDeliveryEnabled() cdktn.IResolvable
	EinvoiceDeliveryPreference() DataAwsccInvoicingProcurementPortalPreferenceEinvoiceDeliveryPreferenceOutputReference
	EinvoiceDeliveryPreferenceStatus() *string
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
	LastUpdateDate() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	ProcurementPortalInstanceEndpoint() *string
	ProcurementPortalName() *string
	ProcurementPortalPreferenceArn() *string
	ProcurementPortalSharedSecret() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	PurchaseOrderRetrievalEnabled() cdktn.IResolvable
	PurchaseOrderRetrievalEndpoint() *string
	PurchaseOrderRetrievalPreferenceStatus() *string
	// Experimental.
	RawOverrides() interface{}
	Selector() DataAwsccInvoicingProcurementPortalPreferenceSelectorOutputReference
	SupplierDomain() *string
	SupplierIdentifier() *string
	Tags() DataAwsccInvoicingProcurementPortalPreferenceTagsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TestEnvPreference() DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference
	Version() *float64
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

// The jsii proxy struct for DataAwsccInvoicingProcurementPortalPreference
type jsiiProxy_DataAwsccInvoicingProcurementPortalPreference struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) BuyerDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) BuyerIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) Contacts() DataAwsccInvoicingProcurementPortalPreferenceContactsList {
	var returns DataAwsccInvoicingProcurementPortalPreferenceContactsList
	_jsii_.Get(
		j,
		"contacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) CreateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) EinvoiceDeliveryEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"einvoiceDeliveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) EinvoiceDeliveryPreference() DataAwsccInvoicingProcurementPortalPreferenceEinvoiceDeliveryPreferenceOutputReference {
	var returns DataAwsccInvoicingProcurementPortalPreferenceEinvoiceDeliveryPreferenceOutputReference
	_jsii_.Get(
		j,
		"einvoiceDeliveryPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) EinvoiceDeliveryPreferenceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"einvoiceDeliveryPreferenceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) LastUpdateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdateDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) ProcurementPortalInstanceEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalInstanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) ProcurementPortalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) ProcurementPortalPreferenceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalPreferenceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) ProcurementPortalSharedSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalSharedSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) PurchaseOrderRetrievalEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"purchaseOrderRetrievalEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) PurchaseOrderRetrievalEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purchaseOrderRetrievalEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) PurchaseOrderRetrievalPreferenceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purchaseOrderRetrievalPreferenceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) Selector() DataAwsccInvoicingProcurementPortalPreferenceSelectorOutputReference {
	var returns DataAwsccInvoicingProcurementPortalPreferenceSelectorOutputReference
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) SupplierDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) SupplierIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) Tags() DataAwsccInvoicingProcurementPortalPreferenceTagsList {
	var returns DataAwsccInvoicingProcurementPortalPreferenceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) TestEnvPreference() DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference {
	var returns DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference
	_jsii_.Get(
		j,
		"testEnvPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/data-sources/invoicing_procurement_portal_preference awscc_invoicing_procurement_portal_preference} Data Source.
func NewDataAwsccInvoicingProcurementPortalPreference(scope constructs.Construct, id *string, config *DataAwsccInvoicingProcurementPortalPreferenceConfig) DataAwsccInvoicingProcurementPortalPreference {
	_init_.Initialize()

	if err := validateNewDataAwsccInvoicingProcurementPortalPreferenceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccInvoicingProcurementPortalPreference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccInvoicingProcurementPortalPreference.DataAwsccInvoicingProcurementPortalPreference",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/data-sources/invoicing_procurement_portal_preference awscc_invoicing_procurement_portal_preference} Data Source.
func NewDataAwsccInvoicingProcurementPortalPreference_Override(d DataAwsccInvoicingProcurementPortalPreference, scope constructs.Construct, id *string, config *DataAwsccInvoicingProcurementPortalPreferenceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccInvoicingProcurementPortalPreference.DataAwsccInvoicingProcurementPortalPreference",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsccInvoicingProcurementPortalPreference resource upon running "cdktn plan <stack-name>".
func DataAwsccInvoicingProcurementPortalPreference_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccInvoicingProcurementPortalPreference_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccInvoicingProcurementPortalPreference.DataAwsccInvoicingProcurementPortalPreference",
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
func DataAwsccInvoicingProcurementPortalPreference_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccInvoicingProcurementPortalPreference_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccInvoicingProcurementPortalPreference.DataAwsccInvoicingProcurementPortalPreference",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccInvoicingProcurementPortalPreference_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccInvoicingProcurementPortalPreference_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccInvoicingProcurementPortalPreference.DataAwsccInvoicingProcurementPortalPreference",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccInvoicingProcurementPortalPreference_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccInvoicingProcurementPortalPreference_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccInvoicingProcurementPortalPreference.DataAwsccInvoicingProcurementPortalPreference",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccInvoicingProcurementPortalPreference_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.dataAwsccInvoicingProcurementPortalPreference.DataAwsccInvoicingProcurementPortalPreference",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreference) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

