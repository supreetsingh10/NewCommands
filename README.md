# NewCommands
I made these command, when I didn't find a good alternative, and I had to stick to shell scripting to automate things.

## **smv**

Selective MoVe: This command move all the items in a directory except the selected ones to a different directory

## **Example** 

```smv ~/Downloads/lz.mp3 ~/Backup```

All the files from the ~/Downloads except for lz.mp3 will move to Backup. It can take multiple folders inputs and will work just fine.

In order to build this, you will need ```go```
1. ```go build smv.go```
2. ```sudo mv smv /usr/local/bin```


