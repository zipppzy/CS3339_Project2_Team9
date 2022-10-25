package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ReadBinary reads text file and makes Instructions and adds them to the InstructionList
func ReadBinary(filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pc uint64
	pc = 96
	for scanner.Scan() {
		InstructionList = append(InstructionList, Instruction{rawInstruction: scanner.Text(), memLoc: pc})
		pc += 4
	}
}

func WriteInstructions(filePath string, list []Instruction) {
	f, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for i := 0; i < len(list); i++ {
		switch list[i].instructionType {
		case "B":
			//write binary with spaces
			_, err := fmt.Fprintf(f, "%s %s\t", list[i].rawInstruction[0:6], list[i].rawInstruction[6:32])
			//write memLoc and opcode
			_, err = fmt.Fprintf(f, "%d\t%s\t", list[i].memLoc, list[i].op)
			//write operands
			_, err = fmt.Fprintf(f, "#%d\n", list[i].offset)
			if err != nil {
				log.Fatal(err)
			}
		case "I":
			//write binary with spaces
			_, err := fmt.Fprintf(f, "%s %s %s %s\t", list[i].rawInstruction[0:10], list[i].rawInstruction[10:22], list[i].rawInstruction[22:27], list[i].rawInstruction[27:32])
			//write memLoc and opcode
			_, err = fmt.Fprintf(f, "%d\t%s\t", list[i].memLoc, list[i].op)
			//write operands
			_, err = fmt.Fprintf(f, "R%d, R%d, #%d\n", list[i].rd, list[i].rn, list[i].immediate)
			if err != nil {
				log.Fatal(err)
			}

		case "CB":
			//write binary with spaces
			_, err := fmt.Fprintf(f, "%s %s %s\t", list[i].rawInstruction[0:8], list[i].rawInstruction[8:27], list[i].rawInstruction[27:32])
			//write memLoc and opcode
			_, err = fmt.Fprintf(f, "%d\t%s\t", list[i].memLoc, list[i].op)
			//write operands
			_, err = fmt.Fprintf(f, "R%d, #%d\n", list[i].conditional, list[i].offset)
			if err != nil {
				log.Fatal(err)
			}
		case "IM":
			//write binary with spaces
			_, err := fmt.Fprintf(f, "%s %s %s %s\t", list[i].rawInstruction[0:9], list[i].rawInstruction[9:12], list[i].rawInstruction[12:27], list[i].rawInstruction[27:32])
			//write memLoc and opcode
			_, err = fmt.Fprintf(f, "%d\t%s\t", list[i].memLoc, list[i].op)
			//write operands
			_, err = fmt.Fprintf(f, "R%d, %d, LSL %d\n", list[i].rd, list[i].field, list[i].shiftCode)
			if err != nil {
				log.Fatal(err)
			}
			// I am not sure about D too
		case "D":
			//write binary with spaces
			_, err := fmt.Fprintf(f, "%s %s %s %s %s\t", list[i].rawInstruction[0:11], list[i].rawInstruction[11:20], list[i].rawInstruction[20:22], list[i].rawInstruction[22:27], list[i].rawInstruction[27:32])
			//write memLoc and opcode
			_, err = fmt.Fprintf(f, "%d\t%s\t", list[i].memLoc, list[i].op)
			//write operands
			_, err = fmt.Fprintf(f, "R%d, [R%d, #%d]\n", list[i].rt, list[i].rn, list[i].address)
			if err != nil {
				log.Fatal(err)
			}
		case "R":
			//write binary with spaces
			_, err := fmt.Fprintf(f, "%s %s %s %s %s\t", list[i].rawInstruction[0:11], list[i].rawInstruction[11:16], list[i].rawInstruction[16:22], list[i].rawInstruction[22:27], list[i].rawInstruction[27:32])
			//write memLoc and opcode
			_, err = fmt.Fprintf(f, "%d\t%s\t", list[i].memLoc, list[i].op)
			//write operands
			_, err = fmt.Fprintf(f, "R%d, R%d, ", list[i].rd, list[i].rn)
			if list[i].op == "LSL" || list[i].op == "ASR" || list[i].op == "LSR" {
				_, err = fmt.Fprintf(f, "#%d\n", list[i].shamt)
			} else {
				_, err = fmt.Fprintf(f, "R%d\n", list[i].rm)
			}
			if err != nil {
				log.Fatal(err)
			}
		case "BREAK":
			_, err := fmt.Fprintf(f, "%s\t%d\tBREAK\n", list[i].rawInstruction, list[i].memLoc)
			if err != nil {
				log.Fatal(err)
			}
		case "MEM":
			_, err := fmt.Fprintf(f, "%s\t%d\t%d\n", list[i].rawInstruction, list[i].memLoc, list[i].memValue)
			if err != nil {
				log.Fatal(err)
			}

		case "NOP":
			_, err := fmt.Fprintf(f, "%s\t%d\t%s\n", list[i].rawInstruction, list[i].memLoc, list[i].op)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
