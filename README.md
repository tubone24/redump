# redump

![redmup_logo](./docs/images/redump_logo.png)

> REDUMP is a tool to migrate data in your Redmine without admin accounts.

You can use the API to retrieve tickets in Redmine, save them in JSON format, and migrate them to another Redmine.

[![license](https://img.shields.io/github/license/tubone24/redump.svg)](LICENSE)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
![TestAndBuild](https://github.com/tubone24/redump/workflows/TestAndBuild/badge.svg)

```
Usage:
  redump migrate [-i|--issue <number>]
  redump list
  redump dump [-c|--concurrency] [-i|--issue <number>]
  redump restore [-i|--issue <number>]
  redump clear [-o|--old]
  redump -h|--help
  redump --version

Options:
  -h --help                  Show this screen.
  -c --concurrency           Concurrency Request Danger!
  -i --issue                 Specify Issues
  -o --old                   Old Server
  --version                  Show version.
```
