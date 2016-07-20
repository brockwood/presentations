	vlp := loopPacketRaw{}
	b := loop[3:99]
	buf := bytes.NewReader(b)
	err = binary.Read(buf, binary.LittleEndian, &vlp)
	if err != nil {
		return nil, err
	}