package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

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

type EslintMigrationStatus string

const (
	EslintNotMigrated       EslintMigrationStatus = "not migrated"
	EslintMigrated          EslintMigrationStatus = "migrated"
	EslintMigratedWithRules EslintMigrationStatus = "migrated with inspired rules"
)

type BiomeConfig struct {
	PackageManager  PackageManager
	InitBiome       bool
	IntegrateVCS    bool
	MigrateEslint   EslintMigrationStatus
	MigratePrettier bool
	Monorepo        bool
}

type BiomeJSON struct {
	Schema          *string    `json:"$schema,omitempty"`
	OrganizeImports *Formatter `json:"organizeImports,omitempty"`
	Vcs             *Vcs       `json:"vcs,omitempty"`
	Linter          *Linter    `json:"linter,omitempty"`
	Formatter       *Formatter `json:"formatter,omitempty"`
}

type Formatter struct {
	Enabled bool `json:"enabled,omitempty"`
}

type Linter struct {
	Enabled bool  `json:"enabled,omitempty"`
	Rules   Rules `json:"rules,omitempty"`
}

type Rules struct {
	Recommended bool `json:"recommended,omitempty"`
}

type Vcs struct {
	Enabled       bool   `json:"enabled,omitempty"`
	ClientKind    string `json:"clientKind,omitempty"`
	UseIgnoreFile bool   `json:"useIgnoreFile,omitempty"`
	DefaultBranch string `json:"defaultBranch,omitempty"`
}

func runCommandWithSpinner(s *spinner.Spinner, cmd *exec.Cmd, title, errMsg string) {
	fmt.Printf("Running command: %s\n", strings.Join(cmd.Args, " "))
	_ = s.Title(title).Action(func() {
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("%s: %s\n%s\n", errMsg, err.Error(), string(output))
			os.Exit(1)
		}
	}).Run()
}

func runEslintMigrateCommand(config BiomeConfig, accessible bool) {
	var eslintCmd *exec.Cmd
	if config.MigrateEslint != EslintNotMigrated {
		switch config.PackageManager {
		case Npm:
			if config.MigrateEslint == EslintMigratedWithRules {
				eslintCmd = exec.Command("npx", "@biomejs/biome", "migrate", "eslint", "--write", "--include-inspired")
			} else {
				eslintCmd = exec.Command("npx", "@biomejs/biome", "migrate", "eslint", "--write")
			}

		case Pnpm:
			if config.MigrateEslint == EslintMigratedWithRules {
				eslintCmd = exec.Command("pnpm", "biome", "migrate", "eslint", "--write", "--include-inspired")
			} else {
				eslintCmd = exec.Command("pnpm", "biome", "migrate", "eslint", "--write")
			}

		case Yarn:
			if config.MigrateEslint == EslintMigratedWithRules {
				eslintCmd = exec.Command("yarn", "biome", "migrate", "eslint", "--write", "--include-inspired")
			} else {
				eslintCmd = exec.Command("yarn", "biome", "migrate", "eslint", "--write")
			}

		case Bun:
			if config.MigrateEslint == EslintMigratedWithRules {
				eslintCmd = exec.Command("bunx", "@biomejs/biome", "migrate", "eslint", "--write", "--include-inspired")
			} else {
				eslintCmd = exec.Command("bunx", "@biomejs/biome", "migrate", "eslint", "--write")
			}
		}
		runCommandWithSpinner(spinner.New().Accessible(accessible), eslintCmd, "Migrating Eslint...", "Error migrating Eslint")
	}
}

func runPrettierMigrateCommand(config BiomeConfig, accessible bool) {
	var prettierCmd *exec.Cmd

	if config.MigratePrettier {
		switch config.PackageManager {
		case Npm:
			prettierCmd = exec.Command("npx", "@biomejs/biome", "migrate", "prettier", "--write")
		case Pnpm:
			prettierCmd = exec.Command("pnpm", "biome", "migrate", "prettier", "--write")
		case Yarn:
			prettierCmd = exec.Command("yarn", "biome", "migrate", "prettier", "--write")
		case Bun:
			prettierCmd = exec.Command("bunx", "@biomejs/biome", "migrate", "prettier", "--write")
		}

		// TODO: Only JSON configurations are supported. Need to warn user before running the migration or convert the Prettier configuration to JSON.
		runCommandWithSpinner(spinner.New().Accessible(accessible), prettierCmd, "Migrating Prettier...", "Error migrating Prettier")
	}
}

