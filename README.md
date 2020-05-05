# QRadar Content Compare Utility

Because the official content migration utility from IBM Qradar is not really reliable, 
I needed a way to verify if the content between two QRadar installations is the same after a migration. 
Because the rest api doesn't cover all content parts, it's not easy to compare everything.
The utility currently checks:
- Tenants
  - Name
  - Description
  - Event Rate Limit
  - Flow Rate Limit
- Domains
  - Name
  - Description
  - assigned Tenant Name
  - assigned Log Source Group Names
- Log Sources
  - Name
  - Description
  - Parent Group Name
  - Child Group Names
- Log Source Groups
  - Name
  - Description
  - Type Name
  - Extension Name
  - Log Source Group Names
  - Enabled Status
  - Credibility
  - Store Event Payload
  - Coalesce Events
- Rules
  - Name
  - Number of Test Definitions
  - Number of Building Blocks in Rule Conditions
- Rule Groups
  - Name
  - Description
  - Parent Name
  - Type
  - Associated Rules
- Network Hierarchy
  - Name
  - Domain Name
  - CIDR
  - Group
- DSM Mappings
  - Log Source Type
  - Log Source Event ID
  - Log Source Event Category
  - QID Name
- QIDs
  - Low Level Category Name
  - Log Source Type Name
  - Severity
  - Description
- Custom Properties
  - Log Source Type
  - Log Source Name
  - Low Level Category
  - QID Name
  - Regex
  - Enabled


