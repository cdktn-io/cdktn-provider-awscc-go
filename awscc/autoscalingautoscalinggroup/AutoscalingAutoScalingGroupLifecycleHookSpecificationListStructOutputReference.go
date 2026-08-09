// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/autoscalingautoscalinggroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference interface {
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
	DefaultResult() *string
	SetDefaultResult(val *string)
	DefaultResultInput() *string
	// Experimental.
	Fqn() *string
	HeartbeatTimeout() *float64
	SetHeartbeatTimeout(val *float64)
	HeartbeatTimeoutInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LifecycleHookName() *string
	SetLifecycleHookName(val *string)
	LifecycleHookNameInput() *string
	LifecycleTransition() *string
	SetLifecycleTransition(val *string)
	LifecycleTransitionInput() *string
	NotificationMetadata() *string
	SetNotificationMetadata(val *string)
	NotificationMetadataInput() *string
	NotificationTargetArn() *string
	SetNotificationTargetArn(val *string)
	NotificationTargetArnInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	ResetDefaultResult()
	ResetHeartbeatTimeout()
	ResetLifecycleHookName()
	ResetLifecycleTransition()
	ResetNotificationMetadata()
	ResetNotificationTargetArn()
	ResetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference
type jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) DefaultResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) DefaultResultInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultResultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) HeartbeatTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) HeartbeatTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heartbeatTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) LifecycleHookName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleHookName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) LifecycleHookNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleHookNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) LifecycleTransition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) LifecycleTransitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) NotificationMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) NotificationMetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) NotificationTargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTargetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) NotificationTargetArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTargetArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference_Override(a AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.autoscalingAutoScalingGroup.AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetDefaultResult(val *string) {
	if err := j.validateSetDefaultResultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultResult",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetHeartbeatTimeout(val *float64) {
	if err := j.validateSetHeartbeatTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"heartbeatTimeout",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetLifecycleHookName(val *string) {
	if err := j.validateSetLifecycleHookNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleHookName",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetLifecycleTransition(val *string) {
	if err := j.validateSetLifecycleTransitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleTransition",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetNotificationMetadata(val *string) {
	if err := j.validateSetNotificationMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationMetadata",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetNotificationTargetArn(val *string) {
	if err := j.validateSetNotificationTargetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationTargetArn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) ResetDefaultResult() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultResult",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) ResetHeartbeatTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetHeartbeatTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) ResetLifecycleHookName() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleHookName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) ResetLifecycleTransition() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleTransition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) ResetNotificationMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) ResetNotificationTargetArn() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationTargetArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingAutoScalingGroupLifecycleHookSpecificationListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

