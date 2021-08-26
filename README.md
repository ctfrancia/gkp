# gkp

## Summary
gkp(Go Kill Process) is aimed to be a simple cli for different OS that will kill a process running on a port so you don't need to google: 'kill process running on port linux/mac/windows"
Instead type `gkp 3000 3001` and the port will be closed

### Installation
TODO: write documentation for this

### Flags
- `-p`: port(s) that is desired to kill **ex** `$ gkp -p "3000 3002 4000"` or `$ gkp -p "3000"` for one port
- `-r`: range of ports **ex** a `$ gkp -r "3000 3005"`

## DISCLAIMER
This is a learning project. While you can use it on your own machine I do not hold any responsibility. This disclaimer will go away after some time and been audited by someone much more competent than my self that nothing serious will happen.

### TODO
- clean up main.go for better legibility
- move flag checker into dedicated module linked to todo above
- ~~check for if both flags are presented, should not allow~~

## License
MIT
