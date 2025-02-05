package filters

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

func ByTaskTheme(theme string) func(any) any {
	return func(a any) any {
		switch v := a.(type) {
		case squirrel.SelectBuilder:
			return v.Where(squirrel.Like{"t.theme": "%" + theme + "%"})
		default:
			panic(fmt.Sprintf("Unexpected filter type: %T", a))
		}
	}
}

func ByTaskNotFinished(is_finished bool) func(any) any {
	return func(a any) any {
		switch v := a.(type) {
		case squirrel.SelectBuilder:
			return v.Where(squirrel.Eq{"t.is_finished": is_finished})
		default:
			panic(fmt.Sprintf("Unexpected filter type: %T", a))
		}
	}
}

func ByTaskComplexity(min, max uint16) func(any) any {
	return func(a any) any {
		switch v := a.(type) {
		case squirrel.SelectBuilder:
			return v.Where("t.complexity BETWEEN ? AND ? ", min, max)
		default:
			panic(fmt.Sprintf("Unexpected filter type: %T", a))
		}
	}
}
//TODO: move to hint filters
func ByUsedHint(is_used bool) func(any) any {
	return func(a any) any {
		switch v := a.(type) {
		case squirrel.SelectBuilder:
			return v.Where(squirrel.Eq{"t.is_used": is_used})
		default:
			panic(fmt.Sprintf("Unexpected filter type: %T", a))
		}
	}
}
