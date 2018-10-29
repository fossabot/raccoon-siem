package sdk

import (
	"fmt"
	"strings"
)

func RegisterFilters(settings []FilterSettings) ([]IFilter, error) {
	uniqueFilters := make(map[string]bool)
	allSimpleFiltersByName := make(map[string]*filter)
	result := make([]IFilter, 0)

	var err error
	var fi IFilter

	for _, setting := range settings {
		if _, ok := uniqueFilters[setting.Name]; ok {
			continue
		}

		if setting.Join {
			fi, err = setting.compileJoin()
		} else {
			f, localErr := setting.compile()
			fi = f
			err = localErr
			allSimpleFiltersByName[f.name] = f
		}

		if err != nil {
			return nil, err
		}

		uniqueFilters[setting.Name] = true
		result = append(result, fi)
	}

	if err := injectSubFilters(allSimpleFiltersByName); err != nil {
		return nil, err
	}

	return result, nil
}

func GetIncludedFilterName(expr string) string {
	return strings.Split(expr[len(FilterIncludeSymbol):], " ")[0]
}

func injectSubFilters(allFilters map[string]*filter) error {
	for _, filter := range allFilters {
		for _, section := range filter.sections {
			for _, cond := range section.conditions {
				if cond.incFilterName != "" {
					incFilter, ok := allFilters[cond.incFilterName]

					if !ok {
						return fmt.Errorf("included filter '%s' does not exist", cond.incFilterName)
					}

					if incFilter.name == filter.name {
						return fmt.Errorf("cyclic dependency in filter '%s'", filter.name)
					}

					cond.incFilter = incFilter
				}
			}
		}
	}
	return nil
}
