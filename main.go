/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
*/

package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	xmlfile string
	masScan Nmaprun
)

func init() {
	flag.StringVar(&xmlfile, "xml", "", "XML file to use")
	flag.Parse()
}

func main() {
	if xmlfile == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	xmlFile, err := os.Open(xmlfile)
	CheckErr(err)
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)

	CheckErr(xml.Unmarshal(byteValue, &masScan))

	fmt.Printf("IP,Host,Proto,Port,Service,Product,NSE Script ID\n")

	for i := 0; i < len(masScan.Host); i++ {

		address := masScan.Host[i].Address.Addr
		hostname := masScan.Host[i].Hostnames.Hostname.Name

		for j := 0; j < len(masScan.Host[i].Ports.Port); j++ {
			portNum := masScan.Host[i].Ports.Port[j].Portid
			protocol := masScan.Host[i].Ports.Port[j].Protocol
			service := masScan.Host[i].Ports.Port[j].Service.Name
			product := masScan.Host[i].Ports.Port[j].Service.Product
			version := masScan.Host[i].Ports.Port[j].Service.Version

			fmt.Printf("%v,%v,%v,%v,%v,%v %v\n", address, hostname, protocol, portNum, service, product, version)
		}
	}
}

// CheckErr will handle errors
// for the entire program
func CheckErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
