// The colors package is for commands and applications that support
// ANSI color escapes. Don't use it unless you are sure that is the only place
// your code will be used. These days ANSI color support is standard.
//
// Rather that incur even the minimal cost of running a method for ever ANSI
// escape this package uses constants and assumes that when you use them you
// are reasonably confident you won't have to *not* use them.
//
// When the option to not use color must be considered, for example when
// piping output directly into vim editing sessions, you can use the Decolor()
// and Strip() functions to strip all colors or ANSI escapes known to this
// package.
//
// Three convience executable commands have been included for testing and can
// be used in shell scripts (although sourcing a native Bash equivalent is
// preferred).
package colors
