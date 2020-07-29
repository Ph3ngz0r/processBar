package processBar

import "fmt"

type Bar struct {
    percent int64  //百分比
    cur     int64  //当前进度位置
    total   int64  //总进度
    rate    string //进度条
    graph   string //显示符号
}

// 创建进度条
func (b *Bar) NewBar (start, total int64) {
    b.cur = start
    b.total =total
    if b.graph == "" {
        b.graph = "█"
    }
    b.percent = b.getPercent()
    for i := 0; i < int(b.percent); i += 2 {
        b.rate += b.graph
    }
}

// 自定义进度条图案
func (b *Bar) NewBarWithCustomGraph(start, total int64, graph string) {
    b.graph = graph
    b.NewBar(start, total)
}

// 进度条百分比
func (b *Bar) getPercent() int64 {
    return int64(float32(b.cur) / float32(b.total) * 100)
}

// 播放进度
func (b *Bar) Play (cur int64) {
    b.cur = cur
    last := b.percent
    b.percent = b.getPercent()
    if last != b.percent && b.percent % 2 == 0 {
        b.rate += b.graph
    }
    fmt.Printf("\r[%-50s]%3d%%  %8d/%d", b.rate, b.percent, b.cur, b.total)
}

func (b *Bar) Finish () {
    fmt.Println()
}
