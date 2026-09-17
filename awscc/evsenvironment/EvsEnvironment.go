// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evsenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/evsenvironment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evs_environment awscc_evs_environment}.
type EvsEnvironment interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Checks() EvsEnvironmentChecksList
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectivityInfo() EvsEnvironmentConnectivityInfoOutputReference
	ConnectivityInfoInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedAt() *string
	Credentials() EvsEnvironmentCredentialsList
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnvironmentArn() *string
	EnvironmentId() *string
	EnvironmentName() *string
	SetEnvironmentName(val *string)
	EnvironmentNameInput() *string
	EnvironmentState() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hosts() EvsEnvironmentHostsList
	HostsInput() interface{}
	Id() *string
	InitialVlans() EvsEnvironmentInitialVlansOutputReference
	InitialVlansInput() interface{}
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	LicenseInfo() EvsEnvironmentLicenseInfoOutputReference
	LicenseInfoInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	ModifiedAt() *string
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
	ServiceAccessSecurityGroups() EvsEnvironmentServiceAccessSecurityGroupsOutputReference
	ServiceAccessSecurityGroupsInput() interface{}
	ServiceAccessSubnetId() *string
	SetServiceAccessSubnetId(val *string)
	ServiceAccessSubnetIdInput() *string
	SiteId() *string
	SetSiteId(val *string)
	SiteIdInput() *string
	StateDetails() *string
	Tags() EvsEnvironmentTagsList
	TagsInput() interface{}
	TermsAccepted() interface{}
	SetTermsAccepted(val interface{})
	TermsAcceptedInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VcfHostnames() EvsEnvironmentVcfHostnamesOutputReference
	VcfHostnamesInput() interface{}
	VcfVersion() *string
	SetVcfVersion(val *string)
	VcfVersionInput() *string
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
	PutConnectivityInfo(value *EvsEnvironmentConnectivityInfo)
	PutHosts(value interface{})
	PutInitialVlans(value *EvsEnvironmentInitialVlans)
	PutLicenseInfo(value *EvsEnvironmentLicenseInfo)
	PutServiceAccessSecurityGroups(value *EvsEnvironmentServiceAccessSecurityGroups)
	PutTags(value interface{})
	PutVcfHostnames(value *EvsEnvironmentVcfHostnames)
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
	ResetConnectivityInfo()
	ResetEnvironmentName()
	ResetHosts()
	ResetInitialVlans()
	ResetKmsKeyId()
	ResetLicenseInfo()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceAccessSecurityGroups()
	ResetSiteId()
	ResetTags()
	ResetVcfHostnames()
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

