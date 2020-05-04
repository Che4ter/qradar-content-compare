package comparator

import (
	"context"
	"fmt"
	"github.com/ilyaglow/go-qradar"
	"qradar-content-compare/qradarenhanced"
	"qradar-content-compare/types"
	"sort"
	"strconv"
	"strings"
)

func CompareTenants(oldQRadar *qradar.Client, newQRadar *qradar.Client) (types.Report, error) {
	oldContent, err := oldQRadar.Tenant.Get(context.Background(), "", "deleted=false", 0, 0)
	if err != nil {
		return types.Report{}, err
	}

	newContent, err := newQRadar.Tenant.Get(context.Background(), "", "deleted=false", 0, 0)
	if err != nil {
		return types.Report{}, err
	}

	var sameCount = 0
	var elementExists = false
	var report = types.Report{}
	report.ElementType = "Tenants"

	for _, oldItem := range oldContent {
		itemName := fmt.Sprintf("Name: %s", *oldItem.Name)
		elementExists = false

		for _, newItem := range newContent {
			if *oldItem.Name == *newItem.Name {
				elementExists = true

				var different = types.DifferentRecord{}

				if *oldItem.Description != *newItem.Description {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Description",
						OldValue: *oldItem.Description,
						NewValue: *newItem.Description,
					})
				}

				if *oldItem.EventRateLimit != *newItem.EventRateLimit {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Event Rate Limit",
						OldValue: strconv.Itoa(*oldItem.EventRateLimit),
						NewValue: strconv.Itoa(*newItem.EventRateLimit),
					})
				}
				if *oldItem.FlowRateLimit != *newItem.FlowRateLimit {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Flow Rate Limit",
						OldValue: strconv.Itoa(*oldItem.FlowRateLimit),
						NewValue: strconv.Itoa(*newItem.FlowRateLimit),
					})
				}
				if len(different.DifferentElements) > 0 {
					different.RecordName = itemName
					report.DifferentRecords = append(report.DifferentRecords, different)
				} else {
					sameCount++
				}
				break
			}
		}
		if !elementExists {
			report.MissingRecords = append(report.MissingRecords, itemName)
		}
	}
	report.SameCount = sameCount

	return report, nil
}

func CompareDomains(oldQRadar *qradar.Client, newQRadar *qradar.Client) (types.Report, error) {
	oldContent, err := qradarenhanced.GetDomainsResolved(oldQRadar)
	if err != nil {
		return types.Report{}, err
	}

	newContent, err := qradarenhanced.GetDomainsResolved(newQRadar)
	if err != nil {
		return types.Report{}, err
	}

	var sameCount = 0
	var elementExists = false
	var report = types.Report{}
	report.ElementType = "Domains"

	for _, oldItem := range oldContent {
		itemName := fmt.Sprintf("Name: %s (%s)", *oldItem.Name, *oldItem.Description)
		elementExists = false

		for _, newItem := range newContent {
			if *oldItem.Name == *newItem.Name {
				elementExists = true

				var different = types.DifferentRecord{}
				if *oldItem.Description != *newItem.Description {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Description",
						OldValue: *oldItem.Description,
						NewValue: *newItem.Description,
					})
				}
				if oldItem.TenantName != newItem.TenantName {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Tenant Name",
						OldValue: oldItem.TenantName,
						NewValue: newItem.TenantName,
					})
				}
				if !stringSliceEqual(oldItem.LogSourceGroupNames, newItem.LogSourceGroupNames) {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Log Source Group Names",
						OldValue: strings.Join(oldItem.LogSourceGroupNames, ", "),
						NewValue: strings.Join(newItem.LogSourceGroupNames, ", "),
					})
				}

				if len(different.DifferentElements) > 0 {
					different.RecordName = itemName
					report.DifferentRecords = append(report.DifferentRecords, different)
				} else {
					sameCount++
				}
				break
			}
		}
		if !elementExists {
			report.MissingRecords = append(report.MissingRecords, itemName)
		}
	}
	report.SameCount = sameCount

	return report, nil
}

