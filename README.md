AAH, or the Annoyingly Advanced Helper, is a basic command-line program to help me do basic things I should already know.

For example:

```
~$ AAH disk mount
mount: 		mkdir /mnt/disk && mount <device> /mnt/disk
~$
```

## Installation

Installation is fairly simple. I will not provide a prebuilt binary for this project, however I can't see many people using it if I did, and I would need to write a post-install (or first-run, idk, both are technically applicable) script to download the aahelp.yaml file (explained below). Regardless, installing this project from source is fairly simple, so I doubt that there is anything lost in making this decision.

### Prerequisites

- [git](https://git-scm.org/downloads)
- [go](https://golang.org/dl/)

### Setup

If you have `go` set up already, you can skip this. It is basically just a simplified version of the instructions on golang's website.

Follow the link above for instructions to download and install the go binary. After installing go, you can either keep following these instructions or follow the installation instructions linked on the download page. Either way, you will end up with a similar result.

Next, you will want to create a "workspace" for your go projects. In this instructions, I will be creating the workspace at `~/go`, but feel free to change this directory to your liking.

```shell
mkdir go && cd go
mkdir {bin,src,pkg}
```

Now that go's fancy "workspace" directory has been created, it needs to be added as an environment variable. Go ahead and add the following to the end of `~/.bashrc`...

```shell
export GOPATH="/home/<you sir>/go"
export PATH="$PATH:$GOPATH/bin"
```

After saving the file, run `source ~/.bashrc` and you're good to go.

### Installing

This bit is fairly simple.

```shell
go get github.com/TheAndroidMaster/AAH
cd ~/go/src/github.com/TheAndroidMaster/AAH
go get && go build && go install
```

After this, you can try running `AAH` to check that it has installed properly.

## Configuration

Configuration is fairly basic. The program will use the [`aahelp.yaml`](./aahelp.yaml) file in the repository by default. You can override this file by creating your own located at `~/.aahelp.yaml`.

The file uses basic YAML syntax (which can be easily learnt by just taking a glance at the default file in the repository), though it should be noted that there is no support for arrays as they do not make much sense in this context.

The program takes any number of arguments, allowing you to create as many nested maps as you would like and type either some or all of the keys as arguments when running the command to access their values.

For example, with the following file:

```yaml
thing:
    something:
        yes:
            stuff: this is a thing
```

You could type `AAH thing something yes stuff` to get the value `this is a thing`. Alternatively, you could just type `AAH thing something yes` to get a list of everything below the key `yes` in the YAML hierarchy.

## Contributing

See this project's [CONTRIBUTING.md](./.github/CONTRIBUTING.md) for instructions on how to contribute to this project.
