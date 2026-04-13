package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/moby/moby/api/types/image"
	"github.com/moby/moby/client"
)

func main() {
	ctx := context.Background()
	fmt.Println("docker-clean-tool starting...")

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("Could not connect to Docker:", err)
		os.Exit(1)
	}
	defer cli.Close()

	// List images
	result, err := cli.ImageList(ctx, client.ImageListOptions{All: true})
	if err != nil {
		fmt.Println("Could not list images:", err)
		os.Exit(1)
	}

	// Print each image
	for _, img := range result.Items {
		name := "<none>"
		if len(img.RepoTags) > 0 {
			name = img.RepoTags[0]
		}
		fmt.Println(name, img.ID[7:19])
	}

	deleteImages(ctx, cli, result.Items)
}

func deleteImages(ctx context.Context, cli *client.Client, images []image.Summary) {
	removed := 0
	failed := 0

	for _, img := range images {
		_, err := cli.ImageRemove(ctx, img.ID, client.ImageRemoveOptions{
			Force:         true,
			PruneChildren: true,
		})
		if err != nil {
			name := img.ID[7:19]
			if len(img.RepoTags) > 0 {
				name = img.RepoTags[0]
			}
			color.Red("  ✗ Failed to remove: %s — %v", name, err)
			failed++
			continue
		}

		name := "<none>"
		if len(img.RepoTags) > 0 {
			name = img.RepoTags[0]
		}
		color.Green("  ✓ Deleted: %s", name)
		removed++
	}

	fmt.Println()
	color.Green("Done! Removed: %d", removed)
	if failed > 0 {
		color.Red("Failed: %d", failed)
	}
}