func CompareLogSourceGroups(oldQRadar *qradar.Client, newQRadar *qradar.Client) (types.Report, error) {
	oldContent, err := qradarenhanced.GetLogSourceGroupsResolved(oldQRadar)
	if err != nil {
		return types.Report{}, err
	}

	newContent, err := qradarenhanced.GetLogSourceGroupsResolved(newQRadar)
	if err != nil {
		return types.Report{}, err
	}

	var sameCount = 0
	var elementExists = false
	var report = types.Report{}
	report.ElementType = "Log Source Goups"

	for _, oldItem := range oldContent {
		itemName := fmt.Sprintf("Group Name: %s (Parent: %s)", *oldItem.Name, oldItem.ParentGroupName)
		elementExists = false
		for _, newItem := range newContent {
			if *oldItem.Name == *newItem.Name && (oldItem.ParentGroupName == newItem.ParentGroupName || *oldItem.ParentID == 1) {
				elementExists = true
				var different = types.DifferentRecord{}

				if *oldItem.Description != *newItem.Description {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Description",
						OldValue: *oldItem.Description,
						NewValue: *newItem.Description,
					})
				}

				if oldItem.ParentGroupName != newItem.ParentGroupName {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Parent Group Name",
						OldValue: oldItem.ParentGroupName,
						NewValue: newItem.ParentGroupName,
					})
				}

				if !stringSliceEqual(oldItem.ChildGroupNames, newItem.ChildGroupNames) {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Child Group Names",
						OldValue: strings.Join(oldItem.ChildGroupNames, ", "),
						NewValue: strings.Join(newItem.ChildGroupNames, ", "),
					})
				}

				if len(different.DifferentElements) > 0 {
					different.RecordName = itemName
					report.DifferentRecords = append(report.DifferentRecords, different)
				} else {
					sameCount++
				}
				break
			}
		}
		if !elementExists {
			report.MissingRecords = append(report.MissingRecords, itemName)
		}
	}
	report.SameCount = sameCount

	return report, nil
}

func CompareLogSources(oldQRadar *qradar.Client, newQRadar *qradar.Client) (types.Report, error) {
	oldContent, err := qradarenhanced.GetLogSourcesResolved(oldQRadar)
	if err != nil {
		return types.Report{}, err
	}

	newContent, err := qradarenhanced.GetLogSourcesResolved(newQRadar)
	if err != nil {
		return types.Report{}, err
	}

	var sameCount = 0
	var elementExists = false
	var report = types.Report{}
	report.ElementType = "Log Sources"

	for _, oldItem := range oldContent {
		itemName := fmt.Sprintf("Name: %s", *oldItem.Name)

		elementExists = false
		for _, newItem := range newContent {
			if *oldItem.Name == *newItem.Name {
				elementExists = true

				var different = types.DifferentRecord{}
				if *oldItem.Description != *newItem.Description {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Description",
						OldValue: *oldItem.Description,
						NewValue: *newItem.Description,
					})
				}
				if oldItem.TypeName != newItem.TypeName {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Type Name",
						OldValue: oldItem.TypeName,
						NewValue: newItem.TypeName,
					})
				}
				if oldItem.ExtensionName != newItem.ExtensionName {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Extension Name",
						OldValue: oldItem.ExtensionName,
						NewValue: newItem.ExtensionName,
					})
				}
				if !stringSliceEqual(oldItem.LogSourceGroupNames, newItem.LogSourceGroupNames) {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Log Source Group Names",
						OldValue: strings.Join(oldItem.LogSourceGroupNames, ", "),
						NewValue: strings.Join(newItem.LogSourceGroupNames, ", "),
					})
				}
				if *oldItem.Enabled != *newItem.Enabled {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Enabled",
						OldValue: strconv.FormatBool(*oldItem.Enabled),
						NewValue: strconv.FormatBool(*newItem.Enabled),
					})
				}
				if *oldItem.Credibility != *newItem.Credibility {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Credibility",
						OldValue: strconv.Itoa(*oldItem.Credibility),
						NewValue: strconv.Itoa(*newItem.Credibility),
					})
				}
				if *oldItem.StoreEventPayload != *newItem.StoreEventPayload {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Store Event Payload",
						OldValue: strconv.FormatBool(*oldItem.StoreEventPayload),
						NewValue: strconv.FormatBool(*newItem.StoreEventPayload),
					})
				}
				if *oldItem.CoalesceEvents != *newItem.CoalesceEvents {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Coalesce Events",
						OldValue: strconv.FormatBool(*oldItem.CoalesceEvents),
						NewValue: strconv.FormatBool(*newItem.CoalesceEvents),
					})
				}
				if len(different.DifferentElements) > 0 {
					different.RecordName = itemName
					report.DifferentRecords = append(report.DifferentRecords, different)
				} else {
					sameCount++
				}
				break
			}
		}
		if !elementExists {
			report.MissingRecords = append(report.MissingRecords, itemName)
		}
	}
	report.SameCount = sameCount

	return report, nil
}

