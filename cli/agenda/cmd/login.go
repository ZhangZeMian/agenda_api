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
//	"agenda_api/cli/entity"
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/spf13/cobra"
	"strings"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "log in the agenda",
	Long: `you need to input the username and password,for example:
	./agenda login -u=zhangzemian -p=12345678`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		client := &http.Client{}
		reqBody := "{"+"\"Username\":\"" + username +"\", \"Password\":\"" + password + "\"}"
		req, err := http.NewRequest("POST", "http://localhost:8080/api/agenda/user/login", strings.NewReader(reqBody))
		if err != nil {
			fmt.Println("err when create the login request")
		}
	
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//		req.Header.Set("Cookie", "name=anny")
		file1, error1 := os.Open("cookies.txt");
		if error1 != nil {
			fmt.Println(error1);
		}
		defer file1.Close();
		buf := make([]byte, 4024);
		byteNum, err1 := file1.Read(buf)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		// var cookieSlice interface{}
		var cookieSlice []http.Cookie
		fmt.Println(string(buf[0:byteNum]))
		if (byteNum != 0) {
			err := json.Unmarshal(buf[0:byteNum], &cookieSlice)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		
		for _, cookie := range cookieSlice {
			fmt.Println(cookie)
			if cookie.Name == "LoginId" && cookie.Path == "/api/agenda/" {
				req.Header.Set("Cookie", "LoginId=" + cookie.Value)
				fmt.Println("find cookie")
				break
			}
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		//output the response body
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))

		//save cookies pass by the server
		cookies := resp.Cookies()
		if len(cookies) != 0 {
			byteData, _ := json.Marshal(cookies)
			//写入文件
			file, error := os.Create("cookies.txt");
			defer file.Close()
			if error != nil {
				fmt.Println(error);
			}
			file.Write(byteData)
		}
		
	//	fmt.Println(byteData)

	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "Username")
	loginCmd.Flags().StringP("password", "p", "", "User password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}