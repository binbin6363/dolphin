package dolphin



type Processor struct {
    ProcessorConfig *Config
    Connections map[string] *Conn[]
    Handlers map[string] CallBack
}

func NewProcessor(config Config) (processor *Processor, err error) {
    return &Processor{
        ProcessorConfig: config,
        Connections : make(map[string]*Conn[]),
        Handlers : make(map[string] CallBack)
    }
}

func (self *Processor)ConnServer(addr, protocol, name, net string) (conn *Conn, err error){

}

func (self *Processor)NewServer(addr, protocol, name, net string) (err error){

}

func (self *Processor)RequestData(name string, request byte[], hash string) (result byte[], err error){

}

func (self *Processor)SendData(name string, request byte[], hash string)(err error){

}

func (self *Processor)SetHandler(name string, callback CallBack){
    Handlers[name] = callback
}