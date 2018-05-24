package chinapnr

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/fromiuan/chinapnr/encoding"
)

const (
	Mertest = "http://mertest.chinapnr.com/npay/merchantRequest"
	Finance = "https://finance.chinapnr.com/npay/merchantRequest"
)

type CheckValue struct {
	Value string `json:"check_value"`
}

type ChinapnrPay struct {
	version    string
	merCustId  string
	apiDomain  string
	pfxPass    string
	pfxContent []byte
	Client     *http.Client
}

func New(version, mercustid, pfxPass string, pfxContent []byte, dev bool) (client *ChinapnrPay) {
	client = &ChinapnrPay{}
	client.version = version
	client.merCustId = mercustid
	client.Client = http.DefaultClient
	client.pfxContent = pfxContent
	client.pfxPass = pfxPass
	if dev {
		client.apiDomain = Mertest
	} else {
		client.apiDomain = Finance
	}
	return client
}

func (this *ChinapnrPay) URLValues(param ChinapnrParam) url.Values {
	var p = url.Values{}
	p.Add("version", this.version)
	p.Add("cmd_id", param.APICmId())
	p.Add("mer_cust_id", this.merCustId)

	ps := param.Params()
	if ps != nil {
		for key, value := range ps {
			p.Add(key, value)
		}
	}

	exps := param.ExtParams()
	if exps != nil {
		for key, value := range exps {
			p.Add(key, value)
		}
	}
	return p
}

func (this *ChinapnrPay) ParseURL(p url.Values, cmid string) (string, error) {
	var val = url.Values{}
	val.Add("version", this.version)
	val.Add("cmd_id", cmid)
	val.Add("mer_cust_id", this.merCustId)

	var keys = make([]string, 0, 0)
	for key, _ := range p {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	jsbyte, err := sing(keys, p)
	if err != nil {
		return "", err
	}

	s := encoding.GetP7SignMessageAttach(this.pfxContent, this.pfxPass, jsbyte)
	if s.Err != nil {
		return "", s.Err
	}
	msg, err := s.P7SignMessageAttach()
	if err != nil {
		return "", err
	}
	fmt.Println("signMessage", msg.Code)
	if msg.Code != "90000000" {
		return "", errors.New("SignMessage error" + msg.Code)
	}
	val.Add("check_value", msg.Base64CertString)
	return val.Encode(), nil
}

func (this *ChinapnrPay) doRequest(method string, param ChinapnrParam, results interface{}) (err error) {
	var buf io.Reader
	var apiDomain string
	if param != nil {
		p := this.URLValues(param)
		buf = strings.NewReader(p.Encode())
		fmt.Println("Encode", p.Encode())
		apiDomain, err = this.ParseURL(p, param.APICmId())
		if err != nil {
			return err
		}
	}
	apiDomain = this.apiDomain + "?" + apiDomain

	fmt.Println("apiDomain", apiDomain)

	req, err := http.NewRequest(method, apiDomain, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	resp, err := this.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ckValue := new(CheckValue)
	err = encoding.UnMarshal(data, ckValue)
	if err != nil {
		return errors.New(string(data))
	}

	s := encoding.GetP7VerifyMessageAttach(ckValue.Value)
	if s.Err != nil {
		return s.Err
	}
	verify, err := s.P7VerifyMessageAttach()
	if err != nil {
		return err
	}

	if verify.Result == "True" {
		if verify.Code == "90000000" {
			fbasede, err := encoding.Base64Decode(verify.Base64Source)
			if err != nil {
				return err
			}
			tbasede, err := encoding.Base64Decode(string(fbasede))
			if err != nil {
				return err
			}
			fmt.Println("tbasede is", string(tbasede))
			err = encoding.UnMarshal(tbasede, &results)
			if err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("VerifyMessage error" + verify.Code)
		}
	} else {
		return errors.New(verify.Code)
	}
}

func (this *ChinapnrPay) getURL(param ChinapnrParam) (apiDomain string, err error) {
	if param != nil {
		p := this.URLValues(param)
		apiDomain, err = this.ParseURL(p, param.APICmId())
		if err != nil {
			return "", err
		}
	}
	apiDomain = this.apiDomain + "?" + apiDomain
	return apiDomain, nil
}

func (this *ChinapnrPay) UnParseData(data string, results interface{}) (err error) {
	s := encoding.GetP7VerifyMessageAttach(data)
	if s.Err != nil {
		return s.Err
	}
	verify, err := s.P7VerifyMessageAttach()
	if err != nil {
		return err
	}

	if verify.Result == "True" {
		if verify.Code == "90000000" {
			fbasede, err := encoding.Base64Decode(verify.Base64Source)
			if err != nil {
				return err
			}
			tbasede, err := encoding.Base64Decode(string(fbasede))
			if err != nil {
				return err
			}
			fmt.Println("tbasede is", string(tbasede))
			err = encoding.UnMarshal(tbasede, &results)
			if err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("VerifyMessage error" + verify.Code)
		}
	} else {
		return errors.New(verify.Code)
	}
}

func sing(keys []string, param url.Values) ([]byte, error) {
	if param == nil {
		return nil, errors.New("param nil")
	}
	var pList = make(map[string]string, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			pList[key] = value
		}
	}
	jsstr, err := encoding.Marshal(pList)
	return jsstr, err
}
