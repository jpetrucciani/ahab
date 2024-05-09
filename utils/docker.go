package utils

import (
	"context"
	"io"
	"log"

	"github.com/docker/docker/pkg/stdcopy"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerTailer struct {
	cli   *client.Client
	since string
}

func NewLocalDockerTailer(since string) (Tailer, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}
	return &DockerTailer{
		cli:   cli,
		since: since,
	}, nil
}

func (t *DockerTailer) Tail(c string, writer io.Writer) error {
	logs, err := t.cli.ContainerLogs(context.Background(), c, containertypes.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Since:      t.since,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer logs.Close()

	container, err := t.cli.ContainerInspect(context.Background(), c)
	if err != nil {
		return err
	}

	if container.Config.Tty {
		_, err := io.Copy(writer, logs)
		if err != nil {
			return err
		}
	} else {
		_, err = stdcopy.StdCopy(writer, writer, logs)
		if err != nil {
			return err
		}
	}
	return nil
}

type Tailer interface {
	Tail(container string, writer io.Writer) error
}
