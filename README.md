# gh-api-demo

GitHub API Demo is a simple [Golang](https://go.dev) CLI application displaying public [Gists](https://gists.github.com) for a given [GitHub](https://github.com) username.

## Installation

To install GitHub API Demo, please use the following methods (sorted by ease of use, from top to bottom). Each installation method requires an access to a system terminal (ie. `Terminal.app` on MacOS Intel/AppleSilicon) and some terminal proficiency.

### Installation - GitHub Releases

Downloading a precompiled binary is the easiest way of installation, given it doesn't have any external dependencies.

1. Open the [releases](https://github.com/bartekrutkowski/gh-api-demo/releases) page of the GitHub API Demo repository
2. Select and download the binary relevant for your operating system (ie. `gh-api-demo_darwin_amd64` for MacOS Intel or `gh-api-demo_freebsd_arm` for FreeBSD ARM)
3. Either place the downloaded binary in your executable path of your operating system (ie. `/usr/local/bin` for MacOS Intel/AppleSilicon) or make sure you are executing it directly from the folder where you have downloaded the binary
4. Ensure the binary is executable, ie. invoking (for MacOS Intel/AppleSilicon)):

        chmod +x path_to_your_binary

### Installation - Docker

Installation via usage of a [Docker](https://www.docker.com) container is the second easiest way of installation, only because it relies on external software being installed on the system.

1. Pull the [Docker](https://www.docker.com) image from the [DockerHub](https://hub.docker.com/r/bartekrutkowski/gh-api-demo):

        docker pull bartekrutkowski/gh-api-demo:latest

### Installation - Compile Go source code

Compiling the binary is the most complicated way of installation, given it requires [Git](https://git-scm.com) and [Golang](https://go.dev) installed on the system.

1. Clone the gh-api-demo [repository](https://github.com/bartekrutkowski/gh-api-demo):

        git clone git@github.com:bartekrutkowski/gh-api-demo.git

2. Compile the source code into a binary:

        cd gh-api-demo && go build -o gh-api-demo -a ./cmd/cli

3. Ensure the binary is executable, ie. invoking (for MacOS Intel/AppleSilicon)):

        chmod +x path_to_your_binary

## Usage

The app is used directly from the terminal. As a minimum, it requires a flag with a [GitHub](https://github.com) account username.

Note, how different installation methods are used in examples below.

### Basic usage

        $ ./gh-api-demo -acc bartekrutkowski
        https://gist.github.com/291f2837c2a6e8eeae29
        https://gist.github.com/9ef3dd3ed40867717fec
        $

### Advanced usage

Additionally, either `-raw` flag for complete JSON output of the Gists objects or `-table` flag for selected fields output in ASCII table format can be passed.

        $ docker run bartekrutkowski/gh-api-demo:latest -acc bartekrutkowski -table
        +----------------------------------------------------------------------------------------------------------+
        | GitHub Gists for user bartekrutkowski                                                                    |
        +---+----------------------------------------------+-----------------------+-------------------------------+
        | # | URL                                          | DESCRIPTION           | DATE                          |
        +---+----------------------------------------------+-----------------------+-------------------------------+
        | 1 | https://gist.github.com/291f2837c2a6e8eeae29 | OS X Homebrew cleanup | 2016-02-07 18:53:11 +0000 UTC |
        | 2 | https://gist.github.com/9ef3dd3ed40867717fec | JSON2PDNS             | 2015-12-09 15:05:27 +0000 UTC |
        +---+----------------------------------------------+-----------------------+-------------------------------+
        $

### All available flags

Currently there is one required and two optional flags. The optional flags are mutually exclusive, only one of them can be used at once.

        -acc string
                GitHub account username (required)
        -raw
                Display full output in raw JSON format (optional)
        -table
                Display stripped output in table format (optional)

## Testing

None.

## Frequently Asked Questions

Q: *Why Golang, why Docker, why so much code, can't it be done differently?*

A: Yes, it can, it can be done in many, many ways. For the record, the simplest functionality could be delivered as this, but I was given a freedom of choice and I wanted to have some fun with it:

        $ curl -s -H "Accept: application/vnd.github.v3+json" https://api.github.com/users/bartekrutkowski/gists | grep "https://gist.github.com/[^.]*$" | awk  -F'"' '{print $4}'
        https://gist.github.com/291f2837c2a6e8eeae29
        https://gist.github.com/9ef3dd3ed40867717fec
        $

Q: *Wait, what? No testing?*

A: Obviously, there could (and in fact, should) be unit and perhaps integration testing added to the source code. However, the delivery time of the entire GH API Demo application was also a constraint, so as a result there is no unit testing available with the source code.

## License

While I retain the author's rights for the application, the entire repository and all its artifacts are relased under the [BSD 3-Clause license](https://github.com/bartekrutkowski/gh-api-demo/blob/main/LICENSE) and is free to use.
