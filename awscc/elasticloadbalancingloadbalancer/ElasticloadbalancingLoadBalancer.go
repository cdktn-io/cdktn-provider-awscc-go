// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingloadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticloadbalancingloadbalancer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticloadbalancing_load_balancer awscc_elasticloadbalancing_load_balancer}.
type ElasticloadbalancingLoadBalancer interface {
	cdktn.TerraformResource
	AccessLoggingPolicy() ElasticloadbalancingLoadBalancerAccessLoggingPolicyOutputReference
	AccessLoggingPolicyInput() interface{}
	AppCookieStickinessPolicy() ElasticloadbalancingLoadBalancerAppCookieStickinessPolicyList
	AppCookieStickinessPolicyInput() interface{}
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	CanonicalHostedZoneName() *string
	CanonicalHostedZoneNameId() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionDrainingPolicy() ElasticloadbalancingLoadBalancerConnectionDrainingPolicyOutputReference
	ConnectionDrainingPolicyInput() interface{}
	ConnectionSettings() ElasticloadbalancingLoadBalancerConnectionSettingsOutputReference
	ConnectionSettingsInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CrossZone() interface{}
	SetCrossZone(val interface{})
	CrossZoneInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DnsName() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HealthCheck() ElasticloadbalancingLoadBalancerHealthCheckOutputReference
	HealthCheckInput() interface{}
	Id() *string
	Instances() *[]*string
	SetInstances(val *[]*string)
	InstancesInput() *[]*string
	LbCookieStickinessPolicy() ElasticloadbalancingLoadBalancerLbCookieStickinessPolicyList
	LbCookieStickinessPolicyInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Listeners() ElasticloadbalancingLoadBalancerListenersList
	ListenersInput() interface{}
	LoadBalancerName() *string
	SetLoadBalancerName(val *string)
	LoadBalancerNameInput() *string
	// The tree node.
	Node() constructs.Node
	Policies() ElasticloadbalancingLoadBalancerPoliciesList
	PoliciesInput() interface{}
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
	Scheme() *string
	SetScheme(val *string)
	SchemeInput() *string
	SecurityGroups() *[]*string
	SetSecurityGroups(val *[]*string)
	SecurityGroupsInput() *[]*string
	SourceSecurityGroup() ElasticloadbalancingLoadBalancerSourceSecurityGroupOutputReference
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	SubnetsInput() *[]*string
	Tags() ElasticloadbalancingLoadBalancerTagsList
	TagsInput() interface{}
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
	PutAccessLoggingPolicy(value *ElasticloadbalancingLoadBalancerAccessLoggingPolicy)
	PutAppCookieStickinessPolicy(value interface{})
	PutConnectionDrainingPolicy(value *ElasticloadbalancingLoadBalancerConnectionDrainingPolicy)
	PutConnectionSettings(value *ElasticloadbalancingLoadBalancerConnectionSettings)
	PutHealthCheck(value *ElasticloadbalancingLoadBalancerHealthCheck)
	PutLbCookieStickinessPolicy(value interface{})
	PutListeners(value interface{})
	PutPolicies(value interface{})
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
	ResetAccessLoggingPolicy()
	ResetAppCookieStickinessPolicy()
	ResetAvailabilityZones()
	ResetConnectionDrainingPolicy()
	ResetConnectionSettings()
	ResetCrossZone()
	ResetHealthCheck()
	ResetInstances()
	ResetLbCookieStickinessPolicy()
	ResetLoadBalancerName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicies()
	ResetScheme()
	ResetSecurityGroups()
	ResetSubnets()
	ResetTags()
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

// The jsii proxy struct for ElasticloadbalancingLoadBalancer
type jsiiProxy_ElasticloadbalancingLoadBalancer struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) AccessLoggingPolicy() ElasticloadbalancingLoadBalancerAccessLoggingPolicyOutputReference {
	var returns ElasticloadbalancingLoadBalancerAccessLoggingPolicyOutputReference
	_jsii_.Get(
		j,
		"accessLoggingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) AccessLoggingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLoggingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) AppCookieStickinessPolicy() ElasticloadbalancingLoadBalancerAppCookieStickinessPolicyList {
	var returns ElasticloadbalancingLoadBalancerAppCookieStickinessPolicyList
	_jsii_.Get(
		j,
		"appCookieStickinessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) AppCookieStickinessPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appCookieStickinessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) CanonicalHostedZoneName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canonicalHostedZoneName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) CanonicalHostedZoneNameId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canonicalHostedZoneNameId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) ConnectionDrainingPolicy() ElasticloadbalancingLoadBalancerConnectionDrainingPolicyOutputReference {
	var returns ElasticloadbalancingLoadBalancerConnectionDrainingPolicyOutputReference
	_jsii_.Get(
		j,
		"connectionDrainingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) ConnectionDrainingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionDrainingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) ConnectionSettings() ElasticloadbalancingLoadBalancerConnectionSettingsOutputReference {
	var returns ElasticloadbalancingLoadBalancerConnectionSettingsOutputReference
	_jsii_.Get(
		j,
		"connectionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) ConnectionSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) CrossZone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) CrossZoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) HealthCheck() ElasticloadbalancingLoadBalancerHealthCheckOutputReference {
	var returns ElasticloadbalancingLoadBalancerHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) HealthCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Instances() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) InstancesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) LbCookieStickinessPolicy() ElasticloadbalancingLoadBalancerLbCookieStickinessPolicyList {
	var returns ElasticloadbalancingLoadBalancerLbCookieStickinessPolicyList
	_jsii_.Get(
		j,
		"lbCookieStickinessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) LbCookieStickinessPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbCookieStickinessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Listeners() ElasticloadbalancingLoadBalancerListenersList {
	var returns ElasticloadbalancingLoadBalancerListenersList
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) ListenersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"listenersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) LoadBalancerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Policies() ElasticloadbalancingLoadBalancerPoliciesList {
	var returns ElasticloadbalancingLoadBalancerPoliciesList
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) PoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Scheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) SchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) SourceSecurityGroup() ElasticloadbalancingLoadBalancerSourceSecurityGroupOutputReference {
	var returns ElasticloadbalancingLoadBalancerSourceSecurityGroupOutputReference
	_jsii_.Get(
		j,
		"sourceSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) Tags() ElasticloadbalancingLoadBalancerTagsList {
	var returns ElasticloadbalancingLoadBalancerTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticloadbalancing_load_balancer awscc_elasticloadbalancing_load_balancer} Resource.
func NewElasticloadbalancingLoadBalancer(scope constructs.Construct, id *string, config *ElasticloadbalancingLoadBalancerConfig) ElasticloadbalancingLoadBalancer {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingLoadBalancerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticloadbalancingLoadBalancer{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingLoadBalancer.ElasticloadbalancingLoadBalancer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/elasticloadbalancing_load_balancer awscc_elasticloadbalancing_load_balancer} Resource.
func NewElasticloadbalancingLoadBalancer_Override(e ElasticloadbalancingLoadBalancer, scope constructs.Construct, id *string, config *ElasticloadbalancingLoadBalancerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingLoadBalancer.ElasticloadbalancingLoadBalancer",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetCrossZone(val interface{}) {
	if err := j.validateSetCrossZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossZone",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetInstances(val *[]*string) {
	if err := j.validateSetInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetLoadBalancerName(val *string) {
	if err := j.validateSetLoadBalancerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerName",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetScheme(val *string) {
	if err := j.validateSetSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancer)SetSubnets(val *[]*string) {
	if err := j.validateSetSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

// Generates CDKTN code for importing a ElasticloadbalancingLoadBalancer resource upon running "cdktn plan <stack-name>".
func ElasticloadbalancingLoadBalancer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateElasticloadbalancingLoadBalancer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingLoadBalancer.ElasticloadbalancingLoadBalancer",
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
func ElasticloadbalancingLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingLoadBalancer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingLoadBalancer.ElasticloadbalancingLoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticloadbalancingLoadBalancer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingLoadBalancer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingLoadBalancer.ElasticloadbalancingLoadBalancer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticloadbalancingLoadBalancer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticloadbalancingLoadBalancer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.elasticloadbalancingLoadBalancer.ElasticloadbalancingLoadBalancer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElasticloadbalancingLoadBalancer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.elasticloadbalancingLoadBalancer.ElasticloadbalancingLoadBalancer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) PutAccessLoggingPolicy(value *ElasticloadbalancingLoadBalancerAccessLoggingPolicy) {
	if err := e.validatePutAccessLoggingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAccessLoggingPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) PutAppCookieStickinessPolicy(value interface{}) {
	if err := e.validatePutAppCookieStickinessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAppCookieStickinessPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) PutConnectionDrainingPolicy(value *ElasticloadbalancingLoadBalancerConnectionDrainingPolicy) {
	if err := e.validatePutConnectionDrainingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putConnectionDrainingPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) PutConnectionSettings(value *ElasticloadbalancingLoadBalancerConnectionSettings) {
	if err := e.validatePutConnectionSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putConnectionSettings",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) PutHealthCheck(value *ElasticloadbalancingLoadBalancerHealthCheck) {
	if err := e.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) PutLbCookieStickinessPolicy(value interface{}) {
	if err := e.validatePutLbCookieStickinessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLbCookieStickinessPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) PutListeners(value interface{}) {
	if err := e.validatePutListenersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putListeners",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) PutPolicies(value interface{}) {
	if err := e.validatePutPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPolicies",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetAccessLoggingPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetAccessLoggingPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetAppCookieStickinessPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetAppCookieStickinessPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetConnectionDrainingPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetConnectionDrainingPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetConnectionSettings() {
	_jsii_.InvokeVoid(
		e,
		"resetConnectionSettings",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetCrossZone() {
	_jsii_.InvokeVoid(
		e,
		"resetCrossZone",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		e,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetInstances() {
	_jsii_.InvokeVoid(
		e,
		"resetInstances",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetLbCookieStickinessPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetLbCookieStickinessPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetLoadBalancerName() {
	_jsii_.InvokeVoid(
		e,
		"resetLoadBalancerName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetPolicies() {
	_jsii_.InvokeVoid(
		e,
		"resetPolicies",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetScheme() {
	_jsii_.InvokeVoid(
		e,
		"resetScheme",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetSubnets() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnets",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancer) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

