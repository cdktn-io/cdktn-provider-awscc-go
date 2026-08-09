// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kendradatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	CrawlAttachments() interface{}
	SetCrawlAttachments(val interface{})
	CrawlAttachmentsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableLocalGroups() interface{}
	SetDisableLocalGroups(val interface{})
	DisableLocalGroupsInput() interface{}
	DocumentTitleFieldName() *string
	SetDocumentTitleFieldName(val *string)
	DocumentTitleFieldNameInput() *string
	ExclusionPatterns() *[]*string
	SetExclusionPatterns(val *[]*string)
	ExclusionPatternsInput() *[]*string
	FieldMappings() KendraDataSourceDataSourceConfigurationSharePointConfigurationFieldMappingsList
	FieldMappingsInput() interface{}
	// Experimental.
	Fqn() *string
	InclusionPatterns() *[]*string
	SetInclusionPatterns(val *[]*string)
	InclusionPatternsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SecretArn() *string
	SetSecretArn(val *string)
	SecretArnInput() *string
	SharePointVersion() *string
	SetSharePointVersion(val *string)
	SharePointVersionInput() *string
	SslCertificateS3Path() KendraDataSourceDataSourceConfigurationSharePointConfigurationSslCertificateS3PathOutputReference
	SslCertificateS3PathInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Urls() *[]*string
	SetUrls(val *[]*string)
	UrlsInput() *[]*string
	UseChangeLog() interface{}
	SetUseChangeLog(val interface{})
	UseChangeLogInput() interface{}
	VpcConfiguration() KendraDataSourceDataSourceConfigurationSharePointConfigurationVpcConfigurationOutputReference
	VpcConfigurationInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutFieldMappings(value interface{})
	PutSslCertificateS3Path(value *KendraDataSourceDataSourceConfigurationSharePointConfigurationSslCertificateS3Path)
	PutVpcConfiguration(value *KendraDataSourceDataSourceConfigurationSharePointConfigurationVpcConfiguration)
	ResetCrawlAttachments()
	ResetDisableLocalGroups()
	ResetDocumentTitleFieldName()
	ResetExclusionPatterns()
	ResetFieldMappings()
	ResetInclusionPatterns()
	ResetSecretArn()
	ResetSharePointVersion()
	ResetSslCertificateS3Path()
	ResetUrls()
	ResetUseChangeLog()
	ResetVpcConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference
type jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) CrawlAttachments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) CrawlAttachmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crawlAttachmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) DisableLocalGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableLocalGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) DisableLocalGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableLocalGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) DocumentTitleFieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTitleFieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) DocumentTitleFieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTitleFieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ExclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ExclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) FieldMappings() KendraDataSourceDataSourceConfigurationSharePointConfigurationFieldMappingsList {
	var returns KendraDataSourceDataSourceConfigurationSharePointConfigurationFieldMappingsList
	_jsii_.Get(
		j,
		"fieldMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) FieldMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) InclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) InclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) SecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) SharePointVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePointVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) SharePointVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePointVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) SslCertificateS3Path() KendraDataSourceDataSourceConfigurationSharePointConfigurationSslCertificateS3PathOutputReference {
	var returns KendraDataSourceDataSourceConfigurationSharePointConfigurationSslCertificateS3PathOutputReference
	_jsii_.Get(
		j,
		"sslCertificateS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) SslCertificateS3PathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslCertificateS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) Urls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) UrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) UseChangeLog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useChangeLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) UseChangeLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useChangeLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) VpcConfiguration() KendraDataSourceDataSourceConfigurationSharePointConfigurationVpcConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationSharePointConfigurationVpcConfigurationOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) VpcConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


func NewKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference_Override(k KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetCrawlAttachments(val interface{}) {
	if err := j.validateSetCrawlAttachmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crawlAttachments",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetDisableLocalGroups(val interface{}) {
	if err := j.validateSetDisableLocalGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableLocalGroups",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetDocumentTitleFieldName(val *string) {
	if err := j.validateSetDocumentTitleFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentTitleFieldName",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetExclusionPatterns(val *[]*string) {
	if err := j.validateSetExclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionPatterns",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetInclusionPatterns(val *[]*string) {
	if err := j.validateSetInclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inclusionPatterns",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetSecretArn(val *string) {
	if err := j.validateSetSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetSharePointVersion(val *string) {
	if err := j.validateSetSharePointVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharePointVersion",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetUrls(val *[]*string) {
	if err := j.validateSetUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urls",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference)SetUseChangeLog(val interface{}) {
	if err := j.validateSetUseChangeLogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useChangeLog",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) PutFieldMappings(value interface{}) {
	if err := k.validatePutFieldMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putFieldMappings",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) PutSslCertificateS3Path(value *KendraDataSourceDataSourceConfigurationSharePointConfigurationSslCertificateS3Path) {
	if err := k.validatePutSslCertificateS3PathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSslCertificateS3Path",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) PutVpcConfiguration(value *KendraDataSourceDataSourceConfigurationSharePointConfigurationVpcConfiguration) {
	if err := k.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetCrawlAttachments() {
	_jsii_.InvokeVoid(
		k,
		"resetCrawlAttachments",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetDisableLocalGroups() {
	_jsii_.InvokeVoid(
		k,
		"resetDisableLocalGroups",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetDocumentTitleFieldName() {
	_jsii_.InvokeVoid(
		k,
		"resetDocumentTitleFieldName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetExclusionPatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetExclusionPatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetFieldMappings() {
	_jsii_.InvokeVoid(
		k,
		"resetFieldMappings",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetInclusionPatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetInclusionPatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetSecretArn() {
	_jsii_.InvokeVoid(
		k,
		"resetSecretArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetSharePointVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetSharePointVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetSslCertificateS3Path() {
	_jsii_.InvokeVoid(
		k,
		"resetSslCertificateS3Path",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetUrls() {
	_jsii_.InvokeVoid(
		k,
		"resetUrls",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetUseChangeLog() {
	_jsii_.InvokeVoid(
		k,
		"resetUseChangeLog",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ResetVpcConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetVpcConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

