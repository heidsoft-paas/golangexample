package main

//go run glog_ex.go -logtostderr=true
//go run glog_ex.go -logtostderr=true -v=2
//see https://github.com/golang/glog/blob/master/glog.go#L996

import(
    "flag" 
    "github.com/golang/glog"
)
func main(){
        flag.Parse()
        defer glog.Flush()
        glog.Info("Prepare to repel boarders")
        if glog.V(2) {
            glog.Info("in level 2")
        } 
}
