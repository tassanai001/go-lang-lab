#!/bin/bash
go build .\src\joox\main.go
.\main.exe

# https://api-jooxtt.sanook.com/web-fcgi-bin/web_all_singer_list?country=th&lang=th&sin=<start>&ein=<end>&is_all=1

# https://api-jooxtt.sanook.com/web-fcgi-bin/web_album_singer?country=th&lang=th&cmd=2&sin=0&ein=29&singerid=<4088>

# https://api-jooxtt.sanook.com/web-fcgi-bin/web_get_songinfo?country=th&lang=th&songid=<nERfX+awhEeXyKn6IN4PGw==>&https_only=1