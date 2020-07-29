package nasfilesapi

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/nasfiles/nasfilesapi/terminal"
)

//Config ...
type Config struct {
	Development bool
	Host        string
	Port        int
	Secure      bool

	PrivateKey string

	StorageRoot string
	Services    *Services
}

//Log prints all the coniguration values the API is running under
func (c *Config) Log() {
	width := terminal.TerminalSize()

	// Start printing configuration values
	color.HiYellow("Configuration")

	// Beginning separator
	terminal.LineSeparator("-", color.New(color.FgHiCyan), width)

	// Development mode
	fmt.Printf("Development mode: ")
	terminal.YesNoColored(c.Development)

	// Host
	fmt.Printf("Host: ")
	color.HiRed("%s:%d", c.Host, c.Port)

	// Secure
	fmt.Printf("Secure: ")
	terminal.YesNoColored(c.Secure)

	// Storage
	fmt.Printf("Storage path: ")
	color.HiBlue(c.StorageRoot)

	// Ending separator
	terminal.LineSeparator("-", color.New(color.FgHiCyan), width)
}

//Services is a struct to tie all the services into a unified struct
type Services struct {
	User UserService
	Auth AuthService
}
