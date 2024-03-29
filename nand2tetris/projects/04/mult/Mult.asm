// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
//
// This program only needs to handle arguments that satisfy
// R0 >= 0, R1 >= 0, and R0*R1 < 32768.

// Put your code here.

    // Idea for loop 0..RAM[0] summing up RAM[1]
    // first take the infer smaller value and init, n = R0 < R1 ? R0 : R1
    @R0
    D=M

    @R1
    D=D-M

    @OVER_R1
    D;JGE

    @OVER_R0
    D;JLT

    (OVER_R1)
        @R1
        D=M

        @n
        M=D

        @R0
        D=M

        @inc
        M=D

        @SUM
        0;JMP

    (OVER_R0)
        @R0
        D=M

        @n
        M=D

        @R1
        D=M

        @inc
        M=D


        @SUM
        0;JMP

        (SUM)
            @R2
            M=0

            (SUM_LOOP)
                @n
                D=M

                @END
                D;JLE

                @inc
                D=M

                @R2
                M=M+D

                @n
                M=M-1

                @SUM_LOOP
                0;JMP

    (END)
    @END
    0;JMP
