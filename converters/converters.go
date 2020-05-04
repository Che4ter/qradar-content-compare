package converters

import "github.com/ilyaglow/go-qradar"

func LogSourceTypesToMap(itemList []qradar.LogSourceType) (map[int]string, error) {
	resultMap := make(map[int]string)
	for _, item := range itemList {
		resultMap[*item.ID] = *item.Name
	}
	return resultMap, nil
}

func LogSourcesToMap(itemList []qradar.LogSource) (map[int]string, error) {
	resultMap := make(map[int]string)
	for _, item := range itemList {
		resultMap[*item.ID] = *item.Name
	}
	return resultMap, nil
}

func LowLevelCategoriesToMap(itemList []qradar.LowLevelCategory) (map[int]string, error) {
	resultMap := make(map[int]string)
	for _, item := range itemList {
		resultMap[*item.ID] = *item.Name
	}
	return resultMap, nil
}

func QIDsToMap(itemList []qradar.QID) (map[int]string, error) {
	resultMap := make(map[int]string)
	for _, item := range itemList {
		resultMap[*item.ID] = *item.Name
	}
	return resultMap, nil
}

func RulesToMap(itemList []qradar.Rule) (map[int]string, error) {
	resultMap := make(map[int]string)
	for _, item := range itemList {
		resultMap[*item.ID] = *item.Name
	}
	return resultMap, nil
}

func BuildingBlocksToMap(itemList []qradar.BuildingBlock) (map[int]string, error) {
	resultMap := make(map[int]string)
	for _, item := range itemList {
		resultMap[*item.ID] = *item.Name
	}
	return resultMap, nil
}

func DomainsToMap(itemList []qradar.Domain) (map[int]string, error) {
	resultMap := make(map[int]string)
	for _, item := range itemList {
		resultMap[*item.ID] = *item.Name
	}
	return resultMap, nil
}

func LogSourceExtensionsToMap(itemList []qradar.LogSourceExtension) (map[int]string, error) {
	resultMap := make(map[int]string)
	for _, item := range itemList {
		resultMap[*item.ID] = *item.Name
	}
	return resultMap, nil
}

func LogSourceGroupsToMap(itemList []qradar.LogSourceGroup) (map[int]string, error) {
	resultMap := make(map[int]string)
	for _, item := range itemList {
		resultMap[*item.ID] = *item.Name
	}
	return resultMap, nil
}

func TenantsToMap(itemList []qradar.Tenant) (map[int]string, error) {
	resultMap := make(map[int]string)
	for _, item := range itemList {
		resultMap[*item.ID] = *item.Name
	}
	return resultMap, nil
}
