package main

import (
	"fmt"
	"os"
	"strings"

	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type MainObj struct {
	commands []string
	cli      *client.Client
	img      types.ImageSummary
}

func (m *MainObj) printCommands() {
	for i := len(m.commands) - 1; i >= 0; i-- {
		fmt.Println(m.commands[i])
	}
}

func (m *MainObj) getImage(repoTagOrID string) error {
	images, err := m.cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return err
	}

	for _, i := range images {
		if strings.Contains(i.ID, repoTagOrID) {
			m.img = i
			return nil
		}

		for _, rt := range i.RepoTags {
			if rt == repoTagOrID {
				m.img = i
				return nil
			}
		}
	}

	return fmt.Errorf("image %s Not Found", repoTagOrID)
}

func (m *MainObj) insertStep(step string) {
	toAdd := strings.ReplaceAll(step, "&&", "\\\n    &&")
	if strings.Contains(toAdd, "#(nop)") {
		toAdd = strings.Split(toAdd, "#(nop) ")[1]
	} else if strings.Contains(toAdd, "CMD") {
		toAdd = toAdd
	} else {
		toAdd = "RUN " + toAdd
	}
	m.commands = append(m.commands, strings.TrimSpace(toAdd))
}

func (m *MainObj) parseHistory() error {
	hist, err := m.cli.ImageHistory(context.Background(), m.img.ID)
	if err != nil {
		return err
	}

	for _, h := range hist {
		m.insertStep(h.CreatedBy)
	}
	//return nil

	firstTag := false
	var actualTag string
	for _, layer := range hist {
		if len(layer.Tags) > 0 {
			actualTag = layer.Tags[0]
			if firstTag {
				break
			}
			firstTag = true
		}
	}
	if firstTag {
		m.commands = append(m.commands, "FROM "+actualTag)
	}
	return nil

}

func main() {
	m := MainObj{}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
		return
	}
	m.cli = cli

	imageID := os.Args[len(os.Args)-1]
	err = m.getImage(imageID)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = m.parseHistory()
	if err != nil {
		fmt.Println(err)
		return
	}

	m.printCommands()
}
