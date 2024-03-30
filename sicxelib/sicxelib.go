package sicxelib

import (
    "fmt"
)

type DataType int
    
const (
    Integer DataType = iota
    FloatingPoint
    Characters
    Hexadecimal
)

type SymtabEntry struct {
    LabelName string
    Address uint32
    Dtype DataType
}

type Symtab map[string]SymtabEntry

type OptabEntry struct {
    Mnemonic string
    Opcode uint8
}

type Optab map[string]OptabEntry

func GenerateOptab() *Optab {
    optab := make(Optab)

    optab["ADD"] = OptabEntry{"ADD", 0x18}
    optab["LDA"] = OptabEntry{"LDA", 0x00}
    optab["STA"] = OptabEntry{"STA", 0x0C}

    return &optab
}

func (optab Optab) LookUp(symbol string) (OptabEntry, bool) {
    v, ok := optab[symbol]
    return v, ok
}

func (optab Optab) PrintOptab() {
    fmt.Println("Mnemonic\tOpcode")
    fmt.Println("-----------------------")
    for _, value := range optab {
        fmt.Printf("%v\t\t%0.2X\n", value.Mnemonic, value.Opcode)  
    }
}

func NewSymtab() *Symtab {
    return &Symtab{}
}

func (symtab Symtab) PrintSymtab() {
    fmt.Println("LabelName\tAddress\tDataType")
    fmt.Println("----------------------------")
    for _, value := range symtab {
        fmt.Printf("%v\t\t%0.5X\t", value.LabelName, value.Address)  
        switch value.Dtype {
        case Integer:
            fmt.Println("INTEGER")
        default:
            fmt.Println("Unknown")
        }
    }
}
