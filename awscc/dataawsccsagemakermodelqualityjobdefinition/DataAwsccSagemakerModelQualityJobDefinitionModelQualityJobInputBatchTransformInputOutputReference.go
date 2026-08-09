// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccsagemakermodelqualityjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccsagemakermodelqualityjobdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference interface {
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
	DataCapturedDestinationS3Uri() *string
	DatasetFormat() DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputDatasetFormatOutputReference
	EndTimeOffset() *string
	// Experimental.
	Fqn() *string
	InferenceAttribute() *string
	InternalValue() *DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInput
	SetInternalValue(val *DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInput)
	LocalPath() *string
	ProbabilityAttribute() *string
	ProbabilityThresholdAttribute() *float64
	S3DataDistributionType() *string
	S3InputMode() *string
	StartTimeOffset() *string
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

// The jsii proxy struct for DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference
type jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) DataCapturedDestinationS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCapturedDestinationS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) DatasetFormat() DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputDatasetFormatOutputReference {
	var returns DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputDatasetFormatOutputReference
	_jsii_.Get(
		j,
		"datasetFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) EndTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) InferenceAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) InternalValue() *DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInput {
	var returns *DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) ProbabilityAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probabilityAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) ProbabilityThresholdAttribute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"probabilityThresholdAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) S3DataDistributionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataDistributionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) S3InputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3InputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) StartTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSagemakerModelQualityJobDefinition.DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference_Override(d DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSagemakerModelQualityJobDefinition.DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference)SetInternalValue(val *DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccSagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

