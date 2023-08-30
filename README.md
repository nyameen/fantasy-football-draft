# Fantasy Football Draft Tool

> The tool is WIP

This ia an in-person fantasy football draft tool. The tool is not meant to replace a draft board. Instead it will help manage a draft.

The tool will provide a ranking of the remaining players. People can have the option to select player positions to see the highest ranked player in that position. The tool will also have a custom timer to show how much time is left for that pick. Lastly, when a player is drafted, you can select their name to remove him from the list

## Prerequisites  

This tools uses [Fyne](https://fyne.io/) to build the Graphical User Interface. In order to run the tool GoLang must be installed. To install it, please follow this [link](https://go.dev/doc/install).

Finally, this tool is meant to be built on a Linux system. Please install the following packages.

```sh
sudo apt-get update
sudo apt-get install gcc make libgl1-mesa-dev xorg-dev
```

If you want to build a Windows executable you must install a Windows cross-compiler.

```sh
sudo apt-get install gcc-mingw-w64
```

## Build

To build this tool run the following command

```sh
make linux
```

This will perform a `go run` and start the application.

To build for Windows, run the following command

```sh
make exe
```

This will create a Windows executable.
