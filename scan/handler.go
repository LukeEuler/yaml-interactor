package scan

import "yaml-interactor/utils"

// Handle XXX
func Handle(file string) []byte {
	content := utils.ReadFile(file)
	data := utils.Unmarshal(content)
	data = yamlScan(make([]string, 0), make([]string, 0), data)
	return utils.Marshal(data)
}
