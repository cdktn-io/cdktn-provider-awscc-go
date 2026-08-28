// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package mediatailorprefetchschedule

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindow:
		val := val.(*MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindow)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindow:
		val_ := val.(MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindow)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindow; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateSetRetrievalWindowDurationSecondsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewMediatailorPrefetchScheduleRecurringPrefetchConfigurationRecurringRetrievalTrafficShapingRetrievalWindowOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

