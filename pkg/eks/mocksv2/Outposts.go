// Code generated by mockery v2.30.17. DO NOT EDIT.

package mocksv2

import (
	context "context"

	outposts "github.com/aws/aws-sdk-go-v2/service/outposts"
	mock "github.com/stretchr/testify/mock"
)

// Outposts is an autogenerated mock type for the Outposts type
type Outposts struct {
	mock.Mock
}

// CancelOrder provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) CancelOrder(ctx context.Context, params *outposts.CancelOrderInput, optFns ...func(*outposts.Options)) (*outposts.CancelOrderOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.CancelOrderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.CancelOrderInput, ...func(*outposts.Options)) (*outposts.CancelOrderOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.CancelOrderInput, ...func(*outposts.Options)) *outposts.CancelOrderOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.CancelOrderOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.CancelOrderInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrder provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) CreateOrder(ctx context.Context, params *outposts.CreateOrderInput, optFns ...func(*outposts.Options)) (*outposts.CreateOrderOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.CreateOrderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.CreateOrderInput, ...func(*outposts.Options)) (*outposts.CreateOrderOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.CreateOrderInput, ...func(*outposts.Options)) *outposts.CreateOrderOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.CreateOrderOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.CreateOrderInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOutpost provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) CreateOutpost(ctx context.Context, params *outposts.CreateOutpostInput, optFns ...func(*outposts.Options)) (*outposts.CreateOutpostOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.CreateOutpostOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.CreateOutpostInput, ...func(*outposts.Options)) (*outposts.CreateOutpostOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.CreateOutpostInput, ...func(*outposts.Options)) *outposts.CreateOutpostOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.CreateOutpostOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.CreateOutpostInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSite provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) CreateSite(ctx context.Context, params *outposts.CreateSiteInput, optFns ...func(*outposts.Options)) (*outposts.CreateSiteOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.CreateSiteOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.CreateSiteInput, ...func(*outposts.Options)) (*outposts.CreateSiteOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.CreateSiteInput, ...func(*outposts.Options)) *outposts.CreateSiteOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.CreateSiteOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.CreateSiteInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOutpost provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) DeleteOutpost(ctx context.Context, params *outposts.DeleteOutpostInput, optFns ...func(*outposts.Options)) (*outposts.DeleteOutpostOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.DeleteOutpostOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.DeleteOutpostInput, ...func(*outposts.Options)) (*outposts.DeleteOutpostOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.DeleteOutpostInput, ...func(*outposts.Options)) *outposts.DeleteOutpostOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.DeleteOutpostOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.DeleteOutpostInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSite provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) DeleteSite(ctx context.Context, params *outposts.DeleteSiteInput, optFns ...func(*outposts.Options)) (*outposts.DeleteSiteOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.DeleteSiteOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.DeleteSiteInput, ...func(*outposts.Options)) (*outposts.DeleteSiteOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.DeleteSiteInput, ...func(*outposts.Options)) *outposts.DeleteSiteOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.DeleteSiteOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.DeleteSiteInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCatalogItem provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) GetCatalogItem(ctx context.Context, params *outposts.GetCatalogItemInput, optFns ...func(*outposts.Options)) (*outposts.GetCatalogItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.GetCatalogItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetCatalogItemInput, ...func(*outposts.Options)) (*outposts.GetCatalogItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetCatalogItemInput, ...func(*outposts.Options)) *outposts.GetCatalogItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.GetCatalogItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.GetCatalogItemInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnection provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) GetConnection(ctx context.Context, params *outposts.GetConnectionInput, optFns ...func(*outposts.Options)) (*outposts.GetConnectionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.GetConnectionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetConnectionInput, ...func(*outposts.Options)) (*outposts.GetConnectionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetConnectionInput, ...func(*outposts.Options)) *outposts.GetConnectionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.GetConnectionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.GetConnectionInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrder provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) GetOrder(ctx context.Context, params *outposts.GetOrderInput, optFns ...func(*outposts.Options)) (*outposts.GetOrderOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.GetOrderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetOrderInput, ...func(*outposts.Options)) (*outposts.GetOrderOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetOrderInput, ...func(*outposts.Options)) *outposts.GetOrderOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.GetOrderOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.GetOrderInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOutpost provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) GetOutpost(ctx context.Context, params *outposts.GetOutpostInput, optFns ...func(*outposts.Options)) (*outposts.GetOutpostOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.GetOutpostOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetOutpostInput, ...func(*outposts.Options)) (*outposts.GetOutpostOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetOutpostInput, ...func(*outposts.Options)) *outposts.GetOutpostOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.GetOutpostOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.GetOutpostInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOutpostInstanceTypes provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) GetOutpostInstanceTypes(ctx context.Context, params *outposts.GetOutpostInstanceTypesInput, optFns ...func(*outposts.Options)) (*outposts.GetOutpostInstanceTypesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.GetOutpostInstanceTypesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetOutpostInstanceTypesInput, ...func(*outposts.Options)) (*outposts.GetOutpostInstanceTypesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetOutpostInstanceTypesInput, ...func(*outposts.Options)) *outposts.GetOutpostInstanceTypesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.GetOutpostInstanceTypesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.GetOutpostInstanceTypesInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSite provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) GetSite(ctx context.Context, params *outposts.GetSiteInput, optFns ...func(*outposts.Options)) (*outposts.GetSiteOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.GetSiteOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetSiteInput, ...func(*outposts.Options)) (*outposts.GetSiteOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetSiteInput, ...func(*outposts.Options)) *outposts.GetSiteOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.GetSiteOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.GetSiteInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSiteAddress provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) GetSiteAddress(ctx context.Context, params *outposts.GetSiteAddressInput, optFns ...func(*outposts.Options)) (*outposts.GetSiteAddressOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.GetSiteAddressOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetSiteAddressInput, ...func(*outposts.Options)) (*outposts.GetSiteAddressOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.GetSiteAddressInput, ...func(*outposts.Options)) *outposts.GetSiteAddressOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.GetSiteAddressOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.GetSiteAddressInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAssets provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) ListAssets(ctx context.Context, params *outposts.ListAssetsInput, optFns ...func(*outposts.Options)) (*outposts.ListAssetsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.ListAssetsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListAssetsInput, ...func(*outposts.Options)) (*outposts.ListAssetsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListAssetsInput, ...func(*outposts.Options)) *outposts.ListAssetsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.ListAssetsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.ListAssetsInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCatalogItems provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) ListCatalogItems(ctx context.Context, params *outposts.ListCatalogItemsInput, optFns ...func(*outposts.Options)) (*outposts.ListCatalogItemsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.ListCatalogItemsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListCatalogItemsInput, ...func(*outposts.Options)) (*outposts.ListCatalogItemsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListCatalogItemsInput, ...func(*outposts.Options)) *outposts.ListCatalogItemsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.ListCatalogItemsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.ListCatalogItemsInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOrders provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) ListOrders(ctx context.Context, params *outposts.ListOrdersInput, optFns ...func(*outposts.Options)) (*outposts.ListOrdersOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.ListOrdersOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListOrdersInput, ...func(*outposts.Options)) (*outposts.ListOrdersOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListOrdersInput, ...func(*outposts.Options)) *outposts.ListOrdersOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.ListOrdersOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.ListOrdersInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOutposts provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) ListOutposts(ctx context.Context, params *outposts.ListOutpostsInput, optFns ...func(*outposts.Options)) (*outposts.ListOutpostsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.ListOutpostsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListOutpostsInput, ...func(*outposts.Options)) (*outposts.ListOutpostsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListOutpostsInput, ...func(*outposts.Options)) *outposts.ListOutpostsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.ListOutpostsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.ListOutpostsInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSites provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) ListSites(ctx context.Context, params *outposts.ListSitesInput, optFns ...func(*outposts.Options)) (*outposts.ListSitesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.ListSitesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListSitesInput, ...func(*outposts.Options)) (*outposts.ListSitesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListSitesInput, ...func(*outposts.Options)) *outposts.ListSitesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.ListSitesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.ListSitesInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForResource provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) ListTagsForResource(ctx context.Context, params *outposts.ListTagsForResourceInput, optFns ...func(*outposts.Options)) (*outposts.ListTagsForResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.ListTagsForResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListTagsForResourceInput, ...func(*outposts.Options)) (*outposts.ListTagsForResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.ListTagsForResourceInput, ...func(*outposts.Options)) *outposts.ListTagsForResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.ListTagsForResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.ListTagsForResourceInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartConnection provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) StartConnection(ctx context.Context, params *outposts.StartConnectionInput, optFns ...func(*outposts.Options)) (*outposts.StartConnectionOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.StartConnectionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.StartConnectionInput, ...func(*outposts.Options)) (*outposts.StartConnectionOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.StartConnectionInput, ...func(*outposts.Options)) *outposts.StartConnectionOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.StartConnectionOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.StartConnectionInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagResource provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) TagResource(ctx context.Context, params *outposts.TagResourceInput, optFns ...func(*outposts.Options)) (*outposts.TagResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.TagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.TagResourceInput, ...func(*outposts.Options)) (*outposts.TagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.TagResourceInput, ...func(*outposts.Options)) *outposts.TagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.TagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.TagResourceInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UntagResource provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) UntagResource(ctx context.Context, params *outposts.UntagResourceInput, optFns ...func(*outposts.Options)) (*outposts.UntagResourceOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.UntagResourceOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.UntagResourceInput, ...func(*outposts.Options)) (*outposts.UntagResourceOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.UntagResourceInput, ...func(*outposts.Options)) *outposts.UntagResourceOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.UntagResourceOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.UntagResourceInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOutpost provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) UpdateOutpost(ctx context.Context, params *outposts.UpdateOutpostInput, optFns ...func(*outposts.Options)) (*outposts.UpdateOutpostOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.UpdateOutpostOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.UpdateOutpostInput, ...func(*outposts.Options)) (*outposts.UpdateOutpostOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.UpdateOutpostInput, ...func(*outposts.Options)) *outposts.UpdateOutpostOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.UpdateOutpostOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.UpdateOutpostInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSite provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) UpdateSite(ctx context.Context, params *outposts.UpdateSiteInput, optFns ...func(*outposts.Options)) (*outposts.UpdateSiteOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.UpdateSiteOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.UpdateSiteInput, ...func(*outposts.Options)) (*outposts.UpdateSiteOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.UpdateSiteInput, ...func(*outposts.Options)) *outposts.UpdateSiteOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.UpdateSiteOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.UpdateSiteInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSiteAddress provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) UpdateSiteAddress(ctx context.Context, params *outposts.UpdateSiteAddressInput, optFns ...func(*outposts.Options)) (*outposts.UpdateSiteAddressOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.UpdateSiteAddressOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.UpdateSiteAddressInput, ...func(*outposts.Options)) (*outposts.UpdateSiteAddressOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.UpdateSiteAddressInput, ...func(*outposts.Options)) *outposts.UpdateSiteAddressOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.UpdateSiteAddressOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.UpdateSiteAddressInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSiteRackPhysicalProperties provides a mock function with given fields: ctx, params, optFns
func (_m *Outposts) UpdateSiteRackPhysicalProperties(ctx context.Context, params *outposts.UpdateSiteRackPhysicalPropertiesInput, optFns ...func(*outposts.Options)) (*outposts.UpdateSiteRackPhysicalPropertiesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *outposts.UpdateSiteRackPhysicalPropertiesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.UpdateSiteRackPhysicalPropertiesInput, ...func(*outposts.Options)) (*outposts.UpdateSiteRackPhysicalPropertiesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *outposts.UpdateSiteRackPhysicalPropertiesInput, ...func(*outposts.Options)) *outposts.UpdateSiteRackPhysicalPropertiesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*outposts.UpdateSiteRackPhysicalPropertiesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *outposts.UpdateSiteRackPhysicalPropertiesInput, ...func(*outposts.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOutposts creates a new instance of Outposts. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOutposts(t interface {
	mock.TestingT
	Cleanup(func())
}) *Outposts {
	mock := &Outposts{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
