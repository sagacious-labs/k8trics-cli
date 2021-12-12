package plugin

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
	sse "github.com/r3labs/sse/v2"
	"github.com/sagacious-labs/kcli/pkg/k8trics/formatter"
	"github.com/sagacious-labs/kcli/pkg/utils"
)

type Plugin struct {
	baseURL string
}

func New(baseURL string) *Plugin {
	return &Plugin{
		baseURL: utils.ConstructURL(baseURL, "module"),
	}
}

func (p *Plugin) Apply(locations []string) error {
	for _, loc := range locations {
		file, err := os.Open(loc)
		if err != nil {
			return err
		}

		resp, err := http.Post(p.baseURL, "application/json", file)
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
			color.Green("✅ Successfully applied plugin: %s", loc)
		} else {
			color.Red("❌ Failed to apply plugin: ", loc)
		}
	}

	return nil
}

func (p *Plugin) Delete(names []string) error {
	for _, name := range names {
		req, err := http.NewRequest(http.MethodDelete, utils.ConstructURL(p.baseURL, name), nil)
		if err != nil {
			return err
		}

		c := &http.Client{}
		resp, err := c.Do(req)
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {
			color.Green("✅ Successfully deleted plugin: %s", name)
		} else {
			color.Red("❌ Failed to deleted plugin: %s", name)
		}
	}

	return nil
}

func (p *Plugin) Get(name string) error {
	resp, err := http.Get(utils.ConstructURL(p.baseURL, name))
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, resp.Body)

	return nil
}

func (p *Plugin) StreamLog(name string) error {
	client := sse.NewClient(utils.ConstructURL(p.baseURL, name, "log"))

	client.Subscribe("log", func(msg *sse.Event) {
		fmt.Println(string(msg.Data))
	})

	return nil
}

func (p *Plugin) StreamData(name string) error {
	client := sse.NewClient(utils.ConstructURL(p.baseURL, name, "data"))

	client.Subscribe("data", func(msg *sse.Event) {
		formatter.Write(msg.Data, "container_id")
	})

	return nil
}
