// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccsesmailmanagerruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccsesmailmanagerruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccSesMailManagerRuleSetRulesActionsOutputReference interface {
	cdktn.ComplexObject
	AddHeader() DataAwsccSesMailManagerRuleSetRulesActionsAddHeaderOutputReference
	Archive() DataAwsccSesMailManagerRuleSetRulesActionsArchiveOutputReference
	Bounce() DataAwsccSesMailManagerRuleSetRulesActionsBounceOutputReference
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
	DeliverToMailbox() DataAwsccSesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference
	DeliverToQBusiness() DataAwsccSesMailManagerRuleSetRulesActionsDeliverToQBusinessOutputReference
	Drop() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccSesMailManagerRuleSetRulesActions
	SetInternalValue(val *DataAwsccSesMailManagerRuleSetRulesActions)
	InvokeLambda() DataAwsccSesMailManagerRuleSetRulesActionsInvokeLambdaOutputReference
	PublishToSns() DataAwsccSesMailManagerRuleSetRulesActionsPublishToSnsOutputReference
	Relay() DataAwsccSesMailManagerRuleSetRulesActionsRelayOutputReference
	ReplaceRecipient() DataAwsccSesMailManagerRuleSetRulesActionsReplaceRecipientOutputReference
	Send() DataAwsccSesMailManagerRuleSetRulesActionsSendOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WriteToS3() DataAwsccSesMailManagerRuleSetRulesActionsWriteToS3OutputReference
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

// The jsii proxy struct for DataAwsccSesMailManagerRuleSetRulesActionsOutputReference
type jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) AddHeader() DataAwsccSesMailManagerRuleSetRulesActionsAddHeaderOutputReference {
	var returns DataAwsccSesMailManagerRuleSetRulesActionsAddHeaderOutputReference
	_jsii_.Get(
		j,
		"addHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) Archive() DataAwsccSesMailManagerRuleSetRulesActionsArchiveOutputReference {
	var returns DataAwsccSesMailManagerRuleSetRulesActionsArchiveOutputReference
	_jsii_.Get(
		j,
		"archive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) Bounce() DataAwsccSesMailManagerRuleSetRulesActionsBounceOutputReference {
	var returns DataAwsccSesMailManagerRuleSetRulesActionsBounceOutputReference
	_jsii_.Get(
		j,
		"bounce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) DeliverToMailbox() DataAwsccSesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference {
	var returns DataAwsccSesMailManagerRuleSetRulesActionsDeliverToMailboxOutputReference
	_jsii_.Get(
		j,
		"deliverToMailbox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) DeliverToQBusiness() DataAwsccSesMailManagerRuleSetRulesActionsDeliverToQBusinessOutputReference {
	var returns DataAwsccSesMailManagerRuleSetRulesActionsDeliverToQBusinessOutputReference
	_jsii_.Get(
		j,
		"deliverToQBusiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) Drop() *string {
	var returns *string
	_jsii_.Get(
		j,
		"drop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) InternalValue() *DataAwsccSesMailManagerRuleSetRulesActions {
	var returns *DataAwsccSesMailManagerRuleSetRulesActions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) InvokeLambda() DataAwsccSesMailManagerRuleSetRulesActionsInvokeLambdaOutputReference {
	var returns DataAwsccSesMailManagerRuleSetRulesActionsInvokeLambdaOutputReference
	_jsii_.Get(
		j,
		"invokeLambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) PublishToSns() DataAwsccSesMailManagerRuleSetRulesActionsPublishToSnsOutputReference {
	var returns DataAwsccSesMailManagerRuleSetRulesActionsPublishToSnsOutputReference
	_jsii_.Get(
		j,
		"publishToSns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) Relay() DataAwsccSesMailManagerRuleSetRulesActionsRelayOutputReference {
	var returns DataAwsccSesMailManagerRuleSetRulesActionsRelayOutputReference
	_jsii_.Get(
		j,
		"relay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) ReplaceRecipient() DataAwsccSesMailManagerRuleSetRulesActionsReplaceRecipientOutputReference {
	var returns DataAwsccSesMailManagerRuleSetRulesActionsReplaceRecipientOutputReference
	_jsii_.Get(
		j,
		"replaceRecipient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) Send() DataAwsccSesMailManagerRuleSetRulesActionsSendOutputReference {
	var returns DataAwsccSesMailManagerRuleSetRulesActionsSendOutputReference
	_jsii_.Get(
		j,
		"send",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) WriteToS3() DataAwsccSesMailManagerRuleSetRulesActionsWriteToS3OutputReference {
	var returns DataAwsccSesMailManagerRuleSetRulesActionsWriteToS3OutputReference
	_jsii_.Get(
		j,
		"writeToS3",
		&returns,
	)
	return returns
}


func NewDataAwsccSesMailManagerRuleSetRulesActionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccSesMailManagerRuleSetRulesActionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccSesMailManagerRuleSetRulesActionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSesMailManagerRuleSet.DataAwsccSesMailManagerRuleSetRulesActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccSesMailManagerRuleSetRulesActionsOutputReference_Override(d DataAwsccSesMailManagerRuleSetRulesActionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSesMailManagerRuleSet.DataAwsccSesMailManagerRuleSetRulesActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference)SetInternalValue(val *DataAwsccSesMailManagerRuleSetRulesActions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccSesMailManagerRuleSetRulesActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

