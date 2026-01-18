package todoist

import "context"

type Section struct {
	HaveID
	HaveProjectID
	Collapsed    bool   `json:"collapsed"`
	Name         string `json:"name"`
	IsArchived   bool   `json:"is_archived"`
	IsDeleted    bool   `json:"is_deleted"`
	SectionOrder int    `json:"section_order"`
}

type Sections []Section

func (section Section) AddParam() interface{} {
	param := map[string]interface{}{}
	if section.Name != "" {
		param["name"] = section.Name
	}
	if section.ProjectID != "" {
		param["project_id"] = section.ProjectID
	}
	if section.SectionOrder != 0 {
		param["order"] = section.SectionOrder
	}
	return param
}

func (c *Client) AddSection(ctx context.Context, section Section) error {
	var res Section
	return c.doRestV2Api(ctx, "POST", "sections", section.AddParam(), &res)
}
