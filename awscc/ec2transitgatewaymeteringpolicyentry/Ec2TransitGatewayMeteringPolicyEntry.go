// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewaymeteringpolicyentry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2transitgatewaymeteringpolicyentry/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_transit_gateway_metering_policy_entry awscc_ec2_transit_gateway_metering_policy_entry}.
type Ec2TransitGatewayMeteringPolicyEntry interface {
	cdktn.TerraformResource
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
	DestinationCidrBlock() *string
	SetDestinationCidrBlock(val *string)
	DestinationCidrBlockInput() *string
	DestinationPortRange() *string
	SetDestinationPortRange(val *string)
	DestinationPortRangeInput() *string
	DestinationTransitGatewayAttachmentId() *string
	SetDestinationTransitGatewayAttachmentId(val *string)
	DestinationTransitGatewayAttachmentIdInput() *string
	DestinationTransitGatewayAttachmentType() *string
	SetDestinationTransitGatewayAttachmentType(val *string)
	DestinationTransitGatewayAttachmentTypeInput() *string
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
	MeteredAccount() *string
	SetMeteredAccount(val *string)
	MeteredAccountInput() *string
	// The tree node.
	Node() constructs.Node
	PolicyRuleNumber() *float64
	SetPolicyRuleNumber(val *float64)
	PolicyRuleNumberInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	SourceCidrBlock() *string
	SetSourceCidrBlock(val *string)
	SourceCidrBlockInput() *string
	SourcePortRange() *string
	SetSourcePortRange(val *string)
	SourcePortRangeInput() *string
	SourceTransitGatewayAttachmentId() *string
	SetSourceTransitGatewayAttachmentId(val *string)
	SourceTransitGatewayAttachmentIdInput() *string
	SourceTransitGatewayAttachmentType() *string
	SetSourceTransitGatewayAttachmentType(val *string)
	SourceTransitGatewayAttachmentTypeInput() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TransitGatewayMeteringPolicyId() *string
	SetTransitGatewayMeteringPolicyId(val *string)
	TransitGatewayMeteringPolicyIdInput() *string
	UpdateEffectiveAt() *string
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
	ResetDestinationCidrBlock()
	ResetDestinationPortRange()
	ResetDestinationTransitGatewayAttachmentId()
	ResetDestinationTransitGatewayAttachmentType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtocol()
	ResetSourceCidrBlock()
	ResetSourcePortRange()
	ResetSourceTransitGatewayAttachmentId()
	ResetSourceTransitGatewayAttachmentType()
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

// The jsii proxy struct for Ec2TransitGatewayMeteringPolicyEntry
type jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) DestinationCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) DestinationCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) DestinationPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) DestinationPortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) DestinationTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) DestinationTransitGatewayAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) DestinationTransitGatewayAttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) DestinationTransitGatewayAttachmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTransitGatewayAttachmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) MeteredAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meteredAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) MeteredAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"meteredAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) PolicyRuleNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyRuleNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) PolicyRuleNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyRuleNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) SourceCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) SourceCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) SourcePortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) SourcePortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) SourceTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) SourceTransitGatewayAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) SourceTransitGatewayAttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) SourceTransitGatewayAttachmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitGatewayAttachmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) TransitGatewayMeteringPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayMeteringPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) TransitGatewayMeteringPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayMeteringPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) UpdateEffectiveAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateEffectiveAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_transit_gateway_metering_policy_entry awscc_ec2_transit_gateway_metering_policy_entry} Resource.
func NewEc2TransitGatewayMeteringPolicyEntry(scope constructs.Construct, id *string, config *Ec2TransitGatewayMeteringPolicyEntryConfig) Ec2TransitGatewayMeteringPolicyEntry {
	_init_.Initialize()

	if err := validateNewEc2TransitGatewayMeteringPolicyEntryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2TransitGatewayMeteringPolicyEntry.Ec2TransitGatewayMeteringPolicyEntry",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_transit_gateway_metering_policy_entry awscc_ec2_transit_gateway_metering_policy_entry} Resource.
func NewEc2TransitGatewayMeteringPolicyEntry_Override(e Ec2TransitGatewayMeteringPolicyEntry, scope constructs.Construct, id *string, config *Ec2TransitGatewayMeteringPolicyEntryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2TransitGatewayMeteringPolicyEntry.Ec2TransitGatewayMeteringPolicyEntry",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetDestinationCidrBlock(val *string) {
	if err := j.validateSetDestinationCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationCidrBlock",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetDestinationPortRange(val *string) {
	if err := j.validateSetDestinationPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPortRange",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetDestinationTransitGatewayAttachmentId(val *string) {
	if err := j.validateSetDestinationTransitGatewayAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetDestinationTransitGatewayAttachmentType(val *string) {
	if err := j.validateSetDestinationTransitGatewayAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationTransitGatewayAttachmentType",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetMeteredAccount(val *string) {
	if err := j.validateSetMeteredAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"meteredAccount",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetPolicyRuleNumber(val *float64) {
	if err := j.validateSetPolicyRuleNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyRuleNumber",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetSourceCidrBlock(val *string) {
	if err := j.validateSetSourceCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCidrBlock",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetSourcePortRange(val *string) {
	if err := j.validateSetSourcePortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourcePortRange",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetSourceTransitGatewayAttachmentId(val *string) {
	if err := j.validateSetSourceTransitGatewayAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetSourceTransitGatewayAttachmentType(val *string) {
	if err := j.validateSetSourceTransitGatewayAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTransitGatewayAttachmentType",
		val,
	)
}

func (j *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry)SetTransitGatewayMeteringPolicyId(val *string) {
	if err := j.validateSetTransitGatewayMeteringPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayMeteringPolicyId",
		val,
	)
}

// Generates CDKTN code for importing a Ec2TransitGatewayMeteringPolicyEntry resource upon running "cdktn plan <stack-name>".
func Ec2TransitGatewayMeteringPolicyEntry_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateEc2TransitGatewayMeteringPolicyEntry_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ec2TransitGatewayMeteringPolicyEntry.Ec2TransitGatewayMeteringPolicyEntry",
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
func Ec2TransitGatewayMeteringPolicyEntry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2TransitGatewayMeteringPolicyEntry_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ec2TransitGatewayMeteringPolicyEntry.Ec2TransitGatewayMeteringPolicyEntry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2TransitGatewayMeteringPolicyEntry_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2TransitGatewayMeteringPolicyEntry_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ec2TransitGatewayMeteringPolicyEntry.Ec2TransitGatewayMeteringPolicyEntry",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2TransitGatewayMeteringPolicyEntry_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2TransitGatewayMeteringPolicyEntry_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.ec2TransitGatewayMeteringPolicyEntry.Ec2TransitGatewayMeteringPolicyEntry",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2TransitGatewayMeteringPolicyEntry_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.ec2TransitGatewayMeteringPolicyEntry.Ec2TransitGatewayMeteringPolicyEntry",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := e.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ResetDestinationCidrBlock() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationCidrBlock",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ResetDestinationPortRange() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationPortRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ResetDestinationTransitGatewayAttachmentId() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationTransitGatewayAttachmentId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ResetDestinationTransitGatewayAttachmentType() {
	_jsii_.InvokeVoid(
		e,
		"resetDestinationTransitGatewayAttachmentType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ResetProtocol() {
	_jsii_.InvokeVoid(
		e,
		"resetProtocol",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ResetSourceCidrBlock() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceCidrBlock",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ResetSourcePortRange() {
	_jsii_.InvokeVoid(
		e,
		"resetSourcePortRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ResetSourceTransitGatewayAttachmentId() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceTransitGatewayAttachmentId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ResetSourceTransitGatewayAttachmentType() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceTransitGatewayAttachmentType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

