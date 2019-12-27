# Learn to "Heap ballast" in Golang GC

```shell script
$ go build .

# normal case
$ GODEBUG=gctrace=1 .\golang-heap-ballast -ballast-mb=0
gc 1 @0.002s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 2 @0.007s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
gc 3 @0.012s 0%: 0+0.99+0 ms clock, 0+0/0/0+0 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
gc 4 @0.018s 0%: 0+0.99+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 5 @0.022s 0%: 0+0.99+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 6 @0.027s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 7 @0.032s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 8 @0.035s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 9 @0.039s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 10 @0.043s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 11 @0.046s 0%: 0+0.99+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 12 @0.051s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 13 @0.055s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 14 @0.060s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 15 @0.066s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 16 @0.070s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 17 @0.073s 1%: 0.99+0+0 ms clock, 7.9+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 18 @0.077s 1%: 0+0.99+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 19 @0.081s 1%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 20 @0.086s 1%: 0+0.99+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 21 @0.091s 1%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 22 @0.094s 1%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 23 @0.098s 1%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
gc 24 @0.104s 0%: 0+0.99+0 ms clock, 0+0/0/0+0 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 25 @0.109s 0%: 0
elapsed 106.7172ms

# add heap ballast
$ GODEBUG=gctrace=1 .\golang-heap-ballast -ballast-mb=1000

gc 1 @0.017s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 1000->1001->1001 MB, 1001 MB goal, 8 P
gc 2 @0.021s 0%: 0+0+0 ms clock, 0+0/0/0+0 ms cpu, 1002->1017->1015 MB, 2002 MB goal, 8 P

elapsed 6.982ms
```

- ref: https://github.com/golang/go/issues/23044
