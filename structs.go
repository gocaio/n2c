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

import "encoding/xml"

// Nmaprun was generated 2020-06-02 16:02:50 by joanbono on MacBook.local.
type Nmaprun struct {
	XMLName          xml.Name `xml:"nmaprun"`
	Text             string   `xml:",chardata"`
	Args             string   `xml:"args,attr"`
	Scanner          string   `xml:"scanner,attr"`
	Start            string   `xml:"start,attr"`
	Startstr         string   `xml:"startstr,attr"`
	Version          string   `xml:"version,attr"`
	Xmloutputversion string   `xml:"xmloutputversion,attr"`
	Scaninfo         []struct {
		Text        string `xml:",chardata"`
		Numservices string `xml:"numservices,attr"`
		Protocol    string `xml:"protocol,attr"`
		Services    string `xml:"services,attr"`
		Type        string `xml:"type,attr"`
	} `xml:"scaninfo"`
	Verbose []struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level,attr"`
	} `xml:"verbose"`
	Debugging []struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level,attr"`
	} `xml:"debugging"`
	Host []struct {
		Text      string `xml:",chardata"`
		Endtime   string `xml:"endtime,attr"`
		Starttime string `xml:"starttime,attr"`
		Status    struct {
			Text      string `xml:",chardata"`
			Reason    string `xml:"reason,attr"`
			ReasonTtl string `xml:"reason_ttl,attr"`
			State     string `xml:"state,attr"`
		} `xml:"status"`
		Address struct {
			Text     string `xml:",chardata"`
			Addr     string `xml:"addr,attr"`
			Addrtype string `xml:"addrtype,attr"`
		} `xml:"address"`
		Hostnames struct {
			Text     string `xml:",chardata"`
			Hostname struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
				Type string `xml:"type,attr"`
			} `xml:"hostname"`
		} `xml:"hostnames"`
		Ports struct {
			Text string `xml:",chardata"`
			Port []struct {
				Text     string `xml:",chardata"`
				Portid   string `xml:"portid,attr"`
				Protocol string `xml:"protocol,attr"`
				State    struct {
					Text      string `xml:",chardata"`
					Reason    string `xml:"reason,attr"`
					ReasonTtl string `xml:"reason_ttl,attr"`
					State     string `xml:"state,attr"`
				} `xml:"state"`
				Service struct {
					Text      string   `xml:",chardata"`
					Conf      string   `xml:"conf,attr"`
					Extrainfo string   `xml:"extrainfo,attr"`
					Method    string   `xml:"method,attr"`
					Name      string   `xml:"name,attr"`
					Product   string   `xml:"product,attr"`
					Version   string   `xml:"version,attr"`
					Tunnel    string   `xml:"tunnel,attr"`
					Servicefp string   `xml:"servicefp,attr"`
					Ostype    string   `xml:"ostype,attr"`
					Cpe       []string `xml:"cpe"`
				} `xml:"service"`
				Script []struct {
					Text   string `xml:",chardata"`
					ID     string `xml:"id,attr"`
					Output string `xml:"output,attr"`
					Table  []struct {
						Text string `xml:",chardata"`
						Key  string `xml:"key,attr"`
						Elem []struct {
							Text string `xml:",chardata"`
							Key  string `xml:"key,attr"`
						} `xml:"elem"`
						Table []struct {
							Text string `xml:",chardata"`
							Key  string `xml:"key,attr"`
							Elem []struct {
								Text string `xml:",chardata"`
								Key  string `xml:"key,attr"`
							} `xml:"elem"`
							Table struct {
								Text string   `xml:",chardata"`
								Key  string   `xml:"key,attr"`
								Elem []string `xml:"elem"`
							} `xml:"table"`
						} `xml:"table"`
					} `xml:"table"`
					Elem []struct {
						Text string `xml:",chardata"`
						Key  string `xml:"key,attr"`
					} `xml:"elem"`
				} `xml:"script"`
			} `xml:"port"`
			Extraports struct {
				Text         string `xml:",chardata"`
				Count        string `xml:"count,attr"`
				State        string `xml:"state,attr"`
				Extrareasons struct {
					Text   string `xml:",chardata"`
					Count  string `xml:"count,attr"`
					Reason string `xml:"reason,attr"`
				} `xml:"extrareasons"`
			} `xml:"extraports"`
		} `xml:"ports"`
		Times struct {
			Text   string `xml:",chardata"`
			Rttvar string `xml:"rttvar,attr"`
			Srtt   string `xml:"srtt,attr"`
			To     string `xml:"to,attr"`
		} `xml:"times"`
	} `xml:"host"`
	Runstats struct {
		Text     string `xml:",chardata"`
		Finished struct {
			Text    string `xml:",chardata"`
			Elapsed string `xml:"elapsed,attr"`
			Exit    string `xml:"exit,attr"`
			Summary string `xml:"summary,attr"`
			Time    string `xml:"time,attr"`
			Timestr string `xml:"timestr,attr"`
		} `xml:"finished"`
		Hosts struct {
			Text  string `xml:",chardata"`
			Down  string `xml:"down,attr"`
			Total string `xml:"total,attr"`
			Up    string `xml:"up,attr"`
		} `xml:"hosts"`
	} `xml:"runstats"`
}
