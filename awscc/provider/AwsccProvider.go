// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/provider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs awscc}.
type AwsccProvider interface {
	cdktn.TerraformProvider
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AssumeRole() *AwsccProviderAssumeRole
	SetAssumeRole(val *AwsccProviderAssumeRole)
	AssumeRoleInput() *AwsccProviderAssumeRole
	AssumeRoleWithWebIdentity() *AwsccProviderAssumeRoleWithWebIdentity
	SetAssumeRoleWithWebIdentity(val *AwsccProviderAssumeRoleWithWebIdentity)
	AssumeRoleWithWebIdentityInput() *AwsccProviderAssumeRoleWithWebIdentity
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Endpoints() *AwsccProviderEndpoints
	SetEndpoints(val *AwsccProviderEndpoints)
	EndpointsInput() *AwsccProviderEndpoints
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpProxy() *string
	SetHttpProxy(val *string)
	HttpProxyInput() *string
	HttpsProxy() *string
	SetHttpsProxy(val *string)
	HttpsProxyInput() *string
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	NoProxy() *string
	SetNoProxy(val *string)
	NoProxyInput() *string
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
	SharedConfigFiles() *[]*string
	SetSharedConfigFiles(val *[]*string)
	SharedConfigFilesInput() *[]*string
	SharedCredentialsFiles() *[]*string
	SetSharedCredentialsFiles(val *[]*string)
	SharedCredentialsFilesInput() *[]*string
	SkipMedatadataApiCheck() interface{}
	SetSkipMedatadataApiCheck(val interface{})
	SkipMedatadataApiCheckInput() interface{}
	SkipMetadataApiCheck() interface{}
	SetSkipMetadataApiCheck(val interface{})
	SkipMetadataApiCheckInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	UserAgent() interface{}
	SetUserAgent(val interface{})
	UserAgentInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
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
	ResetAccessKey()
	ResetAlias()
	ResetAssumeRole()
	ResetAssumeRoleWithWebIdentity()
	ResetEndpoints()
	ResetHttpProxy()
	ResetHttpsProxy()
	ResetInsecure()
	ResetMaxRetries()
	ResetNoProxy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProfile()
	ResetRegion()
	ResetRoleArn()
	ResetSecretKey()
	ResetSharedConfigFiles()
	ResetSharedCredentialsFiles()
	ResetSkipMedatadataApiCheck()
	ResetSkipMetadataApiCheck()
	ResetToken()
	ResetUserAgent()
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

// The jsii proxy struct for AwsccProvider
type jsiiProxy_AwsccProvider struct {
	internal.Type__cdktnTerraformProvider
}

func (j *jsiiProxy_AwsccProvider) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) AssumeRole() *AwsccProviderAssumeRole {
	var returns *AwsccProviderAssumeRole
	_jsii_.Get(
		j,
		"assumeRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) AssumeRoleInput() *AwsccProviderAssumeRole {
	var returns *AwsccProviderAssumeRole
	_jsii_.Get(
		j,
		"assumeRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) AssumeRoleWithWebIdentity() *AwsccProviderAssumeRoleWithWebIdentity {
	var returns *AwsccProviderAssumeRoleWithWebIdentity
	_jsii_.Get(
		j,
		"assumeRoleWithWebIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) AssumeRoleWithWebIdentityInput() *AwsccProviderAssumeRoleWithWebIdentity {
	var returns *AwsccProviderAssumeRoleWithWebIdentity
	_jsii_.Get(
		j,
		"assumeRoleWithWebIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) Endpoints() *AwsccProviderEndpoints {
	var returns *AwsccProviderEndpoints
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) EndpointsInput() *AwsccProviderEndpoints {
	var returns *AwsccProviderEndpoints
	_jsii_.Get(
		j,
		"endpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) HttpProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) HttpProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) HttpsProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) HttpsProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) NoProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) NoProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) SharedConfigFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedConfigFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) SharedConfigFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedConfigFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) SharedCredentialsFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedCredentialsFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) SharedCredentialsFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedCredentialsFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) SkipMedatadataApiCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMedatadataApiCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) SkipMedatadataApiCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMedatadataApiCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) SkipMetadataApiCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMetadataApiCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) SkipMetadataApiCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMetadataApiCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) UserAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsccProvider) UserAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAgentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs awscc} Resource.
func NewAwsccProvider(scope constructs.Construct, id *string, config *AwsccProviderConfig) AwsccProvider {
	_init_.Initialize()

	if err := validateNewAwsccProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsccProvider{}

	_jsii_.Create(
		"@cdktn/provider-awscc.provider.AwsccProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs awscc} Resource.
func NewAwsccProvider_Override(a AwsccProvider, scope constructs.Construct, id *string, config *AwsccProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.provider.AwsccProvider",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsccProvider)SetAccessKey(val *string) {
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetAssumeRole(val *AwsccProviderAssumeRole) {
	if err := j.validateSetAssumeRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assumeRole",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetAssumeRoleWithWebIdentity(val *AwsccProviderAssumeRoleWithWebIdentity) {
	if err := j.validateSetAssumeRoleWithWebIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assumeRoleWithWebIdentity",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetEndpoints(val *AwsccProviderEndpoints) {
	if err := j.validateSetEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoints",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetHttpProxy(val *string) {
	_jsii_.Set(
		j,
		"httpProxy",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetHttpsProxy(val *string) {
	_jsii_.Set(
		j,
		"httpsProxy",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetInsecure(val interface{}) {
	if err := j.validateSetInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetNoProxy(val *string) {
	_jsii_.Set(
		j,
		"noProxy",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetProfile(val *string) {
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetSecretKey(val *string) {
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetSharedConfigFiles(val *[]*string) {
	_jsii_.Set(
		j,
		"sharedConfigFiles",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetSharedCredentialsFiles(val *[]*string) {
	_jsii_.Set(
		j,
		"sharedCredentialsFiles",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetSkipMedatadataApiCheck(val interface{}) {
	if err := j.validateSetSkipMedatadataApiCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipMedatadataApiCheck",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetSkipMetadataApiCheck(val interface{}) {
	if err := j.validateSetSkipMetadataApiCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipMetadataApiCheck",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_AwsccProvider)SetUserAgent(val interface{}) {
	if err := j.validateSetUserAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userAgent",
		val,
	)
}

// Generates CDKTN code for importing a AwsccProvider resource upon running "cdktn plan <stack-name>".
func AwsccProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsccProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.provider.AwsccProvider",
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
func AwsccProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsccProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.provider.AwsccProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsccProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsccProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.provider.AwsccProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsccProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsccProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.provider.AwsccProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsccProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.provider.AwsccProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsccProvider) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsccProvider) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsccProvider) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsccProvider) ResetAccessKey() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetAssumeRole() {
	_jsii_.InvokeVoid(
		a,
		"resetAssumeRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetAssumeRoleWithWebIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetAssumeRoleWithWebIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetEndpoints() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpoints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetHttpProxy() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpProxy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetHttpsProxy() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpsProxy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		a,
		"resetInsecure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetNoProxy() {
	_jsii_.InvokeVoid(
		a,
		"resetNoProxy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetSecretKey() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetSharedConfigFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedConfigFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetSharedCredentialsFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedCredentialsFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetSkipMedatadataApiCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipMedatadataApiCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetSkipMetadataApiCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipMetadataApiCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetToken() {
	_jsii_.InvokeVoid(
		a,
		"resetToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) ResetUserAgent() {
	_jsii_.InvokeVoid(
		a,
		"resetUserAgent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsccProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsccProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsccProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsccProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsccProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsccProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsccProvider) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

