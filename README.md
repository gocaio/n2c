# n2c

[![Go Report Card](https://goreportcard.com/badge/github.com/gocaio/n2c)](https://goreportcard.com/report/github.com/gocaio/n2c)

Tool to convert `nmap` XML output to CSV:


## USAGE

Options:

```sh
$ n2c -h
  -xml string
    	XML file to use
```

Convert XML file to CSV:

```sh
$ n2c -xml testdata.xml

IP,Host,Proto,Port,Service,Product,NSE Script ID
192.168.1.100,nuc.int,tcp,22,ssh,OpenSSH 7.4p1 Debian 10+deb9u7
192.168.1.100,nuc.int,tcp,2222,ssh,OpenSSH 7.7
192.168.1.100,nuc.int,tcp,3000,ppp,
192.168.1.100,nuc.int,tcp,3306,mysql,MySQL 5.7.26
192.168.1.100,nuc.int,tcp,8080,http-proxy,
192.168.1.100,nuc.int,tcp,9090,http,Node.js Express framework
```