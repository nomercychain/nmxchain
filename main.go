package main

import (
"os"
"github.com/spf13/cobra"
)

func main() {
rootCmd := &cobra.Command{
Use:   "nmxchaind",
Short: "NoMercyChain App",
Long:  "NoMercyChain is a fully functioning, scalable, AI-powered Layer 1 blockchain built on Cosmos SDK.",
Run: func(cmd *cobra.Command, args []string) {
cmd.Println("NoMercyChain daemon - placeholder implementation")
},
}

// Add version command
versionCmd := &cobra.Command{
Use:   "version",
Short: "Print the application version",
Run: func(cmd *cobra.Command, args []string) {
cmd.Println("v0.1.0-alpha")
},
}
rootCmd.AddCommand(versionCmd)

// Execute the root command
if err := rootCmd.Execute(); err != nil {
os.Exit(1)
}
}
