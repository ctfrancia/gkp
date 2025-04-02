# gkp

## Summary
gkp(Go Kill Process) is aimed to be a simple cli for different OS that will kill a process running on a port so you don't need to google: 'kill process running on port linux/mac/windows"
Instead type `gkp 3000 3001` and in the example given will kill ports 3000 and 3001 on linux/mac/windows

### Installation
TODO: write documentation for this

### Use
- `-p`: port(s) that is desired to kill **ex** `$ gkp -p "3000 3002 4000"` or `$ gkp -p "3000"` for one port
- `-r`: range of ports **ex** a `$ gkp -r "3000 3005"` will kill ports 3000 **to** 3005

## DISCLAIMER
This is a work in progress and is not ready for production use.

### TODO
- ~~clean up main.go for better legibility~~
- ~~move flag checker into dedicated module linked to todo above~~
- ~~check for if both flags are presented, should not allow~~
- Terminate multiple ports with the `-p` flag on Windows
- Terminate multiple ports with the `-p` flag on Unix-like systems
- Terminate range of flags with `-r` flag on Windows
- Terminate range of flags with `-r` flag on Unix-like systems

## License
MIT
