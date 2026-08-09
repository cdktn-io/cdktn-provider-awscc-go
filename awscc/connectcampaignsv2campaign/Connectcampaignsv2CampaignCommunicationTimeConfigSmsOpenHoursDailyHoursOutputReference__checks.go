// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package connectcampaignsv2campaign

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validatePutValueParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursValue:
		value := value.(*[]*Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursValue)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursValue:
		value_ := value.([]*Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursValue)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursValue; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHours:
		val := val.(*Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHours)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHours:
		val_ := val.(Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHours)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHours; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateSetKeyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Connectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewConnectcampaignsv2CampaignCommunicationTimeConfigSmsOpenHoursDailyHoursOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

