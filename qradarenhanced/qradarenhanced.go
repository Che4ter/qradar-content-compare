package qradarenhanced

import (
	"context"
	"encoding/xml"
	"github.com/ilyaglow/go-qradar"
	"qradar-content-compare/converters"
	"qradar-content-compare/types"
	"sort"
	"strconv"
)

func GetPropertiesRegexExpressionResolved(qRadar *qradar.Client) ([]types.PropertyExpressionRegexResolved, error) {
	customProperties, err := qRadar.PropertyExpression.Get(context.Background(), "", "", 0, 0)
	if err != nil {
		return nil, err
	}

	logSourceTypes, err := getLogSourcesTypeMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	logSources, err := getLogSourcesMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	lowLevelCategories, err := getLogLowLevelCategoryMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	qids, err := getQIDsMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	var propertiesResolved []types.PropertyExpressionRegexResolved

	for _, property := range customProperties {
		propertyResolved := types.PropertyExpressionRegexResolved{
			PropertyExpression: property,
		}

		if property.LogSourceTypeID != nil {
			propertyResolved.LogSourceTypeName = logSourceTypes[*property.LogSourceTypeID]
		}
		if property.LogSourceID != nil {
			propertyResolved.LogSourceName = logSources[*property.LogSourceID]
		}
		if property.QID != nil {
			propertyResolved.QidName = qids[*property.QID]
		}
		if property.LowLevelCategoryID != nil {
			propertyResolved.LowLevelCategoryName = lowLevelCategories[*property.LowLevelCategoryID]
		}

		propertiesResolved = append(propertiesResolved, propertyResolved)
	}

	return propertiesResolved, nil
}

func GetRuleGroupsResolved(qRadar *qradar.Client) ([]types.RuleGroupResolved, error) {
	ruleGroups, err := qRadar.RuleGroup.Get(context.Background(), "", "", 0, 0)
	if err != nil {
		return nil, err
	}

	rules, err := getRulesMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	buildingBlocks, err := getBuildingBlocksMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	var ruleGroupsResolved []types.RuleGroupResolved

	for _, ruleGroup := range ruleGroups {
		ruleGroupResolved := types.RuleGroupResolved{
			RuleGroup: ruleGroup,
		}

		for _, parentRuleGroup := range ruleGroups {
			if ruleGroup.ParentID != nil {
				if *ruleGroup.ParentID == *parentRuleGroup.ID {
					ruleGroupResolved.ParentName = *parentRuleGroup.Name
					break
				}
			}
		}

		var ruleName []string
		for _, ruleID := range ruleGroup.ChildItems {
			var id, _ = strconv.Atoi(ruleID)
			var name = rules[id]
			if name == "" {
				name = buildingBlocks[id]
			}
			if name != "" {
				ruleName = append(ruleName, name)
			}
		}
		sort.Strings(ruleName)
		ruleGroupResolved.RuleNamesAssociated = ruleName

		ruleGroupsResolved = append(ruleGroupsResolved, ruleGroupResolved)
	}

	return ruleGroupsResolved, nil
}

func GetNetworkHierarchyResolved(qRadar *qradar.Client) ([]types.NetworkHierarchyResolved, error) {
	networkHierarchies, err := qRadar.NetworkHierarchy.Get(context.Background(), "")
	if err != nil {
		return nil, err
	}

	domains, err := getDomainsMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	var networkHierarchiesResolved []types.NetworkHierarchyResolved
	for _, networkHierarchy := range networkHierarchies {
		networkHierarchyResolved := types.NetworkHierarchyResolved{
			NetworkHierarchy: networkHierarchy,
		}

		if networkHierarchy.DomainID != nil {
			if *networkHierarchy.DomainID == 0 {
				networkHierarchyResolved.DomainName = "Default Domain"
			} else {
				networkHierarchyResolved.DomainName = domains[*networkHierarchy.DomainID]
			}
		}

		networkHierarchiesResolved = append(networkHierarchiesResolved, networkHierarchyResolved)
	}

	return networkHierarchiesResolved, nil
}

func GetQIDsResolved(qRadar *qradar.Client) (map[string]types.QIDsResolved, error) {
	qids, err := qRadar.QID.Get(context.Background(), "", "", 0, 0)
	if err != nil {
		return nil, err
	}

	logSourceTypes, err := getLogSourcesTypeMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	lowLevelCategories, err := getLogLowLevelCategoryMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	qIDsResolved := make(map[string]types.QIDsResolved)

	for _, qid := range qids {
		qidResolved := types.QIDsResolved{
			QID:                  qid,
			LowLevelCategoryName: lowLevelCategories[*qid.LowLevelCategoryID],
		}

		logSourceTypeName := "empty"
		if qid.LogSourceTypeID != nil {
			logSourceTypeName = logSourceTypes[*qid.LogSourceTypeID]
		}
		qidResolved.LogSourceTypeName = logSourceTypeName

		if qid.LogSourceTypeID != nil {
			qidResolved.LowLevelCategoryName = lowLevelCategories[*qid.LowLevelCategoryID]
		}

		qIDsResolved[*qid.Name] = qidResolved
	}

	return qIDsResolved, nil
}

