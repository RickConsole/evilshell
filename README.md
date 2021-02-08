# Evil Shell

*This shell has to be rebuilt so it can mimic a real shell in every way except it captures the root password. I will be working on that soon. *

Evil Shell is a social engineering tool to capture user passwords

Running the bianry will present you with a bash-like prompt that mimics the format and colors of a default ubuntu system. 

Ideally, an alias is placed in the compromised user's bashrc that will cause this binary to run when they execute a certain command. The victim will hopefully not notice a difference and continue their tasks. 

Once they run a command using `sudo` they will recieve the standard password prompt, but that password is actually being recorded and stored in a secret location specified by the attacker before compilation.  

Feel free to modify this code to send the password to a cnc server or anything else

## Features

- Prompt made to look like standard ubuntu bash prompt
- Prompt shows the current directory the victim was in when they ran the binary
- Basic commands are fully functional while in the fake shell
- CTRL+C and `exit` cannot exit the fake shell
- Secret developer passphrase that can be used to exit the fake shell (default: devexit)

## Current Limitations
- Cannot change directories or mimic changing of directories
- Cannot pipe commands together (results in unknown command error)
- No tab completion