func configureVersionControl() {
	// Read biome.json file
	biomeJsonData, readError := os.ReadFile("biome.json")
	if readError != nil {
		fmt.Println("Error reading biome.json:", readError)
		os.Exit(1)
	}

	// Unmarshal JSON data
	var biomeConfigJson BiomeJSON
	parseError := json.Unmarshal(biomeJsonData, &biomeConfigJson)
	if parseError != nil {
		fmt.Println("Error parsing biome.json:", parseError)
		os.Exit(1)
	}

	// Append the VCS configuration
	biomeConfigJson.Vcs = &Vcs{
		Enabled:       bool(true),
		ClientKind:    "git",
		UseIgnoreFile: bool(true),
		DefaultBranch: "main",
	}

	// Marshal the updated configuration
	updatedBiomeJsonData, marshalError := json.MarshalIndent(biomeConfigJson, "", "  ")
	if marshalError != nil {
		fmt.Println("Error generating updated biome.json:", marshalError)
		os.Exit(1)
	}

	// Write the updated configuration back to biome.json
	writeError := os.WriteFile("biome.json", updatedBiomeJsonData, 0644)
	if writeError != nil {
		fmt.Println("Error writing updated biome.json:", writeError)
		os.Exit(1)
	}
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
				Title("Is this a monorepo?").
				Value(&config.Monorepo),
			huh.NewConfirm().
				Title("Initialize Biome?").
				Value(&config.InitBiome),
			huh.NewConfirm().
				Title("Integrate with Version Control System?").
				Value(&config.IntegrateVCS),
			huh.NewSelect[EslintMigrationStatus]().
				Title("Migrate ESlint?").
				Options(
					huh.NewOptions(EslintNotMigrated, EslintMigrated, EslintMigratedWithRules)...,
				).
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
			if config.Monorepo {
				installCmd = exec.Command("npm", "install", "-W", "--save-dev", "--save-exact", "@biomejs/biome")
			} else {
				installCmd = exec.Command("npm", "install", "--save-dev", "--save-exact", "@biomejs/biome")
			}
			initCmd = exec.Command("npx", "@biomejs/biome", "init")
		case Pnpm:
			if config.Monorepo {
				installCmd = exec.Command("pnpm", "add", "-r", "--save-dev", "--save-exact", "@biomejs/biome")
			} else {
				installCmd = exec.Command("pnpm", "add", "--save-dev", "--save-exact", "@biomejs/biome")
			}
			initCmd = exec.Command("pnpm", "biome", "init")
		case Yarn:
			if config.Monorepo {
				installCmd = exec.Command("yarn", "add", "-W", "--dev", "--exact", "@biomejs/biome")
			} else {
				installCmd = exec.Command("yarn", "add", "--dev", "--exact", "@biomejs/biome")
			}
			initCmd = exec.Command("yarn", "biome", "init")
		case Bun:
			installCmd = exec.Command("bun", "add", "--dev", "--exact", "@biomejs/biome")
			initCmd = exec.Command("bunx", "@biomejs/biome", "init")
		}

		runCommandWithSpinner(spinner.New().Accessible(accessible), installCmd, "Installing Biome...", "Error installing Biome")
		runCommandWithSpinner(spinner.New().Accessible(accessible), initCmd, "Initializing Biome...", "Error initializing Biome")
	}

	if config.IntegrateVCS {
		configureVersionControl()
	}

	runEslintMigrateCommand(config, accessible)
	runPrettierMigrateCommand(config, accessible)

	fmt.Println("\nBiome setup is now complete. For more information, please visit:")
	fmt.Println("\t- Get started: https://biomejs.dev/guides/getting-started/")
	fmt.Println("\t- Migrate Eslint and Prettier: https://biomejs.dev/guides/migrate-eslint-prettier/")
	fmt.Println("\nYou should check these recipes:")
	fmt.Println("\t- Continuous Integration: https://biomejs.dev/recipes/continuous-integration/")
	fmt.Println("\t- Git Hooks: https://biomejs.dev/recipes/git-hooks/")
	fmt.Println("\nIf you encounter any issues, please open issue them at:")
	fmt.Println("\t- https://github.com/jellydn/biome-interactive")
	fmt.Println("\nContributions are welcome! If you would like to improve the project, feel free to open a pull request.")
	fmt.Println("\nIt's time to remove EsLint and Prettier from your devDependencies and its config files. Enjoy using Biome!")
}
