// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotdomainconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotdomainconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_domain_configuration awscc_iot_domain_configuration}.
type IotDomainConfiguration interface {
	cdktn.TerraformResource
	ApplicationProtocol() *string
	SetApplicationProtocol(val *string)
	ApplicationProtocolInput() *string
	Arn() *string
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
	AuthorizerConfig() IotDomainConfigurationAuthorizerConfigOutputReference
	AuthorizerConfigInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientCertificateConfig() IotDomainConfigurationClientCertificateConfigOutputReference
	ClientCertificateConfigInput() interface{}
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
	DomainConfigurationName() *string
	SetDomainConfigurationName(val *string)
	DomainConfigurationNameInput() *string
	DomainConfigurationStatus() *string
	SetDomainConfigurationStatus(val *string)
	DomainConfigurationStatusInput() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	DomainType() *string
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
	ServerCertificateArns() *[]*string
	SetServerCertificateArns(val *[]*string)
	ServerCertificateArnsInput() *[]*string
	ServerCertificateConfig() IotDomainConfigurationServerCertificateConfigOutputReference
	ServerCertificateConfigInput() interface{}
	ServerCertificates() IotDomainConfigurationServerCertificatesList
	ServiceType() *string
	SetServiceType(val *string)
	ServiceTypeInput() *string
	Tags() IotDomainConfigurationTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsConfig() IotDomainConfigurationTlsConfigOutputReference
	TlsConfigInput() interface{}
	ValidationCertificateArn() *string
	SetValidationCertificateArn(val *string)
	ValidationCertificateArnInput() *string
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
	PutAuthorizerConfig(value *IotDomainConfigurationAuthorizerConfig)
	PutClientCertificateConfig(value *IotDomainConfigurationClientCertificateConfig)
	PutServerCertificateConfig(value *IotDomainConfigurationServerCertificateConfig)
	PutTags(value interface{})
	PutTlsConfig(value *IotDomainConfigurationTlsConfig)
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
	ResetApplicationProtocol()
	ResetAuthenticationType()
	ResetAuthorizerConfig()
	ResetClientCertificateConfig()
	ResetDomainConfigurationName()
	ResetDomainConfigurationStatus()
	ResetDomainName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServerCertificateArns()
	ResetServerCertificateConfig()
	ResetServiceType()
	ResetTags()
	ResetTlsConfig()
	ResetValidationCertificateArn()
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

// The jsii proxy struct for IotDomainConfiguration
type jsiiProxy_IotDomainConfiguration struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_IotDomainConfiguration) ApplicationProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ApplicationProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) AuthorizerConfig() IotDomainConfigurationAuthorizerConfigOutputReference {
	var returns IotDomainConfigurationAuthorizerConfigOutputReference
	_jsii_.Get(
		j,
		"authorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) AuthorizerConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ClientCertificateConfig() IotDomainConfigurationClientCertificateConfigOutputReference {
	var returns IotDomainConfigurationClientCertificateConfigOutputReference
	_jsii_.Get(
		j,
		"clientCertificateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ClientCertificateConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainConfigurationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainConfigurationStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) DomainType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServerCertificateArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverCertificateArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServerCertificateArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverCertificateArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServerCertificateConfig() IotDomainConfigurationServerCertificateConfigOutputReference {
	var returns IotDomainConfigurationServerCertificateConfigOutputReference
	_jsii_.Get(
		j,
		"serverCertificateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServerCertificateConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverCertificateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServerCertificates() IotDomainConfigurationServerCertificatesList {
	var returns IotDomainConfigurationServerCertificatesList
	_jsii_.Get(
		j,
		"serverCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServiceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ServiceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) Tags() IotDomainConfigurationTagsList {
	var returns IotDomainConfigurationTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TlsConfig() IotDomainConfigurationTlsConfigOutputReference {
	var returns IotDomainConfigurationTlsConfigOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) TlsConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ValidationCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotDomainConfiguration) ValidationCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationCertificateArnInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_domain_configuration awscc_iot_domain_configuration} Resource.
func NewIotDomainConfiguration(scope constructs.Construct, id *string, config *IotDomainConfigurationConfig) IotDomainConfiguration {
	_init_.Initialize()

	if err := validateNewIotDomainConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotDomainConfiguration{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotDomainConfiguration.IotDomainConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_domain_configuration awscc_iot_domain_configuration} Resource.
func NewIotDomainConfiguration_Override(i IotDomainConfiguration, scope constructs.Construct, id *string, config *IotDomainConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotDomainConfiguration.IotDomainConfiguration",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetApplicationProtocol(val *string) {
	if err := j.validateSetApplicationProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationProtocol",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetDomainConfigurationName(val *string) {
	if err := j.validateSetDomainConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainConfigurationName",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetDomainConfigurationStatus(val *string) {
	if err := j.validateSetDomainConfigurationStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainConfigurationStatus",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetServerCertificateArns(val *[]*string) {
	if err := j.validateSetServerCertificateArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertificateArns",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetServiceType(val *string) {
	if err := j.validateSetServiceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceType",
		val,
	)
}

func (j *jsiiProxy_IotDomainConfiguration)SetValidationCertificateArn(val *string) {
	if err := j.validateSetValidationCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationCertificateArn",
		val,
	)
}

// Generates CDKTN code for importing a IotDomainConfiguration resource upon running "cdktn plan <stack-name>".
func IotDomainConfiguration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateIotDomainConfiguration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotDomainConfiguration.IotDomainConfiguration",
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
func IotDomainConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotDomainConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotDomainConfiguration.IotDomainConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotDomainConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotDomainConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotDomainConfiguration.IotDomainConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotDomainConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotDomainConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotDomainConfiguration.IotDomainConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotDomainConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.iotDomainConfiguration.IotDomainConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotDomainConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotDomainConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotDomainConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotDomainConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotDomainConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotDomainConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotDomainConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotDomainConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotDomainConfiguration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotDomainConfiguration) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (i *jsiiProxy_IotDomainConfiguration) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) PutAuthorizerConfig(value *IotDomainConfigurationAuthorizerConfig) {
	if err := i.validatePutAuthorizerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAuthorizerConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) PutClientCertificateConfig(value *IotDomainConfigurationClientCertificateConfig) {
	if err := i.validatePutClientCertificateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClientCertificateConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) PutServerCertificateConfig(value *IotDomainConfigurationServerCertificateConfig) {
	if err := i.validatePutServerCertificateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putServerCertificateConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) PutTlsConfig(value *IotDomainConfigurationTlsConfig) {
	if err := i.validatePutTlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTlsConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := i.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetApplicationProtocol() {
	_jsii_.InvokeVoid(
		i,
		"resetApplicationProtocol",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetAuthorizerConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthorizerConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetClientCertificateConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetClientCertificateConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetDomainConfigurationName() {
	_jsii_.InvokeVoid(
		i,
		"resetDomainConfigurationName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetDomainConfigurationStatus() {
	_jsii_.InvokeVoid(
		i,
		"resetDomainConfigurationStatus",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetDomainName() {
	_jsii_.InvokeVoid(
		i,
		"resetDomainName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetServerCertificateArns() {
	_jsii_.InvokeVoid(
		i,
		"resetServerCertificateArns",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetServerCertificateConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetServerCertificateConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetServiceType() {
	_jsii_.InvokeVoid(
		i,
		"resetServiceType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) ResetValidationCertificateArn() {
	_jsii_.InvokeVoid(
		i,
		"resetValidationCertificateArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotDomainConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotDomainConfiguration) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

