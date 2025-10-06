package replacer

import (
	"strings"
	"wgetNew/external/fs"
)

// wrap for default strings.Replacer
type Replacer struct {
	rpl *strings.Replacer
}

func NewReplacer(data []string) *Replacer {
	return &Replacer{
		rpl: strings.NewReplacer(data...),
	}
}

func (r *Replacer) Do(path, name string) error {
	data, err := fs.ReadString(path + "/" + name)
	if err != nil {
		return err
	}

	f, err := fs.Descriptor(path, name)
	if err != nil {
		return err
	}

	written, err := r.rpl.WriteString(f, data)
	if err != nil {
		return err
	}
	if written == 0 {
		return err
	}

	return nil
}
