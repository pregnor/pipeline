// +build !ignore_autogenerated

// Code generated by mga tool. DO NOT EDIT.

package cluster

import (
	"context"
	"github.com/stretchr/testify/mock"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
)

// MockdynamicInterface is an autogenerated mock for the dynamicInterface type.
type MockdynamicInterface struct {
	mock.Mock
}

// Resource provides a mock function.
func (_m *MockdynamicInterface) Resource(resource schema.GroupVersionResource) (_result_0 dynamic.NamespaceableResourceInterface) {
	ret := _m.Called(resource)

	var r0 dynamic.NamespaceableResourceInterface
	if rf, ok := ret.Get(0).(func(schema.GroupVersionResource) dynamic.NamespaceableResourceInterface); ok {
		r0 = rf(resource)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynamic.NamespaceableResourceInterface)
		}
	}

	return r0
}

// MockdynamicNamespaceableResourceInterface is an autogenerated mock for the dynamicNamespaceableResourceInterface type.
type MockdynamicNamespaceableResourceInterface struct {
	mock.Mock
}

// Create provides a mock function.
func (_m *MockdynamicNamespaceableResourceInterface) Create(ctx context.Context, obj *unstructured.Unstructured, options v1.CreateOptions, subresources ...string) (_result_0 *unstructured.Unstructured, _result_1 error) {
	varParams := make([]interface{}, 3+len(subresources))
	varParams[0] = ctx
	varParams[1] = obj
	varParams[2] = options
	for varIndex, varParam := range subresources {
		varParams[3+varIndex] = varParam
	}

	ret := _m.Called(varParams...)

	var r0 *unstructured.Unstructured
	if rf, ok := ret.Get(0).(func(context.Context, *unstructured.Unstructured, v1.CreateOptions, ...string) *unstructured.Unstructured); ok {
		r0 = rf(ctx, obj, options, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *unstructured.Unstructured, v1.CreateOptions, ...string) error); ok {
		r1 = rf(ctx, obj, options, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function.
func (_m *MockdynamicNamespaceableResourceInterface) Delete(ctx context.Context, name string, options v1.DeleteOptions, subresources ...string) (_result_0 error) {
	varParams := make([]interface{}, 3+len(subresources))
	varParams[0] = ctx
	varParams[1] = name
	varParams[2] = options
	for varIndex, varParam := range subresources {
		varParams[3+varIndex] = varParam
	}

	ret := _m.Called(varParams...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions, ...string) error); ok {
		r0 = rf(ctx, name, options, subresources...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function.
func (_m *MockdynamicNamespaceableResourceInterface) DeleteCollection(ctx context.Context, options v1.DeleteOptions, listOptions v1.ListOptions) (_result_0 error) {
	ret := _m.Called(ctx, options, listOptions)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, options, listOptions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function.
func (_m *MockdynamicNamespaceableResourceInterface) Get(ctx context.Context, name string, options v1.GetOptions, subresources ...string) (_result_0 *unstructured.Unstructured, _result_1 error) {
	varParams := make([]interface{}, 3+len(subresources))
	varParams[0] = ctx
	varParams[1] = name
	varParams[2] = options
	for varIndex, varParam := range subresources {
		varParams[3+varIndex] = varParam
	}

	ret := _m.Called(varParams...)

	var r0 *unstructured.Unstructured
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions, ...string) *unstructured.Unstructured); ok {
		r0 = rf(ctx, name, options, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions, ...string) error); ok {
		r1 = rf(ctx, name, options, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function.
func (_m *MockdynamicNamespaceableResourceInterface) List(ctx context.Context, opts v1.ListOptions) (_result_0 *unstructured.UnstructuredList, _result_1 error) {
	ret := _m.Called(ctx, opts)

	var r0 *unstructured.UnstructuredList
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *unstructured.UnstructuredList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.UnstructuredList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Namespace provides a mock function.
func (_m *MockdynamicNamespaceableResourceInterface) Namespace(_parameter_0 string) (_result_0 dynamic.ResourceInterface) {
	ret := _m.Called(_parameter_0)

	var r0 dynamic.ResourceInterface
	if rf, ok := ret.Get(0).(func(string) dynamic.ResourceInterface); ok {
		r0 = rf(_parameter_0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynamic.ResourceInterface)
		}
	}

	return r0
}

// Patch provides a mock function.
func (_m *MockdynamicNamespaceableResourceInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []uint8, options v1.PatchOptions, subresources ...string) (_result_0 *unstructured.Unstructured, _result_1 error) {
	varParams := make([]interface{}, 5+len(subresources))
	varParams[0] = ctx
	varParams[1] = name
	varParams[2] = pt
	varParams[3] = data
	varParams[4] = options
	for varIndex, varParam := range subresources {
		varParams[5+varIndex] = varParam
	}

	ret := _m.Called(varParams...)

	var r0 *unstructured.Unstructured
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []uint8, v1.PatchOptions, ...string) *unstructured.Unstructured); ok {
		r0 = rf(ctx, name, pt, data, options, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []uint8, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, options, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function.
func (_m *MockdynamicNamespaceableResourceInterface) Update(ctx context.Context, obj *unstructured.Unstructured, options v1.UpdateOptions, subresources ...string) (_result_0 *unstructured.Unstructured, _result_1 error) {
	varParams := make([]interface{}, 3+len(subresources))
	varParams[0] = ctx
	varParams[1] = obj
	varParams[2] = options
	for varIndex, varParam := range subresources {
		varParams[3+varIndex] = varParam
	}

	ret := _m.Called(varParams...)

	var r0 *unstructured.Unstructured
	if rf, ok := ret.Get(0).(func(context.Context, *unstructured.Unstructured, v1.UpdateOptions, ...string) *unstructured.Unstructured); ok {
		r0 = rf(ctx, obj, options, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *unstructured.Unstructured, v1.UpdateOptions, ...string) error); ok {
		r1 = rf(ctx, obj, options, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function.
func (_m *MockdynamicNamespaceableResourceInterface) UpdateStatus(ctx context.Context, obj *unstructured.Unstructured, options v1.UpdateOptions) (_result_0 *unstructured.Unstructured, _result_1 error) {
	ret := _m.Called(ctx, obj, options)

	var r0 *unstructured.Unstructured
	if rf, ok := ret.Get(0).(func(context.Context, *unstructured.Unstructured, v1.UpdateOptions) *unstructured.Unstructured); ok {
		r0 = rf(ctx, obj, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *unstructured.Unstructured, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, obj, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function.
func (_m *MockdynamicNamespaceableResourceInterface) Watch(ctx context.Context, opts v1.ListOptions) (_result_0 watch.Interface, _result_1 error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDynamicKubeClientFactory is an autogenerated mock for the DynamicKubeClientFactory type.
type MockDynamicKubeClientFactory struct {
	mock.Mock
}

// FromSecret provides a mock function.
func (_m *MockDynamicKubeClientFactory) FromSecret(ctx context.Context, secretID string) (_result_0 dynamic.Interface, _result_1 error) {
	ret := _m.Called(ctx, secretID)

	var r0 dynamic.Interface
	if rf, ok := ret.Get(0).(func(context.Context, string) dynamic.Interface); ok {
		r0 = rf(ctx, secretID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynamic.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, secretID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore is an autogenerated mock for the Store type.
type MockStore struct {
	mock.Mock
}

// GetCluster provides a mock function.
func (_m *MockStore) GetCluster(ctx context.Context, id uint) (_result_0 Cluster, _result_1 error) {
	ret := _m.Called(ctx, id)

	var r0 Cluster
	if rf, ok := ret.Get(0).(func(context.Context, uint) Cluster); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(Cluster)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClusterByName provides a mock function.
func (_m *MockStore) GetClusterByName(ctx context.Context, orgID uint, clusterName string) (_result_0 Cluster, _result_1 error) {
	ret := _m.Called(ctx, orgID, clusterName)

	var r0 Cluster
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) Cluster); ok {
		r0 = rf(ctx, orgID, clusterName)
	} else {
		r0 = ret.Get(0).(Cluster)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, string) error); ok {
		r1 = rf(ctx, orgID, clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetStatus provides a mock function.
func (_m *MockStore) SetStatus(ctx context.Context, id uint, status string, statusMessage string) (_result_0 error) {
	ret := _m.Called(ctx, id, status, statusMessage)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, string, string) error); ok {
		r0 = rf(ctx, id, status, statusMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockService is an autogenerated mock for the Service type.
type MockService struct {
	mock.Mock
}

// CreateNodePools provides a mock function.
func (_m *MockService) CreateNodePools(ctx context.Context, clusterID uint, rawNodePools map[string]NewRawNodePool) (_result_0 error) {
	ret := _m.Called(ctx, clusterID, rawNodePools)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, map[string]NewRawNodePool) error); ok {
		r0 = rf(ctx, clusterID, rawNodePools)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCluster provides a mock function.
func (_m *MockService) DeleteCluster(ctx context.Context, clusterIdentifier Identifier, options DeleteClusterOptions) (deleted bool, err error) {
	ret := _m.Called(ctx, clusterIdentifier, options)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, Identifier, DeleteClusterOptions) bool); ok {
		r0 = rf(ctx, clusterIdentifier, options)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, Identifier, DeleteClusterOptions) error); ok {
		r1 = rf(ctx, clusterIdentifier, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteNodePool provides a mock function.
func (_m *MockService) DeleteNodePool(ctx context.Context, clusterID uint, name string) (deleted bool, err error) {
	ret := _m.Called(ctx, clusterID, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, uint, string) bool); ok {
		r0 = rf(ctx, clusterID, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, string) error); ok {
		r1 = rf(ctx, clusterID, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNodePools provides a mock function.
func (_m *MockService) ListNodePools(ctx context.Context, clusterID uint) (nodePoolList RawNodePoolList, err error) {
	ret := _m.Called(ctx, clusterID)

	var r0 RawNodePoolList
	if rf, ok := ret.Get(0).(func(context.Context, uint) RawNodePoolList); ok {
		r0 = rf(ctx, clusterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(RawNodePoolList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, clusterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCluster provides a mock function.
func (_m *MockService) UpdateCluster(ctx context.Context, clusterIdentifier Identifier, clusterUpdate ClusterUpdate) (err error) {
	ret := _m.Called(ctx, clusterIdentifier, clusterUpdate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Identifier, ClusterUpdate) error); ok {
		r0 = rf(ctx, clusterIdentifier, clusterUpdate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNodePool provides a mock function.
func (_m *MockService) UpdateNodePool(ctx context.Context, clusterID uint, nodePoolName string, rawNodePoolUpdate RawNodePoolUpdate) (processID string, err error) {
	ret := _m.Called(ctx, clusterID, nodePoolName, rawNodePoolUpdate)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, uint, string, RawNodePoolUpdate) string); ok {
		r0 = rf(ctx, clusterID, nodePoolName, rawNodePoolUpdate)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint, string, RawNodePoolUpdate) error); ok {
		r1 = rf(ctx, clusterID, nodePoolName, rawNodePoolUpdate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
