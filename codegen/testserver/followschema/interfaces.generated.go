// Code generated by github.com/norbux/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/norbux/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type BackedByInterfaceResolver interface {
	ID(ctx context.Context, obj BackedByInterface) (string, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _BackedByInterface_id(ctx context.Context, field graphql.CollectedField, obj BackedByInterface) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "BackedByInterface",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.BackedByInterface().ID(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _BackedByInterface_thisShouldBind(ctx context.Context, field graphql.CollectedField, obj BackedByInterface) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "BackedByInterface",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ThisShouldBind(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _BackedByInterface_thisShouldBindWithError(ctx context.Context, field graphql.CollectedField, obj BackedByInterface) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "BackedByInterface",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ThisShouldBindWithError()
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Cat_species(ctx context.Context, field graphql.CollectedField, obj *Cat) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Cat",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Species, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Cat_catBreed(ctx context.Context, field graphql.CollectedField, obj *Cat) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Cat",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.CatBreed, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Circle_radius(ctx context.Context, field graphql.CollectedField, obj *Circle) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Circle",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Radius, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalOFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Circle_area(ctx context.Context, field graphql.CollectedField, obj *Circle) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Circle",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Area(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalOFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Circle_coordinates(ctx context.Context, field graphql.CollectedField, obj *Circle) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Circle",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Coordinates, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(Coordinates)
	fc.Result = res
	return ec.marshalOCoordinates2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCoordinates(ctx, field.Selections, res)
}

func (ec *executionContext) _ConcreteNodeA_id(ctx context.Context, field graphql.CollectedField, obj *ConcreteNodeA) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "ConcreteNodeA",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _ConcreteNodeA_child(ctx context.Context, field graphql.CollectedField, obj *ConcreteNodeA) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "ConcreteNodeA",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Child()
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(Node)
	fc.Result = res
	return ec.marshalNNode2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐNode(ctx, field.Selections, res)
}

func (ec *executionContext) _ConcreteNodeA_name(ctx context.Context, field graphql.CollectedField, obj *ConcreteNodeA) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "ConcreteNodeA",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _ConcreteNodeInterface_id(ctx context.Context, field graphql.CollectedField, obj ConcreteNodeInterface) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "ConcreteNodeInterface",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) _ConcreteNodeInterface_child(ctx context.Context, field graphql.CollectedField, obj ConcreteNodeInterface) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "ConcreteNodeInterface",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Child()
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(Node)
	fc.Result = res
	return ec.marshalNNode2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐNode(ctx, field.Selections, res)
}

func (ec *executionContext) _Coordinates_x(ctx context.Context, field graphql.CollectedField, obj *Coordinates) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Coordinates",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.X, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalNFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Coordinates_y(ctx context.Context, field graphql.CollectedField, obj *Coordinates) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Coordinates",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Y, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalNFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Dog_species(ctx context.Context, field graphql.CollectedField, obj *Dog) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Dog",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Species, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Dog_dogBreed(ctx context.Context, field graphql.CollectedField, obj *Dog) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Dog",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DogBreed, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _Rectangle_length(ctx context.Context, field graphql.CollectedField, obj *Rectangle) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Rectangle",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Length, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalOFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Rectangle_width(ctx context.Context, field graphql.CollectedField, obj *Rectangle) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Rectangle",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Width, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalOFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Rectangle_area(ctx context.Context, field graphql.CollectedField, obj *Rectangle) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Rectangle",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Area(), nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(float64)
	fc.Result = res
	return ec.marshalOFloat2float64(ctx, field.Selections, res)
}

