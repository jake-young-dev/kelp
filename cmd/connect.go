/*
Copyright © 2024 jake-young-dev
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jake-young-dev/mcr"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

const (
	//command argument values
	SRVR_ARG     = "server"
	SRVR_ARG_SH  = "s"
	SRVR_ARG_DSC = "server rcon address"
	PORT_ARG     = "port"
	PORT_ARG_SH  = "p"
	PORT_ARG_DSC = "server rcon port"
)

// command flags
var (
	address string //server address/ip
	port    int    //rcon port
)

// connectCmd represents the connect subcommand for kelp
var connectCmd = &cobra.Command{
	Use:          "connect",
	Short:        "Connect to a Minecraft server",
	Long:         `Make a connection request to the Minecraft server, the connection can then be used to send rcon commands to the server in typical Minecraft fashion`,
	SilenceUsage: true, //don't show usage/help message when erroring to make error easier to see
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Attempting to connect to %s on port %d\n", address, port)

		//create rcon client
		connectStr := fmt.Sprintf("%s:%d", address, port)
		client := mcr.NewClient(connectStr)

		//read in password
		fmt.Printf("Password: ")
		ps, err := term.ReadPassword(int(os.Stdin.Fd())) //hides password from terminal
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
		//read in commands until the user disconnects
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
	//add and enforce flags
	connectCmd.Flags().StringVarP(&address, SRVR_ARG, SRVR_ARG_SH, "", SRVR_ARG_DSC)
	connectCmd.Flags().IntVarP(&port, PORT_ARG, PORT_ARG_SH, 0, PORT_ARG_DSC)
	connectCmd.MarkFlagRequired(SRVR_ARG)
	connectCmd.MarkFlagRequired(PORT_ARG)
}
