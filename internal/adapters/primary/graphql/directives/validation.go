package directives

import (
	"context"
	"regexp"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ValidationError はバリデーションエラーを表す構造体
type ValidationError struct {
	Field   string
	Message string
}

// Error はエラーメッセージを返す
func (e *ValidationError) Error() string {
	return e.Message
}

// createValidationError は標準的なGraphQLエラーを生成する
func createValidationError(ctx context.Context, path ast.Path, value interface{}, message string) *gqlerror.Error {
	return &gqlerror.Error{
		Path:    path,
		Message: message,
		Extensions: map[string]interface{}{
			"code":  "GRAPHQL_VALIDATION_FAILED",
			"value": value,
		},
		Locations: []gqlerror.Location{},
	}
}

// ValidationDirective はバリデーションディレクティブの実装
func ValidationDirective(ctx context.Context, obj interface{}, next graphql.Resolver, required *bool, minLength *int, maxLength *int, min *int, max *int, email *bool, pattern *string) (interface{}, error) {
	val, err := next(ctx)
	if err != nil {
		return nil, err
	}

	path := graphql.GetPath(ctx)
	var fieldName string
	if len(path) > 0 {
		if step, ok := path[len(path)-1].(ast.PathElement); ok {
			switch s := step.(type) {
			case ast.PathName:
				fieldName = string(s)
			default:
				fieldName = "unknown"
			}
		}
	}

	if val == nil {
		if required != nil && *required {
			return nil, createValidationError(ctx, path, val,
				"Field \""+fieldName+"\" of required type was not provided")
		}
		return val, nil
	}

	validate := validator.New()

	// 文字列のバリデーション
	if str, ok := val.(string); ok {
		if minLength != nil {
			if len(str) < *minLength {
				return nil, createValidationError(ctx, path, val,
					"String length must be at least "+strconv.Itoa(*minLength)+" characters")
			}
		}
		if maxLength != nil {
			if len(str) > *maxLength {
				return nil, createValidationError(ctx, path, val,
					"String length must be at most "+strconv.Itoa(*maxLength)+" characters")
			}
		}
		if email != nil && *email {
			if err := validate.Var(str, "email"); err != nil {
				return nil, createValidationError(ctx, path, val, "Invalid email format")
			}
		}
		if pattern != nil {
			re, err := regexp.Compile(*pattern)
			if err != nil {
				return nil, createValidationError(ctx, path, val, "Invalid pattern")
			}
			if !re.MatchString(str) {
				return nil, createValidationError(ctx, path, val,
					"String does not match required pattern")
			}
		}
	}

	// 数値のバリデーション
	if num, ok := val.(int); ok {
		if min != nil {
			if num < *min {
				return nil, createValidationError(ctx, path, val,
					"Number must be at least "+strconv.Itoa(*min))
			}
		}
		if max != nil {
			if num > *max {
				return nil, createValidationError(ctx, path, val,
					"Number must be at most "+strconv.Itoa(*max))
			}
		}
	}

	// 配列のバリデーション
	if arr, ok := val.([]interface{}); ok {
		if minLength != nil {
			if len(arr) < *minLength {
				return nil, createValidationError(ctx, path, val,
					"Array length must be at least "+strconv.Itoa(*minLength))
			}
		}
		if maxLength != nil {
			if len(arr) > *maxLength {
				return nil, createValidationError(ctx, path, val,
					"Array length must be at most "+strconv.Itoa(*maxLength))
			}
		}
	}

	return val, nil
}
