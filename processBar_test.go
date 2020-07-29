package processBar

import (
    "testing"
    "time"
)

func TestBar_Play(t *testing.T) {
    var bar Bar
    bar.NewBar(0, 100)
    for i:= 0; i<=100; i++{
        time.Sleep(1*time.Second)
        bar.Play(i)
    }
    bar.Finish()
    //
    // bar.NewBarWithCustomGraph(0,100,"#")
    // for i:= 0; i<=100; i++{
    //     time.Sleep(1*time.Second)
    //     bar.Play(i)
    // }
    // bar.Finish()
    //
    // bar.NewBarWithCustomGraphAndAccuracy(0,100,5,"#")
    // for i:= 0; i<=100; i++{
    //     time.Sleep(1*time.Second)
    //     bar.Play(i)
    // }
    // bar.Finish()
}
