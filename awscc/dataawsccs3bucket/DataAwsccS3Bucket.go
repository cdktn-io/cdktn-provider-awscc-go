// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccs3bucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccs3bucket/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/data-sources/s3_bucket awscc_s3_bucket}.
type DataAwsccS3Bucket interface {
	cdktn.TerraformDataSource
	AccelerateConfiguration() DataAwsccS3BucketAccelerateConfigurationOutputReference
	AccessControl() *string
	AnalyticsConfigurations() DataAwsccS3BucketAnalyticsConfigurationsList
	Arn() *string
	BucketEncryption() DataAwsccS3BucketBucketEncryptionOutputReference
	BucketName() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CorsConfiguration() DataAwsccS3BucketCorsConfigurationOutputReference
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DomainName() *string
	DualStackDomainName() *string
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
	IntelligentTieringConfigurations() DataAwsccS3BucketIntelligentTieringConfigurationsList
	InventoryConfigurations() DataAwsccS3BucketInventoryConfigurationsList
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LifecycleConfiguration() DataAwsccS3BucketLifecycleConfigurationOutputReference
	LoggingConfiguration() DataAwsccS3BucketLoggingConfigurationOutputReference
	MetadataTableConfiguration() DataAwsccS3BucketMetadataTableConfigurationOutputReference
	MetricsConfigurations() DataAwsccS3BucketMetricsConfigurationsList
	// The tree node.
	Node() constructs.Node
	NotificationConfiguration() DataAwsccS3BucketNotificationConfigurationOutputReference
	ObjectLockConfiguration() DataAwsccS3BucketObjectLockConfigurationOutputReference
	ObjectLockEnabled() cdktn.IResolvable
	OwnershipControls() DataAwsccS3BucketOwnershipControlsOutputReference
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	PublicAccessBlockConfiguration() DataAwsccS3BucketPublicAccessBlockConfigurationOutputReference
	// Experimental.
	RawOverrides() interface{}
	RegionalDomainName() *string
	ReplicationConfiguration() DataAwsccS3BucketReplicationConfigurationOutputReference
	Tags() DataAwsccS3BucketTagsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VersioningConfiguration() DataAwsccS3BucketVersioningConfigurationOutputReference
	WebsiteConfiguration() DataAwsccS3BucketWebsiteConfigurationOutputReference
	WebsiteUrl() *string
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

// The jsii proxy struct for DataAwsccS3Bucket
type jsiiProxy_DataAwsccS3Bucket struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAwsccS3Bucket) AccelerateConfiguration() DataAwsccS3BucketAccelerateConfigurationOutputReference {
	var returns DataAwsccS3BucketAccelerateConfigurationOutputReference
	_jsii_.Get(
		j,
		"accelerateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) AccessControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) AnalyticsConfigurations() DataAwsccS3BucketAnalyticsConfigurationsList {
	var returns DataAwsccS3BucketAnalyticsConfigurationsList
	_jsii_.Get(
		j,
		"analyticsConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) BucketEncryption() DataAwsccS3BucketBucketEncryptionOutputReference {
	var returns DataAwsccS3BucketBucketEncryptionOutputReference
	_jsii_.Get(
		j,
		"bucketEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) CorsConfiguration() DataAwsccS3BucketCorsConfigurationOutputReference {
	var returns DataAwsccS3BucketCorsConfigurationOutputReference
	_jsii_.Get(
		j,
		"corsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) DualStackDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dualStackDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) IntelligentTieringConfigurations() DataAwsccS3BucketIntelligentTieringConfigurationsList {
	var returns DataAwsccS3BucketIntelligentTieringConfigurationsList
	_jsii_.Get(
		j,
		"intelligentTieringConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) InventoryConfigurations() DataAwsccS3BucketInventoryConfigurationsList {
	var returns DataAwsccS3BucketInventoryConfigurationsList
	_jsii_.Get(
		j,
		"inventoryConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) LifecycleConfiguration() DataAwsccS3BucketLifecycleConfigurationOutputReference {
	var returns DataAwsccS3BucketLifecycleConfigurationOutputReference
	_jsii_.Get(
		j,
		"lifecycleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) LoggingConfiguration() DataAwsccS3BucketLoggingConfigurationOutputReference {
	var returns DataAwsccS3BucketLoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) MetadataTableConfiguration() DataAwsccS3BucketMetadataTableConfigurationOutputReference {
	var returns DataAwsccS3BucketMetadataTableConfigurationOutputReference
	_jsii_.Get(
		j,
		"metadataTableConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) MetricsConfigurations() DataAwsccS3BucketMetricsConfigurationsList {
	var returns DataAwsccS3BucketMetricsConfigurationsList
	_jsii_.Get(
		j,
		"metricsConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) NotificationConfiguration() DataAwsccS3BucketNotificationConfigurationOutputReference {
	var returns DataAwsccS3BucketNotificationConfigurationOutputReference
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) ObjectLockConfiguration() DataAwsccS3BucketObjectLockConfigurationOutputReference {
	var returns DataAwsccS3BucketObjectLockConfigurationOutputReference
	_jsii_.Get(
		j,
		"objectLockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) ObjectLockEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"objectLockEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) OwnershipControls() DataAwsccS3BucketOwnershipControlsOutputReference {
	var returns DataAwsccS3BucketOwnershipControlsOutputReference
	_jsii_.Get(
		j,
		"ownershipControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) PublicAccessBlockConfiguration() DataAwsccS3BucketPublicAccessBlockConfigurationOutputReference {
	var returns DataAwsccS3BucketPublicAccessBlockConfigurationOutputReference
	_jsii_.Get(
		j,
		"publicAccessBlockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) ReplicationConfiguration() DataAwsccS3BucketReplicationConfigurationOutputReference {
	var returns DataAwsccS3BucketReplicationConfigurationOutputReference
	_jsii_.Get(
		j,
		"replicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) Tags() DataAwsccS3BucketTagsList {
	var returns DataAwsccS3BucketTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) VersioningConfiguration() DataAwsccS3BucketVersioningConfigurationOutputReference {
	var returns DataAwsccS3BucketVersioningConfigurationOutputReference
	_jsii_.Get(
		j,
		"versioningConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) WebsiteConfiguration() DataAwsccS3BucketWebsiteConfigurationOutputReference {
	var returns DataAwsccS3BucketWebsiteConfigurationOutputReference
	_jsii_.Get(
		j,
		"websiteConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccS3Bucket) WebsiteUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteUrl",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/data-sources/s3_bucket awscc_s3_bucket} Data Source.
func NewDataAwsccS3Bucket(scope constructs.Construct, id *string, config *DataAwsccS3BucketConfig) DataAwsccS3Bucket {
	_init_.Initialize()

	if err := validateNewDataAwsccS3BucketParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccS3Bucket{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccS3Bucket.DataAwsccS3Bucket",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/data-sources/s3_bucket awscc_s3_bucket} Data Source.
func NewDataAwsccS3Bucket_Override(d DataAwsccS3Bucket, scope constructs.Construct, id *string, config *DataAwsccS3BucketConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccS3Bucket.DataAwsccS3Bucket",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccS3Bucket)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3Bucket)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3Bucket)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3Bucket)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3Bucket)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccS3Bucket)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTN code for importing a DataAwsccS3Bucket resource upon running "cdktn plan <stack-name>".
func DataAwsccS3Bucket_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccS3Bucket_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccS3Bucket.DataAwsccS3Bucket",
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
func DataAwsccS3Bucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccS3Bucket_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccS3Bucket.DataAwsccS3Bucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccS3Bucket_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccS3Bucket_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccS3Bucket.DataAwsccS3Bucket",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccS3Bucket_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccS3Bucket_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dataAwsccS3Bucket.DataAwsccS3Bucket",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccS3Bucket_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.dataAwsccS3Bucket.DataAwsccS3Bucket",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccS3Bucket) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccS3Bucket) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccS3Bucket) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccS3Bucket) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccS3Bucket) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccS3Bucket) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccS3Bucket) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccS3Bucket) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccS3Bucket) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccS3Bucket) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccS3Bucket) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccS3Bucket) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccS3Bucket) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAwsccS3Bucket) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccS3Bucket) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3Bucket) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3Bucket) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3Bucket) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3Bucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3Bucket) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccS3Bucket) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

