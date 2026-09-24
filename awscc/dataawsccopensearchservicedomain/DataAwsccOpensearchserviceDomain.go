// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccopensearchservicedomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccopensearchservicedomain/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/data-sources/opensearchservice_domain awscc_opensearchservice_domain}.
type DataAwsccOpensearchserviceDomain interface {
	cdktn.TerraformDataSource
	AccessPolicies() *string
	AdvancedOptions() cdktn.StringMap
	AdvancedSecurityOptions() DataAwsccOpensearchserviceDomainAdvancedSecurityOptionsOutputReference
	AimlOptions() DataAwsccOpensearchserviceDomainAimlOptionsOutputReference
	Arn() *string
	AutomatedSnapshotPauseOptions() DataAwsccOpensearchserviceDomainAutomatedSnapshotPauseOptionsOutputReference
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClusterConfig() DataAwsccOpensearchserviceDomainClusterConfigOutputReference
	CognitoOptions() DataAwsccOpensearchserviceDomainCognitoOptionsOutputReference
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
	DeploymentStrategyOptions() DataAwsccOpensearchserviceDomainDeploymentStrategyOptionsOutputReference
	DomainArn() *string
	DomainEndpoint() *string
	DomainEndpointOptions() DataAwsccOpensearchserviceDomainDomainEndpointOptionsOutputReference
	DomainEndpoints() cdktn.StringMap
	DomainEndpointV2() *string
	DomainId() *string
	DomainName() *string
	EbsOptions() DataAwsccOpensearchserviceDomainEbsOptionsOutputReference
	EncryptionAtRestOptions() DataAwsccOpensearchserviceDomainEncryptionAtRestOptionsOutputReference
	EngineMode() *string
	EngineVersion() *string
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
	IdentityCenterOptions() DataAwsccOpensearchserviceDomainIdentityCenterOptionsOutputReference
	IdInput() *string
	IpAddressType() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogPublishingOptions() DataAwsccOpensearchserviceDomainLogPublishingOptionsMap
	// The tree node.
	Node() constructs.Node
	NodeToNodeEncryptionOptions() DataAwsccOpensearchserviceDomainNodeToNodeEncryptionOptionsOutputReference
	OffPeakWindowOptions() DataAwsccOpensearchserviceDomainOffPeakWindowOptionsOutputReference
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ServiceSoftwareOptions() DataAwsccOpensearchserviceDomainServiceSoftwareOptionsOutputReference
	SkipShardMigrationWait() cdktn.IResolvable
	SnapshotOptions() DataAwsccOpensearchserviceDomainSnapshotOptionsOutputReference
	SoftwareUpdateOptions() DataAwsccOpensearchserviceDomainSoftwareUpdateOptionsOutputReference
	Tags() DataAwsccOpensearchserviceDomainTagsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UseCase() *string
	VpcOptions() DataAwsccOpensearchserviceDomainVpcOptionsOutputReference
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

// The jsii proxy struct for DataAwsccOpensearchserviceDomain
type jsiiProxy_DataAwsccOpensearchserviceDomain struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) AccessPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) AdvancedOptions() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"advancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) AdvancedSecurityOptions() DataAwsccOpensearchserviceDomainAdvancedSecurityOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainAdvancedSecurityOptionsOutputReference
	_jsii_.Get(
		j,
		"advancedSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) AimlOptions() DataAwsccOpensearchserviceDomainAimlOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainAimlOptionsOutputReference
	_jsii_.Get(
		j,
		"aimlOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) AutomatedSnapshotPauseOptions() DataAwsccOpensearchserviceDomainAutomatedSnapshotPauseOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainAutomatedSnapshotPauseOptionsOutputReference
	_jsii_.Get(
		j,
		"automatedSnapshotPauseOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) ClusterConfig() DataAwsccOpensearchserviceDomainClusterConfigOutputReference {
	var returns DataAwsccOpensearchserviceDomainClusterConfigOutputReference
	_jsii_.Get(
		j,
		"clusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) CognitoOptions() DataAwsccOpensearchserviceDomainCognitoOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainCognitoOptionsOutputReference
	_jsii_.Get(
		j,
		"cognitoOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) DeploymentStrategyOptions() DataAwsccOpensearchserviceDomainDeploymentStrategyOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainDeploymentStrategyOptionsOutputReference
	_jsii_.Get(
		j,
		"deploymentStrategyOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) DomainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) DomainEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) DomainEndpointOptions() DataAwsccOpensearchserviceDomainDomainEndpointOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainDomainEndpointOptionsOutputReference
	_jsii_.Get(
		j,
		"domainEndpointOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) DomainEndpoints() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"domainEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) DomainEndpointV2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpointV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) EbsOptions() DataAwsccOpensearchserviceDomainEbsOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainEbsOptionsOutputReference
	_jsii_.Get(
		j,
		"ebsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) EncryptionAtRestOptions() DataAwsccOpensearchserviceDomainEncryptionAtRestOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainEncryptionAtRestOptionsOutputReference
	_jsii_.Get(
		j,
		"encryptionAtRestOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) EngineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) IdentityCenterOptions() DataAwsccOpensearchserviceDomainIdentityCenterOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainIdentityCenterOptionsOutputReference
	_jsii_.Get(
		j,
		"identityCenterOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) LogPublishingOptions() DataAwsccOpensearchserviceDomainLogPublishingOptionsMap {
	var returns DataAwsccOpensearchserviceDomainLogPublishingOptionsMap
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) NodeToNodeEncryptionOptions() DataAwsccOpensearchserviceDomainNodeToNodeEncryptionOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainNodeToNodeEncryptionOptionsOutputReference
	_jsii_.Get(
		j,
		"nodeToNodeEncryptionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) OffPeakWindowOptions() DataAwsccOpensearchserviceDomainOffPeakWindowOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainOffPeakWindowOptionsOutputReference
	_jsii_.Get(
		j,
		"offPeakWindowOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) ServiceSoftwareOptions() DataAwsccOpensearchserviceDomainServiceSoftwareOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainServiceSoftwareOptionsOutputReference
	_jsii_.Get(
		j,
		"serviceSoftwareOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) SkipShardMigrationWait() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"skipShardMigrationWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) SnapshotOptions() DataAwsccOpensearchserviceDomainSnapshotOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainSnapshotOptionsOutputReference
	_jsii_.Get(
		j,
		"snapshotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) SoftwareUpdateOptions() DataAwsccOpensearchserviceDomainSoftwareUpdateOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainSoftwareUpdateOptionsOutputReference
	_jsii_.Get(
		j,
		"softwareUpdateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) Tags() DataAwsccOpensearchserviceDomainTagsList {
	var returns DataAwsccOpensearchserviceDomainTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) UseCase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain) VpcOptions() DataAwsccOpensearchserviceDomainVpcOptionsOutputReference {
	var returns DataAwsccOpensearchserviceDomainVpcOptionsOutputReference
	_jsii_.Get(
		j,
		"vpcOptions",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/data-sources/opensearchservice_domain awscc_opensearchservice_domain} Data Source.
func NewDataAwsccOpensearchserviceDomain(scope constructs.Construct, id *string, config *DataAwsccOpensearchserviceDomainConfig) DataAwsccOpensearchserviceDomain {
	_init_.Initialize()

	if err := validateNewDataAwsccOpensearchserviceDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccOpensearchserviceDomain{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccOpensearchserviceDomain.DataAwsccOpensearchserviceDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/data-sources/opensearchservice_domain awscc_opensearchservice_domain} Data Source.
func NewDataAwsccOpensearchserviceDomain_Override(d DataAwsccOpensearchserviceDomain, scope constructs.Construct, id *string, config *DataAwsccOpensearchserviceDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccOpensearchserviceDomain.DataAwsccOpensearchserviceDomain",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccOpensearchserviceDomain)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsccOpensearchserviceDomain resource upon running "cdktn plan <stack-name>".
func DataAwsccOpensearchserviceDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccOpensearchserviceDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOpensearchserviceDomain.DataAwsccOpensearchserviceDomain",
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
func DataAwsccOpensearchserviceDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccOpensearchserviceDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOpensearchserviceDomain.DataAwsccOpensearchserviceDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccOpensearchserviceDomain_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccOpensearchserviceDomain_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOpensearchserviceDomain.DataAwsccOpensearchserviceDomain",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccOpensearchserviceDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccOpensearchserviceDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccOpensearchserviceDomain.DataAwsccOpensearchserviceDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccOpensearchserviceDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.dataAwsccOpensearchserviceDomain.DataAwsccOpensearchserviceDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccOpensearchserviceDomain) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

