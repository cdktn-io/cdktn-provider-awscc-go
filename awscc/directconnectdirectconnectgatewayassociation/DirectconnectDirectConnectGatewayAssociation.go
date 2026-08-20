// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directconnectdirectconnectgatewayassociation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/directconnectdirectconnectgatewayassociation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_direct_connect_gateway_association awscc_directconnect_direct_connect_gateway_association}.
type DirectconnectDirectConnectGatewayAssociation interface {
	cdktn.TerraformResource
	AcceptDirectConnectGatewayAssociationProposalRoleArn() *string
	SetAcceptDirectConnectGatewayAssociationProposalRoleArn(val *string)
	AcceptDirectConnectGatewayAssociationProposalRoleArnInput() *string
	AllowedPrefixesToDirectConnectGateway() *[]*string
	SetAllowedPrefixesToDirectConnectGateway(val *[]*string)
	AllowedPrefixesToDirectConnectGatewayInput() *[]*string
	AssociatedGatewayId() *string
	SetAssociatedGatewayId(val *string)
	AssociatedGatewayIdInput() *string
	AssociationId() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirectConnectGatewayId() *string
	SetDirectConnectGatewayId(val *string)
	DirectConnectGatewayIdInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAcceptDirectConnectGatewayAssociationProposalRoleArn()
	ResetAllowedPrefixesToDirectConnectGateway()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DirectconnectDirectConnectGatewayAssociation
type jsiiProxy_DirectconnectDirectConnectGatewayAssociation struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) AcceptDirectConnectGatewayAssociationProposalRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptDirectConnectGatewayAssociationProposalRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) AcceptDirectConnectGatewayAssociationProposalRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceptDirectConnectGatewayAssociationProposalRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) AllowedPrefixesToDirectConnectGateway() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPrefixesToDirectConnectGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) AllowedPrefixesToDirectConnectGatewayInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPrefixesToDirectConnectGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) AssociatedGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) AssociatedGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associatedGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) DirectConnectGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directConnectGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) DirectConnectGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directConnectGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_direct_connect_gateway_association awscc_directconnect_direct_connect_gateway_association} Resource.
func NewDirectconnectDirectConnectGatewayAssociation(scope constructs.Construct, id *string, config *DirectconnectDirectConnectGatewayAssociationConfig) DirectconnectDirectConnectGatewayAssociation {
	_init_.Initialize()

	if err := validateNewDirectconnectDirectConnectGatewayAssociationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DirectconnectDirectConnectGatewayAssociation{}

	_jsii_.Create(
		"@cdktn/provider-awscc.directconnectDirectConnectGatewayAssociation.DirectconnectDirectConnectGatewayAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/directconnect_direct_connect_gateway_association awscc_directconnect_direct_connect_gateway_association} Resource.
func NewDirectconnectDirectConnectGatewayAssociation_Override(d DirectconnectDirectConnectGatewayAssociation, scope constructs.Construct, id *string, config *DirectconnectDirectConnectGatewayAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.directconnectDirectConnectGatewayAssociation.DirectconnectDirectConnectGatewayAssociation",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation)SetAcceptDirectConnectGatewayAssociationProposalRoleArn(val *string) {
	if err := j.validateSetAcceptDirectConnectGatewayAssociationProposalRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptDirectConnectGatewayAssociationProposalRoleArn",
		val,
	)
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation)SetAllowedPrefixesToDirectConnectGateway(val *[]*string) {
	if err := j.validateSetAllowedPrefixesToDirectConnectGatewayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedPrefixesToDirectConnectGateway",
		val,
	)
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation)SetAssociatedGatewayId(val *string) {
	if err := j.validateSetAssociatedGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associatedGatewayId",
		val,
	)
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation)SetDirectConnectGatewayId(val *string) {
	if err := j.validateSetDirectConnectGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directConnectGatewayId",
		val,
	)
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DirectconnectDirectConnectGatewayAssociation)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a DirectconnectDirectConnectGatewayAssociation resource upon running "cdktn plan <stack-name>".
func DirectconnectDirectConnectGatewayAssociation_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDirectconnectDirectConnectGatewayAssociation_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.directconnectDirectConnectGatewayAssociation.DirectconnectDirectConnectGatewayAssociation",
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
func DirectconnectDirectConnectGatewayAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectconnectDirectConnectGatewayAssociation_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.directconnectDirectConnectGatewayAssociation.DirectconnectDirectConnectGatewayAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DirectconnectDirectConnectGatewayAssociation_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectconnectDirectConnectGatewayAssociation_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.directconnectDirectConnectGatewayAssociation.DirectconnectDirectConnectGatewayAssociation",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DirectconnectDirectConnectGatewayAssociation_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDirectconnectDirectConnectGatewayAssociation_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.directconnectDirectConnectGatewayAssociation.DirectconnectDirectConnectGatewayAssociation",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DirectconnectDirectConnectGatewayAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.directconnectDirectConnectGatewayAssociation.DirectconnectDirectConnectGatewayAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) ResetAcceptDirectConnectGatewayAssociationProposalRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetAcceptDirectConnectGatewayAssociationProposalRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) ResetAllowedPrefixesToDirectConnectGateway() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedPrefixesToDirectConnectGateway",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DirectconnectDirectConnectGatewayAssociation) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

