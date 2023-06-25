package markdown

import (
	"errors"
	"fmt"
	"io"
)

var (
	h1Prefix              = []byte{'#', ' '}
	h2Prefix              = []byte{'#', '#', ' '}
	h3Prefix              = []byte{'#', '#', '#', ' '}
	h4Prefix              = []byte{'#', '#', '#', '#', ' '}
	h5Prefix              = []byte{'#', '#', '#', '#', '#', ' '}
	italicsSentinel       = []byte{'_', '_'}
	boldSentinel          = []byte{'*', '*'}
	strikethroughSentinel = []byte{'~', '~'}
	codeBlockSentinel     = []byte{'`', '`', '`'}
	codeInlineSentinel    = []byte{'`'}
	tableSep              = []byte{' ', '|', ' '}
	tableHeaderSepPart    = []byte{'-', '-', '-'}
	blockQuotePrefix      = []byte{'>', ' '}
	hr                    = []byte{'*', '*', '*', '\n'}
	br                    = []byte{'\n', '\n'}
	newLine               = []byte{'\n'}
)

type Generator struct {
	writer io.Writer // TODO: use StringWriter interface instead
}

func NewGenerator(writer io.Writer) *Generator {
	return &Generator{
		writer: writer,
	}
}

func (g *Generator) H1(s string) error {
	if err := g.Write("<h1>"); err != nil {
		return err
	}
	if err := g.Write(s); err != nil {
		return err
	}
	if err := g.Write("</h1>"); err != nil {
		return err
	}
	return nil
}

func (g *Generator) H2(s string) error {
	if err := g.Write("<h2>"); err != nil {
		return err
	}
	if err := g.Write(s); err != nil {
		return err
	}
	if err := g.Write("</h2>"); err != nil {
		return err
	}
	return nil
}

func (g *Generator) H3(s string) error {
	if err := g.Write("<h3>"); err != nil {
		return err
	}
	if err := g.Write(s); err != nil {
		return err
	}
	if err := g.Write("</h3>"); err != nil {
		return err
	}
	return nil
}

func (g *Generator) H4(s string) error {
	if err := g.Write("<h4>"); err != nil {
		return err
	}
	if err := g.Write(s); err != nil {
		return err
	}
	if err := g.Write("</h4>"); err != nil {
		return err
	}
	return nil
}

func (g *Generator) H5(s string) error {
	if err := g.Write("<h5>"); err != nil {
		return err
	}
	if err := g.Write(s); err != nil {
		return err
	}
	if err := g.Write("</h5>"); err != nil {
		return err
	}
	return nil
}

func (g *Generator) Italics(s string) error {
	if err := g.Write("<i>"); err != nil {
		return err
	}
	if err := g.Write(s); err != nil {
		return err
	}
	if err := g.Write("</i>"); err != nil {
		return err
	}
	return nil
}

func (g *Generator) Bold(s string) error {
	if err := g.Write("<b>"); err != nil {
		return err
	}
	if err := g.Write(s); err != nil {
		return err
	}
	if err := g.Write("</b>"); err != nil {
		return err
	}
	return nil

	//if err := g.WriteBytes(boldSentinel); err != nil {
	//	return err
	//}
	//if err := g.Write(s); err != nil {
	//	return err
	//}
	//if err := g.WriteBytes(boldSentinel); err != nil {
	//	return err
	//}
	//return nil
}

func (g *Generator) Strikethrough(s string) error {
	if err := g.WriteBytes(strikethroughSentinel); err != nil {
		return err
	}
	if err := g.Write(s); err != nil {
		return err
	}
	if err := g.WriteBytes(strikethroughSentinel); err != nil {
		return err
	}
	return nil
}

