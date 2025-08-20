# Installation

Install and configure your Go toolchain.

## From Source

- Download the latest Go binary for your OS from https://go.dev/dl
- Choose based on your system's OS and architecture type

**Linux example**:

```sh
wget https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz
```

Make sure your shell's `PATH` environment variable points to add the installation directory to your shell environment,
by updating your `.bashrc` / `.bash_profile` / `.profile` file.
[If you don't get this part, don't worry, you can read more about it here.](/topics/programming/foundation/shell/cli/EnvironmentVariables.md)
[If shell itself is new to you, then you can read more about it here.](/topics/programming/foundation/CommandLineInterface.md)

```sh
PATH="$PATH:/usr/local/go/bin"
export PATH
```

If you have a previous golang installation or you want to uninstall it:

```sh
sudo rm -rf /usr/local/go
```

## Homebrew

```sh
brew install --formula golang
```

## Snap

```sh
sudo snap install go --classic
```

## Debian/Ubuntu via APT

```sh
sudo apt update
sudo apt install golang-go
```

## Fedora/CentOS/RHEL via DNF/YUM

**Fedora**:

```sh
sudo dnf install golang
```

**CentOS/RHEL**:

```sh
sudo yum install golang
```

## Verify installation

```sh
go version
go env GOPATH GOROOT
```

## Cloning this project for the exercises

```sh
git clone "https://github.com/adamluzsi/learn-to-code.git"
```

Or If you have SSH configured with GitHub:

```sh
git clone git@github.com:adamluzsi/learn-to-code.git
```

You will find the exercise within the `language/golang/exercise/` directory.
