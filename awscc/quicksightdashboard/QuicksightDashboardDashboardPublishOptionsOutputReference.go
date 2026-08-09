// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdashboard/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDashboardDashboardPublishOptionsOutputReference interface {
	cdktn.ComplexObject
	AdHocFilteringOption() QuicksightDashboardDashboardPublishOptionsAdHocFilteringOptionOutputReference
	AdHocFilteringOptionInput() interface{}
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
	ExportToCsvOption() QuicksightDashboardDashboardPublishOptionsExportToCsvOptionOutputReference
	ExportToCsvOptionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SheetControlsOption() QuicksightDashboardDashboardPublishOptionsSheetControlsOptionOutputReference
	SheetControlsOptionInput() interface{}
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
	PutAdHocFilteringOption(value *QuicksightDashboardDashboardPublishOptionsAdHocFilteringOption)
	PutExportToCsvOption(value *QuicksightDashboardDashboardPublishOptionsExportToCsvOption)
	PutSheetControlsOption(value *QuicksightDashboardDashboardPublishOptionsSheetControlsOption)
	ResetAdHocFilteringOption()
	ResetExportToCsvOption()
	ResetSheetControlsOption()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDashboardDashboardPublishOptionsOutputReference
type jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) AdHocFilteringOption() QuicksightDashboardDashboardPublishOptionsAdHocFilteringOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsAdHocFilteringOptionOutputReference
	_jsii_.Get(
		j,
		"adHocFilteringOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) AdHocFilteringOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adHocFilteringOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ExportToCsvOption() QuicksightDashboardDashboardPublishOptionsExportToCsvOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsExportToCsvOptionOutputReference
	_jsii_.Get(
		j,
		"exportToCsvOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ExportToCsvOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportToCsvOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) SheetControlsOption() QuicksightDashboardDashboardPublishOptionsSheetControlsOptionOutputReference {
	var returns QuicksightDashboardDashboardPublishOptionsSheetControlsOptionOutputReference
	_jsii_.Get(
		j,
		"sheetControlsOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) SheetControlsOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sheetControlsOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightDashboardDashboardPublishOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightDashboardDashboardPublishOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDashboardDashboardPublishOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDashboard.QuicksightDashboardDashboardPublishOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDashboardDashboardPublishOptionsOutputReference_Override(q QuicksightDashboardDashboardPublishOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDashboard.QuicksightDashboardDashboardPublishOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutAdHocFilteringOption(value *QuicksightDashboardDashboardPublishOptionsAdHocFilteringOption) {
	if err := q.validatePutAdHocFilteringOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAdHocFilteringOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutExportToCsvOption(value *QuicksightDashboardDashboardPublishOptionsExportToCsvOption) {
	if err := q.validatePutExportToCsvOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putExportToCsvOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) PutSheetControlsOption(value *QuicksightDashboardDashboardPublishOptionsSheetControlsOption) {
	if err := q.validatePutSheetControlsOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putSheetControlsOption",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetAdHocFilteringOption() {
	_jsii_.InvokeVoid(
		q,
		"resetAdHocFilteringOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetExportToCsvOption() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToCsvOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ResetSheetControlsOption() {
	_jsii_.InvokeVoid(
		q,
		"resetSheetControlsOption",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDashboardDashboardPublishOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