func CompareRules(oldQRadar *qradar.Client, newQRadar *qradar.Client) (types.Report, error) {
	oldContent, err := qradarenhanced.GetRulesResolved(oldQRadar)
	if err != nil {
		return types.Report{}, err
	}

	newContent, err := qradarenhanced.GetRulesResolved(newQRadar)
	if err != nil {
		return types.Report{}, err
	}

	var sameCount = 0
	var elementExists = false
	var report = types.Report{}
	report.ElementType = "Rules"

	for _, oldItem := range oldContent {
		itemName := fmt.Sprintf("Rule Name: %s", oldItem.Name)

		elementExists = false
		for _, newItem := range newContent {
			if oldItem.Name == newItem.Name {
				elementExists = true

				var different = types.DifferentRecord{}
				if len(oldItem.TestDefinitions.Test) == len(newItem.TestDefinitions.Test) {
					for _, testOld := range oldItem.TestDefinitions.Test {
						if testOld.Name == "com.q1labs.semsources.cre.tests.RuleMatch_Test" {
							for _, testNew := range newItem.TestDefinitions.Test {
								if testNew.Name == "com.q1labs.semsources.cre.tests.RuleMatch_Test" && testOld.Uid == testNew.Uid {
									if testOld.Parameter[1].UserSelection != testNew.Parameter[1].UserSelection {
										different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
											Name:     "Has different Number Building Blocks",
											OldValue: strconv.Itoa(len(strings.Split(testOld.Parameter[1].UserSelection, ", "))),
											NewValue: strconv.Itoa(len(strings.Split(testNew.Parameter[1].UserSelection, ", "))),
										})
									}
									break
								}
							}
						}
					}
				} else {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Has Different Number of Conditions",
						OldValue: strconv.Itoa(len(oldItem.TestDefinitions.Test)),
						NewValue: strconv.Itoa(len(newItem.TestDefinitions.Test)),
					})
				}
				if len(different.DifferentElements) > 0 {
					different.RecordName = itemName
					report.DifferentRecords = append(report.DifferentRecords, different)
				} else {
					sameCount++
				}
				break
			}
		}
		if !elementExists {
			report.MissingRecords = append(report.MissingRecords, itemName)
		}
	}
	report.SameCount = sameCount

	return report, nil
}

func CompareDSMMappings(oldQRadar *qradar.Client, newQRadar *qradar.Client) (types.Report, error) {
	oldContent, err := qradarenhanced.GetDSMMappingsResolved(oldQRadar)
	if err != nil {
		return types.Report{}, err
	}

	newContent, err := qradarenhanced.GetDSMMappingsResolved(newQRadar)
	if err != nil {
		return types.Report{}, err
	}

	var sameCount = 0
	var report = types.Report{}
	report.ElementType = "DSM Mappings"

	for searchString, oldItem := range oldContent {
		if _, ok := newContent[searchString]; !ok {
			var missingInformation = ""
			missingInformation += fmt.Sprintf("Log Source Type: %s\n", oldItem.LogSourceTypeName)
			missingInformation += fmt.Sprintf("Log Source Event ID: %s\n", *oldItem.LogSourceEventID)
			missingInformation += fmt.Sprintf("Log Source Event Category: %s\n", *oldItem.LogSourceEventCategory)
			missingInformation += fmt.Sprintf("QID Name: %s\n", oldItem.QidName)
			missingInformation += fmt.Sprintf("Is Custom Mapping: %s\n", strconv.FormatBool(*oldItem.CustomEvent))

			report.MissingRecords = append(report.MissingRecords, missingInformation)
		} else {
			sameCount++
		}
	}

	report.SameCount = sameCount

	return report, nil
}

