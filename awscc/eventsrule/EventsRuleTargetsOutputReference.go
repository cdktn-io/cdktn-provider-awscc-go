// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/eventsrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EventsRuleTargetsOutputReference interface {
	cdktn.ComplexObject
	AppSyncParameters() EventsRuleTargetsAppSyncParametersOutputReference
	AppSyncParametersInput() interface{}
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
	BatchParameters() EventsRuleTargetsBatchParametersOutputReference
	BatchParametersInput() interface{}
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
	DeadLetterConfig() EventsRuleTargetsDeadLetterConfigOutputReference
	DeadLetterConfigInput() interface{}
	EcsParameters() EventsRuleTargetsEcsParametersOutputReference
	EcsParametersInput() interface{}
	// Experimental.
	Fqn() *string
	HttpParameters() EventsRuleTargetsHttpParametersOutputReference
	HttpParametersInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Input() *string
	SetInput(val *string)
	InputInput() *string
	InputPath() *string
	SetInputPath(val *string)
	InputPathInput() *string
	InputTransformer() EventsRuleTargetsInputTransformerOutputReference
	InputTransformerInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KinesisParameters() EventsRuleTargetsKinesisParametersOutputReference
	KinesisParametersInput() interface{}
	RedshiftDataParameters() EventsRuleTargetsRedshiftDataParametersOutputReference
	RedshiftDataParametersInput() interface{}
	RetryPolicy() EventsRuleTargetsRetryPolicyOutputReference
	RetryPolicyInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	RunCommandParameters() EventsRuleTargetsRunCommandParametersOutputReference
	RunCommandParametersInput() interface{}
	SageMakerPipelineParameters() EventsRuleTargetsSageMakerPipelineParametersOutputReference
	SageMakerPipelineParametersInput() interface{}
	SqsParameters() EventsRuleTargetsSqsParametersOutputReference
	SqsParametersInput() interface{}
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
	PutAppSyncParameters(value *EventsRuleTargetsAppSyncParameters)
	PutBatchParameters(value *EventsRuleTargetsBatchParameters)
	PutDeadLetterConfig(value *EventsRuleTargetsDeadLetterConfig)
	PutEcsParameters(value *EventsRuleTargetsEcsParameters)
	PutHttpParameters(value *EventsRuleTargetsHttpParameters)
	PutInputTransformer(value *EventsRuleTargetsInputTransformer)
	PutKinesisParameters(value *EventsRuleTargetsKinesisParameters)
	PutRedshiftDataParameters(value *EventsRuleTargetsRedshiftDataParameters)
	PutRetryPolicy(value *EventsRuleTargetsRetryPolicy)
	PutRunCommandParameters(value *EventsRuleTargetsRunCommandParameters)
	PutSageMakerPipelineParameters(value *EventsRuleTargetsSageMakerPipelineParameters)
	PutSqsParameters(value *EventsRuleTargetsSqsParameters)
	ResetAppSyncParameters()
	ResetArn()
	ResetBatchParameters()
	ResetDeadLetterConfig()
	ResetEcsParameters()
	ResetHttpParameters()
	ResetId()
	ResetInput()
	ResetInputPath()
	ResetInputTransformer()
	ResetKinesisParameters()
	ResetRedshiftDataParameters()
	ResetRetryPolicy()
	ResetRoleArn()
	ResetRunCommandParameters()
	ResetSageMakerPipelineParameters()
	ResetSqsParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventsRuleTargetsOutputReference
