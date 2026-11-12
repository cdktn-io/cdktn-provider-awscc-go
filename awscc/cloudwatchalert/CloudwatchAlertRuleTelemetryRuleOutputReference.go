// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/cloudwatchalert/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudwatchAlertRuleTelemetryRuleOutputReference interface {
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
	Condition() CloudwatchAlertRuleTelemetryRuleConditionOutputReference
	ConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Evaluation() CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference
	EvaluationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoData() CloudwatchAlertRuleTelemetryRuleNoDataOutputReference
	NoDataInput() interface{}
	Query() CloudwatchAlertRuleTelemetryRuleQueryOutputReference
	QueryInput() interface{}
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
	PutCondition(value *CloudwatchAlertRuleTelemetryRuleCondition)
	PutEvaluation(value *CloudwatchAlertRuleTelemetryRuleEvaluation)
	PutNoData(value *CloudwatchAlertRuleTelemetryRuleNoData)
	PutQuery(value *CloudwatchAlertRuleTelemetryRuleQuery)
	ResetNoData()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudwatchAlertRuleTelemetryRuleOutputReference
type jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) Condition() CloudwatchAlertRuleTelemetryRuleConditionOutputReference {
	var returns CloudwatchAlertRuleTelemetryRuleConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) ConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) Evaluation() CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference {
	var returns CloudwatchAlertRuleTelemetryRuleEvaluationOutputReference
	_jsii_.Get(
		j,
		"evaluation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) EvaluationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) NoData() CloudwatchAlertRuleTelemetryRuleNoDataOutputReference {
	var returns CloudwatchAlertRuleTelemetryRuleNoDataOutputReference
	_jsii_.Get(
		j,
		"noData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) NoDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) Query() CloudwatchAlertRuleTelemetryRuleQueryOutputReference {
	var returns CloudwatchAlertRuleTelemetryRuleQueryOutputReference
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) QueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudwatchAlertRuleTelemetryRuleOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) CloudwatchAlertRuleTelemetryRuleOutputReference {
	_init_.Initialize()

	if err := validateNewCloudwatchAlertRuleTelemetryRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudwatchAlert.CloudwatchAlertRuleTelemetryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudwatchAlertRuleTelemetryRuleOutputReference_Override(c CloudwatchAlertRuleTelemetryRuleOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.cloudwatchAlert.CloudwatchAlertRuleTelemetryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) PutCondition(value *CloudwatchAlertRuleTelemetryRuleCondition) {
	if err := c.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) PutEvaluation(value *CloudwatchAlertRuleTelemetryRuleEvaluation) {
	if err := c.validatePutEvaluationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEvaluation",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) PutNoData(value *CloudwatchAlertRuleTelemetryRuleNoData) {
	if err := c.validatePutNoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNoData",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) PutQuery(value *CloudwatchAlertRuleTelemetryRuleQuery) {
	if err := c.validatePutQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQuery",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) ResetNoData() {
	_jsii_.InvokeVoid(
		c,
		"resetNoData",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudwatchAlertRuleTelemetryRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

