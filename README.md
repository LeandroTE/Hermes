# Hermes
Hermes is serial terminal for interface with IoT devices. Hermes is focusing in helping developing device that have UART interfaces.

# Prepare enviroment

## Installing Golang

```
 rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
```
After add /usr/local/go/bin to the PATH environment variable

```
export PATH=$PATH:/usr/local/go/bin
```

## Installing additional components

Fyne require gcc and other components components for Debian/Ubuntu use:

```
sudo apt-get install gcc libgl1-mesa-dev xorg-dev
```
For other OS follow the instructions on the [link](https://docs.fyne.io/started/)

## Running the app

To run the app just execute the following comand

```
go run .
```

For the first time it might take a little while to build the app, especially in Windows, it cantake a fews minutes to build. For sub sequent build it will be much faster.

