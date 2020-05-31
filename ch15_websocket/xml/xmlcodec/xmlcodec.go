package xmlcodec

import (
	"encoding/xml"
	"golang.org/x/net/websocket"
)

func xmlMarshal(v interface{}) (data []byte, payloadType byte, err error) {
	data, err = xml.Marshal(v)
	return data, websocket.TextFrame, err
}

func xmlUnmarshal(data []byte, payloadType byte, v interface{}) (err error) {
	err = xml.Unmarshal(data, v)
	return err
}

var XMLCodec = websocket.Codec{
	Marshal:   xmlMarshal,
	Unmarshal: xmlUnmarshal,
}
