// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kendradatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraDataSourceDataSourceConfigurationOutputReference interface {
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
	ConfluenceConfiguration() KendraDataSourceDataSourceConfigurationConfluenceConfigurationOutputReference
	ConfluenceConfigurationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference
	DatabaseConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	GoogleDriveConfiguration() KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference
	GoogleDriveConfigurationInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OneDriveConfiguration() KendraDataSourceDataSourceConfigurationOneDriveConfigurationOutputReference
	OneDriveConfigurationInput() interface{}
	S3Configuration() KendraDataSourceDataSourceConfigurationS3ConfigurationOutputReference
	S3ConfigurationInput() interface{}
	SalesforceConfiguration() KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference
	SalesforceConfigurationInput() interface{}
	ServiceNowConfiguration() KendraDataSourceDataSourceConfigurationServiceNowConfigurationOutputReference
	ServiceNowConfigurationInput() interface{}
	SharePointConfiguration() KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference
	SharePointConfigurationInput() interface{}
	TemplateConfiguration() KendraDataSourceDataSourceConfigurationTemplateConfigurationOutputReference
	TemplateConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WebCrawlerConfiguration() KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference
	WebCrawlerConfigurationInput() interface{}
	WorkDocsConfiguration() KendraDataSourceDataSourceConfigurationWorkDocsConfigurationOutputReference
	WorkDocsConfigurationInput() interface{}
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
	PutConfluenceConfiguration(value *KendraDataSourceDataSourceConfigurationConfluenceConfiguration)
	PutDatabaseConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfiguration)
	PutGoogleDriveConfiguration(value *KendraDataSourceDataSourceConfigurationGoogleDriveConfiguration)
	PutOneDriveConfiguration(value *KendraDataSourceDataSourceConfigurationOneDriveConfiguration)
	PutS3Configuration(value *KendraDataSourceDataSourceConfigurationS3Configuration)
	PutSalesforceConfiguration(value *KendraDataSourceDataSourceConfigurationSalesforceConfiguration)
	PutServiceNowConfiguration(value *KendraDataSourceDataSourceConfigurationServiceNowConfiguration)
	PutSharePointConfiguration(value *KendraDataSourceDataSourceConfigurationSharePointConfiguration)
	PutTemplateConfiguration(value *KendraDataSourceDataSourceConfigurationTemplateConfiguration)
	PutWebCrawlerConfiguration(value *KendraDataSourceDataSourceConfigurationWebCrawlerConfiguration)
	PutWorkDocsConfiguration(value *KendraDataSourceDataSourceConfigurationWorkDocsConfiguration)
	ResetConfluenceConfiguration()
	ResetDatabaseConfiguration()
	ResetGoogleDriveConfiguration()
	ResetOneDriveConfiguration()
	ResetS3Configuration()
	ResetSalesforceConfiguration()
	ResetServiceNowConfiguration()
	ResetSharePointConfiguration()
	ResetTemplateConfiguration()
	ResetWebCrawlerConfiguration()
	ResetWorkDocsConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceDataSourceConfigurationOutputReference
type jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ConfluenceConfiguration() KendraDataSourceDataSourceConfigurationConfluenceConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationConfluenceConfigurationOutputReference
	_jsii_.Get(
		j,
		"confluenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ConfluenceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confluenceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) DatabaseConfiguration() KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference
	_jsii_.Get(
		j,
		"databaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) DatabaseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) GoogleDriveConfiguration() KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference
	_jsii_.Get(
		j,
		"googleDriveConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) GoogleDriveConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleDriveConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) OneDriveConfiguration() KendraDataSourceDataSourceConfigurationOneDriveConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationOneDriveConfigurationOutputReference
	_jsii_.Get(
		j,
		"oneDriveConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) OneDriveConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oneDriveConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) S3Configuration() KendraDataSourceDataSourceConfigurationS3ConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) S3ConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) SalesforceConfiguration() KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference
	_jsii_.Get(
		j,
		"salesforceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) SalesforceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"salesforceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ServiceNowConfiguration() KendraDataSourceDataSourceConfigurationServiceNowConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationServiceNowConfigurationOutputReference
	_jsii_.Get(
		j,
		"serviceNowConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ServiceNowConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceNowConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) SharePointConfiguration() KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference
	_jsii_.Get(
		j,
		"sharePointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) SharePointConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharePointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) TemplateConfiguration() KendraDataSourceDataSourceConfigurationTemplateConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationTemplateConfigurationOutputReference
	_jsii_.Get(
		j,
		"templateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) TemplateConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) WebCrawlerConfiguration() KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference
	_jsii_.Get(
		j,
		"webCrawlerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) WebCrawlerConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webCrawlerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) WorkDocsConfiguration() KendraDataSourceDataSourceConfigurationWorkDocsConfigurationOutputReference {
	var returns KendraDataSourceDataSourceConfigurationWorkDocsConfigurationOutputReference
	_jsii_.Get(
		j,
		"workDocsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) WorkDocsConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workDocsConfigurationInput",
		&returns,
	)
	return returns
}


func NewKendraDataSourceDataSourceConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KendraDataSourceDataSourceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceDataSourceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceDataSourceConfigurationOutputReference_Override(k KendraDataSourceDataSourceConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) PutConfluenceConfiguration(value *KendraDataSourceDataSourceConfigurationConfluenceConfiguration) {
	if err := k.validatePutConfluenceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putConfluenceConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) PutDatabaseConfiguration(value *KendraDataSourceDataSourceConfigurationDatabaseConfiguration) {
	if err := k.validatePutDatabaseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDatabaseConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) PutGoogleDriveConfiguration(value *KendraDataSourceDataSourceConfigurationGoogleDriveConfiguration) {
	if err := k.validatePutGoogleDriveConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putGoogleDriveConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) PutOneDriveConfiguration(value *KendraDataSourceDataSourceConfigurationOneDriveConfiguration) {
	if err := k.validatePutOneDriveConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putOneDriveConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) PutS3Configuration(value *KendraDataSourceDataSourceConfigurationS3Configuration) {
	if err := k.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) PutSalesforceConfiguration(value *KendraDataSourceDataSourceConfigurationSalesforceConfiguration) {
	if err := k.validatePutSalesforceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSalesforceConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) PutServiceNowConfiguration(value *KendraDataSourceDataSourceConfigurationServiceNowConfiguration) {
	if err := k.validatePutServiceNowConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putServiceNowConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) PutSharePointConfiguration(value *KendraDataSourceDataSourceConfigurationSharePointConfiguration) {
	if err := k.validatePutSharePointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSharePointConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) PutTemplateConfiguration(value *KendraDataSourceDataSourceConfigurationTemplateConfiguration) {
	if err := k.validatePutTemplateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTemplateConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) PutWebCrawlerConfiguration(value *KendraDataSourceDataSourceConfigurationWebCrawlerConfiguration) {
	if err := k.validatePutWebCrawlerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putWebCrawlerConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) PutWorkDocsConfiguration(value *KendraDataSourceDataSourceConfigurationWorkDocsConfiguration) {
	if err := k.validatePutWorkDocsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putWorkDocsConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ResetConfluenceConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetConfluenceConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ResetDatabaseConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetDatabaseConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ResetGoogleDriveConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetGoogleDriveConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ResetOneDriveConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetOneDriveConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ResetS3Configuration() {
	_jsii_.InvokeVoid(
		k,
		"resetS3Configuration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ResetSalesforceConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSalesforceConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ResetServiceNowConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetServiceNowConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ResetSharePointConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSharePointConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ResetTemplateConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetTemplateConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ResetWebCrawlerConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetWebCrawlerConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ResetWorkDocsConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetWorkDocsConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

