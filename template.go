package zabbix

// For `TemplateGetParams` field: `Evaltype`
const (
	TemplateEvaltypeAndOr = 0
	TemplateEvaltypeOr    = 2
)

// For `TemplateTag` field: `Operator`
const (
	TemplateTagOperatorContains = 0
	TemplateTagOperatorEquals   = 1
)

// Template struct is used to store template operations results
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/template/object
type Template struct {
	TemplateID  string `json:"templateid,omitempty"`
	Host        string `json:"host,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`

	Groups          []Hostgroup   `json:"groups,omitempty"`
	Tags            []TemplateTag `json:"tags,omitempty"`
	Templates       []Template    `json:"templates,omitempty"`
	ParentTemplates []Template    `json:"parentTemplates,omitempty"`
	Hosts           []Host        `json:"hosts,omitempty"`
}

// TemplateTag struct is used to store template tag data
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/template/object#template_tag
type TemplateTag struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`

	Operator int `json:"operator,omitempty"` // Used for `get` operations, has defined consts, see above
}

// TemplateGetParams struct is used for template get requests
//
// see: https://www.zabbix.com/documentation/5.4/manual/api/reference/template/get#parameters
type TemplateGetParams struct {
	GetParameters

	TemplateIDs       []int `json:"templateids,omitempty"`
	GroupIDs          []int `json:"groupids,omitempty"`
	ParentTemplateIDs []int `json:"parentTemplateids,omitempty"`
	HostIDs           []int `json:"hostids,omitempty"`
	GraphIDs          []int `json:"graphids,omitempty"`
	ItemIDs           []int `json:"itemids,omitempty"`
	TriggerIDs        []int `json:"triggerids,omitempty"`

	WithItems     bool          `json:"with_items,omitempty"`
	WithTriggers  bool          `json:"with_triggers,omitempty"`
	WithGraphs    bool          `json:"with_graphs,omitempty"`
	WithHttptests bool          `json:"with_httptests,omitempty"`
	Evaltype      int           `json:"evaltype,omitempty"` // has defined consts, see above
	Tags          []TemplateTag `json:"tags,omitempty"`

	SelectGroups          SelectQuery `json:"selectGroups,omitempty"`
	SelectTags            SelectQuery `json:"selectTags,omitempty"`
	SelectHosts           SelectQuery `json:"selectHosts,omitempty"`
	SelectTemplates       SelectQuery `json:"selectTemplates,omitempty"`
	SelectParentTemplates SelectQuery `json:"selectParentTemplates,omitempty"`
	SelectMacros          SelectQuery `json:"selectMacros,omitempty"`

	// SelectHttpTests       SelectQuery `json:"selectHttpTests,omitempty"` // not implemented yet
	// SelectItems           SelectQuery `json:"selectItems,omitempty"` // not implemented yet
	// SelectDiscoveries     SelectQuery `json:"selectDiscoveries,omitempty"` // not implemented yet
	// SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"` // not implemented yet
	// SelectGraphs          SelectQuery `json:"selectGraphs,omitempty"` // not implemented yet
	// SelectApplications    SelectQuery `json:"selectApplications,omitempty"` // not implemented yet
	// SelectScreens         SelectQuery `json:"selectScreens,omitempty"` // not implemented yet
}

// GetTemplates queries the Zabbix API for Templates matching the given search
// parameters.
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetTemplates(params TemplateGetParams) ([]Template, error) {
	templates := make([]Template, 0)
	err := c.Get("template.get", params, &templates)
	if err != nil {
		return nil, err
	}

	if len(templates) == 0 {
		return nil, ErrNotFound
	}

	return templates, nil
}
