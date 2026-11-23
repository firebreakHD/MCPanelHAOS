package docker

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	models "github.com/nickheyer/discopanel/internal/db"
)

const HAOSStopTimeout = 120 * time.Second

const haosLabel = "io.mcpanelhaos.installation"

func HAOSMode() bool { return os.Getenv("MCPANEL_HAOS") == "true" }

func HAOSMemory(memory int) int { return max(5120, min(7168, memory)) }

func haosEnvironment(env []string, memory int) []string {
	filtered := make([]string, 0, len(env)+4)
	for _, value := range env {
		name, _, _ := strings.Cut(value, "=")
		switch name {
		case "MEMORY", "INIT_MEMORY", "MAX_MEMORY", "SERVER_PORT", "STOP_DURATION":
			continue
		}
		filtered = append(filtered, value)
	}
	heap := strconv.Itoa(HAOSMemory(memory)-1024) + "M"
	return append(filtered, "INIT_MEMORY="+heap, "MAX_MEMORY="+heap, "SERVER_PORT=25565", "STOP_DURATION=105")
}

type HAOSBusyError struct {
	ServerID string
	Name     string
}

func (e *HAOSBusyError) Error() string {
	return fmt.Sprintf("Server %q läuft bereits.", e.Name)
}

var ErrHAOSSwitchChanged = errors.New("Der aktive Server hat sich geändert. Bitte versuche es erneut.")

// SwitchHAOSContainer permits a switch only from the server named in the
// user's confirmation. Automated callers always take the strict start path.
func (c *Client) SwitchHAOSContainer(ctx context.Context, targetID, fromServerID string) error {
	if !HAOSMode() || fromServerID == "" {
		return ErrHAOSSwitchChanged
	}
	return c.startContainer(ctx, targetID, fromServerID)
}

func (c *Client) prepareHAOSStart(ctx context.Context, targetID, fromServerID string) error {
	target, err := c.docker.ContainerInspect(ctx, targetID)
	if err != nil {
		return err
	}
	installation := os.Getenv("MCPANEL_HAOS_INSTALLATION")
	if installation == "" || target.Config == nil || target.Config.Labels[haosLabel] != installation {
		return errors.New("Dieser Container gehört nicht zu dieser MCPanelHAOS-Installation.")
	}

	related, err := c.docker.ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("label", haosLabel+"="+installation)),
	})
	if err != nil {
		return err
	}
	for _, other := range related {
		switch other.State {
		case "created", "exited", "dead":
			continue
		}
		if other.ID == target.ID {
			continue
		}
		if fromServerID == "" {
			return &HAOSBusyError{
				ServerID: other.Labels["discopanel.server.id"],
				Name:     other.Labels["discopanel.server.name"],
			}
		}
		if other.Labels["discopanel.server.id"] != fromServerID {
			return ErrHAOSSwitchChanged
		}
		if err := c.stopHAOSContainer(ctx, other.ID); err != nil {
			return fmt.Errorf("Wechsel abgebrochen: %w", err)
		}
		fromServerID = ""
	}

	if fromServerID != "" {
		return ErrHAOSSwitchChanged
	}
	return nil
}

func (c *Client) stopHAOSContainer(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, HAOSStopTimeout)
	defer cancel()

	info, err := c.docker.ContainerInspect(ctx, id)
	if err != nil {
		return err
	}
	if info.Config == nil || info.Config.Labels[haosLabel] != os.Getenv("MCPANEL_HAOS_INSTALLATION") || info.Config.Labels[haosLabel] == "" {
		return errors.New("Dieser Container gehört nicht zu dieser MCPanelHAOS-Installation.")
	}
	if info.State == nil {
		return errors.New("Docker meldet keinen Containerstatus.")
	}
	if !info.State.Running && !info.State.Restarting && !info.State.Paused {
		return nil
	}
	if info.State.Running && !info.State.Paused && info.State.Health != nil && info.State.Health.Status == "healthy" {
		if _, err := c.ExecCommand(ctx, id, "save-all flush"); err != nil {
			// The image's TERM handler also saves and stops Java. Keep this
			// path usable when RCON is disabled or unavailable.
			c.log.Warn("Explizites Speichern über RCON fehlgeschlagen. Sende reguläres Stoppsignal: %v", err)
		}
	}

	// Wait for the itzg image's shutdown handler to reap Java. Disable
	// Docker's forced kill so a slow world save cannot be cut short.
	timeout := -1
	if info.State.Paused {
		if err := c.docker.ContainerUnpause(ctx, id); err != nil {
			return err
		}
	}
	if err := c.docker.ContainerStop(ctx, id, container.StopOptions{Timeout: &timeout}); err != nil {
		return err
	}
	final, err := c.docker.ContainerInspect(ctx, id)
	if err != nil {
		return err
	}
	if final.State == nil || final.State.Running || final.State.Restarting || final.State.Paused {
		return errors.New("Der vorherige Minecraft-Container wurde noch nicht vollständig beendet.")
	}
	if final.State.ExitCode != 0 && final.State.ExitCode != 143 {
		return fmt.Errorf("Minecraft wurde mit Status %d beendet. Prüfe das Serverprotokoll.", final.State.ExitCode)
	}
	return nil
}

func (c *Client) SetHAOSCurseForgeKey(loader func(context.Context) (*models.ServerConfig, bool, error)) {
	c.haosCurseForgeSettings = loader
}
