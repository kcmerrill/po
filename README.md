# Po

[![Build Status](https://travis-ci.org/kcmerrill/po.svg?branch=master)](https://travis-ci.org/kcmerrill/po)

![po](po.jpg)

## What is it

Besides a work in progress, vaporware and a pie in the sky idea, `po` is a way for your side scripts to send in and aggregate a wide variety of metrics. Text, words, numbers, dates, time intervals and heartbeats to name a few. Data can be validated, based on it's type. Once the data is collected, create dashboards based off of yaml configuration files to display metrics that you care about. 

## Collect Data

As a package, you can configure PO to create entity objects to track. As a standalone app, you can use http web requests. A few examples:

```golang
user := po.Attributes{
  "email": "email-address",
  "username": "string",
  "phone": "phone-number",
  "visits": "number",
}
po.E("kcmerrill", user).A("email").Set("kcmerrill@gmail.com")
```

```sh
# create an entity object
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com
# add new attributes
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/status/string/ok
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/response.code/number/200
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/total.checks/number/0
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/status.404/number/0
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/status.200/number/0
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/elapsed.time.in.ms/list/1
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/last.checked/date/now
$ curl -v http://kcmerrill.com | curl -d @- -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/page.contents/text

# update attributes
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/load.time/list/1s
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/response.code/302
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/total.checks/increment
$ curl -X POST http://po.kcmerrill.com/http.check.kcmerrill.com/elapsed.time.in.ms/3

# grab entity object
# curl http://po.kcmerrill.com/http.check.kcmerrill.com
```

At this point, we have an entity that we can create and evolve over time. Now, lets layer on a customizable dashboard. 

```yaml

http.check.kcmerrill:
  text: 
    width: 20%
    innerHTML: status
    bg-color: {{ if status == "ok" }} green {{ else }} red {{ end }}
  pie-chart:
    width: 20%
    data: status.404 status.200 status.500
  bar-chart:
    width: 20%
    x-axis: Date
    y-axis: Time in Milliseconds
    data: elapsed.time.in.ms
  text: 
    width: 40%
    innerHTML: page.contents
```

## Binaries || Installation

[![MacOSX](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/apple_logo.png "Mac OSX")](http://go-dist.kcmerrill.com/kcmerrill/po/mac/amd64) [![Linux](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/linux_logo.png "Linux")](http://go-dist.kcmerrill.com/kcmerrill/po/linux/amd64)

via golang:

`$ go get -u github.com/kcmerrill/po`
