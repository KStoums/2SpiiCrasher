package commands

import (
	"2SpiiCrasher/messages"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var pingCommand = &cobra.Command{
	Use:     "ping",
	Short:   "Sends a defined ping count to the target server IP",
	Long:    "Sends a defined ping count to the target server IP",
	Example: "2spiicrasher ping 100 play.myserver.net",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		mountPing, err := strconv.ParseInt(args[0], 6, 12)
		if err != nil {
			log.Fatalln(err)
		}

		if mountPing > 20000 || mountPing < 1 {
			log.Fatal(messages.MountPingTooHigh)
		}

		serverAddress := args[1]

		ping := fastping.NewPinger()
		resolveAddress, err := net.ResolveIPAddr("ip4:icmp", serverAddress)
		if err != nil {
			log.Fatalln(err)
		}

		ping.AddIPAddr(resolveAddress)

		waitGroup := sync.WaitGroup{}
		waitGroup.Add(int(mountPing))

		for i := 0; i < int(mountPing); i++ {
			go func() {
				ping.OnRecv = func(ipAddr *net.IPAddr, duration time.Duration) {
					fmt.Println(fmt.Sprintf("IP Address: %s receive, RTT: %v", ipAddr.String(), duration))
				}

				ping.OnIdle = func() {
					fmt.Println(messages.PositiveAnswer)
				}

				err = ping.Run()
				if err != nil {
					fmt.Println(err)
				}

				waitGroup.Done()
			}()
		}

		waitGroup.Wait()
		return

	},
}

func init() {
	rootCommand.AddCommand(pingCommand)
}
