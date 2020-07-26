// 《The Go Programming Language》: http://gopl.io/
// 《Go语言圣经》: https://github.com/golang-china/gopl-zh
// Copyright © 2020 yysfire. All Rights Reserved.

// 练习7.8：
// 很多图形界面提供了一个有状态的多重排序表格插件：主要的排序键是最近一次点击过列头的列，
// 第二个排序键是第二最近点击过列头的列，等等。定义一个sort.Interface的实现用在这样的表格中。
// 比较这个实现方式和重复使用sort.Stable来排序的方式。

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-printTracks

//!+artistcode
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-artistcode

//!+customcode
type keyDirection struct {
	name      string
	isReverse bool
}

type statusSort struct {
	t      []*Track
	status []keyDirection
}

func (s statusSort) Len() int           { return len(s.t) }
func (s statusSort) Less(i, j int) bool { return s.less(s.t[i], s.t[j]) }
func (s statusSort) Swap(i, j int)      { s.t[i], s.t[j] = s.t[j], s.t[i] }

func (s *statusSort) addKeyDirection(key string) {
	key = strings.ToLower(key)
	var index int = -1
	for i, v := range s.status {
		if v.name == key {
			index = i
		}
	}
	if index == -1 {
		s.status = append(s.status, keyDirection{key, false})
	} else {
		s.status[index].isReverse = !s.status[index].isReverse
	}
}
func (s *statusSort) removeKeyDirection(key string) {
	key = strings.ToLower(key)
	var index int = -1
	for i, v := range s.status {
		if v.name == key {
			index = i
		}
	}
	if index <= -1 {
		return
	}
	if index == 0 {
		s.status = s.status[1:]
	} else {
		if index < len(s.status)-1 {
			copy(s.status[index:], s.status[index+1:])
		}
		s.status = s.status[:len(s.status)-1]
	}
}

func (s *statusSort) less(x, y *Track) bool {
	if len(s.status) == 0 {
		return x.Year < y.Year
	}

	for _, d := range s.status {
		switch d.name {
		case "title":
			xValue, yValue := x.Title, y.Title
			if xValue != yValue {
				if d.isReverse {
					return xValue > yValue
				}
				return xValue < yValue
			}
		case "artist":
			xValue, yValue := x.Artist, y.Artist
			if xValue != yValue {
				if d.isReverse {
					return xValue > yValue
				}
				return xValue < yValue
			}
		case "album":
			xValue, yValue := x.Album, y.Album
			if xValue != yValue {
				if d.isReverse {
					return xValue > yValue
				}
				return xValue < yValue
			}
		case "year":
			xValue, yValue := x.Year, y.Year
			if xValue != yValue {
				if d.isReverse {
					return xValue > yValue
				}
				return xValue < yValue
			}
		case "length":
			xValue, yValue := x.Length, y.Length
			if xValue != yValue {
				if d.isReverse {
					return xValue > yValue
				}
				return xValue < yValue
			}
		}
	}
	return false
}

//!-customcode

func main() {
	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nCustom:")
	fmt.Println("byYear(Default):")
	ss := statusSort{tracks, make([]keyDirection, 0, 5)}
	sort.Sort(ss)
	printTracks(tracks)

	fmt.Println("\nbyTitle first, byYear(Reverse) second:")
	ss.addKeyDirection("title")
	ss.addKeyDirection("year")
	ss.addKeyDirection("Year")
	ss.addKeyDirection("star") // 不存在的字段, 没有影响
	sort.Sort(ss)
	printTracks(tracks)

	fmt.Println("\nbyYear(Reverse):")
	ss.removeKeyDirection("Title")
	sort.Sort(ss)
	printTracks(tracks)
}
