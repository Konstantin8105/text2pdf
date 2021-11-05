//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// appname - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:41
//
//This ugly, sparsely-commented program is the source for text2pdf version
//1.1.  It should be ANSI-conforming and compile on most platforms.  You'll
//need to change LF_EXTRA to 1 for machines which write 2 characters for \n.
//These include PCs, of course.
//
//You may distribute the source or compiled versions free of charge.  You may
//not alter the source in any way other than those mentioned above without
//the permission of the author, Phil Smith <phil@bagobytes.co.uk>.
//
//Please send any comments to the author.
//
//Copyright (c) Phil Smith, 1996
//
//REVISION HISTORY
//
//Version 1.1
//11 Oct 96 Added handling of form-feed characters, removed need for tmp file,
//          put reference to resources in each page (avoid bug in Acrobat),
//	  changed date format to PDF-1.1 standard.
//12 Jun 96 Added check to avoid blank last page
//12 Jun 96 Added LINE_END def to get round platform-specific \r, \n etc.
//18 Mar 96 Added ISOLatin1Encoding option
//
// change to 1 for PCs (where \n => <CR><LF>)
var appname = ("text2pdf v1.1")

// progname - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:42
var progname = ("text2pdf")

// infile - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:44
// var infile *noarch.File

// pageNo - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:45
var pageNo int

// pageObs - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:46
var pageObs []int = make([]int, 500)

// curObj - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:47
// object number being or last written
var curObj int = 5

// locations - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:48
var locations []int = make([]int, 100000)

// font - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:50
var font = "/Arial"//"/Courier" //  make([]byte, 256)

// defaultFont - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:51
// var defaultFont = ("Courier")

// ISOEnc - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:52
var ISOEnc int

// doFFs - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:53
var doFFs int = 1

// tab - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:54
var tab int = 8

// pointSize - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:55
var pointSize int = 10

// vertSpace - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:56
var vertSpace int = 12

// lines - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:57
var lines int

// cols - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:58
// max chars per output line
var cols int = 80

// columns - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:59
// number of columns
var columns int = 1

// pageHeight - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:62
// Default paper is Letter size, as in distiller
var pageHeight int = 792

// pageWidth - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:63
var pageWidth int = 612

// buf - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:65
var buf []uint8 = make([]uint8, 1024)

// fpos - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:66
var fpos int

// prn - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:68
func prn(format string, a ...interface{}) {
	n, err := fmt.Printf(format, a...)
	if err != nil {
		panic(err)
	}
	fpos += n

	// Everything written to the PDF file goes through this function.
	// This means we can keep track of the file position without using
	// ftell on a real (tmp) file.  However, PCs write out 2 characters
	// for \n, so we need this ugly loop to keep fpos correct
	// fpos += len(str)
	// 	for str[0] != 0 {
	// 		if int(str[0]) == int('\n') {
	// 			fpos += 0
	// 		}
	// 		prn(">", str)
	// 		// noarch.Putchar(int((func()  {
	// 		// 	defer func() {
	// 		// 		str = str[0+1:]
	// 		// 	}()
	// 		// 	return str
	// 		// }())[0]))
	// 	}
}