type jsiiProxy_EventsRuleTargetsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) AppSyncParameters() EventsRuleTargetsAppSyncParametersOutputReference {
	var returns EventsRuleTargetsAppSyncParametersOutputReference
	_jsii_.Get(
		j,
		"appSyncParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) AppSyncParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appSyncParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) BatchParameters() EventsRuleTargetsBatchParametersOutputReference {
	var returns EventsRuleTargetsBatchParametersOutputReference
	_jsii_.Get(
		j,
		"batchParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) BatchParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) DeadLetterConfig() EventsRuleTargetsDeadLetterConfigOutputReference {
	var returns EventsRuleTargetsDeadLetterConfigOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) DeadLetterConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) EcsParameters() EventsRuleTargetsEcsParametersOutputReference {
	var returns EventsRuleTargetsEcsParametersOutputReference
	_jsii_.Get(
		j,
		"ecsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) EcsParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) HttpParameters() EventsRuleTargetsHttpParametersOutputReference {
	var returns EventsRuleTargetsHttpParametersOutputReference
	_jsii_.Get(
		j,
		"httpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) HttpParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) InputPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) InputTransformer() EventsRuleTargetsInputTransformerOutputReference {
	var returns EventsRuleTargetsInputTransformerOutputReference
	_jsii_.Get(
		j,
		"inputTransformer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) InputTransformerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputTransformerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) KinesisParameters() EventsRuleTargetsKinesisParametersOutputReference {
	var returns EventsRuleTargetsKinesisParametersOutputReference
	_jsii_.Get(
		j,
		"kinesisParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) KinesisParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) RedshiftDataParameters() EventsRuleTargetsRedshiftDataParametersOutputReference {
	var returns EventsRuleTargetsRedshiftDataParametersOutputReference
	_jsii_.Get(
		j,
		"redshiftDataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) RedshiftDataParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftDataParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) RetryPolicy() EventsRuleTargetsRetryPolicyOutputReference {
	var returns EventsRuleTargetsRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) RetryPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) RunCommandParameters() EventsRuleTargetsRunCommandParametersOutputReference {
	var returns EventsRuleTargetsRunCommandParametersOutputReference
	_jsii_.Get(
		j,
		"runCommandParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) RunCommandParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runCommandParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) SageMakerPipelineParameters() EventsRuleTargetsSageMakerPipelineParametersOutputReference {
	var returns EventsRuleTargetsSageMakerPipelineParametersOutputReference
	_jsii_.Get(
		j,
		"sageMakerPipelineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) SageMakerPipelineParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sageMakerPipelineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) SqsParameters() EventsRuleTargetsSqsParametersOutputReference {
	var returns EventsRuleTargetsSqsParametersOutputReference
	_jsii_.Get(
		j,
		"sqsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) SqsParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventsRuleTargetsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventsRuleTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewEventsRuleTargetsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventsRuleTargetsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsRule.EventsRuleTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventsRuleTargetsOutputReference_Override(e EventsRuleTargetsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsRule.EventsRuleTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference)SetInputPath(val *string) {
	if err := j.validateSetInputPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputPath",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutAppSyncParameters(value *EventsRuleTargetsAppSyncParameters) {
	if err := e.validatePutAppSyncParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAppSyncParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutBatchParameters(value *EventsRuleTargetsBatchParameters) {
	if err := e.validatePutBatchParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBatchParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutDeadLetterConfig(value *EventsRuleTargetsDeadLetterConfig) {
	if err := e.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutEcsParameters(value *EventsRuleTargetsEcsParameters) {
	if err := e.validatePutEcsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEcsParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutHttpParameters(value *EventsRuleTargetsHttpParameters) {
	if err := e.validatePutHttpParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHttpParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutInputTransformer(value *EventsRuleTargetsInputTransformer) {
	if err := e.validatePutInputTransformerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInputTransformer",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutKinesisParameters(value *EventsRuleTargetsKinesisParameters) {
	if err := e.validatePutKinesisParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putKinesisParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutRedshiftDataParameters(value *EventsRuleTargetsRedshiftDataParameters) {
	if err := e.validatePutRedshiftDataParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRedshiftDataParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutRetryPolicy(value *EventsRuleTargetsRetryPolicy) {
	if err := e.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutRunCommandParameters(value *EventsRuleTargetsRunCommandParameters) {
	if err := e.validatePutRunCommandParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRunCommandParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutSageMakerPipelineParameters(value *EventsRuleTargetsSageMakerPipelineParameters) {
	if err := e.validatePutSageMakerPipelineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSageMakerPipelineParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) PutSqsParameters(value *EventsRuleTargetsSqsParameters) {
	if err := e.validatePutSqsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSqsParameters",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetAppSyncParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetAppSyncParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetArn() {
	_jsii_.InvokeVoid(
		e,
		"resetArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetBatchParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetBatchParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetEcsParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetEcsParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetHttpParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		e,
		"resetInput",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetInputPath() {
	_jsii_.InvokeVoid(
		e,
		"resetInputPath",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetInputTransformer() {
	_jsii_.InvokeVoid(
		e,
		"resetInputTransformer",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetKinesisParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetKinesisParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetRedshiftDataParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetRedshiftDataParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		e,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetRunCommandParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetRunCommandParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetSageMakerPipelineParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetSageMakerPipelineParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ResetSqsParameters() {
	_jsii_.InvokeVoid(
		e,
		"resetSqsParameters",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