func GetDSMMappingsResolved(qRadar *qradar.Client) (map[string]types.DsmResolved, error) {
	dsms, err := qRadar.DSM.Get(context.Background(), "", "custom_event=true", 0, 0)
	if err != nil {
		return nil, err
	}

	logSourceTypes, err := getLogSourcesTypeMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	qids, err := getQIDsMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	dsmsResolved := make(map[string]types.DsmResolved)

	for _, dsm := range dsms {
		dsmResolved := types.DsmResolved{
			DSM: dsm,
		}

		if dsm.LogSourceTypeID != nil {
			dsmResolved.LogSourceTypeName = logSourceTypes[*dsm.LogSourceTypeID]
		}

		if dsm.QIDRecordID != nil {
			dsmResolved.QidName = qids[*dsm.QIDRecordID]
		}

		eventCategory := ""
		if dsm.LogSourceEventCategory != nil {
			eventCategory = *dsm.LogSourceEventCategory
		}

		logSourceEventID := ""
		if dsm.LogSourceEventID != nil {
			eventCategory = *dsm.LogSourceEventID
		}
		searchString := dsmResolved.LogSourceTypeName + dsmResolved.QidName + eventCategory + logSourceEventID
		dsmsResolved[searchString] = dsmResolved

	}

	return dsmsResolved, nil
}

func GetRulesResolved(qRadar *qradar.Client) ([]types.RulesWithDataResolved, error) {
	rules, err := qRadar.RuleWithData.Get(context.Background(), "", "", 0, 0)
	if err != nil {
		return nil, err
	}
	var rulesResolved []types.RulesWithDataResolved

	for _, rule := range rules {
		ruleXML := types.RuleXML{}
		err = xml.Unmarshal([]byte(*rule.RuleXML), &ruleXML)
		if err != nil {
			return nil, err
		}
		ruleResolved := types.RulesWithDataResolved{
			RuleWithData: rule,
			RuleXML:      ruleXML,
		}
		rulesResolved = append(rulesResolved, ruleResolved)
	}

	return rulesResolved, nil
}

func GetLogSourcesResolved(qRadar *qradar.Client) ([]types.LogSourcesResolved, error) {
	logSources, err := qRadar.LogSource.Get(context.Background(), "", "", 0, 0)
	if err != nil {
		return nil, err
	}

	logSourceExtensions, err := getLogSourceExtensionsMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	logSourceTypes, err := getLogSourcesTypeMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	logSourceGroups, err := getLogSourceGroupsMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	var logSourcesResolved []types.LogSourcesResolved
	for _, logSource := range logSources {
		logSourceResolved := types.LogSourcesResolved{
			LogSource: logSource,
		}

		if logSource.LogSourceExtensionID != nil {
			logSourceResolved.ExtensionName = logSourceExtensions[*logSource.LogSourceExtensionID]
		}

		if logSource.TypeID != nil {
			logSourceResolved.TypeName = logSourceTypes[*logSource.TypeID]
		}

		for _, logSourceGroupId := range logSource.GroupIDs {
			logSourceResolved.LogSourceGroupNames = append(logSourceResolved.LogSourceGroupNames, logSourceGroups[logSourceGroupId])
		}
		sort.Strings(logSourceResolved.LogSourceGroupNames)

		logSourcesResolved = append(logSourcesResolved, logSourceResolved)
	}

	return logSourcesResolved, nil
}

func GetLogSourceGroupsResolved(qRadar *qradar.Client) ([]types.LogSourceGroupsResolved, error) {
	logSourceGroups, err := qRadar.LogSourceGroup.Get(context.Background(), "", "", 0, 0)
	if err != nil {
		return nil, err
	}

	var logSourceGroupsResolved []types.LogSourceGroupsResolved
	for _, logSourceGroup := range logSourceGroups {
		logSourceGroupResolved := types.LogSourceGroupsResolved{
			LogSourceGroup: logSourceGroup,
		}

		for _, logSourceGroupToCompare := range logSourceGroups {
			if *logSourceGroup.ParentID == *logSourceGroupToCompare.ID {
				logSourceGroupResolved.ParentGroupName = *logSourceGroupToCompare.Name
				break
			}
		}
		for _, logSourceGroupChildId := range logSourceGroup.ChildGroupIDs {
			for _, logSourceGroupToCompare := range logSourceGroups {
				if logSourceGroupChildId == *logSourceGroupToCompare.ID {
					logSourceGroupResolved.ChildGroupNames = append(logSourceGroupResolved.ChildGroupNames, *logSourceGroupToCompare.Name)
				}
			}
		}
		sort.Strings(logSourceGroupResolved.ChildGroupNames)

		logSourceGroupsResolved = append(logSourceGroupsResolved, logSourceGroupResolved)
	}

	return logSourceGroupsResolved, nil
}

