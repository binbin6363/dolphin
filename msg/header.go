package msg

type BaseHead struct {
}

func (this *BaseHead) GetHead(packet *gotcp.Packet) (error err) {
	return nil
}

func (this *BaseHead) SetHead()(packet *gotcp.Packet) (error err) {
    return nil
}
