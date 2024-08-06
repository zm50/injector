package injector

import (
	"fmt"
	"reflect"
)

type ApplicationContext interface {
	Inject(provider BeanProvider, name string) error
	Autowise(obj any, name string) error
}

type BeanProvider func() reflect.Value

type context struct {
	namedConatiner map[string]BeanProvider
	typedContainer map[reflect.Type]BeanProvider
}

var instance ApplicationContext

func init() {
	instance = &context{
		namedConatiner: make(map[string]BeanProvider),
        typedContainer: make(map[reflect.Type]BeanProvider),
	}
}

func (c *context) Inject(provider BeanProvider, name string) error {
	if provider == nil {
		return fmt.Errorf("inject: provider can not be nil")
	}

	if name == "" {
		// type inject
		ty := provider().Type()

		if _, ok := c.typedContainer[ty];ok {
			return fmt.Errorf("inject: %v is ambiguous", ty)
		}

		c.typedContainer[ty] = provider
	} else {
		// name inject
		if _, ok := c.namedConatiner[name];ok {
			return fmt.Errorf("inject: %v is ambiguous", name)
		}

		c.namedConatiner[name] = provider
	}

	return nil
}

func (c *context) Autowise(val any, name string) error {
	if val == nil {
		return fmt.Errorf("inject: nil value")
	}
	rv := reflect.ValueOf(val)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("inject: %v is not a pointer", rv)
	}
	ri := reflect.Indirect(rv)
	rt := ri.Type()
	var provider BeanProvider
	if name == "" {
		// type autowise
		provider = c.typedContainer[rt]
	} else {
		// name autowise
		provider = c.namedConatiner[name]
	}

	if provider == nil {
		return fmt.Errorf("inject: %v is not found", name)
	}

	obj := provider()
	if obj.CanConvert(rt) {
		ri.Set(obj.Convert(rt))
		return nil
	}

	return fmt.Errorf("inject: value can not convert to %s", ri.Type())
}

func defaultBeanProvider(v any) BeanProvider {
	return func() reflect.Value {
		return reflect.ValueOf(v)
	}
}

// 对外暴露依赖注入的能力，name为空字符串时表示默认使用类型注入
func Inject(obj any, name string) error {
	return instance.Inject(defaultBeanProvider(obj), name)
}

// 对外暴露依赖注入的能力，name为空字符串时表示默认使用类型注入
func DeepInject(provider BeanProvider, name string) error {
	return instance.Inject(provider, name)
}

// 对外暴露自动装配的能力，name为空字符串时表示默认使用类型自动装配
func Autowise[T any](obj *T, name string) error {
	return instance.Autowise(obj, name)
}