// The jsii proxy struct for EvsEnvironment
type jsiiProxy_EvsEnvironment struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_EvsEnvironment) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Checks() EvsEnvironmentChecksList {
	var returns EvsEnvironmentChecksList
	_jsii_.Get(
		j,
		"checks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) ConnectivityInfo() EvsEnvironmentConnectivityInfoOutputReference {
	var returns EvsEnvironmentConnectivityInfoOutputReference
	_jsii_.Get(
		j,
		"connectivityInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) ConnectivityInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectivityInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Credentials() EvsEnvironmentCredentialsList {
	var returns EvsEnvironmentCredentialsList
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) EnvironmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) EnvironmentState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Hosts() EvsEnvironmentHostsList {
	var returns EvsEnvironmentHostsList
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) HostsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) InitialVlans() EvsEnvironmentInitialVlansOutputReference {
	var returns EvsEnvironmentInitialVlansOutputReference
	_jsii_.Get(
		j,
		"initialVlans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) InitialVlansInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialVlansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) LicenseInfo() EvsEnvironmentLicenseInfoOutputReference {
	var returns EvsEnvironmentLicenseInfoOutputReference
	_jsii_.Get(
		j,
		"licenseInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) LicenseInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) ServiceAccessSecurityGroups() EvsEnvironmentServiceAccessSecurityGroupsOutputReference {
	var returns EvsEnvironmentServiceAccessSecurityGroupsOutputReference
	_jsii_.Get(
		j,
		"serviceAccessSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) ServiceAccessSecurityGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceAccessSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) ServiceAccessSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) ServiceAccessSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) SiteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) SiteIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) StateDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) Tags() EvsEnvironmentTagsList {
	var returns EvsEnvironmentTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) TermsAccepted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"termsAccepted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) TermsAcceptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"termsAcceptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) VcfHostnames() EvsEnvironmentVcfHostnamesOutputReference {
	var returns EvsEnvironmentVcfHostnamesOutputReference
	_jsii_.Get(
		j,
		"vcfHostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) VcfHostnamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vcfHostnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) VcfVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcfVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) VcfVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcfVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironment) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evs_environment awscc_evs_environment} Resource.
func NewEvsEnvironment(scope constructs.Construct, id *string, config *EvsEnvironmentConfig) EvsEnvironment {
	_init_.Initialize()

	if err := validateNewEvsEnvironmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EvsEnvironment{}

	_jsii_.Create(
		"@cdktn/provider-awscc.evsEnvironment.EvsEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evs_environment awscc_evs_environment} Resource.
func NewEvsEnvironment_Override(e EvsEnvironment, scope constructs.Construct, id *string, config *EvsEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.evsEnvironment.EvsEnvironment",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetEnvironmentName(val *string) {
	if err := j.validateSetEnvironmentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentName",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetServiceAccessSubnetId(val *string) {
	if err := j.validateSetServiceAccessSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessSubnetId",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetSiteId(val *string) {
	if err := j.validateSetSiteIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"siteId",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetTermsAccepted(val interface{}) {
	if err := j.validateSetTermsAcceptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"termsAccepted",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetVcfVersion(val *string) {
	if err := j.validateSetVcfVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcfVersion",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironment)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTN code for importing a EvsEnvironment resource upon running "cdktn plan <stack-name>".
func EvsEnvironment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateEvsEnvironment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.evsEnvironment.EvsEnvironment",
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
func EvsEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEvsEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.evsEnvironment.EvsEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EvsEnvironment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEvsEnvironment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.evsEnvironment.EvsEnvironment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EvsEnvironment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEvsEnvironment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.evsEnvironment.EvsEnvironment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EvsEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.evsEnvironment.EvsEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EvsEnvironment) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EvsEnvironment) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EvsEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EvsEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvsEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EvsEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EvsEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EvsEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EvsEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EvsEnvironment) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EvsEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EvsEnvironment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvsEnvironment) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EvsEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvsEnvironment) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (e *jsiiProxy_EvsEnvironment) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EvsEnvironment) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EvsEnvironment) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EvsEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EvsEnvironment) PutConnectivityInfo(value *EvsEnvironmentConnectivityInfo) {
	if err := e.validatePutConnectivityInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putConnectivityInfo",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironment) PutHosts(value interface{}) {
	if err := e.validatePutHostsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHosts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironment) PutInitialVlans(value *EvsEnvironmentInitialVlans) {
	if err := e.validatePutInitialVlansParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInitialVlans",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironment) PutLicenseInfo(value *EvsEnvironmentLicenseInfo) {
	if err := e.validatePutLicenseInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLicenseInfo",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironment) PutServiceAccessSecurityGroups(value *EvsEnvironmentServiceAccessSecurityGroups) {
	if err := e.validatePutServiceAccessSecurityGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putServiceAccessSecurityGroups",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironment) PutTags(value interface{}) {
	if err := e.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTags",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironment) PutVcfHostnames(value *EvsEnvironmentVcfHostnames) {
	if err := e.validatePutVcfHostnamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putVcfHostnames",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EvsEnvironment) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EvsEnvironment) ResetConnectivityInfo() {
	_jsii_.InvokeVoid(
		e,
		"resetConnectivityInfo",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironment) ResetEnvironmentName() {
	_jsii_.InvokeVoid(
		e,
		"resetEnvironmentName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironment) ResetHosts() {
	_jsii_.InvokeVoid(
		e,
		"resetHosts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironment) ResetInitialVlans() {
	_jsii_.InvokeVoid(
		e,
		"resetInitialVlans",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironment) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		e,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironment) ResetLicenseInfo() {
	_jsii_.InvokeVoid(
		e,
		"resetLicenseInfo",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironment) ResetServiceAccessSecurityGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceAccessSecurityGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironment) ResetSiteId() {
	_jsii_.InvokeVoid(
		e,
		"resetSiteId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironment) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironment) ResetVcfHostnames() {
	_jsii_.InvokeVoid(
		e,
		"resetVcfHostnames",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvsEnvironment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvsEnvironment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvsEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvsEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvsEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvsEnvironment) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

