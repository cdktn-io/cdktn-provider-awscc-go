// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermonitoringschedule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakermonitoringschedule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference interface {
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
	CreationTime() *string
	SetCreationTime(val *string)
	CreationTimeInput() *string
	EndpointName() *string
	SetEndpointName(val *string)
	EndpointNameInput() *string
	FailureReason() *string
	SetFailureReason(val *string)
	FailureReasonInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LastModifiedTime() *string
	SetLastModifiedTime(val *string)
	LastModifiedTimeInput() *string
	MonitoringExecutionStatus() *string
	SetMonitoringExecutionStatus(val *string)
	MonitoringExecutionStatusInput() *string
	MonitoringScheduleName() *string
	SetMonitoringScheduleName(val *string)
	MonitoringScheduleNameInput() *string
	ProcessingJobArn() *string
	SetProcessingJobArn(val *string)
	ProcessingJobArnInput() *string
	ScheduledTime() *string
	SetScheduledTime(val *string)
	ScheduledTimeInput() *string
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
	ResetCreationTime()
	ResetEndpointName()
	ResetFailureReason()
	ResetLastModifiedTime()
	ResetMonitoringExecutionStatus()
	ResetMonitoringScheduleName()
	ResetProcessingJobArn()
	ResetScheduledTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference
type jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) CreationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) EndpointNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) FailureReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) LastModifiedTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) MonitoringExecutionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringExecutionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) MonitoringExecutionStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringExecutionStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) MonitoringScheduleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringScheduleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) MonitoringScheduleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringScheduleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ProcessingJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ProcessingJobArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processingJobArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ScheduledTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ScheduledTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerMonitoringSchedule.SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference_Override(s SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerMonitoringSchedule.SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetCreationTime(val *string) {
	if err := j.validateSetCreationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationTime",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetEndpointName(val *string) {
	if err := j.validateSetEndpointNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetFailureReason(val *string) {
	if err := j.validateSetFailureReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureReason",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetLastModifiedTime(val *string) {
	if err := j.validateSetLastModifiedTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastModifiedTime",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetMonitoringExecutionStatus(val *string) {
	if err := j.validateSetMonitoringExecutionStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringExecutionStatus",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetMonitoringScheduleName(val *string) {
	if err := j.validateSetMonitoringScheduleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringScheduleName",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetProcessingJobArn(val *string) {
	if err := j.validateSetProcessingJobArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processingJobArn",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetScheduledTime(val *string) {
	if err := j.validateSetScheduledTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledTime",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ResetCreationTime() {
	_jsii_.InvokeVoid(
		s,
		"resetCreationTime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ResetEndpointName() {
	_jsii_.InvokeVoid(
		s,
		"resetEndpointName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ResetFailureReason() {
	_jsii_.InvokeVoid(
		s,
		"resetFailureReason",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ResetLastModifiedTime() {
	_jsii_.InvokeVoid(
		s,
		"resetLastModifiedTime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ResetMonitoringExecutionStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetMonitoringExecutionStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ResetMonitoringScheduleName() {
	_jsii_.InvokeVoid(
		s,
		"resetMonitoringScheduleName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ResetProcessingJobArn() {
	_jsii_.InvokeVoid(
		s,
		"resetProcessingJobArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ResetScheduledTime() {
	_jsii_.InvokeVoid(
		s,
		"resetScheduledTime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerMonitoringScheduleLastMonitoringExecutionSummaryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

