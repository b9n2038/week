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

var (
	formatWithWeekday bool
)

// func parseArgs = (args[] string) string {
// return ""
// }
// fzf version maybe way better
func newOfCmd() *cobra.Command {
	ofCmd := &cobra.Command{
		Use:   "of",
		Short: "translate a date dd [MM] [yy[yyyy]] format into ISOWeek and day.",
		RunE: func(cmd *cobra.Command, args []string) error {

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
				if len(args[2]) <= 2 {
					yr, _ = strconv.Atoi(args[2])
					yr = 2000 + yr
				}
				//else silent use current yr
			}

			target := time.Date(yr, time.Month(mth), day, 0, 0, 0, 0, time.UTC)

			//1 1 27 works, but 1 1 has 24 wk01 as yr, not 25w53 if thats what it is meant to be
			// fmt.Printf("target yr %d, %d, %d\n", yr, mth, day)
			fromYr, fromWeek := isoweek.FromDate(target.Year(), target.Month(), target.Day())
			// fmt.Printf("fromYr %d, %d\n", fromYr, fromWeek)
			//to get the weekday, need more conversion
			startYr, startMth, _ := isoweek.StartDate(fromYr, fromWeek)
			// fmt.Printf("start yr %d, %d \n", startYr, startMth)
			wd := isoweek.ISOWeekday(startYr, startMth, day)
			// shortYr := startYr.Format("06") // Last two digits of year
			// fmt.Printf("weekday %d,\n", wd)

			//todo: check long flag format
			if formatWithWeekday {
				fmt.Printf("%02dw%02d-%d\n", startYr%1000, fromWeek, wd)
			} else {
				fmt.Printf("%02dw%02d\n", startYr%1000, fromWeek)
			}
			return nil
		},
	}
	ofCmd.Flags().BoolVarP(&formatWithWeekday, "weekday", "w", false, "Format with weekday")
	return ofCmd
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
