package sqlbuilder

const (
	alterTableMarkerInit injectionMarker = iota
	alterTableMarkerAfterCreate
	alterTableMarkerAfterDefine
	alterTableMarkerAfterOption
)

func NewAlterTableBuilder() *CreateTableBuilder {
	return DefaultFlavor.NewCreateTableBuilder()
}

func newAlterTableBuilder() *AlterTableBuilder {
	args := &Args{}
	return &AlterTableBuilder{
		verb:      "ALTER TABLE",
		args:      args,
		injection: newInjection(),
		marker:    alterTableMarkerInit,
	}
}

type AlterTableBuilder struct {
	verb        string
	ifNotExists bool
	table       string
	defs        [][]string
	options     [][]string

	args *Args

	injection *injection
	marker    injectionMarker
}

func (ctb *AlterTableBuilder) Define(def ...string) *AlterTableBuilder {
	ctb.defs = append(ctb.defs, def)
	ctb.marker = alterTableMarkerAfterDefine
	return ctb
}

func (ctb *AlterTableBuilder) FieldComment(comment string) *AlterTableBuilder {
	str := "comment '" + comment + "'"
	ctb.defs[len(ctb.defs)-1] = append(ctb.defs[len(ctb.defs)-1], str)
	ctb.marker = alterTableMarkerAfterDefine
	return ctb
}

func (ctb *AlterTableBuilder) Option(opt ...string) *AlterTableBuilder {
	ctb.options = append(ctb.options, opt)
	ctb.marker = alterTableMarkerAfterOption
	return ctb
}

func (ctb *AlterTableBuilder) TableComment(comment string) *AlterTableBuilder {
	str := "comment='" + comment + "'"
	ctb.options[len(ctb.options)-1] = append(ctb.options[len(ctb.options)-1], str)
	ctb.marker = alterTableMarkerAfterOption
	return ctb
}

func (ctb *AlterTableBuilder) SetFlavor(flavor Flavor) (old Flavor) {
	old = ctb.args.Flavor
	ctb.args.Flavor = flavor
	return
}
