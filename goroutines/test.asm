# command-line-arguments
"".main STEXT size=74 args=0x0 locals=0x10
	0x0000 00000 (/home/sergei/Work/go/test/test.go:5)	TEXT	"".main(SB), ABIInternal, $16-0
	0x0000 00000 (/home/sergei/Work/go/test/test.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (/home/sergei/Work/go/test/test.go:5)	CMPQ	SP, 16(CX)
	0x000d 00013 (/home/sergei/Work/go/test/test.go:5)	JLS	67
	0x000f 00015 (/home/sergei/Work/go/test/test.go:5)	SUBQ	$16, SP
	0x0013 00019 (/home/sergei/Work/go/test/test.go:5)	MOVQ	BP, 8(SP)
	0x0018 00024 (/home/sergei/Work/go/test/test.go:5)	LEAQ	8(SP), BP
	0x001d 00029 (/home/sergei/Work/go/test/test.go:5)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/home/sergei/Work/go/test/test.go:5)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/home/sergei/Work/go/test/test.go:5)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/home/sergei/Work/go/test/test.go:6)	PCDATA	$0, $0
	0x001d 00029 (/home/sergei/Work/go/test/test.go:6)	PCDATA	$1, $0
	0x001d 00029 (/home/sergei/Work/go/test/test.go:6)	CALL	runtime.printlock(SB)
	0x0022 00034 (/home/sergei/Work/go/test/test.go:6)	MOVQ	$3, (SP)
	0x002a 00042 (/home/sergei/Work/go/test/test.go:6)	CALL	runtime.printint(SB)
	0x002f 00047 (/home/sergei/Work/go/test/test.go:6)	CALL	runtime.printnl(SB)
	0x0034 00052 (/home/sergei/Work/go/test/test.go:6)	CALL	runtime.printunlock(SB)
	0x0039 00057 (/home/sergei/Work/go/test/test.go:7)	MOVQ	8(SP), BP
	0x003e 00062 (/home/sergei/Work/go/test/test.go:7)	ADDQ	$16, SP
	0x0042 00066 (/home/sergei/Work/go/test/test.go:7)	RET
	0x0043 00067 (/home/sergei/Work/go/test/test.go:7)	NOP
	0x0043 00067 (/home/sergei/Work/go/test/test.go:5)	PCDATA	$1, $-1
	0x0043 00067 (/home/sergei/Work/go/test/test.go:5)	PCDATA	$0, $-1
	0x0043 00067 (/home/sergei/Work/go/test/test.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x0048 00072 (/home/sergei/Work/go/test/test.go:5)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 34 48  dH..%....H;a.v4H
	0x0010 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 e8 00 00  ...H.l$.H.l$....
	0x0020 00 00 48 c7 04 24 03 00 00 00 e8 00 00 00 00 e8  ..H..$..........
	0x0030 00 00 00 00 e8 00 00 00 00 48 8b 6c 24 08 48 83  .........H.l$.H.
	0x0040 c4 10 c3 e8 00 00 00 00 eb b6                    ..........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=8 runtime.printlock+0
	rel 43+4 t=8 runtime.printint+0
	rel 48+4 t=8 runtime.printnl+0
	rel 53+4 t=8 runtime.printunlock+0
	rel 68+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename.main SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
	0x0000 03 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+74
	rel 27+4 t=29 gofile../home/sergei/Work/go/test/test.go+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 0f 04 0e 03 05 01 08 02 20 00                 ......... .
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
