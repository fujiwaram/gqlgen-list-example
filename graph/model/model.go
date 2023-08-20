package model

import (
	"fmt"
	"strconv"
	"strings"
)

func (params *UserListParams) ToQuery() (string, []interface{}, error) {
	if params == nil {
		return "", nil, nil
	}

	var query string
	var args []interface{}

	if params.Filter != nil {
		q, a, err := params.Filter.ToQuery()
		if err != nil {
			return "", nil, fmt.Errorf(": %w", err)
		}
		query += q
		args = a
	}
	if params.Sort != nil {
		q, err := params.Sort.ToQuery()
		if err != nil {
			return "", nil, fmt.Errorf(": %w", err)
		}
		query += q
	}
	if params.Limit != nil {
		query += " LIMIT " + strconv.FormatInt(*params.Limit, 10)
	}
	if params.Offset != nil {
		query += " OFFSET " + strconv.FormatInt(*params.Offset, 10)
	}

	return query, args, nil
}

func (filter UserFilter) ToQuery() (string, []interface{}, error) {
	queries := make([]string, 0)
	var args []interface{}

	if filter.ID != nil {
		q, err := filter.ID.Condition.ToQuery()
		if err != nil {
			return "", nil, fmt.Errorf("failed to convert ID filter condition to query: %w", err)
		}
		queries = append(queries, "id "+q)
		args = append(args, int(filter.ID.Value))
	}
	if filter.Name != nil {
		q, err := filter.Name.Condition.ToQuery()
		if err != nil {
			return "", nil, fmt.Errorf("failed to convert Name filter condition to query: %w", err)
		}
		queries = append(queries, "name "+q)
		args = append(args, filter.Name.Value)
	}
	if filter.Email != nil {
		q, err := filter.Email.Condition.ToQuery()
		if err != nil {
			return "", nil, fmt.Errorf("failed to convert Email filter condition to query: %w", err)
		}
		queries = append(queries, "email "+q)
		args = append(args, filter.Email.Value)
	}
	if filter.Birthday != nil {
		q, err := filter.Birthday.Condition.ToQuery()
		if err != nil {
			return "", nil, fmt.Errorf("failed to convert Birthday filter condition to query: %w", err)
		}
		queries = append(queries, "birthday "+q)
		args = append(args, filter.Birthday.Value.String())
	}
	if len(queries) == 0 {
		return "", nil, nil
	}
	fmt.Println("args=", args)
	return " WHERE " + strings.Join(queries, " AND "), args, nil
}

func (cond IntFilterCondition) ToQuery() (string, error) {
	switch cond {
	case IntFilterConditionEq:
		return "= ?", nil
	case IntFilterConditionGte:
		return ">= ?", nil
	case IntFilterConditionLte:
		return "<= ?", nil
	}
	return "", fmt.Errorf("invalid int filter condition: %v", cond)
}

func (cond StringFilterCondition) ToQuery() (string, error) {
	switch cond {
	case StringFilterConditionEq:
		return "= ?", nil
	case StringFilterConditionContains:
		return "LIKE '%?%'", nil
	}
	return "", fmt.Errorf("invalid string filter condition: %v", cond)
}

func (cond TimeFilterCondition) ToQuery() (string, error) {
	switch cond {
	case TimeFilterConditionEq:
		return "= ?", nil
	case TimeFilterConditionGte:
		return ">= ?", nil
	case TimeFilterConditionLte:
		return "<= ?", nil
	}
	return "", fmt.Errorf("invalid time filter condition: %v", cond)
}

func (sort UserSort) ToQuery() (string, error) {
	queries := make([]string, 0)

	if sort.ID != nil {
		direction, err := sort.ID.ToQuery()
		if err != nil {
			return "", fmt.Errorf("failed to convert ID sort direction to query: %w", err)
		}
		queries = append(queries, "id "+direction)
	}
	if sort.Name != nil {
		direction, err := sort.Name.ToQuery()
		if err != nil {
			return "", fmt.Errorf("failed to convert Name sort direction to query: %w", err)
		}
		queries = append(queries, "name "+direction)
	}
	if sort.Email != nil {
		direction, err := sort.Email.ToQuery()
		if err != nil {
			return "", fmt.Errorf("failed to convert Email sort direction to query: %w", err)
		}
		queries = append(queries, "email "+direction)
	}
	if sort.Birthday != nil {
		direction, err := sort.Birthday.ToQuery()
		if err != nil {
			return "", fmt.Errorf("failed to convert Birthday sort direction to query: %w", err)
		}
		queries = append(queries, "birthday "+direction)
	}
	if len(queries) == 0 {
		return "", nil
	}
	return " ORDER BY " + strings.Join(queries, ", "), nil
}

func (sort SortDirection) ToQuery() (string, error) {
	switch sort {
	case SortDirectionAsc:
		return "ASC", nil
	case SortDirectionDesc:
		return "DESC", nil
	}
	return "", fmt.Errorf("invalid sort direction: %v", sort)
}
