package main
import(
	
)


type SesionListReq struct{
	_msgpack struct{} `msgpack:",asArray"`
	Method string
	Token string
}

type SessionListRes struct{
	// var type `gettingscheme:"name"`
	 ID uint32 `msgpack:",omitempty"`
	 Type string `msgpack:"type"`
	 TunnelLocal string `msgpack:"tunnel_local"`
	 TunnelPeer string `msgpack:"tunnel_peer"`
	 ViaExploit string `msgpack:"via_exploit"`
	 ViaPayload string `msgpack:"via_payload"`
	 Description string `msgpack:"desc"`
	 Info string `msgpack:"info"`
	 Workspace string `msgpack:"workspace"`
	 SessionHost string `msgpack:"target_host"`
	 SessionPort int `msgpack:"session_port"`
	 Username string `msgpack:"username"`
	 UUID string `msgpack:"uuid"`
	 Exploit_UUID string `msgpack:"exploit_uuid"`
}
