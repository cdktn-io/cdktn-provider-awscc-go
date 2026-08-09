// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package lexbot

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtons:
		val := val.(*LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtons)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtons:
		val_ := val.(LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtons)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtons; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateSetTextParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_LexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReference) validateSetValueParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewLexBotBotLocalesIntentsIntentClosingSettingClosingResponseMessageGroupsListVariationsImageResponseCardButtonsOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

