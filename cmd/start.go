/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"
	
)
var focus int
var offtime int
var cycles int
// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a customizable Pomodoro timer",
	Long: `Start a Pomodoro session with customizable focus and break durations.

The Pomodoro Technique is a time management method that uses timed intervals:
- A focus period (default: 25 minutes)
- A short break (default: 5 minutes)
- Multiple cycles (default: 4)

Example usage:

  pomo-cli start --focus 25 --break 5 --cycles 4

You can use shorthand flags too:

  pomo-cli start -f 15 -b 3 -c 2

This will run 2 cycles of 15 minutes focus and 3 minutes break.`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 1; i <= cycles; i++ {
			fmt.Println("Round:", i);
			beginNewCycle(focus, offtime);
		}
		fmt.Println("Session completed with %d cycles",cycles);
	},
}
func timer(d int) {
   duration := time.Duration(d) * time.Minute
    ticker := time.NewTicker(time.Second)
    start := time.Now()
    for t := range ticker.C {
        elapsed := int(t.Sub(start).Seconds())
        if elapsed >= int(duration.Seconds()) {
            break
        }
		fmt.Printf("\r %d/%d seconds elapsed...", elapsed, focus*60)
    }	
	ticker.Stop()
}
func beginNewCycle(focus int, offtime int){
	fmt.Println("Starting focus period");
	timer(focus);
	beep();
	fmt.Println("Starting break period");
	timer(offtime);
	beep();
}
func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntVarP(&focus, "focus", "f", 25, "Focus time in minutes")
	startCmd.Flags().IntVarP(&offtime, "break", "b", 5, "Break time in minutes")
	startCmd.Flags().IntVarP(&cycles, "cycles", "c", 4, "Number of Pomodoro cycles")
}