func CompareQidMappings(oldQRadar *qradar.Client, newQRadar *qradar.Client) (types.Report, error) {
	oldContent, err := qradarenhanced.GetQIDsResolved(oldQRadar)
	if err != nil {
		return types.Report{}, err
	}

	newContent, err := qradarenhanced.GetQIDsResolved(newQRadar)
	if err != nil {
		return types.Report{}, err
	}

	var sameCount = 0
	var elementExists = false
	var report = types.Report{}
	report.ElementType = "QID Mappings"

	for searchString, oldItem := range oldContent {
		itemName := fmt.Sprintf("%s", *oldItem.Name)

		if newItem, ok := newContent[searchString]; ok {
			elementExists = false

			var different = types.DifferentRecord{}
			if oldItem.LowLevelCategoryName != newItem.LowLevelCategoryName {
				different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
					Name:     "Low Level Category Name",
					OldValue: oldItem.LowLevelCategoryName,
					NewValue: newItem.LowLevelCategoryName,
				})
			}
			if oldItem.LogSourceTypeName != newItem.LogSourceTypeName {
				different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
					Name:     "Log Source Type Name",
					OldValue: oldItem.LogSourceTypeName,
					NewValue: newItem.LogSourceTypeName,
				})
			}
			if *oldItem.Severity != *newItem.Severity {
				different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
					Name:     "Severity",
					OldValue: strconv.Itoa(*oldItem.Severity),
					NewValue: strconv.Itoa(*newItem.Severity),
				})
			}
			if *oldItem.Description != *newItem.Description {
				different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
					Name:     "Description",
					OldValue: *oldItem.Description,
					NewValue: *newItem.Description,
				})
			}
			if len(different.DifferentElements) > 0 {
				different.RecordName = itemName
				report.DifferentRecords = append(report.DifferentRecords, different)
			} else {
				sameCount++
			}
		}

		if !elementExists {
			report.MissingRecords = append(report.MissingRecords, itemName)
		}
	}
	report.SameCount = sameCount

	return report, nil
}

func CompareNetworkHierarchy(oldQRadar *qradar.Client, newQRadar *qradar.Client) (types.Report, error) {
	oldContent, err := qradarenhanced.GetNetworkHierarchyResolved(oldQRadar)
	if err != nil {
		return types.Report{}, err
	}

	newContent, err := qradarenhanced.GetNetworkHierarchyResolved(newQRadar)
	if err != nil {
		return types.Report{}, err
	}

	var sameCount = 0
	var elementExists = false
	var report = types.Report{}
	report.ElementType = "Network Hierarchy"

	for _, oldItem := range oldContent {
		itemName := fmt.Sprintf("Name: %s\nCidr: %s\nGroup: %s\nDomain: %s\n", *oldItem.Name, *oldItem.Cidr, *oldItem.Group, oldItem.DomainName)
		elementExists = false

		for _, newItem := range newContent {
			if oldItem.DomainName == newItem.DomainName &&
				*oldItem.Name == *newItem.Name &&
				*oldItem.Cidr == *newItem.Cidr {
				elementExists = true

				var different = types.DifferentRecord{}
				if *oldItem.Description != *newItem.Description {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Description",
						OldValue: *oldItem.Description,
						NewValue: *newItem.Description,
					})
				}
				if *oldItem.Group != *newItem.Group {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Group",
						OldValue: *oldItem.Group,
						NewValue: *newItem.Group,
					})
				}
				if len(different.DifferentElements) > 0 {
					different.RecordName = itemName
					report.DifferentRecords = append(report.DifferentRecords, different)
				} else {
					sameCount++
				}
				break
			}
		}
		if !elementExists {
			report.MissingRecords = append(report.MissingRecords, itemName)
		}
	}
	report.SameCount = sameCount

	return report, nil
}

func CompareRuleGroups(oldQRadar *qradar.Client, newQRadar *qradar.Client) (types.Report, error) {
	oldContent, err := qradarenhanced.GetRuleGroupsResolved(oldQRadar)
	if err != nil {
		return types.Report{}, err
	}

	newContent, err := qradarenhanced.GetRuleGroupsResolved(newQRadar)
	if err != nil {
		return types.Report{}, err
	}

	var sameCount = 0
	var elementExists = false
	var report = types.Report{}
	report.ElementType = "Rule Groups"

	for _, oldItem := range oldContent {
		itemName := fmt.Sprintf("Name: %s (Parent Name: %s)", *oldItem.Name, oldItem.ParentName)
		elementExists = false
		for _, newItem := range newContent {
			if *oldItem.Name == *newItem.Name {
				elementExists = true

				var different = types.DifferentRecord{}
				if *oldItem.Description != *newItem.Description {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Description",
						OldValue: *oldItem.Description,
						NewValue: *newItem.Description,
					})
				}
				if oldItem.ParentName != newItem.ParentName {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Parent Name",
						OldValue: oldItem.ParentName,
						NewValue: newItem.ParentName,
					})
				}
				if *oldItem.Type != *newItem.Type {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Type",
						OldValue: *oldItem.Type,
						NewValue: *newItem.Type,
					})
				}
				missingInOld, missingInNew, isEquals := listCompare(oldItem.RuleNamesAssociated, newItem.RuleNamesAssociated)
				if !isEquals {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Missing Associated Rules (missing in Old/New)",
						OldValue: strings.Join(missingInOld, "\n"),
						NewValue: strings.Join(missingInNew, "\n"),
					})
				}

				if len(different.DifferentElements) > 0 {
					different.RecordName = itemName
					report.DifferentRecords = append(report.DifferentRecords, different)
				} else {
					sameCount++
				}
				break
			}
		}
		if !elementExists {
			report.MissingRecords = append(report.MissingRecords, itemName)
		}
	}
	report.SameCount = sameCount

	return report, nil
}

