# Evil Shell

Evil Shell is a social engineering tool to capture user passwords

Running the bianry will present you with a bash-like prompt that mimics the format and colors of a default ubuntu system. 

The idea is that the binary is placed on a target machine and an admin runs it by mistake or out of curiosity. 

To the victim, nothing happens when it is run, but in reality they are in a fake shell. 

Once they run a command using `sudo` they will recieve the standard password prompt, but that password is actually being recorded and stored in a secret location specified by the attacker before compilation.  

# Will it work?

I dont know, maybe? The chances of an admin running an unknown binary are very slim. But if it is named something believable and put in a directory where the admin is likely to see it, there's a chance they could run it... right?

# Features

- Prompt made to look like standard ubuntu bash prompt
- Prompt shows the current directory the victim was in when they ran the binary
- Basic commands are fully functional while in the fake shell
- CTRL+C and `exit` cannot exit the fake shell
- Secret developer passphrase that can be used to exit the fake shell (default: devexit)

# Current Limitations
- Cannot change directories or mimic changing of directories
- Cannot pipe commands together (results in unknown command error)
- No tab completion
