package sicxelib

import (
    "fmt"
    "sort"
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
    Format [4]uint8
    Opcode uint8
}

type Optab map[string]OptabEntry

func GenerateOptab() *Optab {
    optab := make(Optab)

    optab["ADD"]    = OptabEntry{"ADD",     [4]uint8{0,0,1,1}, 0x18}
    optab["ADDF"]   = OptabEntry{"ADDF",    [4]uint8{0,0,1,1}, 0x58}
    optab["ADDR"]   = OptabEntry{"ADDR",    [4]uint8{0,1,0,0}, 0x90}
    optab["AND"]    = OptabEntry{"AND",     [4]uint8{0,0,1,1}, 0x40}
    optab["CLEAR"]  = OptabEntry{"CLEAR",   [4]uint8{0,1,0,0}, 0xB4}
    optab["COMP"]   = OptabEntry{"COMP",    [4]uint8{0,0,1,1}, 0x28}
    optab["COMPF"]  = OptabEntry{"COMPF",   [4]uint8{0,0,1,1}, 0x88}
    optab["COMPR"]  = OptabEntry{"COMPR",   [4]uint8{0,1,0,0}, 0xA0}
    optab["DIV"]    = OptabEntry{"DIV",     [4]uint8{0,0,1,1}, 0x24}
    optab["DIVF"]   = OptabEntry{"DIVF",    [4]uint8{0,0,1,1}, 0x64}
    optab["DIVR"]   = OptabEntry{"DIVR",    [4]uint8{0,1,0,0}, 0x9C}
    optab["FIX"]    = OptabEntry{"FIX",     [4]uint8{1,0,0,0}, 0xC4}
    optab["FLOAT"]  = OptabEntry{"FLOAT",   [4]uint8{1,0,0,0}, 0xC0}
    optab["HIO"]    = OptabEntry{"HIO",     [4]uint8{1,0,0,0}, 0xF4}
    optab["J"]      = OptabEntry{"J",       [4]uint8{0,0,1,1}, 0x3C}
    optab["JEQ"]    = OptabEntry{"JEQ",     [4]uint8{0,0,1,1}, 0x30}
    optab["JGT"]    = OptabEntry{"JGT",     [4]uint8{0,0,1,1}, 0x34}
    optab["JLT"]    = OptabEntry{"JLT",     [4]uint8{0,0,1,1}, 0x38}
    optab["JSUB"]   = OptabEntry{"JSUB",    [4]uint8{0,0,1,1}, 0x48}
    optab["LDA"]    = OptabEntry{"LDA",     [4]uint8{0,0,1,1}, 0x00}
    optab["LDB"]    = OptabEntry{"LDB",     [4]uint8{0,0,1,1}, 0x68}
    optab["LDCH"]   = OptabEntry{"LDCH",    [4]uint8{0,0,1,1}, 0x50}
    optab["LDF"]    = OptabEntry{"LDF",     [4]uint8{0,0,1,1}, 0x70}
    optab["LDL"]    = OptabEntry{"LDL",     [4]uint8{0,0,1,1}, 0x08}
    optab["LDS"]    = OptabEntry{"LDS",     [4]uint8{0,0,1,1}, 0x6C}
    optab["LDT"]    = OptabEntry{"LDT",     [4]uint8{0,0,1,1}, 0x74}
    optab["LDX"]    = OptabEntry{"LDX",     [4]uint8{0,0,1,1}, 0x04}
    optab["LPS"]    = OptabEntry{"LPS",     [4]uint8{0,0,1,1}, 0xD0}
    optab["MUL"]    = OptabEntry{"MUL",     [4]uint8{0,0,1,1}, 0x20}
    optab["MULF"]   = OptabEntry{"MULF",    [4]uint8{0,0,1,1}, 0x60}
    optab["MULR"]   = OptabEntry{"MULR",    [4]uint8{0,1,0,0}, 0x98}
    optab["NORM"]   = OptabEntry{"NORM",    [4]uint8{1,0,0,0}, 0xC8}
    optab["OR"]     = OptabEntry{"OR",      [4]uint8{0,0,1,1}, 0x44}
    optab["RD"]     = OptabEntry{"RD",      [4]uint8{0,0,1,1}, 0xD8}
    optab["RMO"]    = OptabEntry{"RMO",     [4]uint8{0,1,0,0}, 0xAC}
    optab["RSUB"]   = OptabEntry{"RSUB",    [4]uint8{0,0,1,1}, 0x4C}
    optab["SHIFTL"] = OptabEntry{"SHIFTL",  [4]uint8{0,1,0,0}, 0xA4}
    optab["SHITFR"] = OptabEntry{"SHIFTR",  [4]uint8{0,1,0,0}, 0xA8}
    optab["SIO"]    = OptabEntry{"SIO",     [4]uint8{1,0,0,0}, 0xF0}
    optab["SSK"]    = OptabEntry{"SSK",     [4]uint8{0,0,1,1}, 0xEC}
    optab["STA"]    = OptabEntry{"STA",     [4]uint8{0,0,1,1}, 0x0C}
    optab["STB"]    = OptabEntry{"STB",     [4]uint8{0,0,1,1}, 0x78}
    optab["STCH"]   = OptabEntry{"STCH",    [4]uint8{0,0,1,1}, 0x54}
    optab["STF"]    = OptabEntry{"STF",     [4]uint8{0,0,1,1}, 0x80}
    optab["STI"]    = OptabEntry{"STI",     [4]uint8{0,0,1,1}, 0xD4}
    optab["STL"]    = OptabEntry{"STL",     [4]uint8{0,0,1,1}, 0x14}
    optab["STS"]    = OptabEntry{"STS",     [4]uint8{0,0,1,1}, 0x7C}
    optab["STSW"]   = OptabEntry{"STSW",    [4]uint8{0,0,1,1}, 0xE8}
    optab["STI"]    = OptabEntry{"STI",     [4]uint8{0,0,1,1}, 0x84}
    optab["STX"]    = OptabEntry{"STX",     [4]uint8{0,0,1,1}, 0x10}
    optab["SUB"]    = OptabEntry{"SUB",     [4]uint8{0,0,1,1}, 0x1C}
    optab["SUBF"]   = OptabEntry{"SUBF",    [4]uint8{0,0,1,1}, 0x5C}
    optab["SUBR"]   = OptabEntry{"SUBR",    [4]uint8{0,1,0,0}, 0x94}
    optab["SVC"]    = OptabEntry{"SVC",     [4]uint8{0,1,0,0}, 0xB0}
    optab["TD"]     = OptabEntry{"TD",      [4]uint8{0,0,1,1}, 0xE0}
    optab["TIO"]    = OptabEntry{"TIO",     [4]uint8{1,0,0,0}, 0xF8}
    optab["TIX"]    = OptabEntry{"TIX",     [4]uint8{0,0,1,1}, 0x2C}
    optab["TIXR"]   = OptabEntry{"TIXR",    [4]uint8{0,1,0,0}, 0xB8}
    optab["WD"]     = OptabEntry{"WD",      [4]uint8{0,0,1,1}, 0xDC}

    return &optab
}

func (optab Optab) LookUp(symbol string) (OptabEntry, bool) {
    v, ok := optab[symbol]
    return v, ok
}

func (optab Optab) PrintOptab() {
    mnemonics := make([]string, 0, len(optab))
    for k := range optab {
        mnemonics = append(mnemonics, k)
    }
    sort.Strings(mnemonics)

    fmt.Println("Mnemonic\tFormat\t\tOpcode")
    fmt.Println("---------------------------------------")
    for _, mne := range mnemonics {
        value := optab[mne]
        fmt.Printf("%v\t\t", value.Mnemonic)
        for i, v := range value.Format {
            if v == 1 {
                fmt.Printf("%d ", i+1)
            }
        }
        fmt.Printf("\t\t%.2X\n", value.Opcode)
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
