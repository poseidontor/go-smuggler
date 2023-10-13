## HTTP Request Smuggling Detector

*#### THIS PROJECT IS STILL IN DEVELOPMENT* 

Inspired from [https://github.com/anshumanpattnaik/http-request-smuggling](https://github.com/anshumanpattnaik/http-request-smuggling) which is written in python.

HTTP request smuggling is a high severity vulnerability which is a technique where an attacker smuggles an ambiguous HTTP request to bypass security controls and gain unauthorized access to performs malicious activities, the vulnerability was discovered back in 2005 by [watchfire](https://www.cgisecurity.com/lib/HTTP-Request-Smuggling.pdf) and later in August 2019 it re-discovered by [James Kettle - (albinowax)](https://twitter.com/albinowax) and presented at [DEF CON 27](https://www.youtube.com/watch?v=w-eJM2Pc0KI) and [Black-Hat USA](https://www.youtube.com/watch?v=_A04msdplXs), to know more about this vulnerability you can refer his well-documented research blogs at [Portswigger website](https://portswigger.net/research/http-desync-attacks-request-smuggling-reborn). So the idea behind this security tool is to detect HRS vulnerability for a given host and the detection happens based on the time delay technique with the given permutes.

### Security Consent
It's quite important to know some of the legal disclaimers before scanning any of the targets, you should have proper authorization before scanning any of the targets otherwise I suggest do not use this tool to scan an unauthorized target because to detect the vulnerability it sends multiple payloads for multiple times which means if something goes wrong then there is a possibility that backend socket might get poisoned with the payloads and any genuine visitors of that particular website might end up seeing the poisoned payload rather seeing the actual content of the website. So I'll highly suggest taking proper precautions before scanning any of the target website otherwise you will face some legal issue.

### Installation
```
git clone https://github.com/poseidontor/go-smuggler
go run cmd/go-smuggler/main.go
```

### Options
```
usage: main.go [-u URL] [-f FILE CONTAINING MULTIPLE URLS] [-t TIMEOUT] 

HTTP Request Smuggling vulnerability detector

optional arguments:
  -u URL    set the target url
  -f FILE    provide path to file containing multiple URLS
  -t URL    set timeout
```
