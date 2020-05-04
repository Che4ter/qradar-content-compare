package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/ilyaglow/go-qradar"
	"log"
	"qradar-content-compare/comparator"
	"qradar-content-compare/questions"
	"qradar-content-compare/reporting"
	"qradar-content-compare/types"
)

var Version = ""

var reportTypes = []string{"Tenants", "Domains", "Log Sources",
	"Log Source Groups", "Rules", "Rule Groups",
	"Network Hierarchy", "DSM Mappings", "QIDs", "Custom Properties"}

func main() {
	fmt.Println("Welcome to QRadar Content Compare (Version " + Version + ")")

	baseUrlOldQRadar, securityTokenOldQRadar, baseUrlNewQRadar, securityTokenNewQRadar, err := questions.AskForConnectionDetails()

	oldQradar, err := qradar.NewClient(
		baseUrlOldQRadar,
		qradar.SetSECKey(securityTokenOldQRadar),
	)
	if err != nil {
		log.Fatal(err)
	}

	newQradar, err := qradar.NewClient(
		baseUrlNewQRadar,
		qradar.SetSECKey(securityTokenNewQRadar),
	)
	if err != nil {
		log.Fatal(err)
	}

	fullReport := false
	prompt := &survey.Confirm{
		Message: "Do you want to run a full report?",
	}
	survey.AskOne(prompt, &fullReport)

	answers := reportTypes
	if !fullReport{
		answers, err = questions.ReportSelection(reportTypes)
		if err != nil {
			log.Fatal(err)
		}
	}

	var reports []types.Report

	for _, answer := range answers {
		switch answer {
		case "Tenants":
			fmt.Println("compare tenants...")
			tenantReport, err := comparator.CompareTenants(oldQradar, newQradar)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, tenantReport)
		case "Domains":
			fmt.Println("compare doamins...")
			domainReport, err := comparator.CompareDomains(oldQradar, newQradar)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, domainReport)
		case "Log Sources":
			fmt.Println("compare log sources...")
			logSourceReport, err := comparator.CompareLogSources(oldQradar, newQradar)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, logSourceReport)
		case "Log Source Groups":
			fmt.Println("compare log source groups...")
			logSourceGroupReport, err := comparator.CompareLogSourceGroups(oldQradar, newQradar)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, logSourceGroupReport)
		case "Rules":
			fmt.Println("compare rules...")
			ruleReport, err := comparator.CompareRules(oldQradar, newQradar)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, ruleReport)
		case "Rule Groups":
			fmt.Println("compare rule groups...")
			ruleGroupReport, err := comparator.CompareRuleGroups(oldQradar, newQradar)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, ruleGroupReport)
		case "Network Hierarchy":
			fmt.Println("compare network hierarchy...")
			networkHierarchyReport, err := comparator.CompareNetworkHierarchy(oldQradar, newQradar)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, networkHierarchyReport)

		case "DSM Mappings":
			fmt.Println("compare dsm mappings...")
			dsmMappingReport, err := comparator.CompareDSMMappings(oldQradar, newQradar)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, dsmMappingReport)
		case "QIDs":
			fmt.Println("compare qids...")
			qidReport, err := comparator.CompareQidMappings(oldQradar, newQradar)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, qidReport)
		case "Custom Properties":
			fmt.Println("compare custom properties...")
			customPropertyReport, err := comparator.CompareCustomProperties(oldQradar, newQradar)
			if err != nil {
				log.Fatal(err)
			}
			reports = append(reports, customPropertyReport)
		default:
			log.Fatal("Selection Type not implemented yet")
		}
	}

	reporting.ReportToFiles(reports)
}
