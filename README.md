# Calendar 事务提醒服务

### install

```shell
go get -u github.com/duguying/calendar
```

### configure

- add you `api.key` into `~/.calendar` directory, support yunpian
- add you `sms.tpl` into `~/.calendar` directory

`sms.tpl` as follow

```
xxxxxxx%dxxxxxxx%dxxxxx
```
the first %d will be time, the second %d will be event description.

### service

command `calendar serve` will start calendar as a service. you can run it as a deamon service via systemd from file `calendar.service`.

### client

##### help

```
calendar help
```

##### add

```
calendar add 
```
you can get detail via command `calendar help add`

##### remove

```
calendar remove id
```
you can get id via command `calendar list`

##### list

```
calendar list
```
list all your reminder in your calendar.

### License

MIT License


