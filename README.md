# kelp
Terminal-based remote console for Minecraft servers.

# usage
- Install command with go install
    - go install github.com/jake-young-dev/kelp@latest
- Connect to the server with the kelp command with server address and port number
    - kelp connect -s address -p port
- Enter rcon password when prompted
- Enter commands to send to the server, type "quit" to disconnect

# dependencies
- [cobra](https://github.com/spf13/cobra)
- [mcr](https://github.com/jake-young-dev/mcr)
- [term](https://golang.org/x/term)

# notes
- This repo is working to the best of my knowledge but is still in early stages of development and you may encounter bugs. Please open an issue for any bugs found.
- RCon isn't the most secure protocol and can be abused, ensure you use a VERY strong password and if possible do not expose the RCon port to the open internet. When possible use a VPN or local connection.
- RCon runs commands directly on your Minecraft server, only run commands that you know the expected outcome to prevent altering your server accidentally.
- This repo has only been tested on Linux and may not work on other OS's
