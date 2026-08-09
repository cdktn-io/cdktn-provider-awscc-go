// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccconnectcampaignsv2campaign

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccconnectcampaignsv2campaign/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHours
	SetInternalValue(val *DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHours)
	Key() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Value() DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursValueList
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

// The jsii proxy struct for DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference
type jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) InternalValue() *DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHours {
	var returns *DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHours
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) Value() DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursValueList {
	var returns DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursValueList
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewDataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccConnectcampaignsv2Campaign.DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference_Override(d DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccConnectcampaignsv2Campaign.DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference)SetInternalValue(val *DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHours) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

