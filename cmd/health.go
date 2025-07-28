/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
	"runtime"
)

// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "check health, it will beep",
	Long: `checks system health, if it returns system-health that means everything thing from ourside is fine`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(checkHealth());
	},
}

func init() {
	rootCmd.AddCommand(healthCmd);
}
func beep() {
	fmt.Print("\a") // fallback bell
	switch runtime.GOOS {
	case "windows":
		exec.Command("powershell", "-c", `[console]::beep(300,500)`).Run()
	case "darwin":
		exec.Command("afplay", "/System/Library/Sounds/Glass.aiff").Run()
	case "linux":
		exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga").Run()
	default:
		fmt.Println("🔇 Sound not supported on this OS")
	}
}
func checkHealth() string {
	beep()
	return "system-healthy"
}
