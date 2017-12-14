## dingrun
一个简单的小程序，当程序异常退出时发出DingDing通知

## Install
```
go get -v github.com/codeskyblue/dingrobot/dingrun
```

## Usage
```bash
# 最好直接设置到 ~/.bashrc 里面去
export DINGROBOT_TOKEN="xxxxxxx-robot-token--"

# 直接后面接上运行的命令即可
dingrun sleep 10
# 程序结束后会发出通知，只有当Ctrl+C结束时，通知不会发送。
```