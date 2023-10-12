package codec

type ICodec interface {
	MIMEType() string
	Marshal(v any) (data []byte, err error)
	Unmarshal(data []byte, v any) error
}

// MIME type的缩写为(Multipurpose Internet Mail Extensions)代表互联网媒体类型(Internet media type)
/**type有下面的形式。
Text：用于标准化地表示的文本信息，文本消息可以是多种字符集和或者多种格式的；
Multipart：用于连接消息体的多个部分构成一个消息，这些部分可以是不同类型的数据；
Application：用于传输应用程序数据或者二进制数据；
Message：用于包装一个E-mail消息；
Image：用于传输静态图片数据；
Audio：用于传输音频或者音声数据；
Video：用于传输动态影像数据，可以是与音频编辑在一起的视频数据格式。
*/
type MIMEType string

const (
	MIMEType_Binary   MIMEType = "application/binary"
	MIMEType_Xml      MIMEType = "application/xml"
	MIMEType_Json     MIMEType = "application/json"
	MIMEType_protobuf MIMEType = "application/x-protobuf"
)