func (g *Generator) OrderedList(ss []string) error {
	for i, s := range ss {
		if err := g.Write(fmt.Sprintf("%d. %s\n", i, s)); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) UnorderedList(ss []string) error {
	for _, s := range ss {
		if err := g.WriteBytes([]byte{'*', ' '}); err != nil {
			return err
		}
		if err := g.Write(s); err != nil {
			return err
		}
		if err := g.WriteBytes(newLine); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) Link(s string, url string) error {
	return g.Write(fmt.Sprintf("<a href=\"%s\">%s</a>", url, s))
	//return g.Write(fmt.Sprintf("[%s](%s)", s, url))
}

func (g *Generator) Image(altText string, url string) error {
	return g.Write(fmt.Sprintf("![%s](%s)", altText, url))
}

func (g *Generator) InlineCode(s string) error {
	if err := g.Write("<code>"); err != nil {
		return err
	}
	if err := g.Write(s); err != nil {
		return err
	}
	if err := g.Write("</code>"); err != nil {
		return err
	}
	return nil
}

func (g *Generator) CodeBlock(s string, lang ...string) error {
	// NOTE: don't think we can do syntax highlighting with pre
	if err := g.Write("<pre>"); err != nil {
		return err
	}
	if err := g.Write(s); err != nil {
		return err
	}
	if err := g.Write("</pre>"); err != nil {
		return err
	}
	return nil
	//if err := g.WriteBytes(codeBlockSentinel); err != nil {
	//	return err
	//}
	//if len(lang) > 0 {
	//	if err := g.Write(lang[0]); err != nil {
	//		return err
	//	}
	//}
	//if err := g.WriteBytes(newLine); err != nil {
	//	return err
	//}
	//
	//if err := g.Write(s); err != nil {
	//	return err
	//}
	//
	//if err := g.WriteBytes(codeBlockSentinel); err != nil {
	//	return err
	//}
	//if err := g.WriteBytes(newLine); err != nil {
	//	return err
	//}
	//return nil
}

func (g *Generator) TableWideHeader(s string) error {
	if err := g.Write(fmt.Sprintf("<table><tr><th colspan=\"%s\">%s</th></tr>", "100%", s)); err != nil {
		return err
	}
	return nil
}

func (g *Generator) TableHeader(header ...string) error {
	if err := g.Write("<table>"); err != nil {
		return err
	}

	if len(header) > 0 {
		if err := g.Write("<tr>"); err != nil {
			return err
		}

		for _, colHeader := range header {
			if err := g.Write("<th>"); err != nil {
				return err
			}
			if err := g.Write(colHeader); err != nil {
				return err
			}
			if err := g.Write("</th>"); err != nil {
				return err
			}
		}

		if err := g.Write("</tr>"); err != nil {
			return err
		}
	}

	return nil

	//// Heading row
	//l := len(header)
	//
	//for i, colHeading := range header {
	//	if err := g.Write(colHeading); err != nil {
	//		return err
	//	}
	//	if i != l-1 {
	//		if err := g.WriteBytes(tableSep); err != nil {
	//			return err
	//		}
	//	}
	//}
	//if err := g.WriteBytes(newLine); err != nil {
	//	return err
	//}
	//
	//// Heading separator row
	//for i := 0; i < l; i++ {
	//	if err := g.WriteBytes(tableHeaderSepPart); err != nil {
	//		return err
	//	}
	//	if i != l-1 {
	//		if err := g.WriteBytes(tableSep); err != nil {
	//			return err
	//		}
	//	}
	//}
	//if err := g.WriteBytes(newLine); err != nil {
	//	return err
	//}
	//
	//return nil
}

func (g *Generator) TableRow(fs ...func() error) error {
	if err := g.Write("<tr>"); err != nil {
		return err
	}

	for _, f := range fs {
		if err := g.Write("<td>"); err != nil {
			return err
		}
		if err := f(); err != nil {
			return err
		}
		if err := g.Write("</td>"); err != nil {
			return err
		}
	}

	if err := g.Write("</tr>"); err != nil {
		return err
	}

	return nil

	//// Data rows
	//l := len(fs)
	//
	//for i, f := range fs {
	//	if err := f(); err != nil {
	//		return err
	//	}
	//	if i != l-1 {
	//		if err := g.WriteBytes(tableSep); err != nil {
	//			return err
	//		}
	//	}
	//}
	//if err := g.WriteBytes(newLine); err != nil {
	//	return err
	//}
	//
	//return nil
}

func (g *Generator) TableFooter() error {
	return g.Write("</table>\n\n")
}

func (g *Generator) BlockQuote(s string) error {
	return errors.Join(
		g.WriteBytes(blockQuotePrefix),
		g.Write(s),
	)
}

func (g *Generator) HR() error {
	return g.WriteBytes(hr)
}

func (g *Generator) BR() error {
	return g.WriteBytes(br)
}

func (g *Generator) WriteBytes(bs []byte) error {
	_, err := g.writer.Write(bs)
	return err
}

func (g *Generator) Write(s string) error {
	return g.WriteBytes([]byte(s))
}
