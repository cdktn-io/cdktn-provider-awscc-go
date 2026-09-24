// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package invoicingprocurementportalpreference

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/invoicingprocurementportalpreference/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference awscc_invoicing_procurement_portal_preference}.
type InvoicingProcurementPortalPreference interface {
	cdktn.TerraformResource
	AwsAccountId() *string
	BuyerDomain() *string
	SetBuyerDomain(val *string)
	BuyerDomainInput() *string
	BuyerIdentifier() *string
	SetBuyerIdentifier(val *string)
	BuyerIdentifierInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Contacts() InvoicingProcurementPortalPreferenceContactsList
	ContactsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateDate() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EinvoiceDeliveryEnabled() interface{}
	SetEinvoiceDeliveryEnabled(val interface{})
	EinvoiceDeliveryEnabledInput() interface{}
	EinvoiceDeliveryPreference() InvoicingProcurementPortalPreferenceEinvoiceDeliveryPreferenceOutputReference
	EinvoiceDeliveryPreferenceInput() interface{}
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
	LastUpdateDate() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	ProcurementPortalInstanceEndpoint() *string
	SetProcurementPortalInstanceEndpoint(val *string)
	ProcurementPortalInstanceEndpointInput() *string
	ProcurementPortalName() *string
	SetProcurementPortalName(val *string)
	ProcurementPortalNameInput() *string
	ProcurementPortalPreferenceArn() *string
	ProcurementPortalSharedSecret() *string
	SetProcurementPortalSharedSecret(val *string)
	ProcurementPortalSharedSecretInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PurchaseOrderRetrievalEnabled() interface{}
	SetPurchaseOrderRetrievalEnabled(val interface{})
	PurchaseOrderRetrievalEnabledInput() interface{}
	PurchaseOrderRetrievalEndpoint() *string
	PurchaseOrderRetrievalPreferenceStatus() *string
	// Experimental.
	RawOverrides() interface{}
	Selector() InvoicingProcurementPortalPreferenceSelectorOutputReference
	SelectorInput() interface{}
	SupplierDomain() *string
	SetSupplierDomain(val *string)
	SupplierDomainInput() *string
	SupplierIdentifier() *string
	SetSupplierIdentifier(val *string)
	SupplierIdentifierInput() *string
	Tags() InvoicingProcurementPortalPreferenceTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TestEnvPreference() InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference
	TestEnvPreferenceInput() interface{}
	Version() *float64
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
	PutContacts(value interface{})
	PutEinvoiceDeliveryPreference(value *InvoicingProcurementPortalPreferenceEinvoiceDeliveryPreference)
	PutSelector(value *InvoicingProcurementPortalPreferenceSelector)
	PutTags(value interface{})
	PutTestEnvPreference(value *InvoicingProcurementPortalPreferenceTestEnvPreference)
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
	ResetEinvoiceDeliveryPreference()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProcurementPortalInstanceEndpoint()
	ResetProcurementPortalSharedSecret()
	ResetSelector()
	ResetTags()
	ResetTestEnvPreference()
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

// The jsii proxy struct for InvoicingProcurementPortalPreference
type jsiiProxy_InvoicingProcurementPortalPreference struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) BuyerDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) BuyerDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) BuyerIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) BuyerIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Contacts() InvoicingProcurementPortalPreferenceContactsList {
	var returns InvoicingProcurementPortalPreferenceContactsList
	_jsii_.Get(
		j,
		"contacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) ContactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) CreateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) EinvoiceDeliveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"einvoiceDeliveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) EinvoiceDeliveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"einvoiceDeliveryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) EinvoiceDeliveryPreference() InvoicingProcurementPortalPreferenceEinvoiceDeliveryPreferenceOutputReference {
	var returns InvoicingProcurementPortalPreferenceEinvoiceDeliveryPreferenceOutputReference
	_jsii_.Get(
		j,
		"einvoiceDeliveryPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) EinvoiceDeliveryPreferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"einvoiceDeliveryPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) EinvoiceDeliveryPreferenceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"einvoiceDeliveryPreferenceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) LastUpdateDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdateDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) ProcurementPortalInstanceEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalInstanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) ProcurementPortalInstanceEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalInstanceEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) ProcurementPortalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) ProcurementPortalNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) ProcurementPortalPreferenceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalPreferenceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) ProcurementPortalSharedSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalSharedSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) ProcurementPortalSharedSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalSharedSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) PurchaseOrderRetrievalEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purchaseOrderRetrievalEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) PurchaseOrderRetrievalEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purchaseOrderRetrievalEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) PurchaseOrderRetrievalEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purchaseOrderRetrievalEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) PurchaseOrderRetrievalPreferenceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purchaseOrderRetrievalPreferenceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Selector() InvoicingProcurementPortalPreferenceSelectorOutputReference {
	var returns InvoicingProcurementPortalPreferenceSelectorOutputReference
	_jsii_.Get(
		j,
		"selector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) SelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) SupplierDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) SupplierDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) SupplierIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) SupplierIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Tags() InvoicingProcurementPortalPreferenceTagsList {
	var returns InvoicingProcurementPortalPreferenceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) TestEnvPreference() InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference {
	var returns InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference
	_jsii_.Get(
		j,
		"testEnvPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) TestEnvPreferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"testEnvPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference awscc_invoicing_procurement_portal_preference} Resource.
func NewInvoicingProcurementPortalPreference(scope constructs.Construct, id *string, config *InvoicingProcurementPortalPreferenceConfig) InvoicingProcurementPortalPreference {
	_init_.Initialize()

	if err := validateNewInvoicingProcurementPortalPreferenceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_InvoicingProcurementPortalPreference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.invoicingProcurementPortalPreference.InvoicingProcurementPortalPreference",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/invoicing_procurement_portal_preference awscc_invoicing_procurement_portal_preference} Resource.
func NewInvoicingProcurementPortalPreference_Override(i InvoicingProcurementPortalPreference, scope constructs.Construct, id *string, config *InvoicingProcurementPortalPreferenceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.invoicingProcurementPortalPreference.InvoicingProcurementPortalPreference",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetBuyerDomain(val *string) {
	if err := j.validateSetBuyerDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buyerDomain",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetBuyerIdentifier(val *string) {
	if err := j.validateSetBuyerIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buyerIdentifier",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetEinvoiceDeliveryEnabled(val interface{}) {
	if err := j.validateSetEinvoiceDeliveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"einvoiceDeliveryEnabled",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetProcurementPortalInstanceEndpoint(val *string) {
	if err := j.validateSetProcurementPortalInstanceEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"procurementPortalInstanceEndpoint",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetProcurementPortalName(val *string) {
	if err := j.validateSetProcurementPortalNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"procurementPortalName",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetProcurementPortalSharedSecret(val *string) {
	if err := j.validateSetProcurementPortalSharedSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"procurementPortalSharedSecret",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetPurchaseOrderRetrievalEnabled(val interface{}) {
	if err := j.validateSetPurchaseOrderRetrievalEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purchaseOrderRetrievalEnabled",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetSupplierDomain(val *string) {
	if err := j.validateSetSupplierDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supplierDomain",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreference)SetSupplierIdentifier(val *string) {
	if err := j.validateSetSupplierIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supplierIdentifier",
		val,
	)
}

// Generates CDKTN code for importing a InvoicingProcurementPortalPreference resource upon running "cdktn plan <stack-name>".
func InvoicingProcurementPortalPreference_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateInvoicingProcurementPortalPreference_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.invoicingProcurementPortalPreference.InvoicingProcurementPortalPreference",
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
func InvoicingProcurementPortalPreference_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInvoicingProcurementPortalPreference_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.invoicingProcurementPortalPreference.InvoicingProcurementPortalPreference",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InvoicingProcurementPortalPreference_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInvoicingProcurementPortalPreference_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.invoicingProcurementPortalPreference.InvoicingProcurementPortalPreference",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InvoicingProcurementPortalPreference_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInvoicingProcurementPortalPreference_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.invoicingProcurementPortalPreference.InvoicingProcurementPortalPreference",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func InvoicingProcurementPortalPreference_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.invoicingProcurementPortalPreference.InvoicingProcurementPortalPreference",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := i.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) PutContacts(value interface{}) {
	if err := i.validatePutContactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putContacts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) PutEinvoiceDeliveryPreference(value *InvoicingProcurementPortalPreferenceEinvoiceDeliveryPreference) {
	if err := i.validatePutEinvoiceDeliveryPreferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEinvoiceDeliveryPreference",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) PutSelector(value *InvoicingProcurementPortalPreferenceSelector) {
	if err := i.validatePutSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSelector",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) PutTestEnvPreference(value *InvoicingProcurementPortalPreferenceTestEnvPreference) {
	if err := i.validatePutTestEnvPreferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTestEnvPreference",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := i.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ResetEinvoiceDeliveryPreference() {
	_jsii_.InvokeVoid(
		i,
		"resetEinvoiceDeliveryPreference",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ResetProcurementPortalInstanceEndpoint() {
	_jsii_.InvokeVoid(
		i,
		"resetProcurementPortalInstanceEndpoint",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ResetProcurementPortalSharedSecret() {
	_jsii_.InvokeVoid(
		i,
		"resetProcurementPortalSharedSecret",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ResetSelector() {
	_jsii_.InvokeVoid(
		i,
		"resetSelector",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ResetTestEnvPreference() {
	_jsii_.InvokeVoid(
		i,
		"resetTestEnvPreference",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreference) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

