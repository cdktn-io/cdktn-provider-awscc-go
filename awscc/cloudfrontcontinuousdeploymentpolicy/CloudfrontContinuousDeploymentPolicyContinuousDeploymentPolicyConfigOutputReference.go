// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontcontinuousdeploymentpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/cloudfrontcontinuousdeploymentpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SingleHeaderPolicyConfig() CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleHeaderPolicyConfigOutputReference
	SingleHeaderPolicyConfigInput() interface{}
	SingleWeightPolicyConfig() CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleWeightPolicyConfigOutputReference
	SingleWeightPolicyConfigInput() interface{}
	StagingDistributionDnsNames() *[]*string
	SetStagingDistributionDnsNames(val *[]*string)
	StagingDistributionDnsNamesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrafficConfig() CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigOutputReference
	TrafficConfigInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutSingleHeaderPolicyConfig(value *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleHeaderPolicyConfig)
	PutSingleWeightPolicyConfig(value *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleWeightPolicyConfig)
	PutTrafficConfig(value *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfig)
	ResetSingleHeaderPolicyConfig()
	ResetSingleWeightPolicyConfig()
	ResetTrafficConfig()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference
type jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) SingleHeaderPolicyConfig() CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleHeaderPolicyConfigOutputReference {
	var returns CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleHeaderPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"singleHeaderPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) SingleHeaderPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleHeaderPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) SingleWeightPolicyConfig() CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleWeightPolicyConfigOutputReference {
	var returns CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleWeightPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"singleWeightPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) SingleWeightPolicyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleWeightPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) StagingDistributionDnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stagingDistributionDnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) StagingDistributionDnsNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stagingDistributionDnsNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) TrafficConfig() CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigOutputReference {
	var returns CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfigOutputReference
	_jsii_.Get(
		j,
		"trafficConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) TrafficConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudfrontContinuousDeploymentPolicy.CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference_Override(c CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudfrontContinuousDeploymentPolicy.CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetStagingDistributionDnsNames(val *[]*string) {
	if err := j.validateSetStagingDistributionDnsNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingDistributionDnsNames",
		val,
	)
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) PutSingleHeaderPolicyConfig(value *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleHeaderPolicyConfig) {
	if err := c.validatePutSingleHeaderPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSingleHeaderPolicyConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) PutSingleWeightPolicyConfig(value *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigSingleWeightPolicyConfig) {
	if err := c.validatePutSingleWeightPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSingleWeightPolicyConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) PutTrafficConfig(value *CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigTrafficConfig) {
	if err := c.validatePutTrafficConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTrafficConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ResetSingleHeaderPolicyConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSingleHeaderPolicyConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ResetSingleWeightPolicyConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSingleWeightPolicyConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ResetTrafficConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetTrafficConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontContinuousDeploymentPolicyContinuousDeploymentPolicyConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

