---
title: Platform
description: Online ticketing system, announcements, and resources for students
date: 2020-08-29
colourlight: b3b3b3
colourdark: 101010
tags: Golang, Web Dev
---

# Platform

*Online ticketing system, announcements, and resources for students*

## Others Involved

<a href="https://humaidq.ae/projects/platform/" class="no-raise" target="_blank" rel="noreferrer"><button class="button">Humaid</button></a>

## Links

<a href="https://github.com/hw-cs-reps/platform" class="no-raise" target="_blank" rel="noreferrer"><button class="button buttonHighlight">GitHub</button></a>

<center>

![Screenshot of Platform on Chrome in the dark theme](platform-dark.png)

*Screenshot of Platform on Chrome in the dark theme*
</center>

## Introduction

I have been a Class Representative for three years in the Computer Science course at Heriot-Watt University, now moving into the School Officer role. During my time I found it difficult to gauge the opinions of students on issues and to distribute information.

Thus, my fellow Class Rep, [Humaid](https://humaidq.ae), and I decided to build an online service which allows for students to submit anonymous tickets, complaints, and to see announcements from Class Reps. At the time of writing, there is no system provided by the university or Student Union (the organisation that we volunteer for) that offers any of these features in a convenient manner.

<center>

![Three screenshots of Platform on mobile](mobile.png)

*Platform on mobile in pseudo-PWA (fullscreen) mode*
</center>

## Technology Behind Platform

Platform has been written in [Go](https://golang.org/) and built using [Macaron](https://github.com/go-macaron/macaron) as the backend framework, with [XORM](https://xorm.io/) (object relational mapping) and [SQLite](https://sqlite.org/index.html) handling the databases. The front end uses a template written by Humaid and myself in pure HTML and CSS.

Below are some of the libraries that Platform uses:

- [TOML Parser & Encoder for Go](https://github.com/BurntSushi/toml)
- [durafmt](https://github.com/hako/durafmt) for formatting date and time
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [bluemonday](https://github.com/microcosm-cc/bluemonday) for santising HTML
- [goldmark](https://github.com/yuin/goldmark) for Markdown to HTML
- [Ace](https://ace.c9.io/#nav=about) for editing the TOML file in the browser

## Functionality

Platform's main requirements are that it is fast and easy to use. Due to this we've decided on building it from the ground up using pure HTML and CSS to ensure that it loads quickly, especially on mobile data. We also don't want students to feel frustrated at Platform as this will reduce their willingness to use it.

### Announcements

Class Reps have the ability to write detailed announcements using Markdown on Platform. Tags may also be added to give students more information at a glance, indicating what the post will contain.

<center>

![Screenshot of the announcements page on Platform](announcements.png)

*Screenshot of the announcements page on Platform*
</center>

Humaid wrote a function that takes the first 25 words of the announcement to display as a description, further increasing the ease for students and possibly reducing the number of interactions required.

### Ticketing System

Platform has an anonymous ticketing system where students can create tickets (issues) that others may comment on and upvote.

<center>

![Screenshot of the tickets page on Platform](tickets.png)

*Screenshot of the tickets page on Platform*

![Screenshot of a specific ticket on Platform](ticket.png)

*Screenshot of a specific ticket on Platform*
</center>

In order to keep upvotes anonymous, Platform assigns each user a session identifier, which is a hash of their IP address, user-agent header, and a randomly generated number.

Students that want to create tickets or leave comments are given a temporary username so that they may interact with the comments if needed. This is all done seamlessly with no extra effort required from the student - no logins or passwords! This does open Platform to abuse, unfortunately, but it's a calculated risk; one that gives a large degree of anonymity and seamlessness for the students.

### Miscellaneous

#### Complaints

Different courses within the computer science sphere at Heriot-Watt have their own Reps such as Computer Systems and Information Systems. If a student doesn't know which Rep to contact regarding their concern, Platform has the ability to automatically look up the Rep(s) that is/are responsible for that course and to make sure that the message goes to the right people.

#### Administrative Privileges

Class Reps (admins) have the ability to create, edit, and delete announcements/tickets/comments from within the browser window. I'm fairly certain that I don't need to explain why they require this power.

To login, the admin must be listed in the configuration file and have access to their Heriot-Watt email client. They enter their email address and a code will be sent to them.

<center>

![The login process for a Class Rep](auth.png)

*The login process for a Class Rep*
</center>

#### Moderation Log

Like [Lobste.rs](https://lobste.rs/moderations), Platform has a moderation log to allow students to see exactly what has been edited or removed by the Reps. Hopefully this addition can help build trust.

<center>

![Screenshot of the moderation log on Platform](mod.png)

*Screenshot of the moderation log on Platform*
</center>

#### Online Configurator

It can be quite inconvenient going into the TOML configuration file to update the website's configuration (such as courses, lecturers, links, reps). That's why we've used [Ace](https://ace.c9.io/#nav=about) to allow for editing of the TOML file directly from the browser, so that the less tech savvy don't have to learn what [SSH](https://en.wikipedia.org/wiki/Secure_Shell) is.

<center>

![Screenshot of the online configurator on Platform](config.png)

*Screenshot of the online configurator on Platform*
</center>