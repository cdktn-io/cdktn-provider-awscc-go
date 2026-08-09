// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package batchjobdefinition

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevicesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevicesList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevicesList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevicesList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevices:
		val := val.(*[]*BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevices)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevices:
		val_ := val.([]*BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevices)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevices; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevicesList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevicesList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevicesList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersDevicesListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

