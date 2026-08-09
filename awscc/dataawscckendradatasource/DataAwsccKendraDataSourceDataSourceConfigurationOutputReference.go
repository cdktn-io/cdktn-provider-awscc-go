// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscckendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscckendradatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccKendraDataSourceDataSourceConfigurationOutputReference interface {
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
	ConfluenceConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationConfluenceConfigurationOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference
	// Experimental.
	Fqn() *string
	GoogleDriveConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference
	InternalValue() *DataAwsccKendraDataSourceDataSourceConfiguration
	SetInternalValue(val *DataAwsccKendraDataSourceDataSourceConfiguration)
	OneDriveConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationOneDriveConfigurationOutputReference
	S3Configuration() DataAwsccKendraDataSourceDataSourceConfigurationS3ConfigurationOutputReference
	SalesforceConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference
	ServiceNowConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationServiceNowConfigurationOutputReference
	SharePointConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference
	TemplateConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationTemplateConfigurationOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WebCrawlerConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference
	WorkDocsConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationWorkDocsConfigurationOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccKendraDataSourceDataSourceConfigurationOutputReference
type jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) ConfluenceConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationConfluenceConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationConfluenceConfigurationOutputReference
	_jsii_.Get(
		j,
		"confluenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) DatabaseConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationDatabaseConfigurationOutputReference
	_jsii_.Get(
		j,
		"databaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) GoogleDriveConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference
	_jsii_.Get(
		j,
		"googleDriveConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) InternalValue() *DataAwsccKendraDataSourceDataSourceConfiguration {
	var returns *DataAwsccKendraDataSourceDataSourceConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) OneDriveConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationOneDriveConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationOneDriveConfigurationOutputReference
	_jsii_.Get(
		j,
		"oneDriveConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) S3Configuration() DataAwsccKendraDataSourceDataSourceConfigurationS3ConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationS3ConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) SalesforceConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationSalesforceConfigurationOutputReference
	_jsii_.Get(
		j,
		"salesforceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) ServiceNowConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationServiceNowConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationServiceNowConfigurationOutputReference
	_jsii_.Get(
		j,
		"serviceNowConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) SharePointConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationSharePointConfigurationOutputReference
	_jsii_.Get(
		j,
		"sharePointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) TemplateConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationTemplateConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationTemplateConfigurationOutputReference
	_jsii_.Get(
		j,
		"templateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) WebCrawlerConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationWebCrawlerConfigurationOutputReference
	_jsii_.Get(
		j,
		"webCrawlerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) WorkDocsConfiguration() DataAwsccKendraDataSourceDataSourceConfigurationWorkDocsConfigurationOutputReference {
	var returns DataAwsccKendraDataSourceDataSourceConfigurationWorkDocsConfigurationOutputReference
	_jsii_.Get(
		j,
		"workDocsConfiguration",
		&returns,
	)
	return returns
}


func NewDataAwsccKendraDataSourceDataSourceConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccKendraDataSourceDataSourceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccKendraDataSourceDataSourceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccKendraDataSource.DataAwsccKendraDataSourceDataSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccKendraDataSourceDataSourceConfigurationOutputReference_Override(d DataAwsccKendraDataSourceDataSourceConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccKendraDataSource.DataAwsccKendraDataSourceDataSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference)SetInternalValue(val *DataAwsccKendraDataSourceDataSourceConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKendraDataSourceDataSourceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

