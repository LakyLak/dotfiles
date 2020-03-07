// +build mage

package main

// Setup macOS. Install apps, tools, link dotfiles.
func Setup() {
	installApps()
	installCLI()
	Link()
}

// Link dotfiles.
func Link() {
	deleteFile("")
}

func installApps() {
	// TODO: for apps not part of brew cask, offer to open URLs for download pages of apps
	apps := []string{"vscode-insiders", "iterm"}
	for _, app := range apps {
		caskInstall(app)
	}
}

func caskInstall(app string) {
	// brew cask install <app>
}

// Install CLI tools I use.
func installCLI() {
	brewInstall()
	goInstall()
	rustInstall()
}

func brewInstall() {
	// TODO: install with brew
}

func goInstall() {
	// TODO: install with go
}

func rustInstall() {
	// TODO: install with cargo/brew?
}

// TODO: add cmd with https://github.com/r-darwish/topgrade?
// TODO: combile mage into binary to use globally