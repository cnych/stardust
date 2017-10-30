#!/usr/bin/env python
import urllib2
import re
page = urllib2.urlopen("http://www.useragentstring.com/pages/useragentstring.php?name=All").read()
e = re.compile(r"<a href='/index.php\?id=\d+'>([^<]+)</a>")
l = e.findall(page)
t = """
package useragent

var (
\trawUserAgents = []string{
%s
\t}
)
"""

with open('raw.go', 'w') as f:
    print >>f, t % ("\n".join(["\t\t`%s`," % ua.strip() for ua in l]))
