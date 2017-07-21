package src

import (
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

func scan(prefix, context []string, source yaml.MapSlice) yaml.MapSlice {
	for index, item := range source {
		switch key := item.Key.(type) {
		case string:
			nowPrefix := append(prefix, key)
			value := item.Value
			if value == nil {
				resultStr := interact(strings.Join(nowPrefix, "."), context)
				if num, err := strconv.Atoi(resultStr); err == nil {
					source[index] = yaml.MapItem{Key: key, Value: num}
				} else {
					source[index] = yaml.MapItem{Key: key, Value: resultStr}
				}
				continue
			}

			if _, ok := item.Value.(yaml.MapSlice); ok {
				value := scan(nowPrefix, context, item.Value.(yaml.MapSlice))
				source[index] = yaml.MapItem{Key: key, Value: value}
				continue
			}

			if _, ok := item.Value.([]interface{}); ok {
				value := listHandler(nowPrefix, context, item.Value.([]interface{}))
				source[index] = yaml.MapItem{Key: key, Value: value}
				continue
			}
		}
	}
	return source
}

func listHandler(prefix, context []string, source []interface{}) []interface{} {
	for index, item := range source {
		switch nowItem := item.(type) {
		case yaml.MapSlice:
			nowContext := append(context, getContext(nowItem)...)
			source[index] = scan(prefix, nowContext, nowItem)
		case []interface{}:
			source[index] = listHandler(prefix, context, nowItem)
		}
	}
	return source
}

func getContext(source yaml.MapSlice) []string {
	context := make([]string, 0)
	for _, item := range source {
		switch key := item.Key.(type) {
		case string:
			switch value := item.Value.(type) {
			case string:
				context = append(context, key+":"+value)
			}
		}
	}
	return context
}
