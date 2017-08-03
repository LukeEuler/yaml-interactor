package scan

import (
	"strings"

	"gopkg.in/yaml.v2"
)

func yamlScan(prefix, context []string, source yaml.MapSlice) yaml.MapSlice {
	for index, item := range source {
		switch key := item.Key.(type) {
		case string:
			nowPrefix := append(prefix, key)
			value := item.Value
			if value == nil {
				source[index] = interact(strings.Join(nowPrefix, "."), key, context)
				continue
			}

			if _, ok := item.Value.(yaml.MapSlice); ok {
				value := yamlScan(nowPrefix, context, item.Value.(yaml.MapSlice))
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
			source[index] = yamlScan(prefix, nowContext, nowItem)
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
