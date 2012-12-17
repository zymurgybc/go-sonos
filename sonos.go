// 
// go-sonos
// ========
// 
// Copyright (c) 2012, Ian T. Richards <ianr@panix.com>
// All rights reserved.
// 
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 
//   * Redistributions of source code must retain the above copyright notice,
//     this list of conditions and the following disclaimer.
//   * Redistributions in binary form must reproduce the above copyright
//     notice, this list of conditions and the following disclaimer in the
//     documentation and/or other materials provided with the distribution.
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
// TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
// LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
// NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
// 

package sonos

import (
	"github.com/ianr0bkny/go-sonos/ssdp"
	"github.com/ianr0bkny/go-sonos/upnp"
)

const MUSIC_SERVICES = "schemas-upnp-org-MusicServices"
const SONOS = "Sonos"

type Sonos struct {
	AlarmClock
	AVTransport
	ConnectionManager
	ContentDirectory
	DeviceProperties
	GroupManagement
	MusicServices
	RenderingControl
	SystemProperties
	ZoneGroupTopology
}

var reactor upnp.Reactor

func sonosHandleUpdate(svc *upnp.Service, value string) {
}

func MakeSonos(svc_map upnp.ServiceMap) (sonos *Sonos) {
	sonos = &Sonos{}
	for svc_type, svc_list := range svc_map {
		switch svc_type {
		case "AlarmClock":
			for _, svc := range svc_list {
				sonos.AlarmClock.svc = svc
				upnp.DescribeService(svc)
				reactor.Subscribe(svc, sonosHandleUpdate)
				break
			}
		case "AVTransport":
			for _, svc := range svc_list {
				sonos.AVTransport.svc = svc
				upnp.DescribeService(svc)
				reactor.Subscribe(svc, sonosHandleUpdate)
				break
			}
		case "ConnectionManager":
			for _, svc := range svc_list {
				sonos.ConnectionManager.svc = svc
				upnp.DescribeService(svc)
				reactor.Subscribe(svc, sonosHandleUpdate)
				break
			}
		case "ContentDirectory":
			for _, svc := range svc_list {
				sonos.ContentDirectory.svc = svc
				upnp.DescribeService(svc)
				reactor.Subscribe(svc, sonosHandleUpdate)
				break
			}
		case "DeviceProperties":
			for _, svc := range svc_list {
				sonos.DeviceProperties.svc = svc
				upnp.DescribeService(svc)
				reactor.Subscribe(svc, sonosHandleUpdate)
				break
			}
		case "GroupManagement":
			for _, svc := range svc_list {
				sonos.GroupManagement.svc = svc
				upnp.DescribeService(svc)
				reactor.Subscribe(svc, sonosHandleUpdate)
				break
			}
		case "MusicServices":
			for _, svc := range svc_list {
				sonos.MusicServices.svc = svc
				upnp.DescribeService(svc)
				reactor.Subscribe(svc, sonosHandleUpdate)
				break
			}
		case "RenderingControl":
			for _, svc := range svc_list {
				sonos.RenderingControl.svc = svc
				upnp.DescribeService(svc)
				reactor.Subscribe(svc, sonosHandleUpdate)
				break
			}
		case "SystemProperties":
			for _, svc := range svc_list {
				sonos.SystemProperties.svc = svc
				upnp.DescribeService(svc)
				reactor.Subscribe(svc, sonosHandleUpdate)
				break
			}
		case "ZoneGroupTopology":
			for _, svc := range svc_list {
				sonos.ZoneGroupTopology.svc = svc
				upnp.DescribeService(svc)
				reactor.Subscribe(svc, sonosHandleUpdate)
				break
			}
		}
	}
	return
}

func sonosQuerySSDPResults(mgr ssdp.Manager) (sonos []*Sonos) {
	qry := ssdp.ServiceQueryTerms{
		ssdp.ServiceKey(MUSIC_SERVICES): -1,
	}
	res := mgr.QueryServices(qry)
	if dev_list, has := res[MUSIC_SERVICES]; has {
		for _, dev := range dev_list {
			if SONOS == dev.Product {
				if svc_map, err := upnp.Describe(dev.Location); nil != err {
					panic(err)
				} else {
					sonos = append(sonos, MakeSonos(svc_map))
				}
				break
			}
		}
	}
	return
}

func Discover(ifiname, port string) (sonos []*Sonos, err error) {
	reactor = upnp.MakeReactor()
	reactor.Init(ifiname, "13105")
	mgr := ssdp.MakeManager()
	mgr.Discover(ifiname, port, false)
	sonos = sonosQuerySSDPResults(mgr)
	return
}