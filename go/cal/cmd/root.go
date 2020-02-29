package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	// for flags
	border string
)

func getTitle(m string, w int) string {
	fstr := fmt.Sprintf("%%%ds", w-((w-len(m))/2))
	return fmt.Sprintf(fstr, m)
}

func getCal(n time.Time) string {
	cal := fmt.Sprintf(fmt.Sprintf("%%%ds", ((n.Day()+4)%7)*4), "")
	// loop through the month
	for t := n.AddDate(0, 0, n.Day()*-1+1); t.Month() == n.Month(); t = t.Add(time.Hour * 24) {
		if t.Day() == n.Day() {
			cal += fmt.Sprintf("(%2d)", t.Day())
		} else {
			cal += fmt.Sprintf("%3d ", t.Day())
		}
		if t.Weekday().String() == "Sunday" {
			cal += "\n"
		}
	}
	return cal
}

func applyBorder(t, h, c string) string {
	switch border {
	case "", "None":
		return fmt.Sprintf("%s\n%s\n%s\n", t, h, c)
	case "Outside":
		w := len(h)
		top := fmt.Sprintf("+%s+\n", strings.Repeat("-", w))
		bottom := top
		formatStr := fmt.Sprintf("|%%%ds|\n", w*-1)
		tb := fmt.Sprintf(formatStr, t)
		hb := fmt.Sprintf(formatStr, h)
		cb := ""
		for _, l := range strings.Split(c, "\n") {
			cb += fmt.Sprintf(formatStr, l)
		}
		return fmt.Sprintf("%s%s%s%s%s", top, tb, hb, cb, bottom)
	default:
		return "This should be an error"
	}
}

func cal(c *cobra.Command, args []string) {
	now := time.Now()
	header := "Mon Tue Wed Thu Fri Sat Sun "
	title := getTitle(now.Month().String(), len(header))
	cal := getCal(now)
	out := applyBorder(title, header, cal)
	fmt.Print(out)
	return
}

var rootCmd = &cobra.Command{
	Use:   "cal",
	Short: "Display a calendar",
	Run:   cal,
}

// Execute will run the thing
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&border, "border", "b", "", "border style for the calendar")
}