func GetDomainsResolved(qRadar *qradar.Client) ([]types.DomainResolved, error) {
	domains, err := qRadar.Domain.Get(context.Background(), "", "deleted=false", 0, 0)
	if err != nil {
		return nil, err
	}

	tenants, err := getTenantsMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	logSourceGroups, err := getLogSourceGroupsMinimum(qRadar)
	if err != nil {
		return nil, err
	}

	var domainsResolved []types.DomainResolved
	for _, domain := range domains {
		domainResolved := types.DomainResolved{
			Domain: domain,
		}

		if domain.TenantID != nil {
			domainResolved.TenantName = tenants[*domain.TenantID]
		}

		for _, logSourceGroupId := range domainResolved.LogSourceGroupIds {
			domainResolved.LogSourceGroupNames = append(domainResolved.LogSourceGroupNames, logSourceGroups[logSourceGroupId])
		}
		sort.Strings(domainResolved.LogSourceGroupNames)

		domainsResolved = append(domainsResolved, domainResolved)
	}

	return domainsResolved, nil
}


func getTenantsMinimum(qRadar *qradar.Client) (map[int]string, error) {
	resultItems, err := qRadar.Tenant.Get(context.Background(), "", "deleted=false", 0, 0)
	if err != nil {
		return nil, err
	}

	return converters.TenantsToMap(resultItems)
}
func getLogSourceGroupsMinimum(qRadar *qradar.Client) (map[int]string, error) {
	resultItems, err := qRadar.LogSourceGroup.Get(context.Background(), "name,id", "", 0, 0)
	if err != nil {
		return nil, err
	}

	return converters.LogSourceGroupsToMap(resultItems)
}
func getLogSourceExtensionsMinimum(qRadar *qradar.Client) (map[int]string, error) {
	resultItems, err := qRadar.LogSourceExtension.Get(context.Background(), "name,id", "", 0, 0)
	if err != nil {
		return nil, err
	}

	return converters.LogSourceExtensionsToMap(resultItems)
}
func getDomainsMinimum(qRadar *qradar.Client) (map[int]string, error) {
	resultItems, err := qRadar.Domain.Get(context.Background(), "name,id", "", 0, 0)
	if err != nil {
		return nil, err
	}

	return converters.DomainsToMap(resultItems)
}
func getLogSourcesMinimum(qRadar *qradar.Client) (map[int]string, error) {
	resultItems, err := qRadar.LogSource.Get(context.Background(), "name,id", "", 0, 0)
	if err != nil {
		return nil, err
	}

	return converters.LogSourcesToMap(resultItems)
}
func getLogSourcesTypeMinimum(qRadar *qradar.Client) (map[int]string, error) {
	resultItems, err := qRadar.LogSourceType.Get(context.Background(), "id,name", "", 0, 0)
	if err != nil {
		return nil, err
	}

	return converters.LogSourceTypesToMap(resultItems)
}
func getLogLowLevelCategoryMinimum(qRadar *qradar.Client) (map[int]string, error) {
	resultItems, err := qRadar.LowLevelCategory.Get(context.Background(), "id,name", "", 0, 0)
	if err != nil {
		return nil, err
	}

	return converters.LowLevelCategoriesToMap(resultItems)
}
func getQIDsMinimum(qRadar *qradar.Client) (map[int]string, error) {
	resultItems, err := qRadar.QID.Get(context.Background(), "id,name", "", 0, 0)
	if err != nil {
		return nil, err
	}

	return converters.QIDsToMap(resultItems)
}
func getRulesMinimum(qRadar *qradar.Client) (map[int]string, error) {
	resultItems, err := qRadar.Rule.Get(context.Background(), "id,name", "", 0, 0)
	if err != nil {
		return nil, err
	}

	return converters.RulesToMap(resultItems)
}
func getBuildingBlocksMinimum(qRadar *qradar.Client) (map[int]string, error) {
	resultItems, err := qRadar.BuildingBlock.Get(context.Background(), "id,name", "", 0, 0)
	if err != nil {
		return nil, err
	}

	return converters.BuildingBlocksToMap(resultItems)
}