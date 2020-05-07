package types

import (
	"encoding/xml"
	"github.com/ilyaglow/go-qradar"
)

type DifferentTenants struct {
	OldTenant qradar.Tenant
	NewTenant qradar.Tenant
}

type DifferentDomains struct {
	OldDomains DomainResolved
	NewDomains DomainResolved
}

type DomainResolved struct {
	qradar.Domain
	TenantName          string
	LogSourceGroupNames []string
}

type LogSourceGroupsResolved struct {
	qradar.LogSourceGroup
	ParentGroupName string
	ChildGroupNames []string
}

type DifferentLogSourceGroups struct {
	OldLogSourceGroups LogSourceGroupsResolved
	NewLogSourceGroups LogSourceGroupsResolved
}

type LogSourcesResolved struct {
	qradar.LogSource
	ExtensionName       string
	TypeName            string
	LogSourceGroupNames []string
}

type DifferentLogSources struct {
	OldLogSources LogSourcesResolved
	NewLogSources LogSourcesResolved
}

type DifferentRulesWithData struct {
	OldRulesWithData RulesWithDataResolved
	NewRulesWithData RulesWithDataResolved
	HasDifferentBB   bool
}

type DifferentDSMs struct {
	OldDSMResolved DsmResolved
	NewDSMResolved DsmResolved
}

type RulesWithDataResolved struct {
	qradar.RuleWithData
	RuleXML
}

type DsmResolved struct {
	qradar.DSM
	LogSourceTypeName string
	QidName           string
}

type QIDsResolved struct {
	qradar.QID
	LowLevelCategoryName string
	LogSourceTypeName    string
}

type DifferentQIDs struct {
	OldQIDResolved QIDsResolved
	NewQIDResolved QIDsResolved
}

type DifferentNetworkHierarchy struct {
	OldNetworkHierarchy NetworkHierarchyResolved
	NewNetworkHierarchy NetworkHierarchyResolved
}

type NetworkHierarchyResolved struct {
	qradar.NetworkHierarchy
	DomainName string
}

type RuleGroupResolved struct {
	qradar.RuleGroup
	ParentName          string
	RuleNamesAssociated []string
}

type DifferentRuleGroupsResolved struct {
	OldRuleGroupsResolved RuleGroupResolved
	NewRuleGroupsResolved RuleGroupResolved
}

type RuleXML struct {
	XMLName         xml.Name        `xml:"rule"`
	Text            string          `xml:",chardata"`
	Overrideid      int             `xml:"overrideid,attr"`
	Owner           string          `xml:"owner,attr"`
	Scope           string          `xml:"scope,attr"`
	Type            string          `xml:"type,attr"`
	RoleDefinition  bool            `xml:"roleDefinition,attr"`
	BuildingBlock   bool            `xml:"buildingBlock,attr"`
	Enabled         bool            `xml:"enabled,attr"`
	ID              int             `xml:"id,attr"`
	Name            string          `xml:"name"`
	Notes           string          `xml:"notes"`
	TestDefinitions TestDefinitions `xml:"testDefinitions"`
	Actions         struct {
		Text                          string `xml:",chardata"`
		FlowAnalysisInterval          string `xml:"flowAnalysisInterval,attr"`
		IncludeAttackerEventsInterval string `xml:"includeAttackerEventsInterval,attr"`
		ForceOffenseCreation          string `xml:"forceOffenseCreation,attr"`
		OffenseMapping                string `xml:"offenseMapping,attr"`
	} `xml:"actions"`
	Responses struct {
		Text                     string `xml:",chardata"`
		ReferenceTableRemove     bool   `xml:"referenceTableRemove,attr"`
		ReferenceMapOfMapsRemove bool   `xml:"referenceMapOfMapsRemove,attr"`
		ReferenceMapOfSetsRemove bool   `xml:"referenceMapOfSetsRemove,attr"`
		ReferenceMapRemove       bool   `xml:"referenceMapRemove,attr"`
		ReferenceTable           bool   `xml:"referenceTable,attr"`
		ReferenceMapOfMaps       bool   `xml:"referenceMapOfMaps,attr"`
		ReferenceMapOfSets       bool   `xml:"referenceMapOfSets,attr"`
		ReferenceMap             bool   `xml:"referenceMap,attr"`
		Newevent                 struct {
			Text                  string `xml:",chardata"`
			LowLevelCategory      string `xml:"lowLevelCategory,attr"`
			OffenseMapping        string `xml:"offenseMapping,attr"`
			ForceOffenseCreation  bool   `xml:"forceOffenseCreation,attr"`
			Qid                   int    `xml:"qid,attr"`
			ContributeOffenseName bool   `xml:"contributeOffenseName,attr"`
			OverrideOffenseName   bool   `xml:"overrideOffenseName,attr"`
			DescribeOffense       bool   `xml:"describeOffense,attr"`
			Relevance             string `xml:"relevance,attr"`
			Credibility           string `xml:"credibility,attr"`
			Severity              string `xml:"severity,attr"`
			Description           string `xml:"description,attr"`
			Name                  string `xml:"name,attr"`
		} `xml:"newevent"`
	} `xml:"responses"`
}

type TestDefinitions struct {
	Text string     `xml:",text"`
	Test []RuleTest `xml:"test"`
}

type RuleTest struct {
	RequiredCapabilities string `xml:"requiredCapabilities,attr"`
	Group                string `xml:"group,attr"`
	Uid                  int    `xml:"uid,attr"`
	Name                 string `xml:"name,attr"`
	ID                   int    `xml:"id,attr"`
	GroupId              int    `xml:"groupId,attr"`
	Negate               string `xml:"negate,attr"`
	Text                 string `xml:"text"`
	Parameter            []struct {
		Text           string `xml:",chardata"`
		ID             int    `xml:"id,attr"`
		InitialText    string `xml:"initialText"`
		SelectionLabel string `xml:"selectionLabel"`
		UserOptions    struct {
			Text        string `xml:",chardata"`
			Multiselect bool   `xml:"multiselect,attr"`
			Method      string `xml:"method,attr"`
			Source      string `xml:"source,attr"`
			Format      string `xml:"format,attr"`
			Errorkey    string `xml:"errorkey,attr"`
			Validation  string `xml:"validation,attr"`
			Option      []struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"option"`
		} `xml:"userOptions"`
		UserSelection      string `xml:"userSelection"`
		UserSelectionTypes string `xml:"userSelectionTypes"`
		UserSelectionId    int    `xml:"userSelectionId"`
		Name               string `xml:"name"`
	} `xml:"parameter"`
}

type PropertyExpressionRegexResolved struct {
	qradar.PropertyExpression
	LogSourceTypeName    string
	QidName              string
	LogSourceName        string
	LowLevelCategoryName string
}

type DifferentPropertyExpressionRegex struct {
	OldPropertyExpressionRegexResolved PropertyExpressionRegexResolved
	NewPropertyExpressionRegexResolved PropertyExpressionRegexResolved
}

type Report struct {
	ElementType      string
	SameCount        int
	OldCount		 int
	NewCount		 int
	MissingRecords   []string
	DifferentRecords []DifferentRecord
}

type DifferentRecord struct {
	RecordName        string
	DifferentElements []DifferentElement
}

type DifferentElement struct {
	Name     string
	OldValue string
	NewValue string
}
