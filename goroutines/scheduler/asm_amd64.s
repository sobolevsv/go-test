
TEXT runtime·rt0_go(SB),NOSPLIT,$0

	LEAQ	runtime·m0+m_tls(SB), DI
	CALL	runtime·settls(SB)


	get_tls(BX)
	LEAQ	runtime·g0(SB), CX
	MOVQ	CX, g(BX)
	LEAQ	runtime·m0(SB), AX

	// save m->g0 = g0
	MOVQ	CX, m_g0(AX)
	// save m0 to g0->m
	MOVQ	AX, g_m(CX)


    CALL	runtime·schedinit(SB)

    // create a new goroutine to start program
    MOVQ	$runtime·mainPC(SB), AX		// entry
    PUSHQ	AX
    PUSHQ	$0			// arg size
    CALL	runtime·newproc(SB)

	// start this M
	CALL	runtime·mstart(SB)


DATA	runtime·mainPC+0(SB)/8,$runtime·main(SB)
GLOBL	runtime·mainPC(SB),RODATA,$8
