// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package dlmlifecyclepolicy

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRule:
		val := val.(*DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRule)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRule:
		val_ := val.(DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRule)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRule; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateSetIntervalParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateSetIntervalUnitParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDlmLifecyclePolicyPolicyDetailsSchedulesCrossRegionCopyRulesDeprecateRuleOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

