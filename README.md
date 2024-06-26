<h1 align="center">Welcome to biome-interactive 👋</h1>
<p>
  A simple interactive CLI to install Biome to your project. Migrate from ESLint and Prettier with a Single Command.
</p>

[![ITMan - Biome Interactive CLI - Migrate from ESLint and Prettier with a single command](https://i.ytimg.com/vi/fruvriN-Fpc/hqdefault.jpg)](https://www.youtube.com/watch?v=fruvriN-Fpc)

## Table of Contents

<!--toc:start-->

- [Motivation](#motivation)
- [Install](#install)
- [Built with](#built-with)
- [Author](#author)
- [Show your support](#show-your-support)
- [📝 License](#📝-license)
<!--toc:end-->

## Motivation

Biome is a high-performance code formatter supporting JavaScript, TypeScript, JSX, and JSON, boasting a 97% compatibility rate with Prettier. To streamline the repetitive and time-consuming process of migrating from Eslint/Prettier to Biome, `biome-interactive` was developed. This CLI tool facilitates a consistent and error-free migration, enhancing both CI and development workflows.

## Install

### Using goblin.run

```bash
# goblin.run will build the binary and place it in PATH
curl -sf http://goblin.run/github.com/jellydn/biome-interactive | sh
```

### Using Go

```bash
go install github.com/jellydn/biome-interactive@latest
```

This will install the `biome-interactive` binary to your `$GOPATH/bin` directory.

### Manual

Download the binary for your system from the [releases page](https://github.com/jellydn/biome-interactive/releases).

## Usage

Run the following command in your terminal:

```bash
biome-interactive
```

Follow the interactive prompts to install Biome, initialize it, and migrate configurations from Eslint and Prettier.

[![Demo](https://i.gyazo.com/f0fa4c62b5614ca6e263766ad71774ac.gif)](https://gyazo.com/f0fa4c62b5614ca6e263766ad71774ac)

## Built with

- [charmbracelet/huh](https://github.com/charmbracelet/huh)
- [GoReleaser](https://goreleaser.com/quick-start/)
- [Biome](https://biomejs.dev/blog/biome-v1-7/)

## Resources

- [Integrate Biome with your VCS](https://biomejs.dev/guides/integrate-in-vcs/)
- [Migrate from ESLint and Prettier](https://biomejs.dev/guides/migrate-eslint-prettier/)
- [Continuous Integration](https://biomejs.dev/recipes/continuous-integration/)
- [Git Hooks](https://biomejs.dev/recipes/git-hooks/)

## Author

👤 **Dung Huynh Duc <dung@productsway.com>**

- Github: [@jellydn](https://github.com/jellydn)

## Show your support

Give a ⭐️ if this project helped you!

[![kofi](https://img.shields.io/badge/Ko--fi-F16061?style=for-the-badge&logo=ko-fi&logoColor=white)](https://ko-fi.com/dunghd)
[![paypal](https://img.shields.io/badge/PayPal-00457C?style=for-the-badge&logo=paypal&logoColor=white)](https://paypal.me/dunghd)
[![buymeacoffee](https://img.shields.io/badge/Buy_Me_A_Coffee-FFDD00?style=for-the-badge&logo=buy-me-a-coffee&logoColor=black)](https://www.buymeacoffee.com/dunghd)

## 📝 License

Copyright © 2024 [Dung Huynh Duc <dung@productsway.com>](https://github.com/jelydn).<br />
This project is [MIT](https://github.com/jelydn/moleculer-connect/blob/master/LICENSE) licensed.
