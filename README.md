# falcongo

*falcongo* is personal project to create a custom home automation system using the Go programming language and the 9P protocol. I used a 9P implementation written in Go called [go9p](https://github.com/lionkov/go9p).

The project consists of several components:
- An arduino connected to various sensors and actuators, communicating with a Go program over a serial connection.
- A Go program that reads data from the arduino and exposes it via a 9P server.
- 9P clients that can connect to the server and interact with the sensors and actuators:
  - A command line client to read sensor data or actuators state.
  - A command line client to set actuators state.
  - A fs-based client that mounts the 9P server as a filesystem, allowing users to interact with sensors and actuators using standard file operations.

## Components

- `sketch/`: Arduino code for reading sensors and controlling actuators.

- `cmd/falcongo/`: Go program that communicates with the Arduino and exposes a 9P server.

- `cmd/falcongo-read/`: Command line client to read sensor data or actuators state.

- `cmd/falcongo-write/`: Command line client to set actuators state.

- `cmd/falcongo-fs/`: fs-based client that mounts the 9P server as a filesystem.

