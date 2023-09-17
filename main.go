/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	// "Hertz/hertz_gen/domprinter"
	// "Hertz/hz_client/domprinter_service"
	"context"
	"fmt"

	"github.com/Dup4/domprinter/hertz_client/domprinter"
	"github.com/Dup4/domprinter/hertz_client/domprinter_service"
)

// const hostname = "http://172.22.237.2:8881"
const hostname = "http://127.0.0.1:8888"

func main() {
	//nowclient = domprinter_service.DefaultClient
	//res,rsp,err := domprinter_service.FetchPrintTask(context.Background(),,"POST")
	defaultClient, err := domprinter_service.NewDOMPrinterServiceClient(hostname)
	if err != nil {
		fmt.Println(err)
		return
	}

	// var task domprinter.FetchPrintTaskReq
	task := domprinter.NewFetchPrintTaskReq()
	task.TaskState = domprinter.TaskStateEnum_Submitted
	task.BaseReq = domprinter.NewBaseReq()
	task.BaseReq.AuthToken = ""

	rsp, res, err := defaultClient.FetchPrintTask(context.Background(), task)
	if err != nil {
		fmt.Println("err Client", err)
		return
	}

	fmt.Println("namomo", rsp, res)
}
