# SCM
简体字&amp;繁体字笔画映射工具包（stroke count mapping），以及笔画库编辑工具。

## How to use

	1. go get github.com/JaveCai/SCM.git
	2. copy "utf8_Stroke.txt" to the same path with the user(other wise will have a open err)
	3. import "github.com/JaveCai/scm"
	4. call scm.GetStrokeCount(word) then you get the stroke count of "word"

## NOTICE

The current version of "utf8_Stroke.txt" was adjust with the rule below:

https://raw.githubusercontent.com/JaveCai/SCM/master/image/adjust.png 

If you don't need these change, just edit it with "strokeCountEditer".