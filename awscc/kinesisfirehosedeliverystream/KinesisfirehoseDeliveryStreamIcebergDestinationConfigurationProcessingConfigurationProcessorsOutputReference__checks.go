// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package kinesisfirehosedeliverystream

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validatePutParametersParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsParameters:
		value := value.(*[]*KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsParameters)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsParameters:
		value_ := value.([]*KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsParameters)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsParameters; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessors:
		val := val.(*KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessors)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessors:
		val_ := val.(KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessors)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessors; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReference) validateSetTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewKinesisfirehoseDeliveryStreamIcebergDestinationConfigurationProcessingConfigurationProcessorsOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

