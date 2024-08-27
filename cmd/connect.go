/*
Copyright © 2024 jake-young-dev
*/
package cmd

/*
this library has not been tested yet
*/

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jake-young-dev/mcr"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// flags
var (
	address string
	port    int
)

// connectCmd represents the connect subcommand for kelp
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a Minecraft server",
	Long:  `Attempt to connect to a Minecraft server, password is required on connection. Once authenticated, the connection can be used to make rcon commands to the server in typical Minecraft fashion`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Attempting to connect to %s on port %d\n", address, port)

		//create rcon client
		connectStr := fmt.Sprintf("%s:%d", address, port)
		client := mcr.NewClient(connectStr)

		//read in password
		fmt.Printf("Password: ")
		ps, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return err
		}
		fmt.Println() //print newline after password entered for prettier printing

		//connect to server and authenticate
		err = client.Connect(string(ps))
		if err != nil {
			return err
		}
		defer client.Close()
		fmt.Println("Connected!")

		scanner := bufio.NewScanner(os.Stdin)
		var runningCmd string
		for runningCmd != "quit" {
			fmt.Printf("RCON /> ")
			if scanner.Scan() {
				runningCmd = scanner.Text()

				//no need to send exit command
				if runningCmd == "quit" {
					break
				}
				//empty commands have no impact
				if runningCmd == "" {
					continue
				}

				//send command to server and print response
				res, err := client.Command(runningCmd)
				if err != nil {
					return err
				}
				fmt.Printf("\n%s\n", res)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringVarP(&address, "server", "s", "", "server rcon address")
	connectCmd.Flags().IntVarP(&port, "port", "p", 0, "server rcon port")
	connectCmd.MarkFlagRequired("server")
	connectCmd.MarkFlagRequired("port")
}
