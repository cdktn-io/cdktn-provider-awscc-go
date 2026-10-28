// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package autoscalingplansscalingplan

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validatePutDimensionsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationDimensions:
		value := value.(*[]*AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationDimensions)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationDimensions:
		value_ := value.([]*AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationDimensions)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationDimensions; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecification:
		val := val.(*AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecification)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecification:
		val_ := val.(AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecification)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecification; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateSetMetricNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateSetNamespaceParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateSetStatisticParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReference) validateSetUnitParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

