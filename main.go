package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type PackageManager string

const (
	Npm  PackageManager = "npm"
	Pnpm PackageManager = "pnpm"
	Yarn PackageManager = "yarn"
	Bun  PackageManager = "bun"
)

type BiomeConfig struct {
	PackageManager  PackageManager
	InitBiome       bool
	MigrateEslint   bool
	MigratePrettier bool
}

func runCommandWithSpinner(s *spinner.Spinner, cmd *exec.Cmd, title, errMsg string) {
	_ = s.Title(title).Action(func() {
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("%s: %s\n%s\n", errMsg, err.Error(), string(output))
			os.Exit(1)
		}
	}).Run()
}

func main() {
	var config BiomeConfig

	// Should we run in accessible mode?
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[PackageManager]().
				Options(huh.NewOptions(Npm, Pnpm, Yarn, Bun)...).
				Title("Choose your package manager").
				Value(&config.PackageManager),
			huh.NewConfirm().
				Title("Initialize Biome?").
				Value(&config.InitBiome),
			huh.NewConfirm().
				Title("Migrate Eslint?").
				Value(&config.MigrateEslint),
			huh.NewConfirm().
				Title("Migrate Prettier?").
				Value(&config.MigratePrettier),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	if config.InitBiome {

		var installCmd *exec.Cmd
		var initCmd *exec.Cmd

		switch config.PackageManager {
		case Npm:
			installCmd = exec.Command("npm", "install", "--save-dev", "--save-exact", "@biomejs/biome")
			initCmd = exec.Command("npx", "@biomejs/biome", "init")
		case Pnpm:
			installCmd = exec.Command("pnpm", "add", "--save-dev", "--save-exact", "@biomejs/biome")
			initCmd = exec.Command("pnpm", "biome", "init")
		case Yarn:
			installCmd = exec.Command("yarn", "add", "--dev", "--exact", "@biomejs/biome")
			initCmd = exec.Command("yarn", "biome", "init")
		case Bun:
			installCmd = exec.Command("bun", "add", "--dev", "--exact", "@biomejs/biome")
			initCmd = exec.Command("bunx", "@biomejs/biome", "init")
		}

		runCommandWithSpinner(spinner.New().Accessible(accessible), installCmd, "Installing Biome...", "Error installing Biome")
		runCommandWithSpinner(spinner.New().Accessible(accessible), initCmd, "Initializing Biome...", "Error initializing Biome")
	}

	if config.MigrateEslint {
		cmd := exec.Command("biome", "migrate", "eslint", "--write")
		runCommandWithSpinner(spinner.New().Accessible(accessible), cmd, "Migrating Eslint...", "Error migrating Eslint")
	}

	if config.MigratePrettier {
		cmd := exec.Command("biome", "migrate", "prettier", "--write")
		runCommandWithSpinner(spinner.New().Accessible(accessible), cmd, "Migrating Prettier...", "Error migrating Prettier")
	}

	fmt.Println("Biome setup complete.")
}
