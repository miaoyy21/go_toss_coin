package toss

import (
	"bytes"
	"fmt"
)

type Schema byte

func (o Schema) Byte() byte {
	if o == SchemaPositive || o == SchemaNegative {
		return byte(o)
	}

	panic(fmt.Sprintf("unreachable Schema of %q", string(o)))
}

var SchemaPositive Schema = 'P'
var SchemaNegative Schema = 'N'

type Toss struct {
	rows    []int
	records []Schema

	pattern func(int) Schema
}

// NewToss rows: 数据的历史记录，越靠前代表离现在越近
func NewToss(rows []int, pattern func(int) Schema) *Toss {
	toss := &Toss{
		rows:    rows,
		records: make([]Schema, 0, len(rows)),
		pattern: pattern,
	}

	for _, row := range rows {
		toss.records = append(toss.records, pattern(row))
	}

	return toss
}

// Add 新纪录添加到列表的首部
func (o *Toss) Add(row int) {
	rows := make([]int, 0, len(o.rows)+1)
	records := make([]Schema, 0, len(o.records)+1)

	rows = append(rows, row)
	rows = append(rows, o.rows...)

	records = append(records, o.pattern(row))
	records = append(records, o.records...)

	o.rows = rows
	o.records = records
}

func (o *Toss) String() string {
	buf := new(bytes.Buffer)
	buf.WriteByte('\n')

	// 数据记录
	buf.WriteString("  数据记录: ")
	for i, row := range o.rows {
		if i != 0 {
			buf.WriteByte(',')
		}

		buf.WriteString(fmt.Sprintf("%02d", row))
	}
	buf.WriteByte('\n')

	// 抛硬币模式
	buf.WriteString("  模式转换: ")
	for i, record := range o.records {
		if i != 0 {
			buf.WriteByte(',')
		}

		buf.WriteString(fmt.Sprintf("%2s", string(record)))
	}

	buf.WriteByte('\n')
	return buf.String()
}
