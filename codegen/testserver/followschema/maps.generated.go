// Code generated by github.com/norbux/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"fmt"
	"strconv"

	"github.com/norbux/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _MapStringInterfaceType_a(ctx context.Context, field graphql.CollectedField, obj map[string]interface{}) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "MapStringInterfaceType",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		switch v := obj["a"].(type) {
		case *string:
			return v, nil
		case string:
			return &v, nil
		case nil:
			return (*string)(nil), nil
		default:
			return nil, fmt.Errorf("unexpected type %T for field %s", v, "a")
		}
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) _MapStringInterfaceType_b(ctx context.Context, field graphql.CollectedField, obj map[string]interface{}) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "MapStringInterfaceType",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		switch v := obj["b"].(type) {
		case *int:
			return v, nil
		case int:
			return &v, nil
		case nil:
			return (*int)(nil), nil
		default:
			return nil, fmt.Errorf("unexpected type %T for field %s", v, "b")
		}
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*int)
	fc.Result = res
	return ec.marshalOInt2ᚖint(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputNestedMapInput(ctx context.Context, obj interface{}) (NestedMapInput, error) {
	var it NestedMapInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	for k, v := range asMap {
		switch k {
		case "map":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("map"))
			it.Map, err = ec.unmarshalOMapStringInterfaceInput2map(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mapStringInterfaceTypeImplementors = []string{"MapStringInterfaceType"}

func (ec *executionContext) _MapStringInterfaceType(ctx context.Context, sel ast.SelectionSet, obj map[string]interface{}) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mapStringInterfaceTypeImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("MapStringInterfaceType")
		case "a":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._MapStringInterfaceType_a(ctx, field, obj)
			}

			out.Values[i] = innerFunc(ctx)

		case "b":
			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				return ec._MapStringInterfaceType_b(ctx, field, obj)
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

func (ec *executionContext) unmarshalOMapStringInterfaceInput2map(ctx context.Context, v interface{}) (map[string]interface{}, error) {
	if v == nil {
		return nil, nil
	}
	return v.(map[string]interface{}), nil
}

func (ec *executionContext) marshalOMapStringInterfaceType2map(ctx context.Context, sel ast.SelectionSet, v map[string]interface{}) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._MapStringInterfaceType(ctx, sel, v)
}

func (ec *executionContext) unmarshalONestedMapInput2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐNestedMapInput(ctx context.Context, v interface{}) (*NestedMapInput, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputNestedMapInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
