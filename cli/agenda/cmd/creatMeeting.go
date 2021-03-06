// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"net/http"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	//"time"
)

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "to create a meeting for user",
	Long: `the user need to input the title, participants, startTime, endTime of the meeting.For example:
	./agenda createMeeting -t=Work -p=zhangzekun -s=2016-10-10/10:00 -e=2017-10-10/10:00`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createMeeting called")
		title, _ := cmd.Flags().GetString("title")
		participants, _ := cmd.Flags().GetStringSlice("participants")
		startTime, _ := cmd.Flags().GetString("startTime")
		endTime, _ := cmd.Flags().GetString("endTime")
		//deal with the input data
		participantsStr := "["
		for index, participator := range participants {
			if index != 0 {
				participantsStr = participantsStr + ",\"" + participator + "\""
			} else {
				participantsStr = participantsStr + "\"" +participator + "\""
			}
		}
		participantsStr = participantsStr + "]"

		//set request
		client := &http.Client{}
		reqBody := fmt.Sprintf("{\"Title\":\"%s\", \"Participants\":%s, \"StartTime\":\"%s\", \"EndTime\":\"%s\"}", title, participantsStr, startTime, endTime)
		fmt.Println(reqBody)		
		req, err := http.NewRequest("POST", "http://localhost:8080/api/agenda/meeting/createMeeting", strings.NewReader(reqBody))
		if err != nil {
			fmt.Println("err when create creatMeeting request")
		}
		cookieId := GetCookieID()
		if cookieId != "" {
			req.Header.Set("Cookie", "LoginId="+cookieId)
		}

		//send request and output response body
		_, err = SendRequestAndOutputResponseBody(req, client)
		if err != nil {
			fmt.Println(err.Error());
		}
	},
}

func init() {
	RootCmd.AddCommand(createMeetingCmd)
	createMeetingCmd.Flags().StringP("title", "t", "", "title")
	createMeetingCmd.Flags().StringSliceP("participants", "p", make([]string, 0), "participants")
	createMeetingCmd.Flags().StringP("startTime", "s", "", "startTime")
	createMeetingCmd.Flags().StringP("endTime", "e", "", "User endTime")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
