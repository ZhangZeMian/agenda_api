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
)

// cancelMeetingCmd represents the cancelMeeting command
var cancelMeetingCmd = &cobra.Command{
	Use:   "cancelMeeting",
	Short: "to delete the meeting that user creates",
	Long: `user should input the title of the meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cancelMeeting called")
		title, _ := cmd.Flags().GetString("title")
		//set request
		client := &http.Client{}	
		req, err := http.NewRequest("DELETE", "http://localhost:8080/api/agenda/meeting/sponsor/" + title, nil)
		if err != nil {
			fmt.Println("err when create cancleMeeting request")
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
	RootCmd.AddCommand(cancelMeetingCmd)
	cancelMeetingCmd.Flags().StringP("title", "t", "", "title")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
