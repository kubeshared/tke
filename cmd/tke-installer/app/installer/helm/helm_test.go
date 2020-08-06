/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the “License”); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an “AS IS” BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package helm

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Chdir("/Users/chenglong/src/tkestack/tke/helmfile.d")
}

func newValues() *Values {
	return &Values{
		Audit: Audit{
			Eanbled: true,
			AuditElasticSearch: AuditElasticSearch{
				Address:     "http://172.19.1.0:9200",
				ReserveDays: 1,
				Username:    "",
				Password:    "",
			},
		},
		Auth: Auth{
			Eanbled: true,
		},
		Business: Business{
			Eanbled: true,
		},
		Gateway: Gateway{
			Eanbled: true,
		},
		Logagent: Logagent{
			Eanbled: true,
		},
		Monitor: Monitor{
			Eanbled: true,
			InfluxDB: InfluxDB{
				Eanbled:  true,
				NodeName: "172.19.1.0",
			},
			StorageType:     "influxdb",
			StorageAddress:  "http://172.19.1.0:8086",
			StorageUsername: "admin",
			StoragePassword: "password",
		},
		Registry: Registry{
			Eanbled:   true,
			Domain:    "registry.tke.com",
			Namespace: "library",
			NodeName:  "172.19.1.0",
		},

		CACert: `-----BEGIN CERTIFICATE-----
MIIC3jCCAcagAwIBAgIBADANBgkqhkiG9w0BAQsFADAgMRAwDgYDVQQKEwdUZW5j
ZW50MQwwCgYDVQQDEwNUS0UwHhcNMjAwNzIyMDc0NzMyWhcNMzAwNzIwMDc0NzMy
WjAgMRAwDgYDVQQKEwdUZW5jZW50MQwwCgYDVQQDEwNUS0UwggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDQpGK1oeyA/RooxcIpPu7hn4i152Pdn5g3O5y8
qfPE7mx/6voXArWBZ2lvL2Vo1APoT8jShMI/AWvKX/oBO5ORN40hbPSMlYBgP/Es
bIwls7rJQ61AAgj8xfyCiIjnnfAOn7kyLkEiwO3KPywD/FKLz7as3PluHLKqmnoB
FSnNGvXKiK4eJQR8MrtMzlnWYSuOrLH/LrjGiFNoWJq9gmETFD5HJT6WflhCNjaz
mO7sSrDhnC0cmUmq39GMfwPmag6S3b+gM5SJ4KWY5szjDTjfC3ZSuZT7TZfUcLmD
usBzx+nFqhHNqo26OTvOUVLuIL2hFly5e4/XfcV9Y4YbzBVNAgMBAAGjIzAhMA4G
A1UdDwEB/wQEAwICpDAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IB
AQDNSWSN+IBFNocOY0bPZ64Qxr7Q/pTopfArxn1UuPcRsWqe9j8RZbDFDmuQi7hy
EyZ+Wds19hdGovEDoSJZrF+awkuo1egzkQMMWmRYr6HT2scmVfDSsbCFJH449/e/
PBYZC6y6P9+1LCXoZG7dm6arVQUXt3AzJc5N5tKRFGiQq5+EIN+xhZRNa+l1ptHF
jSCrl+5ho1WOqri0spTSR0d2/bQ9IYilyhqQmv9mKLLdZofajrGN9Z9t+TfHybt4
5eVhGemrd6hSN2KoBrJqmX8A143yRqbzQwJe7IqKbd1eD41dVFO5UNEIIYrSidHo
zVV8UE69A7ZnEs0L1LGUCIAy
-----END CERTIFICATE-----
`,
		CAKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA0KRitaHsgP0aKMXCKT7u4Z+Itedj3Z+YNzucvKnzxO5sf+r6
FwK1gWdpby9laNQD6E/I0oTCPwFryl/6ATuTkTeNIWz0jJWAYD/xLGyMJbO6yUOt
QAII/MX8goiI553wDp+5Mi5BIsDtyj8sA/xSi8+2rNz5bhyyqpp6ARUpzRr1yoiu
HiUEfDK7TM5Z1mErjqyx/y64xohTaFiavYJhExQ+RyU+ln5YQjY2s5ju7Eqw4Zwt
HJlJqt/RjH8D5moOkt2/oDOUieClmObM4w043wt2UrmU+02X1HC5g7rAc8fpxaoR
zaqNujk7zlFS7iC9oRZcuXuP133FfWOGG8wVTQIDAQABAoIBAQCBypFBTm6yu0DH
tmYXhHfUUYsZJc2hM8PcMDBLySnFC1DhropPQDcrPep/6SsbsDCSFvflxgKUcUzd
wyDiHW6qOrWH4qCqx+Hpfj2IlioE1i7PP1OKavmuTXfMuCT6eOy6HDB4AAAzH3ON
igjWneO358pK9NgR4LZgnOJixAiR9yDLQEHgBjMYDvF2nGxb0QlVrhYGOPBT6nRZ
oJGuttb5GPsg4Yo3clFO6NzipGTRHelK8uNd+1vhTazYwgDNItAiL0oCT/n0E49e
FDilUBsxpHIuQp6iGm7HOS3lIHqjSw6JtkOCnUhu4hvZvI6n52wJpg9w4JVAWTcM
pACfEp/dAoGBAO/4t8enuunkHWJHu9OEavsfIAyhqJ2KJSx3pSsvOk9WiMYissVY
PntmEbIRGeQMQ64ufYUeZ9yMHAFyS4P7x6Q0oeUfbnRAQqo9xursig8Z1M/0AiBY
szEP5Cd+HdWgeAaSM0i5MppYMnTWnV7Owxm5Sup/AiiOqeH6JfYHiSVvAoGBAN6T
9vu48WceOJ18UQ3QH1yLDmi+Qz9XuO6DnIogv63SK8JToB+U1l9SuDINgYnrq76y
KgapApPQsogIRlcc5wpB58KbR/9KJC+m7jvvdlgOeEaYDdJOEJQH7q1Y8WTYidCf
QWD4YfQ/mAGPfuddzNjAXdJrzdpKOZxOeBWFCSsDAoGBALxnq/KnL8r+fEtzQnZY
2p163HDOY1H86h61eQoktEUiDZRyHaVvGO9NPctuPcOcFW8ltQrUynqPzlUcZUmk
XiP/5rO7L785zJ1Yj1pWiDY+86SpURoQSe/VKC4RiF729AVEt9r6/E3NvR+aeEe9
UdRadAUT3p/1ElAHU9KMLBMlAoGASDClgq9RXKiopCuV/mzqPxG68q4DEaBaEsIN
G5/Ti4UHPFfzL0sO3IvvwgwvX5RYd9lBWDeo5lhiD0zfw5FUPjdx++RxTZO2eN4j
lrzsHVNUH1Fh6jv3lIueFgsrABJbYJbGlbT50EZv/tGTz2bU8dgylfTxJ0O7u76X
pCMp/hECgYEAuzzYMHsPrXeTR1OVVtLWmQzQTorqVaGhFW8Xsf4ooZsLZg/Qs4NK
9SGYNP5XACURRj6lm5/gTtgOkAGtYMlQgJ0YNF4cmPdKG4vS8eUkL+mycXpIcH/T
gJKH7t1f79QxGDc6ipeSJ8iIPjYk7J7zZgBO4BvFTwEVRG87PE03b94=
-----END RSA PRIVATE KEY-----
`,
		ServerCert: `-----BEGIN CERTIFICATE-----
MIIFCzCCA/OgAwIBAgIIQly7Kp+PR5wwDQYJKoZIhvcNAQELBQAwIDEQMA4GA1UE
ChMHVGVuY2VudDEMMAoGA1UEAxMDVEtFMCAXDTIwMDcyMjA3NDczMloYDzIxMjAw
NjI4MDc0NzMzWjAnMRAwDgYDVQQKEwdUZW5jZW50MRMwEQYDVQQDEwpUS0UtU0VS
VkVSMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0To4AN+1Y++ntEQs
VqOgijOmTJDepzsoq5tZCI3TvZbIcGVE+uuaFnNqXIVJiKSPQY44W+xCdChb9ZBB
cPnMFYaRjKXFUw9JCgFaHumoc3WopgDtNncuZbAWgdcgRq92elhmsK0SrWcZdkXd
9O6AYLfCXD4KTOk1pwM8jfGXbN5OQVBwAdoYA78aoTppSfe2T89HoB8gOBvolwyW
MQrhFDA+NLdjkjoFJp6TVqiCog739qOXxxzwF+A//zoeJBCOaYYH3xROJOazxBv1
v4IEwIxS/N2uVazXf9HugsPD41xDe6IG4XSiJNqxVSJDpfhoPZ/If6QiRZNBNsII
8EhwWwIDAQABo4ICPjCCAjowDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsG
AQUFBwMCBggrBgEFBQcDATCCAgcGA1UdEQSCAf4wggH6ggFhggFigglsb2NhbGhv
c3SCEHRrZS1wbGF0Zm9ybS1hcGmCGHRrZS1wbGF0Zm9ybS1hcGkudGtlLnN2Y4IM
cGxhdGZvcm0tYXBpghB0a2UtYnVzaW5lc3MtYXBpghh0a2UtYnVzaW5lc3MtYXBp
LnRrZS5zdmOCDGJ1c2luZXNzLWFwaYIOdGtlLW5vdGlmeS1hcGmCFnRrZS1ub3Rp
ZnktYXBpLnRrZS5zdmOCCm5vdGlmeS1hcGmCDHRrZS1hdXRoLWFwaYIUdGtlLWF1
dGgtYXBpLnRrZS5zdmOCCGF1dGgtYXBpggt0a2UtZ2F0ZXdheYITdGtlLWdhdGV3
YXkudGtlLnN2Y4IHZ2F0ZXdheYIPdGtlLW1vbml0b3ItYXBpghd0a2UtbW9uaXRv
ci1hcGkudGtlLnN2Y4ILbW9uaXRvci1hcGmCEHRrZS1yZWdpc3RyeS1hcGmCGHRr
ZS1yZWdpc3RyeS1hcGkudGtlLnN2Y4IMcmVnaXN0cnktYXBpghB0a2UtbG9nYWdl
bnQtYXBpghh0a2UtbG9nYWdlbnQtYXBpLnRrZS5zdmOCDGxvZ2FnZW50LWFwaYIN
dGtlLWF1ZGl0LWFwaYIVdGtlLWF1ZGl0LWFwaS50a2Uuc3ZjgglhdWRpdC1hcGmH
BMCoAQGHBMCoAQIwDQYJKoZIhvcNAQELBQADggEBAERxawkDwoANxklh/Zg0Lzw1
i2oorBwIpIIwzJOYK9bazmyOJIJWXDyJHviZTj4heYZ82PDw0LIxMD7+n0WPe5Op
si1xrfRZLKUBS//TsK42HU/01HQG6GI1EL2Qj4vzYKl+ITTWcKbuTciOWkzicbFm
9PFjg1fnnNlg/vc3AEr5BWPrf0xPxHcccaPxBTdWfmDXwrK9BY1rJeqRaVervYyU
A0RjqINaPbvLbJs+RGti01qY5KNINsBFkgjI+ezLk1VoHwEFcpbi/UE+zyhIcuQM
LvmOIRgFmGb9JHWQsQMkLSNaZJXegSKOUymFRS4sYugBrrORgO+XsUsP+IC33EU=
-----END CERTIFICATE-----
`,
		ServerKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA0To4AN+1Y++ntEQsVqOgijOmTJDepzsoq5tZCI3TvZbIcGVE
+uuaFnNqXIVJiKSPQY44W+xCdChb9ZBBcPnMFYaRjKXFUw9JCgFaHumoc3WopgDt
NncuZbAWgdcgRq92elhmsK0SrWcZdkXd9O6AYLfCXD4KTOk1pwM8jfGXbN5OQVBw
AdoYA78aoTppSfe2T89HoB8gOBvolwyWMQrhFDA+NLdjkjoFJp6TVqiCog739qOX
xxzwF+A//zoeJBCOaYYH3xROJOazxBv1v4IEwIxS/N2uVazXf9HugsPD41xDe6IG
4XSiJNqxVSJDpfhoPZ/If6QiRZNBNsII8EhwWwIDAQABAoIBAHxSmcB7LW9qkj4d
XDNHGzfDpQdN9d9s2afOsULR84VuOUAADNMEeBTrE8JSOPWGwtDzTy3f30TiTnJF
+3zEEa6gvY/TWMMYDBBz6TcBJGUm/sVVO57nhpxJ27XfPICD9mSQbsSAst5yqsCK
l27Y4VOXHVhuTPSShpzs4Wh+q7m182qxgZSI/i7wZDeBcJ5rwjzLT/OtsXchlFJv
Wuh3yJdev7jee6GsQLEXTh7egkuTeha46GHLVKIw7cWpvvfihGXjwWppIpOlHAs1
0nm+C5kSGRNEnSVqVtuyjTvU3oL4OeayX6O23ZypsdUazBnRU7/yxs8byBAvnDN7
ZbnJ4IECgYEA3BnVsJAQVOQ4S+rwHJ7cabr33apaZatx2qRCz76TjGxKJKCucvoF
nNdHYjSJaaRacpvo7qczAn7KnrVRV6XZxOPljHndZOGHI+ASsHM+ggeiGNdXPLga
Bg0xY4LB0GrfSFsOeLmr/wJFhXvE2PBDbi+7Xtdbu84Jx+C8x6c/52ECgYEA81pe
TrPCg1gEy3p4XlBHiXpq3HG29/hqm2RSOhiSA4AUH39Mr9qdY9FVHqjXqH+NzY+C
xoFajItccbCfa3yWSmgirPl3Chnk6J91WOrPIYo3UxruHHURkuqyQpEG1vsDqS8x
5DuzA3wmWHeiK/OVhFw8fA1EURbCLX4Y9/3sPTsCgYEAt5hJOqWBKf7QwurvbUBK
ZVuTZHC0RPiU2SeJn5a9pzWxRwpNWhUaV3Dv14gxU1DILa/FxkTr9NnklCx09KKg
m27qTJpexcamHpPLoyoOaxKI8SW3XCvJcWDR4MrydrKfG9Dwql1ejjHL4gOS+M65
wJDKP9tuZqMAX50ke3oUgSECgYEAvfMU19PCiakjnPxMhjmdjTbrwIRWcDoMZJyq
/kP9fRP1IRzJbB51aduU2GisZQ6jTget6WpAlqGVF/zogwrteMjOADx86XqKzfTh
sPUflqt/QbQsljPe8tGVMRZow1eNZPq8s8zRj2/Rso43HJ12YKlqYN4uLClaig36
SWB/jM8CgYEAuzh0HF3S7Og66FZESKx07FeUCrHI781c9HDPe0fheuHMzJX++NMV
48D2W5Z6Q7jSJkladVlnn2HD/hbg1ydaSa+1vfqEZEcIi0IcxoZuoV0tABK4FtQW
nW+EuroD3jzFfl+NfyGu1PvOC0O7H1Uc6GdwYNKoUaymBYBrdHkh6Qk=
-----END RSA PRIVATE KEY-----
`,
		FrontProxyCACert: `-----BEGIN CERTIFICATE-----
MIIC0DCCAbigAwIBAgIBADANBgkqhkiG9w0BAQsFADAZMRcwFQYDVQQDEw5mcm9u
dC1wcm94eS1jYTAeFw0yMDA3MDgwNDE1MTFaFw0zMDA3MDYwNDE1MTFaMBkxFzAV
BgNVBAMTDmZyb250LXByb3h5LWNhMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEAuGw9bxniDlzAcfiCfZMliO8ekZxn8v6v+mUFsGP9x/meu6gRn0B+2kIA
9T4f1Tv6FUEEz+wFYyO9qlOsPS/4foQyGadQeDL91GDk1l3U4oTYenY/dgIs0EFt
6lSfjp7bHG9qkAbYXY01rcLo/k922Uziu1al32GixcuNQmN/tME3hRqGNRMLYSrt
1IHGlL0zz6NTl+cM9xXrTzalhLkP2f3rmRnyRz6IMMKHN5eFTbJNod8oD4NWhIg2
AM0+IBb0yoWjoW6jOFXkViH/diYn6GJgW9iyl0WJ5l6zZKvDggTTD0riozLusxI1
lUmzQlF9yQVLgO/4Ok4wpDfWhKCE9QIDAQABoyMwITAOBgNVHQ8BAf8EBAMCAqQw
DwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAjjTVJYWh4Wk6PY/b
7Sres7xdGrJB6elycABViec6GjCzi58Zt4jqTnxssTZSGNJdpQSht+YVDGJcOuU8
9WNT7SeoKw4+6yK+ee6IM+YQy0okjZARQsiqqWvpniJo0CaaXxxZRT+YFlFBqVVc
Ud8d100tjpPlZDjiIrDEGVkJyyy8/nst3/9cChMH5cgVkFOuXizQ8Vv+a2mkuaV2
87EKaq8x2T0gxJnxSNd2GDbn6lGRpqa8T8oJUqtNpq8y9vW5rCDYUoc1GJwVMNpR
O8CoiBzrJEWf4HAsuUfGOHctkfcq1cT5VJgfcqKMzlQ2Cvwe8NKD+xfdX7umgIlg
hqSn7g==
-----END CERTIFICATE-----
`,
		ETCDValues: ETCDValues{
			Hosts: []string{"172.19.1.0", "172.19.1.31"},
			CACert: `-----BEGIN CERTIFICATE-----
MIICwjCCAaqgAwIBAgIBADANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQDEwdldGNk
LWNhMB4XDTIwMDcwODA0MTUxMloXDTMwMDcwNjA0MTUxMlowEjEQMA4GA1UEAxMH
ZXRjZC1jYTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALHny18d3F5W
U+YVgYYX2KpDOqbChoE2OQIZ5EJ7NRFXgvlWVAbycjIwMBk5UPoHXkroERzNml0B
TpSnnWydMxe+ipnsC/pJ9mdjGll9JzZZf752UpDMaIEp1DSAHySkMDvXGk9nKpb8
9qhYgOZCRtRvjy/Oc7FSnnCUqnwZG3wWiK7cffviIwoQxnOlubC5Q/JHSqoaGZdm
RN77JxLOl8b19fgqoGrxg5bXupGeUF5abd0oxS5owC4tZ/QdtmlizOhaQz7rLYFt
m+IOjv3k3x0IPRUMR1X3NIFIDBM8R04vo7FZC36UTwoELwHHUbg+QujfkJSFM59H
PxN2g1mz/lcCAwEAAaMjMCEwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB/wQFMAMB
Af8wDQYJKoZIhvcNAQELBQADggEBACcM/4Y7tvPdMD7nnJkphS/zbRwT5x5IFhk9
2JRZYOtcuVfmmHWiB9X17nkCuMNmem0RpH2ZRjqLQUTC33bLxvWe4QsL/36IlFDE
+toiIEB9gutAnrWCYLtbqCKi8NLShZQrmPQQzN/eQsgi5o71DuSyrOfeXEkvQajB
YDhTs9MyRqJBz0CnZFaiINBXU1kNDI53cFMDH1dLKfwjKFt01+QUa/D7yq6LcQXr
qIjlP8rkarhPPGjcgJsRCexUm5dyCgN1tQGdD8+I882q6OwUCPg4WHqwxz3+n8xd
+IyQgNoP7swDz9+rhRKEvybDYOSo1FKwc4yabLcjZcq2OUyFVMI=
-----END CERTIFICATE-----
`,
			ClientCert: `-----BEGIN CERTIFICATE-----
MIIC+TCCAeGgAwIBAgIICbMVk3tzfdowDQYJKoZIhvcNAQELBQAwEjEQMA4GA1UE
AxMHZXRjZC1jYTAeFw0yMDA3MDgwNDE1MTJaFw0yMTA3MDgwNDE1MTJaMD4xFzAV
BgNVBAoTDnN5c3RlbTptYXN0ZXJzMSMwIQYDVQQDExprdWJlLWFwaXNlcnZlci1l
dGNkLWNsaWVudDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMdX1A1r
QgAde2OlF8gevlPE+tyPZkreXpZFTIvTA7CH/kpgtgpyKjscqsB3uHPExMJjEp3W
fvYLJqzRODj6xSkXwU6eEz/2N2kTB+DAiuGmD6WeYkPeCici9HcaiQBEsQEmNr1x
ClzF9A13jErda3ImxCIGi5Kalrrs+fPo++RZBHtz6GJBfNy3+0B+vgBFqYi9gHd4
rHkPE/3ZF/SRpXDR2iMBs60dXBz6Rq/Y1cKgyNrwhyaGhVMpGzsn0+QhQasFidid
TVvZJ8C0MCcYkpiDuVVOGMcATMJMuGR19/BRFDAubV3NqFVqJd+xiuzG1Alv04xE
MdyDHeIQ1kP+cmkCAwEAAaMnMCUwDgYDVR0PAQH/BAQDAgWgMBMGA1UdJQQMMAoG
CCsGAQUFBwMCMA0GCSqGSIb3DQEBCwUAA4IBAQBDA07Jw51R5hkRb0EbkhREdGf8
PzKedbKZtZVjD4uvhApkKLJ0UD4jnTGqREEONS3uyfuK7FM+54VBz4MuIwpcOo6z
1Wxp9bZqEJH77CiFJm+fOqec8S6Ga8JkHXH7O/2ccNJX+JkLVEYxjq16Rx0Zvnh2
8E0RxgknsCl9rvs0o25y1mJ7NQbK7/8Vch/8tyiLfKt8CoABXiMEbzZ38imRNpod
P1Oor4i/FmkQKy9RCU2gNntgO6pRVIDi5nAT9rMs0tDW6NgW7nbUKGHvfqXcX78E
h3bZ0/IlFuQZdnLW/sYVg+UaIsA4Dp7Qtm27ks/D848uDvAIvoCns7Kl9uWF
-----END CERTIFICATE-----
`,
			ClientKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAx1fUDWtCAB17Y6UXyB6+U8T63I9mSt5elkVMi9MDsIf+SmC2
CnIqOxyqwHe4c8TEwmMSndZ+9gsmrNE4OPrFKRfBTp4TP/Y3aRMH4MCK4aYPpZ5i
Q94KJyL0dxqJAESxASY2vXEKXMX0DXeMSt1rcibEIgaLkpqWuuz58+j75FkEe3Po
YkF83Lf7QH6+AEWpiL2Ad3iseQ8T/dkX9JGlcNHaIwGzrR1cHPpGr9jVwqDI2vCH
JoaFUykbOyfT5CFBqwWJ2J1NW9knwLQwJxiSmIO5VU4YxwBMwky4ZHX38FEUMC5t
Xc2oVWol37GK7MbUCW/TjEQx3IMd4hDWQ/5yaQIDAQABAoIBAEuKUdY/9gPbKD4L
euagF9ngxHu+b7IhcFCbtDUsYrKL5mZDHdj5iF1cMKy88Y4VW8aaCpz7gqM0eJR2
uEqHZR18Xgmvx3WulTLhPlQCQt7dsgI98djsFIbeLSvRR2dtoHHiDE5fvuCHN0O/
MM/R9nyMijimqE+G+2Wcld/DQ2k1UrZt38haEBmpls15joAyR5ltfmnHIiu8pGjA
e5lXpByYJOTkig3WngT4Pea1S8EzSG7/FSTaE+YopOZYedfrK5EPS52PBh/035LB
4LuVviq5WnSSj50EEgAPuLfAJkmtNCIeyjjw3MYYcWZm4cw8kQkhLGTElwZL7I/W
9+dC1QECgYEAz6TnybBHbHcTE9VZgncrybfLqQntRMJpNRQGe6vb74o+GfvR0CTa
A+NPbJ+i7peTbrgkpOyjU7BGt006J1vSjKkLQm3D+T6Wu6EcAeri6Ax0br65875/
t80Bck1u2VNdDUKGbWIUmbq5L6n8xNRG/S1zxt8366WIII1BVJ4jGuECgYEA9cQJ
0v2eZn4ZF2GTbPm1Uk9+1iBU7X4RrlPaLEdbYwrT7wX4WYOeFnQ23f2/MOA0hB4f
Hp8e0XBzq9FXIiyI7qCUFqbN3oqDqq/J8/rPgX7mvNp/htRzMQqCxkaZyszexhus
DsnUxJpzsi2puLbzYRU2bcc7A1ZhFPBqSMFnEIkCgYEAxlxjrVIAOUbba7QByVes
9gvlu/smtev+81zJALXh6kinIM1m1y+NMe+hvUQXM1SL3FZ/cyo4sMcSFyPobRcT
4ZnPjssrTtySxHWZOt+3mjPjurPrEkWm3uAyBI6iWWyA1Rz7LF1pDp7JMBZ0yVd8
5XnZQ1+UnUHJ9Ebc5UgYgyECgYBeZVGwrKEulLhmdp/9d1vhpY7+B6UXIxj/cE2Y
PEAppnsIj654E0mhLjn0wgWxbCw/oqIpTMy1fQ+wqtNg2OIvZZX7rU3dcC5qXOKg
x3D0OB2/49yfmPI1n+fi6DMh5dabGP3ZIESMzxYpFcjencaVP2SMfb145rCbLkxk
bYR8oQKBgD4WzJUympevsEcjXNGKgt0zLNSTnIgImMJ4sIFpDTCMmS/S/e0TAdnJ
7EMj4K59BbWUNL+YS8eFbvtW8UqnqOOwlqKLPiin3DiBeTDflo4fjK7U65sQK79f
15+lXPHGtJpvizJwVREw2uKY94XgYd7w09GgAOQWzw0+W++QxAA3
-----END RSA PRIVATE KEY-----
`,
		},

		CertsName:        "certs",
		OIDCClientSecret: "ab",
		OIDCIssuerURL:    "oidcIssuerURL",
		USEOIDCCA:        false,
		RegistryPreifx:   "docker.io/tkestack",
		Token:            "token",
		TenantID:         "default",
		AdminUsername:    "admin",
		AdminPassword:    "password",
	}
}

func TestSave(t *testing.T) {
	option := newValues()
	err := save(option)
	assert.Nil(t, err)
	fmt.Println(err)
}

func TestInstall(t *testing.T) {
	option := newValues()
	err := Install(option)
	assert.Nil(t, err)
	fmt.Println(err)
}