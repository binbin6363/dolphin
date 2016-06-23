# dolphin
dolphin is a high performance tcp net framework written in golang

It is not ready for production use! just in development

low-level tcp use gotcp(https://github.com/gansidui/gotcp.git).

we are engaged in making it easier to build tcp service using golang,
and we just show some simple examples as follows:

## usage example

    package main

    import (
        "fmt"
        "log"
        "net"
        "time"

    )

    // msgcenter
    func (self *MsgCenterCallBack)OnConnect(conn *Conn) bool {

    }

    func (self *MsgCenterCallBack)OnMessage(conn *Conn, packet *Packet) bool{

    }

    func (self *MsgCenterCallBack) OnClose(conn *Conn) {

    }

    func (self *MsgCenterCallBack) OnTimeOut(conn *Conn) {

    }


    // push server
    func (self *PushCallBack)OnConnect(conn *Conn) bool {

    }

    func (self *PushCallBack)OnMessage(conn *Conn, packet *Packet) bool{

    }

    func (self *PushCallBack) OnClose(conn *Conn) {

    }

    func (self *PushCallBack) OnTimeOut(conn *Conn) {

    }


    // dispatcher 
    func (self *DispCallBack)OnConnect(conn *Conn) bool {

    }

    func (self *DispCallBack)OnMessage(conn *Conn, packet *Packet) bool{

    }

    func (self *DispCallBack) OnClose(conn *Conn) {

    }

    func (self *DispCallBack) OnTimeOut(conn *Conn) {

    }

    // 框架负责维护session
    // 直接能够按一定规则hash
    func main() {

        addr_srv1 := "127.0.0.1:12345"
        addr_srv2 := "127.0.0.1:12346"
        addr_srv3 := "127.0.0.1:12347"
        addr_srv4 := "127.0.0.1:12348"

        config := &Config{
            PacketSendChanLimit : 200000,
            PacketRecvChanLimit : 200000,
            PacketMaxRecvLen    : (1 * 1024 * 1024),
            PacketMaxSendLen    : (1 * 1024 * 1024)
        }

        processor, err := NewProcessor(config)
        if err != nil {
            return
        }

    //    processor.SetHandler("timeout", TimeoutHandler)
        processor.SetHandler("msgcenter", MsgCenterCallBack)
        processor.SetHandler("pushserver", PushCallBack)
        processor.SetHandler("dispatcher", DispCallBack)
        if client1, err := processor.ConnServer(addr_srv1, "binary", "msgcenter", "tcp"); err != nil {
            return
        }
        if client2, err := processor.ConnServer(addr_srv2, "protoc", "msgcenter", "tcp"); err != nil {
            return
        }
        if client3, err := processor.ConnServer(addr_srv3, "binary", "pushserver", "tcp"); err != nil {
            return
        }
        if server1, err := processor.NewServer(addr_srv4, "binary", "dispatcher", "tcp"); err != nil {
            return
        }

        // 启动调度器
        processor.EventLoopStart()
        return
    }
