package encoding

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"

	"github.com/fromiuan/chinapnr/encoding/php_serialize"
)

const (
	server = "127.0.0.1:21230"
)

type Server struct {
	Err  error
	Data []byte
}

// socket
func newSocket(arg []string) ([]byte, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	// PHP serialize
	phpSerializSource := make(php_serialize.PhpSlice, 0)
	for _, item := range arg {
		phpSerializSource = append(phpSerializSource, item)
	}
	value, err := php_serialize.Serialize(phpSerializSource)
	if err != nil {
		return nil, fmt.Errorf("golang php_serialize error")
	}

	msg := fmt.Sprintf("%d,%s", len(value), value)
	_, err = conn.Write([]byte(msg))
	if err != nil {
		conn.Close()
		return nil, err
	}

	data := make([]byte, 4096)
	var dataLen int
	for {
		count, err := conn.Read(data)
		if count > 0 {
			dataLen = dataLen + count
		}
		if err != nil {
			break
		}
	}
	conn.Close()
	dataContent := data[:dataLen]
	if len(dataContent) < 2 {
		return nil, fmt.Errorf("[LAJP Error] Receive Java exception lt 2")
	}

	//返回类型 "S":成功 "F":异常
	if dataContent[0] == 70 {
		return nil, fmt.Errorf("[LAJP Error] Receive Java exception F")
	}

	return data[:dataLen], nil
	// bufRead, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	conn.Close()
	// 	return nil, err
	// }
	// conn.Close()
	// return bufRead, nil
}

// cfca.sadk.api.CertKit::getCertFromPFX
func GetCertFromPFX(pfxContent []byte, pfxPass string) (s *Server) {
	s = &Server{}
	src := Base64Encode(pfxContent)
	arg := []string{"cfca.sadk.api.CertKit::getCertFromPFX", string(src), pfxPass}
	s.Data, s.Err = newSocket(arg)
	return
}

// cfca.sadk.api.KeyKit::getPrivateKeyIndexFromPFX
func GetPrivateKeyIndexFromPFX(pfxContent []byte, pfxPass string) (s *Server) {
	s = &Server{}
	src := Base64Encode(pfxContent)
	arg := []string{"cfca.sadk.api.KeyKit::getPrivateKeyIndexFromPFX", string(src), pfxPass}
	s.Data, s.Err = newSocket(arg)
	return
}

// cfca.sadk.api.SignatureKit::P7SignMessageAttach
func GetP7SignMessageAttach(pfxContent []byte, pfxPass string, src []byte) (s *Server) {
	s = &Server{}

	// data
	dataEncode := Base64Encode([]byte(Base64Encode(src)))

	// keyIndexServer
	keyIndexServer := GetPrivateKeyIndexFromPFX(pfxContent, pfxPass)
	if keyIndexServer.Err != nil {
		s.Err = keyIndexServer.Err
		return
	}
	keyIndex, err := keyIndexServer.PrivateKeyIndexFromPFX()
	if err != nil {
		s.Err = err
		return
	}
	fmt.Println("keyIndex", keyIndex.Code)
	if keyIndex.Code != "90000000" {
		s.Err = fmt.Errorf("keyindex error code" + keyIndex.Code)
		return
	}

	// getCertFromPFX
	certFromPFXServer := GetCertFromPFX(pfxContent, pfxPass)
	if certFromPFXServer.Err != nil {
		s.Err = certFromPFXServer.Err
		return
	}
	certFrom, err := certFromPFXServer.CertFromPFX()
	if err != nil {
		s.Err = err
		return
	}
	fmt.Println("certFrom", certFrom.Code)
	if keyIndex.Code != "90000000" {
		s.Err = fmt.Errorf("certFrom error code" + certFrom.Code)
		return
	}
	arg := []string{"cfca.sadk.api.SignatureKit::P7SignMessageAttach", "sha256WithRSAEncryption", dataEncode, keyIndex.PrivateKeyIndex, certFrom.Base64CertString}
	s.Data, s.Err = newSocket(arg)
	return
}

// cfca.sadk.api.SignatureKit::P7VerifyMessageAttach
func GetP7VerifyMessageAttach(str string) (s *Server) {
	s = &Server{}
	arg := []string{"cfca.sadk.api.SignatureKit::P7VerifyMessageAttach", str}
	s.Data, s.Err = newSocket(arg)
	return s
}

func (s *Server) CertFromPFX() (*CertFromPFX, error) {
	v := new(CertFromPFX)
	data, err := s.getJson()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, v)
	return v, err
}

func (s *Server) PrivateKeyIndexFromPFX() (*PrivateKeyIndexFromPFX, error) {
	v := new(PrivateKeyIndexFromPFX)
	data, err := s.getJson()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, v)
	return v, err
}

func (s *Server) P7SignMessageAttach() (*P7SignMessageAttach, error) {
	v := new(P7SignMessageAttach)
	data, err := s.getJson()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, v)
	return v, err
}

func (s *Server) P7VerifyMessageAttach() (*P7VerifyMessageAttach, error) {
	v := new(P7VerifyMessageAttach)
	data, err := s.getJson()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, v)
	return v, err
}

func (s *Server) getJson() ([]byte, error) {
	decoder := php_serialize.NewUnSerializer(string(s.Data))
	result, err := decoder.Decode()
	if err != nil {
		return nil, fmt.Errorf("php_unserialize error")
	}
	if data, ok := result.(string); ok {
		return []byte(data), nil
	}
	return nil, fmt.Errorf("get serialize error")
}

func Base64Encode(src []byte) string {
	encodeString := base64.StdEncoding.EncodeToString(src)
	return encodeString
}

func Base64Decode(src string) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(src)
	return decodeBytes, err
}

func Marshal(obj interface{}) ([]byte, error) {
	bt, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return bt, nil
}

func UnMarshal(d []byte, m interface{}) error {
	return json.Unmarshal(d, &m)
}
