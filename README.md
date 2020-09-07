# Evil Shell

Evil Shell mimics a bash shell to hopefully capture a user's password.

Running the bianry will present you with a bash-like prompt that mimics the format and colors of a default ubuntu system. 

The idea is that the binary placed on a target machine and a server admin runs it out of curiosity. 

To the victim, once the binary is run nothing seems to happen, but in reality they are in a fake shell. 

Once they run a command using `sudo` they will recieve the standard password prompt, but that password is actually being recorded and stored in a secret location specified by the attacker before compile time. 

# Features

- Prompt made to look like standard ubuntu bash prompt
- Prompt shows the current directory the victim was in when they ran the binary
- Basic commands are fully functional while in the fake shell
- CTRL+C and `exit` are not alble to exit the fake shell
- Secret developer passphrase that can be used to exit the fake shell (default: devexit)

# Current Limitations
- Cannot change directories or mimic changing of directories
- Cannot pipe commands together (results in unknown command error)
- No tab completion
