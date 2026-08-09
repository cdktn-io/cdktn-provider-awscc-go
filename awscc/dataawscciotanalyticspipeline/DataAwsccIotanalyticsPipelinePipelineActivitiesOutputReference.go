// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawscciotanalyticspipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawscciotanalyticspipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference interface {
	cdktn.ComplexObject
	AddAttributes() DataAwsccIotanalyticsPipelinePipelineActivitiesAddAttributesOutputReference
	Channel() DataAwsccIotanalyticsPipelinePipelineActivitiesChannelOutputReference
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
	Datastore() DataAwsccIotanalyticsPipelinePipelineActivitiesDatastoreOutputReference
	DeviceRegistryEnrich() DataAwsccIotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference
	DeviceShadowEnrich() DataAwsccIotanalyticsPipelinePipelineActivitiesDeviceShadowEnrichOutputReference
	Filter() DataAwsccIotanalyticsPipelinePipelineActivitiesFilterOutputReference
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccIotanalyticsPipelinePipelineActivities
	SetInternalValue(val *DataAwsccIotanalyticsPipelinePipelineActivities)
	Lambda() DataAwsccIotanalyticsPipelinePipelineActivitiesLambdaOutputReference
	Math() DataAwsccIotanalyticsPipelinePipelineActivitiesMathOutputReference
	RemoveAttributes() DataAwsccIotanalyticsPipelinePipelineActivitiesRemoveAttributesOutputReference
	SelectAttributes() DataAwsccIotanalyticsPipelinePipelineActivitiesSelectAttributesOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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

// The jsii proxy struct for DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference
type jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) AddAttributes() DataAwsccIotanalyticsPipelinePipelineActivitiesAddAttributesOutputReference {
	var returns DataAwsccIotanalyticsPipelinePipelineActivitiesAddAttributesOutputReference
	_jsii_.Get(
		j,
		"addAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) Channel() DataAwsccIotanalyticsPipelinePipelineActivitiesChannelOutputReference {
	var returns DataAwsccIotanalyticsPipelinePipelineActivitiesChannelOutputReference
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) Datastore() DataAwsccIotanalyticsPipelinePipelineActivitiesDatastoreOutputReference {
	var returns DataAwsccIotanalyticsPipelinePipelineActivitiesDatastoreOutputReference
	_jsii_.Get(
		j,
		"datastore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) DeviceRegistryEnrich() DataAwsccIotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference {
	var returns DataAwsccIotanalyticsPipelinePipelineActivitiesDeviceRegistryEnrichOutputReference
	_jsii_.Get(
		j,
		"deviceRegistryEnrich",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) DeviceShadowEnrich() DataAwsccIotanalyticsPipelinePipelineActivitiesDeviceShadowEnrichOutputReference {
	var returns DataAwsccIotanalyticsPipelinePipelineActivitiesDeviceShadowEnrichOutputReference
	_jsii_.Get(
		j,
		"deviceShadowEnrich",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) Filter() DataAwsccIotanalyticsPipelinePipelineActivitiesFilterOutputReference {
	var returns DataAwsccIotanalyticsPipelinePipelineActivitiesFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) InternalValue() *DataAwsccIotanalyticsPipelinePipelineActivities {
	var returns *DataAwsccIotanalyticsPipelinePipelineActivities
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) Lambda() DataAwsccIotanalyticsPipelinePipelineActivitiesLambdaOutputReference {
	var returns DataAwsccIotanalyticsPipelinePipelineActivitiesLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) Math() DataAwsccIotanalyticsPipelinePipelineActivitiesMathOutputReference {
	var returns DataAwsccIotanalyticsPipelinePipelineActivitiesMathOutputReference
	_jsii_.Get(
		j,
		"math",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) RemoveAttributes() DataAwsccIotanalyticsPipelinePipelineActivitiesRemoveAttributesOutputReference {
	var returns DataAwsccIotanalyticsPipelinePipelineActivitiesRemoveAttributesOutputReference
	_jsii_.Get(
		j,
		"removeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) SelectAttributes() DataAwsccIotanalyticsPipelinePipelineActivitiesSelectAttributesOutputReference {
	var returns DataAwsccIotanalyticsPipelinePipelineActivitiesSelectAttributesOutputReference
	_jsii_.Get(
		j,
		"selectAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccIotanalyticsPipelinePipelineActivitiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccIotanalyticsPipeline.DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference_Override(d DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccIotanalyticsPipeline.DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference)SetInternalValue(val *DataAwsccIotanalyticsPipelinePipelineActivities) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccIotanalyticsPipelinePipelineActivitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

