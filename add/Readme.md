# add 概要流程

## core/commands/add.go
fileAdder, err := coreunix.NewAdder(req.Context, n.Pinning, n.Blockstore, dserv)
fileAdder.AddFile(file)
fileAdder.Finalize()
fileAdder.PinRoot()


## core/coreunix/add.go 
func (adder *Adder) AddFile(file files.File) error
    |
    |
   \ /
    |
func (adder *Adder) addFile(file files.File) error
    adder.add(reader)

func (adder *Adder) add(reader io.Reader) (ipld.Node, error) 
    chnk, err := chunker.FromString(reader, adder.Chunker)
    balanced.Layout chunker.FromString(reader, adder.Chunker)
