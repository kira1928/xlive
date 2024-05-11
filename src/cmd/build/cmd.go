package build

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
)

func RunCmd() int {
	app := kingpin.New("Build tool", "xlive Build tool.")
	app.Command("dev", "Build for development.").Action(devBuild)
	app.Command("release", "Build for release.").Action(releaseBuild)
	app.Command("release-docker", "Build for release docker.").Action(releaseDocker)
	app.Command("build-web", "Build webapp.").Action(buildWeb)

	kingpin.MustParse(app.Parse(os.Args[1:]))

	return 0
}

func devBuild(c *kingpin.ParseContext) error {
	BuildGoBinary(true)
	return nil
}

func releaseBuild(c *kingpin.ParseContext) error {
	BuildGoBinary(false)
	return nil
}

func releaseDocker(c *kingpin.ParseContext) error {
	fmt.Printf("release-docker command\n")
	return nil
}

func buildWeb(c *kingpin.ParseContext) error {
	fmt.Printf("build-web command\n")
	return nil
}
