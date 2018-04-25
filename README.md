# SCM
简体字&amp;繁体字笔画映射工具包（stroke count mapping），以及笔画库编辑工具。共收录20964个汉字，按照偏旁部首排序。

## How to use

	1. go get github.com/JaveCai/scm
	2. copy "utf8_Stroke.txt" to the same path with the user
	3. import "github.com/JaveCai/scm" in you source code
	4. invoke scm.GetStrokeCount(word) then you get the stroke count of "word"

## NOTICE

The current version of "utf8_Stroke.txt" was adjust with the rule below:

![image](http://github.com/JaveCai/scm/raw/master/image/adjust.png 

If you don't need these change, just edit it with "strokeCountEditer".