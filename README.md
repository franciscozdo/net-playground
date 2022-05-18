# Communication on many hosts playground

## Overview
This project consists of two main programs: master and host.

**master** is the main program that controls every other. It gathers information produced by hosts
and schedules tests on hosts.

**host** is running on every machine and is responsible for communication with master and running tests.

They use gRPC and Protocol Buffers to communicate.

## Configuration
If you want to create configuration for host or master here is a simple description:
* master -- you have to specify
  * address
  * maximum number of connections maintained at the moment
  * description of other hosts: address, port of gRPC server, running services and connected hosts
* host -- you have to specify:
  * name
  * port on which it will open gRPC server
  * connected hosts
  * services to run

To see example configuration go to `config` directory.
To specify configuration file for master or host run it with flag -c and give path to configuration.

## Building
To use it in environment created in docker you can use provided script (see section below).

In each directory (host and master) first run `make proto` (note: you don't have to do it unless you changed sth in proto files)
and then `go build`. It will generate binary ready to use. e.g.
```
cd host
make proto
go build
```
To use simple services that are implemented in directory 'services' just run `go build`.  e.g.
```
cd services/echo
go build
```

## Running tests
To run testing environment just type `./start.sh`. After few seconds you will see master host prompt.
To stop, exit the master host (by typing `exit`) and run `docker-compose down` to stop all created containers.

All configuration for this step is created using simple python script (`config/make_configs.py`).
It generates configuration files for: docker compose, master host and all other hosts.

To modify environment (add hosts, modify them) you can see file `config/config.yaml`.

## Host program

Host runs gRPC server that listens for commands from master. At start it spawns few routines: 
* server -- handles connection
* workers -- they perform tests (in case when there will be many possible tests in the same time)
* data -- it manages data generated during tests
That routines communicate via channels.

## Master program

Master is simple command line utility. You can see all commands with short description if you type `help`.
It connects with host servers (as a client) to pass them messages and schedule tests on them.

To connect with host just type `connect host_name` (name comes from configuration). (You can see current config by typing `config`.)

At any time there is one main connection (which is shown in prompt). If you have connected to more than one host
you will see a star in prompt. You can always switch to another host (with `switch id` command).
To see all connections with their ids type `connections`

To run tests you have two options:
* single test -- it runs on one machine at the time
* multiple tests -- it runs on all already connected machines

Results show which hosts are reachable (for given test).

To disconnect with some host switch to it (by using `switch` command) and type `disconnect`.
(If some host becomes unreachable it doesn't automatically disconnects him. When it will be back online master automatically connects with it.)


