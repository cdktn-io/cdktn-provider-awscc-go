// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccpipespipe

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccpipespipe/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccPipesPipeTargetParametersOutputReference interface {
	cdktn.ComplexObject
	BatchJobParameters() DataAwsccPipesPipeTargetParametersBatchJobParametersOutputReference
	CloudwatchLogsParameters() DataAwsccPipesPipeTargetParametersCloudwatchLogsParametersOutputReference
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
	EcsTaskParameters() DataAwsccPipesPipeTargetParametersEcsTaskParametersOutputReference
	EventBridgeEventBusParameters() DataAwsccPipesPipeTargetParametersEventBridgeEventBusParametersOutputReference
	// Experimental.
	Fqn() *string
	HttpParameters() DataAwsccPipesPipeTargetParametersHttpParametersOutputReference
	InputTemplate() *string
	InternalValue() *DataAwsccPipesPipeTargetParameters
	SetInternalValue(val *DataAwsccPipesPipeTargetParameters)
	KinesisStreamParameters() DataAwsccPipesPipeTargetParametersKinesisStreamParametersOutputReference
	LambdaFunctionParameters() DataAwsccPipesPipeTargetParametersLambdaFunctionParametersOutputReference
	RedshiftDataParameters() DataAwsccPipesPipeTargetParametersRedshiftDataParametersOutputReference
	SageMakerPipelineParameters() DataAwsccPipesPipeTargetParametersSageMakerPipelineParametersOutputReference
	SqsQueueParameters() DataAwsccPipesPipeTargetParametersSqsQueueParametersOutputReference
	StepFunctionStateMachineParameters() DataAwsccPipesPipeTargetParametersStepFunctionStateMachineParametersOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimestreamParameters() DataAwsccPipesPipeTargetParametersTimestreamParametersOutputReference
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

// The jsii proxy struct for DataAwsccPipesPipeTargetParametersOutputReference
type jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) BatchJobParameters() DataAwsccPipesPipeTargetParametersBatchJobParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersBatchJobParametersOutputReference
	_jsii_.Get(
		j,
		"batchJobParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) CloudwatchLogsParameters() DataAwsccPipesPipeTargetParametersCloudwatchLogsParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersCloudwatchLogsParametersOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) EcsTaskParameters() DataAwsccPipesPipeTargetParametersEcsTaskParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersEcsTaskParametersOutputReference
	_jsii_.Get(
		j,
		"ecsTaskParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) EventBridgeEventBusParameters() DataAwsccPipesPipeTargetParametersEventBridgeEventBusParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersEventBridgeEventBusParametersOutputReference
	_jsii_.Get(
		j,
		"eventBridgeEventBusParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) HttpParameters() DataAwsccPipesPipeTargetParametersHttpParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersHttpParametersOutputReference
	_jsii_.Get(
		j,
		"httpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) InputTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) InternalValue() *DataAwsccPipesPipeTargetParameters {
	var returns *DataAwsccPipesPipeTargetParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) KinesisStreamParameters() DataAwsccPipesPipeTargetParametersKinesisStreamParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersKinesisStreamParametersOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) LambdaFunctionParameters() DataAwsccPipesPipeTargetParametersLambdaFunctionParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersLambdaFunctionParametersOutputReference
	_jsii_.Get(
		j,
		"lambdaFunctionParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) RedshiftDataParameters() DataAwsccPipesPipeTargetParametersRedshiftDataParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersRedshiftDataParametersOutputReference
	_jsii_.Get(
		j,
		"redshiftDataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) SageMakerPipelineParameters() DataAwsccPipesPipeTargetParametersSageMakerPipelineParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersSageMakerPipelineParametersOutputReference
	_jsii_.Get(
		j,
		"sageMakerPipelineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) SqsQueueParameters() DataAwsccPipesPipeTargetParametersSqsQueueParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersSqsQueueParametersOutputReference
	_jsii_.Get(
		j,
		"sqsQueueParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) StepFunctionStateMachineParameters() DataAwsccPipesPipeTargetParametersStepFunctionStateMachineParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersStepFunctionStateMachineParametersOutputReference
	_jsii_.Get(
		j,
		"stepFunctionStateMachineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) TimestreamParameters() DataAwsccPipesPipeTargetParametersTimestreamParametersOutputReference {
	var returns DataAwsccPipesPipeTargetParametersTimestreamParametersOutputReference
	_jsii_.Get(
		j,
		"timestreamParameters",
		&returns,
	)
	return returns
}


func NewDataAwsccPipesPipeTargetParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccPipesPipeTargetParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccPipesPipeTargetParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccPipesPipe.DataAwsccPipesPipeTargetParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccPipesPipeTargetParametersOutputReference_Override(d DataAwsccPipesPipeTargetParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccPipesPipe.DataAwsccPipesPipeTargetParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference)SetInternalValue(val *DataAwsccPipesPipeTargetParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccPipesPipeTargetParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