// WriteHeader - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:82
func WriteHeader(title string) {

	// 	var ltime []noarch.Tm
	// 	var clock noarch.TimeT
	// 	var datestring  = make(, 30)
	// 	noarch.Time((*[1000000]noarch.TimeT)(unsafe.Pointer(&clock))[:])
	// 	ltime = noarch.LocalTime((*[1000000]noarch.TimeT)(unsafe.Pointer(&clock))[:])
	// 	strftime(datestring, 30, ("D:%Y%m%d%H%M%S"), ltime)

	prn("%%PDF-1.1\n")
	locations[1] = fpos
	prn("1 0 obj\n")
	prn("<<\n")

	// 	noarch.Sprintf((*[1000000]byte)(unsafe.Pointer(uintptr(int64(uintptr(unsafe.Pointer(&buf[0]))) / int64(1))))[:], ("/CreationDate (%s)\n"), datestring)
	prn("/CreationDate (D:20211104163019)\n")

	prn("/Producer (%s (\\251 Phil Smith, 1996))\n", appname)
	if title != "" {
		prn("/Title (%s)\n", title)
	}
	prn(">>\n")
	prn("endobj\n")
	locations[2] = fpos
	prn("2 0 obj\n")
	prn("<<\n")
	prn("/Type /Catalog\n")
	prn("/Pages 3 0 R\n")
	prn(">>\n")
	prn("endobj\n")
	locations[4] = fpos
	prn("4 0 obj\n")
	prn("<<\n")
	prn("/Type /Font\n")
	prn("/Subtype /Type1\n")
	prn("/Name /F1\n")
	prn("/BaseFont %s\n", font)
	if ISOEnc != 0 {
		prn("/Encoding <<\n")
		prn("/Differences [ 0 /.notdef /.notdef /.notdef /.notdef\n")
		prn("/.notdef /.notdef /.notdef /.notdef /.notdef /.notdef\n")
		prn("/.notdef /.notdef /.notdef /.notdef /.notdef /.notdef\n")
		prn("/.notdef /.notdef /.notdef /.notdef /.notdef /.notdef\n")
		prn("/.notdef /.notdef /.notdef /.notdef /.notdef /.notdef\n")
		prn("/.notdef /.notdef /.notdef /.notdef /space /exclam\n")
		prn("/quotedbl /numbersign /dollar /percent /ampersand\n")
		prn("/quoteright /parenleft /parenright /asterisk /plus /comma\n")
		prn("/hyphen /period /slash /zero /one /two /three /four /five\n")
		prn("/six /seven /eight /nine /colon /semicolon /less /equal\n")
		prn("/greater /question /at /A /B /C /D /E /F /G /H /I /J /K /L\n")
		prn("/M /N /O /P /Q /R /S /T /U /V /W /X /Y /Z /bracketleft\n")
		prn("/backslash /bracketright /asciicircum /underscore\n")
		prn("/quoteleft /a /b /c /d /e /f /g /h /i /j /k /l /m /n /o /p\n")
		prn("/q /r /s /t /u /v /w /x /y /z /braceleft /bar /braceright\n")
		prn("/asciitilde /.notdef /.notdef /.notdef /.notdef /.notdef\n")
		prn("/.notdef /.notdef /.notdef /.notdef /.notdef /.notdef\n")
		prn("/.notdef /.notdef /.notdef /.notdef /.notdef /.notdef\n")
		prn("/dotlessi /grave /acute /circumflex /tilde /macron /breve\n")
		prn("/dotaccent /dieresis /.notdef /ring /cedilla /.notdef\n")
		prn("/hungarumlaut /ogonek /caron /space /exclamdown /cent\n")
		prn("/sterling /currency /yen /brokenbar /section /dieresis\n")
		prn("/copyright /ordfeminine /guillemotleft /logicalnot /hyphen\n")
		prn("/registered /macron /degree /plusminus /twosuperior\n")
		prn("/threesuperior /acute /mu /paragraph /periodcentered\n")
		prn("/cedilla /onesuperior /ordmasculine /guillemotright\n")
		prn("/onequarter /onehalf /threequarters /questiondown /Agrave\n")
		prn("/Aacute /Acircumflex /Atilde /Adieresis /Aring /AE\n")
		prn("/Ccedilla /Egrave /Eacute /Ecircumflex /Edieresis /Igrave\n")
		prn("/Iacute /Icircumflex /Idieresis /Eth /Ntilde /Ograve\n")
		prn("/Oacute /Ocircumflex /Otilde /Odieresis /multiply /Oslash\n")
		prn("/Ugrave /Uacute /Ucircumflex /Udieresis /Yacute /Thorn\n")
		prn("/germandbls /agrave /aacute /acircumflex /atilde /adieresis\n")
		prn("/aring /ae /ccedilla /egrave /eacute /ecircumflex\n")
		prn("/edieresis /igrave /iacute /icircumflex /idieresis /eth\n")
		prn("/ntilde /ograve /oacute /ocircumflex /otilde /odieresis\n")
		prn("/divide /oslash /ugrave /uacute /ucircumflex /udieresis\n")
		prn("/yacute /thorn /ydieresis ]\n")
		prn(">>\n")
	}
	prn(">>\n")
	prn("endobj\n")
	locations[5] = fpos
	prn("5 0 obj\n")
	prn("<<\n")
	prn("  /Font << /F1 4 0 R >>\n")
	prn("  /ProcSet [ /PDF /Text ]\n")
	prn(">>\n")
	prn("endobj\n")
}

func StartPage() (strmPos int) {
	curObj++
	locations[curObj] = fpos

	pageNo++
	pageObs[pageNo] = curObj

	prn("%d 0 obj\n", curObj)
	prn("<<\n")
	prn("/Type /Page\n")
	prn("/Parent 3 0 R\n")
	prn("/Resources 5 0 R\n")
	curObj++
	prn("/Contents %d 0 R\n", curObj)
	prn(">>\n")
	prn("endobj\n")

	locations[curObj] = fpos
	prn("%d 0 obj\n", curObj)
	prn("<<\n")
	prn("/Length %d 0 R\n", curObj+1)
	prn(">>\n")
	prn("stream\n")

	strmPos = fpos
	prn("BT\n")
	prn("/F1 %d Tf\n", pointSize)
	prn("1 0 0 1 50 %d Tm\n", pageHeight-40)
	prn("%d TL\n", vertSpace)
	return strmPos
}

