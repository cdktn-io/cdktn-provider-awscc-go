// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilityadmintelemetryrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/observabilityadmintelemetryrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference interface {
	cdktn.ComplexObject
	CloudtrailParameters() ObservabilityadminTelemetryRuleRuleDestinationConfigurationCloudtrailParametersOutputReference
	CloudtrailParametersInput() interface{}
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
	DestinationPattern() *string
	SetDestinationPattern(val *string)
	DestinationPatternInput() *string
	DestinationType() *string
	SetDestinationType(val *string)
	DestinationTypeInput() *string
	ElbLoadBalancerLoggingParameters() ObservabilityadminTelemetryRuleRuleDestinationConfigurationElbLoadBalancerLoggingParametersOutputReference
	ElbLoadBalancerLoggingParametersInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	LogDeliveryParameters() ObservabilityadminTelemetryRuleRuleDestinationConfigurationLogDeliveryParametersOutputReference
	LogDeliveryParametersInput() interface{}
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VpcFlowLogParameters() ObservabilityadminTelemetryRuleRuleDestinationConfigurationVpcFlowLogParametersOutputReference
	VpcFlowLogParametersInput() interface{}
	WafLoggingParameters() ObservabilityadminTelemetryRuleRuleDestinationConfigurationWafLoggingParametersOutputReference
	WafLoggingParametersInput() interface{}
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
	PutCloudtrailParameters(value *ObservabilityadminTelemetryRuleRuleDestinationConfigurationCloudtrailParameters)
	PutElbLoadBalancerLoggingParameters(value *ObservabilityadminTelemetryRuleRuleDestinationConfigurationElbLoadBalancerLoggingParameters)
	PutLogDeliveryParameters(value *ObservabilityadminTelemetryRuleRuleDestinationConfigurationLogDeliveryParameters)
	PutVpcFlowLogParameters(value *ObservabilityadminTelemetryRuleRuleDestinationConfigurationVpcFlowLogParameters)
	PutWafLoggingParameters(value *ObservabilityadminTelemetryRuleRuleDestinationConfigurationWafLoggingParameters)
	ResetCloudtrailParameters()
	ResetDestinationPattern()
	ResetDestinationType()
	ResetElbLoadBalancerLoggingParameters()
	ResetKmsKeyArn()
	ResetLogDeliveryParameters()
	ResetRetentionInDays()
	ResetVpcFlowLogParameters()
	ResetWafLoggingParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference
type jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) CloudtrailParameters() ObservabilityadminTelemetryRuleRuleDestinationConfigurationCloudtrailParametersOutputReference {
	var returns ObservabilityadminTelemetryRuleRuleDestinationConfigurationCloudtrailParametersOutputReference
	_jsii_.Get(
		j,
		"cloudtrailParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) CloudtrailParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudtrailParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) DestinationPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) DestinationPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) DestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) DestinationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ElbLoadBalancerLoggingParameters() ObservabilityadminTelemetryRuleRuleDestinationConfigurationElbLoadBalancerLoggingParametersOutputReference {
	var returns ObservabilityadminTelemetryRuleRuleDestinationConfigurationElbLoadBalancerLoggingParametersOutputReference
	_jsii_.Get(
		j,
		"elbLoadBalancerLoggingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ElbLoadBalancerLoggingParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elbLoadBalancerLoggingParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) LogDeliveryParameters() ObservabilityadminTelemetryRuleRuleDestinationConfigurationLogDeliveryParametersOutputReference {
	var returns ObservabilityadminTelemetryRuleRuleDestinationConfigurationLogDeliveryParametersOutputReference
	_jsii_.Get(
		j,
		"logDeliveryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) LogDeliveryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeliveryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) VpcFlowLogParameters() ObservabilityadminTelemetryRuleRuleDestinationConfigurationVpcFlowLogParametersOutputReference {
	var returns ObservabilityadminTelemetryRuleRuleDestinationConfigurationVpcFlowLogParametersOutputReference
	_jsii_.Get(
		j,
		"vpcFlowLogParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) VpcFlowLogParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcFlowLogParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) WafLoggingParameters() ObservabilityadminTelemetryRuleRuleDestinationConfigurationWafLoggingParametersOutputReference {
	var returns ObservabilityadminTelemetryRuleRuleDestinationConfigurationWafLoggingParametersOutputReference
	_jsii_.Get(
		j,
		"wafLoggingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) WafLoggingParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wafLoggingParametersInput",
		&returns,
	)
	return returns
}


func NewObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.observabilityadminTelemetryRule.ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference_Override(o ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.observabilityadminTelemetryRule.ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference)SetDestinationPattern(val *string) {
	if err := j.validateSetDestinationPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPattern",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference)SetDestinationType(val *string) {
	if err := j.validateSetDestinationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationType",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference)SetRetentionInDays(val *float64) {
	if err := j.validateSetRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) PutCloudtrailParameters(value *ObservabilityadminTelemetryRuleRuleDestinationConfigurationCloudtrailParameters) {
	if err := o.validatePutCloudtrailParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCloudtrailParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) PutElbLoadBalancerLoggingParameters(value *ObservabilityadminTelemetryRuleRuleDestinationConfigurationElbLoadBalancerLoggingParameters) {
	if err := o.validatePutElbLoadBalancerLoggingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putElbLoadBalancerLoggingParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) PutLogDeliveryParameters(value *ObservabilityadminTelemetryRuleRuleDestinationConfigurationLogDeliveryParameters) {
	if err := o.validatePutLogDeliveryParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogDeliveryParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) PutVpcFlowLogParameters(value *ObservabilityadminTelemetryRuleRuleDestinationConfigurationVpcFlowLogParameters) {
	if err := o.validatePutVpcFlowLogParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putVpcFlowLogParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) PutWafLoggingParameters(value *ObservabilityadminTelemetryRuleRuleDestinationConfigurationWafLoggingParameters) {
	if err := o.validatePutWafLoggingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWafLoggingParameters",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ResetCloudtrailParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudtrailParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ResetDestinationPattern() {
	_jsii_.InvokeVoid(
		o,
		"resetDestinationPattern",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ResetDestinationType() {
	_jsii_.InvokeVoid(
		o,
		"resetDestinationType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ResetElbLoadBalancerLoggingParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetElbLoadBalancerLoggingParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		o,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ResetLogDeliveryParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetLogDeliveryParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ResetRetentionInDays() {
	_jsii_.InvokeVoid(
		o,
		"resetRetentionInDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ResetVpcFlowLogParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetVpcFlowLogParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ResetWafLoggingParameters() {
	_jsii_.InvokeVoid(
		o,
		"resetWafLoggingParameters",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityadminTelemetryRuleRuleDestinationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

