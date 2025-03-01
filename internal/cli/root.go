package cli

import (
	"fmt"
	"github.com/snabb/isoweek"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "week",
		Short: "Tools for Weeks",
	}

	rootCmd.AddCommand(
		newOfCmd(),
		newToCmd(),
	)
	return rootCmd
}

// func parseArgs = (args[] string) string {
// return ""
// }
// fzf version maybe way better
func newOfCmd() *cobra.Command {
	toCmd := &cobra.Command{
		Use:   "of",
		Short: "translate a date dd [MM] [yy[yyyy]] format into ISOWeek and day.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("todo parse the dd mm yy\n")

			today := time.Now()
			day := int(today.Day())
			mth := int(today.Month())
			yr := today.Year()
			if len(args) >= 1 {
				day, _ = strconv.Atoi(args[0])
			}
			if len(args) >= 2 {
				mth, _ = strconv.Atoi(args[1])
			}

			//parse yy or yyyy
			if len(args) >= 3 {
				if len(args[2]) == 4 {
					yr, _ = strconv.Atoi(args[2])
				}
				if len(args[2]) == 2 {
					yr, _ = strconv.Atoi("20" + args[2])
				}
				//else silent use current yr
			}

			target := time.Date(yr, time.Month(mth), day, 0, 0, 0, 0, time.UTC)

			_, iweek := isoweek.FromDate(target.Year(), target.Month(), target.Day())
			startYr, startMth, _ := isoweek.StartDate(today.Year(), iweek)
			wd := isoweek.ISOWeekday(startYr, startMth, day)
			shortYr := target.Format("06") // Last two digits of year

			//todo: check long flag format
			fmt.Printf("%sw%02d\n", shortYr, iweek)
			fmt.Printf("%sw%02d-%d\n", shortYr, iweek, wd)
			return nil
		},
	}
	return toCmd
}
func newToCmd() *cobra.Command {
	toCmd := &cobra.Command{
		Use:   "to",
		Short: "weeks to an ISOWeek",
		RunE: func(cmd *cobra.Command, args []string) error {
			//parse the target date
			target := time.Now()
			yr := target.Format("06")
			_, iweek := isoweek.FromDate(target.Year(), target.Month(), target.Day())

			//parse the from date
			diff := target.Sub(time.Now())
			weekDiff := diff.Hours() / 24 / 7

			fmt.Printf("%sw%02d\n", yr, iweek, weekDiff)
			return nil
		},
	}
	return toCmd
}
