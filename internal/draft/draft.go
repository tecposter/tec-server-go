package draft

import (
	"github.com/tecposter/tec-node-go/internal/com/bin"
	"github.com/tecposter/tec-node-go/internal/com/dto"
	"time"
)

const (
	titleLenMax = 100
)

type draft struct {
	PID     dto.ID      `json:"pid"`
	Changed time.Time   `json:"changed"`
	Cont    dto.Content `json:"cont"`
}

type draftItem struct {
	PID     dto.ID    `json:"pid"`
	Changed time.Time `json:"changed"`
	Title   string    `json:"title"`
}

const (
	idSize   = 16
	timeSize = 8
)

func newDrft(pid dto.ID, cont dto.Content) *draft {
	return &draft{
		PID:     pid,
		Changed: time.Now(),
		Cont:    cont}
}

func (d *draft) marshal() ([]byte, error) {
	id, data, err := d.marshalPair()
	return append(id, data...), err
}

func (d *draft) marshalPair() ([]byte, []byte, error) {
	//fmt.Println("marshal:", d.PID.Bytes())
	id := d.PID.Bytes()
	changed := bin.TimeToBytes(d.Changed)

	typ := []byte{byte(d.Cont.Typ)}
	body := []byte(d.Cont.Body)
	cont := append(typ, body...)

	//fmt.Printf("id: %d, changed: %d, typ: %d\n", len(id), len(changed), d.Cont.Typ)
	return id, append(changed, cont...), nil
}

func (d *draft) unmarshal(src []byte) error {
	id := dto.ID(src[0:idSize])
	data := src[idSize:]
	return d.unmarshalPair(id, data)
}

func (d *draft) unmarshalPair(id, data []byte) error {
	changed := bin.BytesToTime(data[:timeSize])
	typ := dto.ContentType(data[timeSize])
	body := string(data[timeSize+1:])

	d.PID = id
	d.Changed = changed
	d.Cont.Typ = typ
	d.Cont.Body = body

	return nil
}

func (d *draft) Title() string {
	return d.Cont.Title(titleLenMax)
}
