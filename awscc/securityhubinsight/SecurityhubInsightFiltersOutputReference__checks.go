// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package securityhubinsight

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutAwsAccountIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersAwsAccountId:
		value := value.(*[]*SecurityhubInsightFiltersAwsAccountId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersAwsAccountId:
		value_ := value.([]*SecurityhubInsightFiltersAwsAccountId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersAwsAccountId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutAwsAccountNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersAwsAccountName:
		value := value.(*[]*SecurityhubInsightFiltersAwsAccountName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersAwsAccountName:
		value_ := value.([]*SecurityhubInsightFiltersAwsAccountName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersAwsAccountName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutCompanyNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersCompanyName:
		value := value.(*[]*SecurityhubInsightFiltersCompanyName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersCompanyName:
		value_ := value.([]*SecurityhubInsightFiltersCompanyName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersCompanyName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutComplianceAssociatedStandardsIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersComplianceAssociatedStandardsId:
		value := value.(*[]*SecurityhubInsightFiltersComplianceAssociatedStandardsId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersComplianceAssociatedStandardsId:
		value_ := value.([]*SecurityhubInsightFiltersComplianceAssociatedStandardsId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersComplianceAssociatedStandardsId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutComplianceSecurityControlIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersComplianceSecurityControlId:
		value := value.(*[]*SecurityhubInsightFiltersComplianceSecurityControlId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersComplianceSecurityControlId:
		value_ := value.([]*SecurityhubInsightFiltersComplianceSecurityControlId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersComplianceSecurityControlId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutComplianceSecurityControlParametersNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersComplianceSecurityControlParametersName:
		value := value.(*[]*SecurityhubInsightFiltersComplianceSecurityControlParametersName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersComplianceSecurityControlParametersName:
		value_ := value.([]*SecurityhubInsightFiltersComplianceSecurityControlParametersName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersComplianceSecurityControlParametersName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutComplianceSecurityControlParametersValueParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersComplianceSecurityControlParametersValue:
		value := value.(*[]*SecurityhubInsightFiltersComplianceSecurityControlParametersValue)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersComplianceSecurityControlParametersValue:
		value_ := value.([]*SecurityhubInsightFiltersComplianceSecurityControlParametersValue)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersComplianceSecurityControlParametersValue; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutComplianceStatusParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersComplianceStatus:
		value := value.(*[]*SecurityhubInsightFiltersComplianceStatus)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersComplianceStatus:
		value_ := value.([]*SecurityhubInsightFiltersComplianceStatus)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersComplianceStatus; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutConfidenceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersConfidence:
		value := value.(*[]*SecurityhubInsightFiltersConfidence)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersConfidence:
		value_ := value.([]*SecurityhubInsightFiltersConfidence)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersConfidence; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutCreatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersCreatedAt:
		value := value.(*[]*SecurityhubInsightFiltersCreatedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersCreatedAt:
		value_ := value.([]*SecurityhubInsightFiltersCreatedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersCreatedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutCriticalityParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersCriticality:
		value := value.(*[]*SecurityhubInsightFiltersCriticality)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersCriticality:
		value_ := value.([]*SecurityhubInsightFiltersCriticality)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersCriticality; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutDescriptionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersDescription:
		value := value.(*[]*SecurityhubInsightFiltersDescription)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersDescription:
		value_ := value.([]*SecurityhubInsightFiltersDescription)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersDescription; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutFindingProviderFieldsConfidenceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence:
		value := value.(*[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersFindingProviderFieldsConfidence:
		value_ := value.([]*SecurityhubInsightFiltersFindingProviderFieldsConfidence)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersFindingProviderFieldsConfidence; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutFindingProviderFieldsCriticalityParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality:
		value := value.(*[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersFindingProviderFieldsCriticality:
		value_ := value.([]*SecurityhubInsightFiltersFindingProviderFieldsCriticality)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersFindingProviderFieldsCriticality; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutFindingProviderFieldsRelatedFindingsIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId:
		value := value.(*[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId:
		value_ := value.([]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutFindingProviderFieldsRelatedFindingsProductArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn:
		value := value.(*[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn:
		value_ := value.([]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutFindingProviderFieldsSeverityLabelParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel:
		value := value.(*[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel:
		value_ := value.([]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutFindingProviderFieldsSeverityOriginalParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal:
		value := value.(*[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal:
		value_ := value.([]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutFindingProviderFieldsTypesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersFindingProviderFieldsTypes:
		value := value.(*[]*SecurityhubInsightFiltersFindingProviderFieldsTypes)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersFindingProviderFieldsTypes:
		value_ := value.([]*SecurityhubInsightFiltersFindingProviderFieldsTypes)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersFindingProviderFieldsTypes; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutFirstObservedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersFirstObservedAt:
		value := value.(*[]*SecurityhubInsightFiltersFirstObservedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersFirstObservedAt:
		value_ := value.([]*SecurityhubInsightFiltersFirstObservedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersFirstObservedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutGeneratorIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersGeneratorId:
		value := value.(*[]*SecurityhubInsightFiltersGeneratorId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersGeneratorId:
		value_ := value.([]*SecurityhubInsightFiltersGeneratorId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersGeneratorId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersId:
		value := value.(*[]*SecurityhubInsightFiltersId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersId:
		value_ := value.([]*SecurityhubInsightFiltersId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutKeywordParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersKeyword:
		value := value.(*[]*SecurityhubInsightFiltersKeyword)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersKeyword:
		value_ := value.([]*SecurityhubInsightFiltersKeyword)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersKeyword; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutLastObservedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersLastObservedAt:
		value := value.(*[]*SecurityhubInsightFiltersLastObservedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersLastObservedAt:
		value_ := value.([]*SecurityhubInsightFiltersLastObservedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersLastObservedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutMalwareNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersMalwareName:
		value := value.(*[]*SecurityhubInsightFiltersMalwareName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersMalwareName:
		value_ := value.([]*SecurityhubInsightFiltersMalwareName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersMalwareName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutMalwarePathParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersMalwarePath:
		value := value.(*[]*SecurityhubInsightFiltersMalwarePath)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersMalwarePath:
		value_ := value.([]*SecurityhubInsightFiltersMalwarePath)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersMalwarePath; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutMalwareStateParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersMalwareState:
		value := value.(*[]*SecurityhubInsightFiltersMalwareState)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersMalwareState:
		value_ := value.([]*SecurityhubInsightFiltersMalwareState)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersMalwareState; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutMalwareTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersMalwareType:
		value := value.(*[]*SecurityhubInsightFiltersMalwareType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersMalwareType:
		value_ := value.([]*SecurityhubInsightFiltersMalwareType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersMalwareType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNetworkDestinationDomainParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNetworkDestinationDomain:
		value := value.(*[]*SecurityhubInsightFiltersNetworkDestinationDomain)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNetworkDestinationDomain:
		value_ := value.([]*SecurityhubInsightFiltersNetworkDestinationDomain)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNetworkDestinationDomain; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNetworkDestinationIpV4Parameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNetworkDestinationIpV4:
		value := value.(*[]*SecurityhubInsightFiltersNetworkDestinationIpV4)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNetworkDestinationIpV4:
		value_ := value.([]*SecurityhubInsightFiltersNetworkDestinationIpV4)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNetworkDestinationIpV4; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNetworkDestinationIpV6Parameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNetworkDestinationIpV6:
		value := value.(*[]*SecurityhubInsightFiltersNetworkDestinationIpV6)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNetworkDestinationIpV6:
		value_ := value.([]*SecurityhubInsightFiltersNetworkDestinationIpV6)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNetworkDestinationIpV6; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNetworkDestinationPortParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNetworkDestinationPort:
		value := value.(*[]*SecurityhubInsightFiltersNetworkDestinationPort)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNetworkDestinationPort:
		value_ := value.([]*SecurityhubInsightFiltersNetworkDestinationPort)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNetworkDestinationPort; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNetworkDirectionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNetworkDirection:
		value := value.(*[]*SecurityhubInsightFiltersNetworkDirection)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNetworkDirection:
		value_ := value.([]*SecurityhubInsightFiltersNetworkDirection)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNetworkDirection; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNetworkProtocolParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNetworkProtocol:
		value := value.(*[]*SecurityhubInsightFiltersNetworkProtocol)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNetworkProtocol:
		value_ := value.([]*SecurityhubInsightFiltersNetworkProtocol)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNetworkProtocol; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNetworkSourceDomainParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNetworkSourceDomain:
		value := value.(*[]*SecurityhubInsightFiltersNetworkSourceDomain)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNetworkSourceDomain:
		value_ := value.([]*SecurityhubInsightFiltersNetworkSourceDomain)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNetworkSourceDomain; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNetworkSourceIpV4Parameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNetworkSourceIpV4:
		value := value.(*[]*SecurityhubInsightFiltersNetworkSourceIpV4)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNetworkSourceIpV4:
		value_ := value.([]*SecurityhubInsightFiltersNetworkSourceIpV4)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNetworkSourceIpV4; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNetworkSourceIpV6Parameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNetworkSourceIpV6:
		value := value.(*[]*SecurityhubInsightFiltersNetworkSourceIpV6)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNetworkSourceIpV6:
		value_ := value.([]*SecurityhubInsightFiltersNetworkSourceIpV6)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNetworkSourceIpV6; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNetworkSourceMacParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNetworkSourceMac:
		value := value.(*[]*SecurityhubInsightFiltersNetworkSourceMac)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNetworkSourceMac:
		value_ := value.([]*SecurityhubInsightFiltersNetworkSourceMac)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNetworkSourceMac; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNetworkSourcePortParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNetworkSourcePort:
		value := value.(*[]*SecurityhubInsightFiltersNetworkSourcePort)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNetworkSourcePort:
		value_ := value.([]*SecurityhubInsightFiltersNetworkSourcePort)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNetworkSourcePort; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNoteTextParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNoteText:
		value := value.(*[]*SecurityhubInsightFiltersNoteText)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNoteText:
		value_ := value.([]*SecurityhubInsightFiltersNoteText)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNoteText; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNoteUpdatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNoteUpdatedAt:
		value := value.(*[]*SecurityhubInsightFiltersNoteUpdatedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNoteUpdatedAt:
		value_ := value.([]*SecurityhubInsightFiltersNoteUpdatedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNoteUpdatedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutNoteUpdatedByParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersNoteUpdatedBy:
		value := value.(*[]*SecurityhubInsightFiltersNoteUpdatedBy)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersNoteUpdatedBy:
		value_ := value.([]*SecurityhubInsightFiltersNoteUpdatedBy)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersNoteUpdatedBy; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutProcessLaunchedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersProcessLaunchedAt:
		value := value.(*[]*SecurityhubInsightFiltersProcessLaunchedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersProcessLaunchedAt:
		value_ := value.([]*SecurityhubInsightFiltersProcessLaunchedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersProcessLaunchedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutProcessNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersProcessName:
		value := value.(*[]*SecurityhubInsightFiltersProcessName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersProcessName:
		value_ := value.([]*SecurityhubInsightFiltersProcessName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersProcessName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutProcessParentPidParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersProcessParentPid:
		value := value.(*[]*SecurityhubInsightFiltersProcessParentPid)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersProcessParentPid:
		value_ := value.([]*SecurityhubInsightFiltersProcessParentPid)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersProcessParentPid; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutProcessPathParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersProcessPath:
		value := value.(*[]*SecurityhubInsightFiltersProcessPath)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersProcessPath:
		value_ := value.([]*SecurityhubInsightFiltersProcessPath)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersProcessPath; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutProcessPidParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersProcessPid:
		value := value.(*[]*SecurityhubInsightFiltersProcessPid)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersProcessPid:
		value_ := value.([]*SecurityhubInsightFiltersProcessPid)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersProcessPid; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutProcessTerminatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersProcessTerminatedAt:
		value := value.(*[]*SecurityhubInsightFiltersProcessTerminatedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersProcessTerminatedAt:
		value_ := value.([]*SecurityhubInsightFiltersProcessTerminatedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersProcessTerminatedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutProductArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersProductArn:
		value := value.(*[]*SecurityhubInsightFiltersProductArn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersProductArn:
		value_ := value.([]*SecurityhubInsightFiltersProductArn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersProductArn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutProductFieldsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersProductFields:
		value := value.(*[]*SecurityhubInsightFiltersProductFields)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersProductFields:
		value_ := value.([]*SecurityhubInsightFiltersProductFields)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersProductFields; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutProductNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersProductName:
		value := value.(*[]*SecurityhubInsightFiltersProductName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersProductName:
		value_ := value.([]*SecurityhubInsightFiltersProductName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersProductName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutRecommendationTextParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersRecommendationText:
		value := value.(*[]*SecurityhubInsightFiltersRecommendationText)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersRecommendationText:
		value_ := value.([]*SecurityhubInsightFiltersRecommendationText)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersRecommendationText; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutRecordStateParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersRecordState:
		value := value.(*[]*SecurityhubInsightFiltersRecordState)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersRecordState:
		value_ := value.([]*SecurityhubInsightFiltersRecordState)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersRecordState; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutRegionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersRegion:
		value := value.(*[]*SecurityhubInsightFiltersRegion)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersRegion:
		value_ := value.([]*SecurityhubInsightFiltersRegion)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersRegion; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutRelatedFindingsIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersRelatedFindingsId:
		value := value.(*[]*SecurityhubInsightFiltersRelatedFindingsId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersRelatedFindingsId:
		value_ := value.([]*SecurityhubInsightFiltersRelatedFindingsId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersRelatedFindingsId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutRelatedFindingsProductArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersRelatedFindingsProductArn:
		value := value.(*[]*SecurityhubInsightFiltersRelatedFindingsProductArn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersRelatedFindingsProductArn:
		value_ := value.([]*SecurityhubInsightFiltersRelatedFindingsProductArn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersRelatedFindingsProductArn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceApplicationArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceApplicationArn:
		value := value.(*[]*SecurityhubInsightFiltersResourceApplicationArn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceApplicationArn:
		value_ := value.([]*SecurityhubInsightFiltersResourceApplicationArn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceApplicationArn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceApplicationNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceApplicationName:
		value := value.(*[]*SecurityhubInsightFiltersResourceApplicationName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceApplicationName:
		value_ := value.([]*SecurityhubInsightFiltersResourceApplicationName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceApplicationName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsEc2InstanceIamInstanceProfileArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsEc2InstanceImageIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsEc2InstanceIpV4AddressesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpV4Addresses:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpV4Addresses)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsEc2InstanceIpV4Addresses:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpV4Addresses)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpV4Addresses; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsEc2InstanceIpV6AddressesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpV6Addresses:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpV6Addresses)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsEc2InstanceIpV6Addresses:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpV6Addresses)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceIpV6Addresses; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsEc2InstanceKeyNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsEc2InstanceLaunchedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsEc2InstanceSubnetIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsEc2InstanceTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsEc2InstanceType:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsEc2InstanceType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsEc2InstanceVpcIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsIamAccessKeyCreatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsIamAccessKeyPrincipalNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyPrincipalName:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyPrincipalName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsIamAccessKeyPrincipalName:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsIamAccessKeyPrincipalName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyPrincipalName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsIamAccessKeyStatusParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsIamAccessKeyUserNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsIamUserUserNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsIamUserUserName:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsIamUserUserName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsIamUserUserName:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsIamUserUserName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsIamUserUserName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsS3BucketOwnerIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceAwsS3BucketOwnerNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName:
		value := value.(*[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName:
		value_ := value.([]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceContainerImageIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceContainerImageId:
		value := value.(*[]*SecurityhubInsightFiltersResourceContainerImageId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceContainerImageId:
		value_ := value.([]*SecurityhubInsightFiltersResourceContainerImageId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceContainerImageId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceContainerImageNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceContainerImageName:
		value := value.(*[]*SecurityhubInsightFiltersResourceContainerImageName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceContainerImageName:
		value_ := value.([]*SecurityhubInsightFiltersResourceContainerImageName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceContainerImageName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceContainerLaunchedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceContainerLaunchedAt:
		value := value.(*[]*SecurityhubInsightFiltersResourceContainerLaunchedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceContainerLaunchedAt:
		value_ := value.([]*SecurityhubInsightFiltersResourceContainerLaunchedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceContainerLaunchedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceContainerNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceContainerName:
		value := value.(*[]*SecurityhubInsightFiltersResourceContainerName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceContainerName:
		value_ := value.([]*SecurityhubInsightFiltersResourceContainerName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceContainerName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceDetailsOtherParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceDetailsOther:
		value := value.(*[]*SecurityhubInsightFiltersResourceDetailsOther)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceDetailsOther:
		value_ := value.([]*SecurityhubInsightFiltersResourceDetailsOther)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceDetailsOther; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceId:
		value := value.(*[]*SecurityhubInsightFiltersResourceId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceId:
		value_ := value.([]*SecurityhubInsightFiltersResourceId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourcePartitionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourcePartition:
		value := value.(*[]*SecurityhubInsightFiltersResourcePartition)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourcePartition:
		value_ := value.([]*SecurityhubInsightFiltersResourcePartition)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourcePartition; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceRegionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceRegion:
		value := value.(*[]*SecurityhubInsightFiltersResourceRegion)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceRegion:
		value_ := value.([]*SecurityhubInsightFiltersResourceRegion)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceRegion; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceTagsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceTags:
		value := value.(*[]*SecurityhubInsightFiltersResourceTags)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceTags:
		value_ := value.([]*SecurityhubInsightFiltersResourceTags)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceTags; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutResourceTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersResourceType:
		value := value.(*[]*SecurityhubInsightFiltersResourceType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersResourceType:
		value_ := value.([]*SecurityhubInsightFiltersResourceType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersResourceType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutSampleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersSample:
		value := value.(*[]*SecurityhubInsightFiltersSample)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersSample:
		value_ := value.([]*SecurityhubInsightFiltersSample)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersSample; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutSeverityLabelParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersSeverityLabel:
		value := value.(*[]*SecurityhubInsightFiltersSeverityLabel)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersSeverityLabel:
		value_ := value.([]*SecurityhubInsightFiltersSeverityLabel)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersSeverityLabel; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutSeverityNormalizedParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersSeverityNormalized:
		value := value.(*[]*SecurityhubInsightFiltersSeverityNormalized)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersSeverityNormalized:
		value_ := value.([]*SecurityhubInsightFiltersSeverityNormalized)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersSeverityNormalized; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutSeverityProductParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersSeverityProduct:
		value := value.(*[]*SecurityhubInsightFiltersSeverityProduct)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersSeverityProduct:
		value_ := value.([]*SecurityhubInsightFiltersSeverityProduct)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersSeverityProduct; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutSourceUrlParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersSourceUrl:
		value := value.(*[]*SecurityhubInsightFiltersSourceUrl)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersSourceUrl:
		value_ := value.([]*SecurityhubInsightFiltersSourceUrl)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersSourceUrl; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutThreatIntelIndicatorCategoryParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory:
		value := value.(*[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersThreatIntelIndicatorCategory:
		value_ := value.([]*SecurityhubInsightFiltersThreatIntelIndicatorCategory)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersThreatIntelIndicatorCategory; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutThreatIntelIndicatorLastObservedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt:
		value := value.(*[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt:
		value_ := value.([]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutThreatIntelIndicatorSourceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersThreatIntelIndicatorSource:
		value := value.(*[]*SecurityhubInsightFiltersThreatIntelIndicatorSource)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersThreatIntelIndicatorSource:
		value_ := value.([]*SecurityhubInsightFiltersThreatIntelIndicatorSource)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersThreatIntelIndicatorSource; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutThreatIntelIndicatorSourceUrlParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl:
		value := value.(*[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl:
		value_ := value.([]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutThreatIntelIndicatorTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersThreatIntelIndicatorType:
		value := value.(*[]*SecurityhubInsightFiltersThreatIntelIndicatorType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersThreatIntelIndicatorType:
		value_ := value.([]*SecurityhubInsightFiltersThreatIntelIndicatorType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersThreatIntelIndicatorType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutThreatIntelIndicatorValueParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersThreatIntelIndicatorValue:
		value := value.(*[]*SecurityhubInsightFiltersThreatIntelIndicatorValue)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersThreatIntelIndicatorValue:
		value_ := value.([]*SecurityhubInsightFiltersThreatIntelIndicatorValue)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersThreatIntelIndicatorValue; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutTitleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersTitle:
		value := value.(*[]*SecurityhubInsightFiltersTitle)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersTitle:
		value_ := value.([]*SecurityhubInsightFiltersTitle)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersTitle; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersType:
		value := value.(*[]*SecurityhubInsightFiltersType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersType:
		value_ := value.([]*SecurityhubInsightFiltersType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutUpdatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersUpdatedAt:
		value := value.(*[]*SecurityhubInsightFiltersUpdatedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersUpdatedAt:
		value_ := value.([]*SecurityhubInsightFiltersUpdatedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersUpdatedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutUserDefinedFieldsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersUserDefinedFields:
		value := value.(*[]*SecurityhubInsightFiltersUserDefinedFields)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersUserDefinedFields:
		value_ := value.([]*SecurityhubInsightFiltersUserDefinedFields)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersUserDefinedFields; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutVerificationStateParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersVerificationState:
		value := value.(*[]*SecurityhubInsightFiltersVerificationState)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersVerificationState:
		value_ := value.([]*SecurityhubInsightFiltersVerificationState)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersVerificationState; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutVulnerabilitiesExploitAvailableParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersVulnerabilitiesExploitAvailable:
		value := value.(*[]*SecurityhubInsightFiltersVulnerabilitiesExploitAvailable)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersVulnerabilitiesExploitAvailable:
		value_ := value.([]*SecurityhubInsightFiltersVulnerabilitiesExploitAvailable)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersVulnerabilitiesExploitAvailable; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutVulnerabilitiesFixAvailableParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersVulnerabilitiesFixAvailable:
		value := value.(*[]*SecurityhubInsightFiltersVulnerabilitiesFixAvailable)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersVulnerabilitiesFixAvailable:
		value_ := value.([]*SecurityhubInsightFiltersVulnerabilitiesFixAvailable)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersVulnerabilitiesFixAvailable; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutWorkflowStateParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersWorkflowState:
		value := value.(*[]*SecurityhubInsightFiltersWorkflowState)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersWorkflowState:
		value_ := value.([]*SecurityhubInsightFiltersWorkflowState)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersWorkflowState; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validatePutWorkflowStatusParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*SecurityhubInsightFiltersWorkflowStatus:
		value := value.(*[]*SecurityhubInsightFiltersWorkflowStatus)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*SecurityhubInsightFiltersWorkflowStatus:
		value_ := value.([]*SecurityhubInsightFiltersWorkflowStatus)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*SecurityhubInsightFiltersWorkflowStatus; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *SecurityhubInsightFilters:
		val := val.(*SecurityhubInsightFilters)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case SecurityhubInsightFilters:
		val_ := val.(SecurityhubInsightFilters)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *SecurityhubInsightFilters; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewSecurityhubInsightFiltersOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

