# go-music

[![Build Status](https://travis-ci.org/gonearewe/go-music.svg?branch=master)](https://travis-ci.org/gonearewe/go-music) 
[![Go Report Card](https://goreportcard.com/badge/github.com/gonearewe/go-music)](https://goreportcard.com/report/github.com/gonearewe/go-music) [![GitHub stars](https://img.shields.io/github/stars/gonearewe/go-music.svg?label=Stars)](https://github.com/gonearewe/go-music) 
[![GitHub forks](https://img.shields.io/github/forks/gonearewe/go-music.svg?label=Fork)](https://github.com/gonearewe/go-music)
[![Documentation](https://godoc.org/github.com/gonearewe/go-music?status.svg)](http://godoc.org/github.com/gonearewe/go-music) 
[![Coverage Status](https://coveralls.io/repos/github/gonearewe/go-music/badge.svg?branch=master)](https://coveralls.io/github/gonearewe/go-music?branch=master) [![GitHub issues](https://img.shields.io/github/issues/gonearewe/go-music.svg?label=Issue)](https://github.com/gonearewe/go-music/issues) [![license](https://img.shields.io/github/license/gonearewe/go-music.svg)](https://github.com/gonearewe/go-music/master/LICENSE)

A music player written in Go

## How to install

This program is based on a CLI player called ``play`` who can play a
single track, I wrapped it with functions like ``library, player mode and colorful user interface.``.

So you need to install ``play`` first, and it is part of `sox`.

> $ sudo apt install sox

Then `play` is automatically included and installed on your compter.
You may type `play` followed by a track's path on your terminal to
launch it alone.

Now it's time to install this program. Seclect a dictory to continue.

> $ git clone https://github.com/gonearewe/go-music.git
>
> $ cd go-music 
>
> $ go install 

## How to use 

If your Go environment has nothing wrong, after a sucessful installation, simply type

> $ go-music

And that's all, actually.

Go-Music records a library on its config file in your HOME dictory,
it starts to play tracks in that library randomly by default on launch.
It will request a library the first time you launch it. A library is 
where you put all your tracks. Simply type its path and choose a name
for it. Go-Music will refuse to play a single track, it always requires
a library.

When the player is working, **press N to play next track** (don't forget
to press to ENTER), **press P to play previous track**, and **press X
for settings menu**.

In setting menu, you can determine the player mode among ***Random ,
Repeat and Sequential***. Besides, you may reset library though coming
into effect requires a reboot.
