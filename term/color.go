package term

// Color is a terminal color.
type Color string

const (
	// Reset reinitialize terminal style
	Reset Color = "\033[0m"
	// Black terminal color
	Black Color = "\033[0;30m"
	// DarkGray terminal color
	DarkGray Color = "\033[1;30m"
	// Red terminal color
	Red Color = "\033[0;31m"
	// LightRed terminal color
	LightRed Color = "\033[1;31m"
	// Green terminal color
	Green Color = "\033[0;32m"
	// LightGreen terminal color
	LightGreen Color = "\033[1;32m"
	// Orange terminal color
	Orange Color = "\033[0;33m"
	// Yellow terminal color
	Yellow Color = "\033[1;33m"
	// Blue terminal color
	Blue Color = "\033[0;34m"
	// LightBlue terminal color
	LightBlue Color = "\033[1;34m"
	// Purple terminal color
	Purple Color = "\033[0;35m"
	// LightPurple terminal color
	LightPurple Color = "\033[1;35m"
	// Cyan terminal color
	Cyan Color = "\033[0;36m"
	// LightCyan terminal color
	LightCyan Color = "\033[1;36m"
	// LightGray terminal color
	LightGray Color = "\033[0;37m"
	// White terminal color
	White Color = "\033[1;37m"
)

// String returns string value of color
func (color Color) String() string {
	return string(color)
}

// Sprint returns message with color information for terminal
func (color Color) Sprint(message string) string {
	return color.String() + message + Reset.String()
}