func CompareCustomProperties(oldQRadar *qradar.Client, newQRadar *qradar.Client) (types.Report, error) {
	oldContent, err := qradarenhanced.GetPropertiesRegexExpressionResolved(oldQRadar)
	if err != nil {
		return types.Report{}, err
	}

	newContent, err := qradarenhanced.GetPropertiesRegexExpressionResolved(newQRadar)
	if err != nil {
		return types.Report{}, err
	}

	var sameCount = 0
	var elementExists = false
	var report = types.Report{}
	report.ElementType = "Custom Properties"
	for _, oldItem := range oldContent {
		itemName := fmt.Sprintf("Identifier: %s (%s)", *oldItem.Identifier, *oldItem.Regex)
		elementExists = false
		for _, newItem := range newContent {
			if *oldItem.Identifier == *newItem.Identifier {
				elementExists = true

				var different = types.DifferentRecord{}
				if oldItem.LogSourceTypeName != newItem.LogSourceTypeName {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Log Source Type",
						OldValue: oldItem.LogSourceTypeName,
						NewValue: newItem.LogSourceTypeName,
					})
				}
				if oldItem.LogSourceName != newItem.LogSourceName {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Log Source",
						OldValue: oldItem.LogSourceName,
						NewValue: newItem.LogSourceName,
					})
				}
				if oldItem.LowLevelCategoryName != newItem.LowLevelCategoryName {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Low Level Category",
						OldValue: oldItem.LowLevelCategoryName,
						NewValue: newItem.LowLevelCategoryName,
					})
				}
				if oldItem.QidName != newItem.QidName {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "QID",
						OldValue: oldItem.QidName,
						NewValue: newItem.QidName,
					})
				}
				if *oldItem.Regex != *newItem.Regex {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Regex",
						OldValue: *oldItem.Regex,
						NewValue: *newItem.Regex,
					})
				}
				if *oldItem.Enabled != *newItem.Enabled {
					different.DifferentElements = append(different.DifferentElements, types.DifferentElement{
						Name:     "Enabled",
						OldValue: strconv.FormatBool(*oldItem.Enabled),
						NewValue: strconv.FormatBool(*newItem.Enabled),
					})
				}

				if len(different.DifferentElements) > 0 {
					different.RecordName = itemName
					report.DifferentRecords = append(report.DifferentRecords, different)
				} else {
					sameCount++
				}
				break
			}
		}
		if !elementExists {
			report.MissingRecords = append(report.MissingRecords, itemName)
		}
	}
	report.SameCount = sameCount

	return report, nil
}

func listCompare(oldList, newList []string) ([]string, []string, bool) {
	sort.Strings(oldList)
	sort.Strings(newList)

	if strings.Join(oldList, "") == strings.Join(newList, "") {
		return nil, nil, true
	}
	var missingInOld []string
	var missingInNew []string

	for _, oldItem := range oldList {
		exists := false
		for _, newItem := range newList {
			if oldItem == newItem {
				exists = true
				break
			}
		}
		if !exists {
			missingInNew = append(missingInNew, oldItem)
		}
	}

	for _, newItem := range newList {
		exists := false
		for _, oldItem := range oldList {
			if newItem == oldItem {
				exists = true
				break
			}
		}
		if !exists {
			missingInOld = append(missingInOld, newItem)
		}
	}

	return missingInOld, missingInNew, false
}
func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
