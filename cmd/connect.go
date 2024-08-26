/*
Copyright © 2024 jake-young-dev
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jake-young-dev/mcr"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	address string
	port    int
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a Minecraft server",
	Long:  `Attempt to connect to a Minecraft server, password is required after connection. Once authenticated the connection can be used to make rcon commands to the server in typical fashion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Attempting to connect to %s on port %d\n", address, port)
		connectStr := fmt.Sprintf("%s:%d", address, port)
		client := mcr.NewClient(connectStr)

		fmt.Printf("Password: ")
		ps, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		fmt.Println()

		err = client.Connect(string(ps))
		if err != nil {
			panic(err)
		}
		defer client.Close()
		fmt.Println("Connected!")

		var runningCmd string
		for runningCmd != "quit" {
			fmt.Printf("RCON /> ")
			_, err := fmt.Scanln(&runningCmd)
			if err != nil {
				panic(err)
			}

			if runningCmd == "quit" {
				break
			}
			if runningCmd == "" {
				continue
			}

			res, err := client.Command(runningCmd)
			if err != nil {
				panic(err)
			}
			fmt.Println()
			fmt.Println(res)

		}

	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringVarP(&address, "server", "s", "", "server rcon address")
	connectCmd.Flags().IntVarP(&port, "port", "p", 0, "server rcon port")
	connectCmd.MarkFlagRequired("server")
	connectCmd.MarkFlagRequired("port")
}