func EndPage(streamStart int) {
	var streamEnd int
	prn("ET\n")
	streamEnd = fpos
	prn("endstream\n")
	prn("endobj\n")
	curObj++
	locations[curObj] = fpos
	prn("%d 0 obj\n", curObj)
	prn("%d\n", streamEnd-streamStart)
	prn("endobj\n")
}

func WritePages(lines [][]rune) {
	// 	var padding int
	// 	var beginstream int

	// 	content, err := ioutil.ReadFile("test.txt")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	aline := 0
	// 	cols := 80

	beginstream := StartPage()
	for l := range lines {
		line := lines[l]
		prn("(")
		for p := range line {
			if line[p] == '(' || line[p] == ')' || line[p] == '\\' {
				prn("\\")
			}
			// prn("\\u%05d ", line[p])
			if line[p] < 127 {
				prn("%s", string(line[p]))
			} else {
				// prn("\\%.3o", line[p])
				bs := []byte(string(line[p]))
				for _, b := range bs {
					prn("\\%o", b)
				}
			}
		}
		prn(")'\n")
	}
	EndPage(beginstream)

	// 	for ic := 0; ic < len(content); ic++ {
	//
	// 		prn("(")
	// 		charNo := 0
	// 		for charNo = 0; ic+charNo < len(content) && charNo < cols && content[ic+charNo] != '\n'; charNo++ {
	// 			ch := content[ic+charNo]
	// 			switch {
	// 			case ch >= 32 && ch <= 127:
	// 				if ch == '(' || ch == ')' || ch == '\\' {
	// 					prn("\\")
	// 				}
	// 				prn("%c", ch)
	// 			case ch == 9:
	// 				padding = tab - (int(ch)-1)%tab
	// 				for i := 1; i <= padding; i++ {
	// 					prn(" ")
	// 				}
	// 				ch += byte(padding - 1)
	// 			default:
	// 				if ch != 12 {
	// 					// write \xxx form for dodgy character
	// 					prn("\\%.3o", ch)
	// 					// 				} else {
	// 					// 					// don't print anything for a FF
	// 					// 					ch--
	// 				}
	// 			}
	// 		}
	// 		ic += charNo
	// 		prn(")'\n")
	//
	// 		aline++
	// 		if aline == 40 {
	// 			EndPage(beginstream)
	// 			beginstream = StartPage()
	// 			aline = 0
	// 		}
	// 	}
	// 	EndPage(beginstream)
}

// WriteRest - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:288
func WriteRest() {
	var xref int
	var i int
	locations[3] = fpos
	prn("3 0 obj\n")
	prn("<<\n")
	prn("/Type /Pages\n")
	prn("/Count %d\n", pageNo)
	prn("/MediaBox [ 0 0 %d %d ]\n", pageWidth, pageHeight)
	prn("/Kids [ ")
	for i = 1; i <= pageNo; i++ {
		prn("%d 0 R ", pageObs[i])
	}
	prn("]\n")
	prn(">>\n")
	prn("endobj\n")
	xref = fpos
	prn("xref\n")
	prn("0 %d\n", curObj+1)
	// note that \n is translated by prn
	prn("0000000000 65535 f %c", '\r')
	for i = 1; i <= curObj; i++ {
		prn("%010d 00000 n %c", locations[i], '\r')
	}
	prn("trailer\n")
	prn("<<\n")
	prn("/Size %d\n", curObj+1)
	prn("/Root 2 0 R\n")
	prn("/Info 1 0 R\n")
	prn(">>\n")
	prn("startxref\n")
	prn("%d\n", xref)
	prn("%%EOF\n")
}

// ShowHelp - transpiled function from  GOPATH/src/github.com/Konstantin8105/text2pdf/text2pdf.c:327
// func ShowHelp() {
// 	noarch.Printf("\n%s [options] [filename]\n\n"), progname)
// 	noarch.Printf("  %s makes a 7-bit clean PDF file (version 1.1) from any input file.\n"), progname)
// 	prn("  It reads from standard input or a named file, and writes the PDF file\n")
// 	prn("  to standard output.\n")
// 	prn("\n  There are various options as follows:\n\n")
// 	prn("  -h\t\tshow this message\n")
// 	prn("  -f<font>\tuse PostScript <font> (must be in standard 14, default: Courier)\n")
// 	prn("  -I\t\tuse ISOLatin1Encoding\n")
// 	noarch.Printf("  -s<size>\tuse font at given pointsize (default %d)\n"), pointSize)
// 	noarch.Printf("  -v<dist>\tuse given line spacing (default %d points)\n"), vertSpace)
// 	prn("  -l<lines>\tlines per page (default 60, determined automatically\n\t\tif unspecified)\n")
// 	prn("  -c<chars>\tmaximum characters per line (default 80)\n")
// 	prn("  -t<spaces>\tspaces per tab character (default 8)\n")
// 	prn("  -F\t\tignore formfeed characters (^L)\n")
// 	prn("  -A4\t\tuse A4 paper (default Letter)\n")
// 	prn("  -A3\t\tuse A3 paper (default Letter)\n")
// 	prn("  -x<width>\tindependent paper width in points\n")
// 	prn("  -y<height>\tindependent paper height in points\n")
// 	prn("  -2\t\tformat in 2 columns\n")
// 	prn("  -L\t\tlandscape mode\n")
// 	prn("\n  Note that where one variable is implied by two options, the second option\n  takes precedence for that variable. (e.g. -A4 -y500)\n")
// 	prn("  In landscape mode, page width and height are simply swapped over before\n  formatting, no matter how or when they were defined.\n")
// 	noarch.Printf("\n%s (c) Phil Smith, 1996\n"), appname)
// }

