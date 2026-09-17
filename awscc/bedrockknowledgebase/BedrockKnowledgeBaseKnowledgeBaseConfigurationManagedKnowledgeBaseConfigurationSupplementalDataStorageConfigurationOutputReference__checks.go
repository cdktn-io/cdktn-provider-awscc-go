// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package bedrockknowledgebase

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validatePutSupplementalDataStorageLocationsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationSupplementalDataStorageLocations:
		value := value.(*[]*BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationSupplementalDataStorageLocations)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationSupplementalDataStorageLocations:
		value_ := value.([]*BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationSupplementalDataStorageLocations)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationSupplementalDataStorageLocations; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (b *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfiguration:
		val := val.(*BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfiguration)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfiguration:
		val_ := val.(BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfiguration)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfiguration; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBedrockKnowledgeBaseKnowledgeBaseConfigurationManagedKnowledgeBaseConfigurationSupplementalDataStorageConfigurationOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

