package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"math"
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

			// fmt.Printf("target yr %d, %d, %d\n", yr, mth, day)
			target := time.Date(yr, time.Month(mth), day, 0, 0, 0, 0, time.UTC)

			startYr, fromWeek := target.ISOWeek()
			// fmt.Printf("yr %v, %d \n", startYr, fromWeek)
			wd := target.Weekday()
			if formatWithWeekday {
				fmt.Printf("%02dw%02d-%d", startYr%1000, fromWeek, wd)
			} else {
				fmt.Printf("%02dw%02d", startYr%1000, fromWeek)
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

			// fmt.Printf("target yr %d, %d, %d\n", yr, mth, day)
			target := time.Date(yr, time.Month(mth), day, 0, 0, 0, 0, time.UTC)

			targetYr, targetWeek := target.ISOWeek() //isoweek.FromDate(target.Year(), target.Month(), target.Day())

			//parse the from date
			diff := time.Until(target)
			// fmt.Printf("%v\n", diff.Hours())
			hoursDiff := diff.Round(time.Hour)
			// fmt.Printf("%v\n", hoursDiff)

			weekDiff := int(math.Trunc(hoursDiff.Hours() / 24 / 7))
			fmt.Printf("%d weeks to %vw%02d", weekDiff, targetYr, targetWeek)
			return nil
		},
	}
	return toCmd
}
