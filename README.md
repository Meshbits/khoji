# Khoji Explorer
 Blockchain Explorer for Komodo Platform and it's ecosystem smartchains


# About codebase of Khoji Explorer
This code is derived from Zcash's blockchain explorer https://zcha.in, and it's initial credit all goes to it's code author with the handle "lustro".
Back in 2017-2018, lustro hosted explorer for Komodo blockchain under domain https://kpx.io, and soon after he took it down because of maintenance issues. He shared the source code of explorer files with some Komodo Platform community members, but the original files recived from lustro has no instructions, documentation, license file(s), or even comments inside the code files. Since the source code was shared in good faith with the community members, it was not shared as open source with anyone in Komodo community except some Komodo community memebers.

I guess since nobody at that moment had required skillset to set this up for Komodo, and the team was already busy with existing work load it sat in our archives until recently. Over these years, I learned and practiced a bit on Go language and recently explored it's codebase, and made few changes to it to make it work.

So, the code in `syncblocks.go` is mostly from the original zcha.in explorer's files, with comments added to the code everywhere possible for whatever I could understand in the code, along with the variable and function name changes through out the file. I changed some parts of checking the new blocks adding to the local blockchain and triggiering the sync function of that data with RethinkDB database.

Through my [kmdgo](https://github.com/satindergrewal/kmdgo) go package for Komodo Platform toolset I also added the ability to specify Komodo and it's ecosystem smartchains right at the command line parameters, and not needing to edit or supply any RPC information inside block explorer's code.

# Requirements

- [Go](https://golang.org/doc/install) 1.14+
- [RethinkDB](https://rethinkdb.com/docs/install/)
- [Git](https://git-scm.com/)
- [verusd](https://github.com/VerusCoin/VerusCoin/releases)

#### Dependencies

- [RethinkDB](https://github.com/rethinkdb/rethinkdb-go) go driver package
- Git
- Make tools (`automake`, `make` etc.)

# Install instructions

### Install Go on your system

On Ubuntu can follow this guide: https://github.com/golang/go/wiki/Ubuntu

```bash
sudo apt update
sudo apt install make git
sudo apt-get install software-properties-common
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt install golang-go
```

On Mac can install Go using brew.

```bash
brew update
brew install git go
```

### Install RethinkDB

Follow RethinkDB install instructions from it's official source: https://rethinkdb.com/docs/install/

On Ubuntu you can install using following instructions:

```shell
source /etc/lsb-release && echo "deb https://download.rethinkdb.com/repository/ubuntu-$DISTRIB_CODENAME $DISTRIB_CODENAME main" | sudo tee /etc/apt/sources.list.d/rethinkdb.list
wget -qO- https://download.rethinkdb.com/repository/raw/pubkey.gpg | sudo apt-key add -
sudo apt-get update
sudo apt-get install rethinkdb
```

On MacOS can install using brew:

```shell
brew update && brew install rethinkdb
```

## Setting up blockchain API

The blockchain networks supported by Khoji usually gets configured automatically with the required settings, but make sure to have the following settings in `.conf` file of the blockchain network you intend to use with Khoji:

```shell
rpcuser=your_rpc_username
rpcpassword=your_rpc_password
rpcport=your_rpc_port
server=1
txindex=1
rpcallowip=127.0.0.1
```

And in case you using remote node for blockchain API, then make sure to update the value for `rpcallowip`. Documentation from Bitcoin says this about it's settings:

```bash
# server=1 tells Bitcoin-Qt to accept JSON-RPC commands.
# it is also read by bitcoind to determine if RPC should be enabled
#rpcallowip=10.1.1.34/255.255.255.0
#rpcallowip=1.2.3.4/24
#rpcallowip=2001:db8:85a3:0:0:8a2e:370:7334/96
```

## Setting up Khoji

Once you are ready with the Golang installed for your OS, and RethinkDB service started, follow these steps to build `khoji`:

```shell
git clone https://github.com/meshbits/khoji.git
cd khoji
make
```

If successfull you will see `khoji` executable binary for your OS inside root directory of `khoji`.

Make copy of `config.ini.sample` as `config.ini`, and configure `CHAIN_NAME`

```shell
cp -av config.ini.sample config.ini
```

Open `config.ini` file in text editor and and change value for `CHAIN_NAME` from `VRSCTEST` to the chain name you are going to run this explorer for.
If testing using VRSCTEST chain, just leave the changes as is and proceed with next steps.
If using for example Veruscoin blockchain's mainnet then change it to `VRSC`.

If you are running RethinkDB locally by executing the command `rethinkdb` in another terminal, then leave the Database section in `config.ini` file as is.
Otherwise, change the IP address and/or port.

You can select the database name for block explorer to setup and use by Khoji by setting up the value for `RDB_DB` under `DATABASE` section of `config.ini`.

So, if for example I'm using `VRSC` mainnet, I will use the following `config.ini` running local instance of RethinkDB:

```ini
[BLOCKCHAIN]

### Define chain name and it's RPC API details. The name value must be the matching value of it's data directory name.
### Example VerusCoin's Test Network's data directory is `VRSCTEST` and so on.
CHAIN_NAME = VRSC

### If you are using remote node for blockchain, then please speciffy it's RPC details
### Make sure to uncomment by removing ; from the variables to enable these settings
###
### Example if you are using remote blockchain API from IP 192.168.1.100 then use that IP for RPC_IP
; RPC_IP = "127.0.0.1"

### Use RPC API access details from config file from the remote blockchain node
; RPC_USER = "CHANGE RPC USER HERE"
; RPC_PASS = "CHANGE RPC PASSWORD HERE"
; RPC_PORT = "CHANGE RPC PORT HERE"

[DATABASE]
### Rethink database name to create and setup with all tables required for explorer
RDB_DB = VRSC
RDB_IP = 127.0.0.1
RDB_PORT = 28015
```

Now you can execute `khoji` to start explorer. It will setup Rethink database and start synchronising explorer data with blockchain network.

```bash
./khoji
```

# TODO

- [-]	Make Explorer's RPC API using database
- [ ]	Web Graphical Interface for Explorer
- [ ]	Make explorer's gRPC API using database


## Khoji logs

If starting Khoji it will only show sync progress in cosole output.
To view the detailed logs you can check `khoji.log` file in the same directory where you executed `khoji` binary from.
Following example command on Linux/OSX will show updated prints being pushed to `khoji.log` file:

```shell
tail -f khoji.log
```

And Windows users can use the following command in PowerShell to check live khoji logs:
```shell
Get-Content .\khoji.log -Wait
```

you can press CTRL+C to cancel `tail` or `Get-Content` command's output.

# Making a release build

Release builds can be made cross platform.
Means you can build Mac OS build on Linux, and Linux builds on Mac OS, thanks to Go's cross-compilation capabilities.

Execute the following `make` command from the root directory of Khoji source code to make distributable archives of Khoji:

```bash
make dist
```

After this command you'll find a zipped copy of Linux, MacOS, MacOS ARM64 and Windows distributable file under `dist/` directory.

### Clean build

To clean all compiled files execute the following command:

```shell
cd $HOME/go/src/github.com/Meshbits/khoji
make clean
```

It will delete all dist and binary files the build commands created.
