# kelp
Terminal-based remote console for Minecraft servers.

# dependencies
- [github.com/spf13/cobra v1.8.1](https://github.com/spf13/cobra)
- [github.com/jake-young-dev/mcr v0.2.4](https://github.com/jake-young-dev/mcr)
- [golang.org/x/term v0.23.0](https://golang.org/x/term)

# usage
- Install command with go install
    - go install github.com/jake-young-dev/kelp@latest
- Connect to the server with the kelp command with server address and port number
    - kelp connect -s address -p port
- Enter rcon password when prompted
- Enter commands to send to the server, type "quit" to disconnect

# notes
- This repo is working to the best of my knowledge but is still in early stages of development and you may encounter bugs. Please open an issue for any bugs found.
- RCon isn't the most secure protocol and can be abused, ensure you use a VERY strong password and if possible do not expose the RCon port to the open internet. When possible use a VPN or local connection.
- RCon runs commands directly on your Minecraft server, only run commands that you know the expected outcome to prevent altering your server accidentally.
- This repo has only been tested on Linux and may not work on other OS's