func main() {
	// 	var (
	// 		helpFlag         = flag.Bool("h", false, "show help")
	// 		fontFlag         = flag.String("f", "Courier", "PostScript font")
	// 		pointsizeFlag    = flag.Int("s", 10, "font pointsize")
	// 		linespaceFlag    = flag.Int("v", 12, "line spacing of point")
	// 		linesPerPageFlag = flag.Int("l", 60, "lines per page")
	// 		charPerLineFlag  = flag.Int("c", 80, "maximum characters per line")
	// 		tabFlag          = flag.Int("t", 8, "spaces per tab character")
	// 		formfeedFlag     = flag.Bool("F", false, "ignore formfeed characters (^L)")
	// 		columnsFlag      = flag.Bool("2", false, "format in 2 columns")
	// 		pageWidthFlag    = flag.Int("x", 595, "independent paper width in points")
	// 		pageHeightFlag   = flag.Int("y", 842, "independent paper height in points")
	// 		landscapeFlag    = flag.Bool("L", false, "landscape mode")
	// 	)
	// 	case '3':
	pageWidth = 842
	pageHeight = 1190
	// case '4':
	// 	pageWidth = 595
	// 	pageHeight = 842

	//   text2pdf makes a 7-bit clean PDF file (version 1.1) from any input file.
	//   It reads from standard input or a named file, and writes the PDF file
	//   to standard output.
	//
	//   There are various options as follows:
	//
	//   -I		use ISOLatin1Encoding
	//   -A4		use A4 paper (default Letter)
	//   -A3		use A3 paper (default Letter)
	//   -x<width>	independent paper width in points
	//   -y<height>	independent paper height in points

	// 	var tmp int
	// 	var landscape int
	// 	var ifilename
	// 	noarch.Strcpy(font, ("/")
	// 	noarch.Strcat(font, defaultFont)
	// default
	// 	infile = noarch.Stdin
	// case 'I':
	// 	ISOEnc = 1
	// case 'F':
	// 	doFFs = 0

	// case '3':
	// 	pageWidth = 842
	// 	pageHeight = 1190
	// case '4':
	// 	pageWidth = 595
	// 	pageHeight = 842

	// 	if *landscapeFlag {
	// 		pageHeightFlag, pageWidthFlag = pageWidthFlag, pageHeightFlag
	// 	}
	// 	if lines == 0 {
	// 		lines = (pageHeight - 72) / vertSpace
	// 	}
	// 	if lines < 1 {
	// 		lines = 1
	// 	}

	// happens to give 60 as default

	// 	pageHeight = 752 + 40
	//

	filename := "test.txt"

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	content = bytes.ReplaceAll(content, []byte{'\r'}, []byte{})
	columns := 80
	ls := split(content, columns)
	// 	for j := range ls {
	// 		if j != 0 {
	// 			fmt.Printf("\n")
	// 		}
	// 		for k := range ls[j] {
	// 			fmt.Printf("%s", string(ls[j][k]))
	// 		}
	// 	}

	WriteHeader(filename)
	WritePages(ls)
	WriteRest()

	return
}

// split line string without newline per lines splitted by columns runes
func split(line []byte, columns int) (ls [][]rune) {
	if index := bytes.Index(line, []byte{'\n'}); 0 <= index {
		ls = append(ls, split(line[:index], columns)...)
		ls = append(ls, split(line[index+1:], columns)...)
		return
	}
	// split into runes as slice of runes
	rs := []rune(string(line))
	// split by lines
	for {
		if len(rs) <= columns {
			ls = append(ls, rs)
			rs = nil
			break
		}
		space := -1
		for i := columns - 1; 1 <= i; i-- {
			if rs[i] == ' ' {
				space = i
			}
		}
		if space < columns/2 {
			ls = append(ls, rs[:columns])
			rs = rs[columns:]
			continue
		}
		ls = append(ls, rs[:space+1]) // space is not include
		rs = rs[space:]
	}
	return
}
