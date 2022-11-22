#!/usr/bin/env python3
import urllib.request

fp = urllib.request.urlopen("http://cont_server:9090")

encodedContent = fp.read()
decodedContent = encodedContent.decode("utf8")

print(decodedContent)

fp.close()
