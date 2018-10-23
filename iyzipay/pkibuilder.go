package iyzipay

type PkiBuilder struct {
    pkiString string
}

func (p *PkiBuilder) append(key string, value string) {
    p.pkiString = p.pkiString + key + "=" + value + ","
}

func (p *PkiBuilder) removeTrailingComma() {
    p.pkiString = p.pkiString[0:len(p.pkiString)-1]
}

func (p *PkiBuilder) getPkiString() string {
    p.removeTrailingComma()
    return "[" + p.pkiString + "]"
}
