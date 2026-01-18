package main

import (
	"context"
	"fmt"

	todoist "github.com/sachaos/todoist/lib"
	"github.com/urfave/cli/v2"
)

func AddSection(c *cli.Context) error {
	client := GetClient(c)

	if !c.Args().Present() {
		return CommandFailed
	}

	name := c.Args().First()

	projectID := c.String("project-id")
	if projectID == "" {
		projectName := c.String("project-name")
		if projectName == "" {
			return fmt.Errorf("project-id or project-name is required")
		}
		projectID = client.Store.Projects.GetIDByName(projectName)
		if projectID == "" {
			return fmt.Errorf("Did not find a project named '%v'", projectName)
		}
	}

	section := todoist.Section{
		Name:         name,
		SectionOrder: c.Int("order"),
	}
	section.ProjectID = projectID

	if err := client.AddSection(context.Background(), section); err != nil {
		return err
	}

	return Sync(c)
}
