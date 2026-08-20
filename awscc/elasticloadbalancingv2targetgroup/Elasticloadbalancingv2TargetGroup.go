// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2targetgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticloadbalancingv2targetgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancingv2_target_group awscc_elasticloadbalancingv2_target_group}.
type Elasticloadbalancingv2TargetGroup interface {
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
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheckEnabled() interface{}
	SetHealthCheckEnabled(val interface{})
	HealthCheckEnabledInput() interface{}
	HealthCheckIntervalSeconds() *float64
	SetHealthCheckIntervalSeconds(val *float64)
	HealthCheckIntervalSecondsInput() *float64
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPathInput() *string
	HealthCheckPort() *string
	SetHealthCheckPort(val *string)
	HealthCheckPortInput() *string
	HealthCheckProtocol() *string
	SetHealthCheckProtocol(val *string)
	HealthCheckProtocolInput() *string
	HealthCheckTimeoutSeconds() *float64
	SetHealthCheckTimeoutSeconds(val *float64)
	HealthCheckTimeoutSecondsInput() *float64
	HealthyThresholdCount() *float64
	SetHealthyThresholdCount(val *float64)
	HealthyThresholdCountInput() *float64
	Id() *string
	IpAddressType() *string
	SetIpAddressType(val *string)
	IpAddressTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LoadBalancerArns() *[]*string
	Matcher() Elasticloadbalancingv2TargetGroupMatcherOutputReference
	MatcherInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	ProtocolVersion() *string
	SetProtocolVersion(val *string)
	ProtocolVersionInput() *string
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
	Tags() Elasticloadbalancingv2TargetGroupTagsList
	TagsInput() interface{}
	TargetControlPort() *float64
	SetTargetControlPort(val *float64)
	TargetControlPortInput() *float64
	TargetGroupArn() *string
	TargetGroupAttributes() Elasticloadbalancingv2TargetGroupTargetGroupAttributesList
	TargetGroupAttributesInput() interface{}
	TargetGroupFullName() *string
	TargetGroupName() *string
	Targets() Elasticloadbalancingv2TargetGroupTargetsList
	TargetsInput() interface{}
	TargetType() *string
	SetTargetType(val *string)
	TargetTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UnhealthyThresholdCount() *float64
	SetUnhealthyThresholdCount(val *float64)
	UnhealthyThresholdCountInput() *float64
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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
	PutMatcher(value *Elasticloadbalancingv2TargetGroupMatcher)
	PutTags(value interface{})
	PutTargetGroupAttributes(value interface{})
	PutTargets(value interface{})
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
	ResetHealthCheckEnabled()
	ResetHealthCheckIntervalSeconds()
	ResetHealthCheckPath()
	ResetHealthCheckPort()
	ResetHealthCheckProtocol()
	ResetHealthCheckTimeoutSeconds()
	ResetHealthyThresholdCount()
	ResetIpAddressType()
	ResetMatcher()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetProtocol()
	ResetProtocolVersion()
	ResetTags()
	ResetTargetControlPort()
	ResetTargetGroupAttributes()
	ResetTargets()
	ResetTargetType()
	ResetUnhealthyThresholdCount()
	ResetVpcId()
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

// The jsii proxy struct for Elasticloadbalancingv2TargetGroup
type jsiiProxy_Elasticloadbalancingv2TargetGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckPortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthCheckTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthyThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) HealthyThresholdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) LoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Matcher() Elasticloadbalancingv2TargetGroupMatcherOutputReference {
	var returns Elasticloadbalancingv2TargetGroupMatcherOutputReference
	_jsii_.Get(
		j,
		"matcher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) MatcherInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matcherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) ProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) ProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Tags() Elasticloadbalancingv2TargetGroupTagsList {
	var returns Elasticloadbalancingv2TargetGroupTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TargetControlPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetControlPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TargetControlPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetControlPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TargetGroupAttributes() Elasticloadbalancingv2TargetGroupTargetGroupAttributesList {
	var returns Elasticloadbalancingv2TargetGroupTargetGroupAttributesList
	_jsii_.Get(
		j,
		"targetGroupAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TargetGroupAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetGroupAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TargetGroupFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TargetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) Targets() Elasticloadbalancingv2TargetGroupTargetsList {
	var returns Elasticloadbalancingv2TargetGroupTargetsList
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TargetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) UnhealthyThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) UnhealthyThresholdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancingv2_target_group awscc_elasticloadbalancingv2_target_group} Resource.
func NewElasticloadbalancingv2TargetGroup(scope constructs.Construct, id *string, config *Elasticloadbalancingv2TargetGroupConfig) Elasticloadbalancingv2TargetGroup {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2TargetGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2TargetGroup{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2TargetGroup.Elasticloadbalancingv2TargetGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticloadbalancingv2_target_group awscc_elasticloadbalancingv2_target_group} Resource.
func NewElasticloadbalancingv2TargetGroup_Override(e Elasticloadbalancingv2TargetGroup, scope constructs.Construct, id *string, config *Elasticloadbalancingv2TargetGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2TargetGroup.Elasticloadbalancingv2TargetGroup",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetHealthCheckEnabled(val interface{}) {
	if err := j.validateSetHealthCheckEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckEnabled",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetHealthCheckIntervalSeconds(val *float64) {
	if err := j.validateSetHealthCheckIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetHealthCheckPath(val *string) {
	if err := j.validateSetHealthCheckPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetHealthCheckPort(val *string) {
	if err := j.validateSetHealthCheckPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPort",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetHealthCheckProtocol(val *string) {
	if err := j.validateSetHealthCheckProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckProtocol",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetHealthCheckTimeoutSeconds(val *float64) {
	if err := j.validateSetHealthCheckTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetHealthyThresholdCount(val *float64) {
	if err := j.validateSetHealthyThresholdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthyThresholdCount",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetProtocolVersion(val *string) {
	if err := j.validateSetProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolVersion",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetTargetControlPort(val *float64) {
	if err := j.validateSetTargetControlPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetControlPort",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetTargetType(val *string) {
	if err := j.validateSetTargetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetUnhealthyThresholdCount(val *float64) {
	if err := j.validateSetUnhealthyThresholdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyThresholdCount",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TargetGroup)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTN code for importing a Elasticloadbalancingv2TargetGroup resource upon running "cdktn plan <stack-name>".
func Elasticloadbalancingv2TargetGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TargetGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingv2TargetGroup.Elasticloadbalancingv2TargetGroup",
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
func Elasticloadbalancingv2TargetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TargetGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingv2TargetGroup.Elasticloadbalancingv2TargetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Elasticloadbalancingv2TargetGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TargetGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingv2TargetGroup.Elasticloadbalancingv2TargetGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Elasticloadbalancingv2TargetGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingv2TargetGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingv2TargetGroup.Elasticloadbalancingv2TargetGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Elasticloadbalancingv2TargetGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.elasticloadbalancingv2TargetGroup.Elasticloadbalancingv2TargetGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) PutMatcher(value *Elasticloadbalancingv2TargetGroupMatcher) {
	if err := e.validatePutMatcherParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMatcher",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) PutTargetGroupAttributes(value interface{}) {
	if err := e.validatePutTargetGroupAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTargetGroupAttributes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) PutTargets(value interface{}) {
	if err := e.validatePutTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTargets",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetHealthCheckEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetHealthCheckIntervalSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckIntervalSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetHealthCheckPort() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetHealthCheckProtocol() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckProtocol",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetHealthCheckTimeoutSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheckTimeoutSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetHealthyThresholdCount() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthyThresholdCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		e,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetMatcher() {
	_jsii_.InvokeVoid(
		e,
		"resetMatcher",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetPort() {
	_jsii_.InvokeVoid(
		e,
		"resetPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetProtocol() {
	_jsii_.InvokeVoid(
		e,
		"resetProtocol",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetProtocolVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetProtocolVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetTargetControlPort() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetControlPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetTargetGroupAttributes() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetGroupAttributes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetTargets() {
	_jsii_.InvokeVoid(
		e,
		"resetTargets",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetTargetType() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetUnhealthyThresholdCount() {
	_jsii_.InvokeVoid(
		e,
		"resetUnhealthyThresholdCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ResetVpcId() {
	_jsii_.InvokeVoid(
		e,
		"resetVpcId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TargetGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