func (ec *executionContext) _Rectangle_coordinates(ctx context.Context, field graphql.CollectedField, obj *Rectangle) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "Rectangle",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Coordinates, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(Coordinates)
	fc.Result = res
	return ec.marshalOCoordinates2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCoordinates(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _Animal(ctx context.Context, sel ast.SelectionSet, obj Animal) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case Dog:
		return ec._Dog(ctx, sel, &obj)
	case *Dog:
		if obj == nil {
			return graphql.Null
		}
		return ec._Dog(ctx, sel, obj)
	case Cat:
		return ec._Cat(ctx, sel, &obj)
	case *Cat:
		if obj == nil {
			return graphql.Null
		}
		return ec._Cat(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func (ec *executionContext) _Node(ctx context.Context, sel ast.SelectionSet, obj Node) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case *ConcreteNodeA:
		if obj == nil {
			return graphql.Null
		}
		return ec._ConcreteNodeA(ctx, sel, obj)
	case ConcreteNodeInterface:
		if obj == nil {
			return graphql.Null
		}
		return ec._ConcreteNodeInterface(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func (ec *executionContext) _Shape(ctx context.Context, sel ast.SelectionSet, obj Shape) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case *Circle:
		if obj == nil {
			return graphql.Null
		}
		return ec._Circle(ctx, sel, obj)
	case *Rectangle:
		if obj == nil {
			return graphql.Null
		}
		return ec._Rectangle(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

func (ec *executionContext) _ShapeUnion(ctx context.Context, sel ast.SelectionSet, obj ShapeUnion) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case *Circle:
		if obj == nil {
			return graphql.Null
		}
		return ec._Circle(ctx, sel, obj)
	case *Rectangle:
		if obj == nil {
			return graphql.Null
		}
		return ec._Rectangle(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var backedByInterfaceImplementors = []string{"BackedByInterface"}

func (ec *executionContext) _BackedByInterface(ctx context.Context, sel ast.SelectionSet, obj BackedByInterface) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, backedByInterfaceImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("BackedByInterface")
		case "id":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._BackedByInterface_id(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		case "thisShouldBind":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._BackedByInterface_thisShouldBind(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		case "thisShouldBindWithError":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._BackedByInterface_thisShouldBindWithError(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				atomic.AddUint32(&invalids, 1)
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var catImplementors = []string{"Cat", "Animal"}

func (ec *executionContext) _Cat(ctx context.Context, sel ast.SelectionSet, obj *Cat) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, catImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Cat")
		case "species":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Cat_species(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "catBreed":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Cat_catBreed(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var circleImplementors = []string{"Circle", "Shape", "ShapeUnion"}

func (ec *executionContext) _Circle(ctx context.Context, sel ast.SelectionSet, obj *Circle) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, circleImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Circle")
		case "radius":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Circle_radius(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		case "area":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Circle_area(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		case "coordinates":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Circle_coordinates(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var concreteNodeAImplementors = []string{"ConcreteNodeA", "Node"}

func (ec *executionContext) _ConcreteNodeA(ctx context.Context, sel ast.SelectionSet, obj *ConcreteNodeA) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, concreteNodeAImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ConcreteNodeA")
		case "id":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._ConcreteNodeA_id(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "child":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._ConcreteNodeA_child(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "name":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._ConcreteNodeA_name(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var concreteNodeInterfaceImplementors = []string{"ConcreteNodeInterface", "Node"}

func (ec *executionContext) _ConcreteNodeInterface(ctx context.Context, sel ast.SelectionSet, obj ConcreteNodeInterface) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, concreteNodeInterfaceImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("ConcreteNodeInterface")
		case "id":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._ConcreteNodeInterface_id(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "child":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._ConcreteNodeInterface_child(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var coordinatesImplementors = []string{"Coordinates"}

func (ec *executionContext) _Coordinates(ctx context.Context, sel ast.SelectionSet, obj *Coordinates) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, coordinatesImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Coordinates")
		case "x":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Coordinates_x(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "y":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Coordinates_y(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var dogImplementors = []string{"Dog", "Animal"}

func (ec *executionContext) _Dog(ctx context.Context, sel ast.SelectionSet, obj *Dog) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, dogImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Dog")
		case "species":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Dog_species(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "dogBreed":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Dog_dogBreed(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var rectangleImplementors = []string{"Rectangle", "Shape", "ShapeUnion"}

func (ec *executionContext) _Rectangle(ctx context.Context, sel ast.SelectionSet, obj *Rectangle) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, rectangleImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Rectangle")
		case "length":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Rectangle_length(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		case "width":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Rectangle_width(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		case "area":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Rectangle_area(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		case "coordinates":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Rectangle_coordinates(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNNode2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐNode(ctx context.Context, sel ast.SelectionSet, v Node) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._Node(ctx, sel, v)
}

func (ec *executionContext) marshalNShapeUnion2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐShapeUnion(ctx context.Context, sel ast.SelectionSet, v ShapeUnion) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._ShapeUnion(ctx, sel, v)
}

func (ec *executionContext) marshalOAnimal2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐAnimal(ctx context.Context, sel ast.SelectionSet, v Animal) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Animal(ctx, sel, v)
}

func (ec *executionContext) marshalOBackedByInterface2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐBackedByInterface(ctx context.Context, sel ast.SelectionSet, v BackedByInterface) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._BackedByInterface(ctx, sel, v)
}

func (ec *executionContext) marshalOCircle2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCircle(ctx context.Context, sel ast.SelectionSet, v *Circle) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Circle(ctx, sel, v)
}

func (ec *executionContext) marshalOCoordinates2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐCoordinates(ctx context.Context, sel ast.SelectionSet, v Coordinates) graphql.Marshaler {
	return ec._Coordinates(ctx, sel, &v)
}

func (ec *executionContext) marshalOShape2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐShape(ctx context.Context, sel ast.SelectionSet, v Shape) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Shape(ctx, sel, v)
}

func (ec *executionContext) marshalOShape2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐShape(ctx context.Context, sel ast.SelectionSet, v []Shape) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOShape2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐShape(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	return ret
}

// endregion ***************************** type.gotpl *****************************
