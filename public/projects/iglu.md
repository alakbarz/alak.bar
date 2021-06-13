---
title: iglü, The Smart Home of the Future by Nacdlow
description: Third year group project spanning two semesters at Heriot-Watt University
date: 2019-10-11
colourlight: b3b3b3
colourdark: 000000
tags: 3D Printing, Golang, Video, Web
---

# Iglü, The Smart Home of the Future by Nacdlow
*Third year group project spanning two semesters at Heriot-Watt University*

<center>

<a href="https://demo.nacdlow.com" class="no-raise" target="_blank" rel="noreferrer"><button class="button buttonHighlight">Demo</button></a>
<a href="https://nacdlow.com" class="no-raise" target="_blank" rel="noreferrer"><button class="button">Website</button></a>
<a href="https://github.com/Nacdlow" class="no-raise" target="_blank" rel="noreferrer"><button class="button">GitHub</button></a>
<a href="https://www.youtube.com/watch?v=KMfItuTf2jQ" class="no-raise" target="_blank" rel="noreferrer"><button class="button">Video</button></a>

</center>

<center>

![The Nacdlow team](team.jpg)

<a href="https://amaanakram.tech/" class="no-raise" target="_blank" rel="noreferrer"><button class="button">Amaan</button></a>
<a href="https://humaidq.ae/" class="no-raise" target="_blank" rel="noreferrer"><button class="button">Humaid</button></a>
<a href="https://github.com/n-ali1" class="no-raise" target="_blank" rel="noreferrer"><button class="button">Numan</button></a>
<a href="https://ruaridhmollica.com/" class="no-raise" target="_blank" rel="noreferrer"><button class="button">Ruaridh</button></a>
<a href="https://www.linkedin.com/in/mark-bird-/" class="no-raise" target="_blank" rel="noreferrer"><button class="button">Mark</button></a>

*The Nacdlow team*
</center>

<center>

![iglü running on Raspberry Pi inside 3D printed enclosure](eink.jpg)

*iglü running on Raspberry Pi inside 3D printed enclosure*
</center>

## Introduction
From October 2019 through to April 2020, myself and five other students were tasked with creating a possible smart home system for the Solar Decathlon Middle East 2020 ([SDME](https://www.solardecathlonme.com/)) project. The SDME is a project in which teams compete to "...[design](https://www.hw.ac.uk/news/articles/2019/SolarDecathlon2020.htm), build and operate a grid-connected, solar powered house."

During those 7 months we built Nacdlow and Iglü (stylised as "iglü") - a fictional company and product, respectively. This project consisted of:

- The smart home system, interfaced with as a responsive web application (PWA)
- [Plugin API](https://github.com/hashicorp/go-plugin) allowing devices and applications to interface with iglü
- [Marketplace](https://market.nacdlow.com) (we'll get back to this) with additional plugins
- Internal wiki with training/learning materials, records of meetings, style guides, member roles, etc.
- Custom DNS server allowing us to test on mobile devices
- Payment gateway with [Stripe](https://stripe.com/) (in testing mode)
- Custom Raspbian distribution to run the system on a Raspberry Pi, enabling iglü as a system service and adding debugging functionality
- Custom designed and 3D printed case for the Pi
- Plugin packager application to speed up compiling to selected operating systems, compressing, and packaging plugins for the marketplace

<center>

![iglü dashboard on desktop](desktopDashboard.png)

*iglü dashboard on desktop*
</center>

## The System

iglüOS runs on a Raspberry Pi Zero W inside of a custom designed and 3D printed enclosure. The enclosure was designed in [Tinkercad](https://www.tinkercad.com/) and rendered in [Blender](https://www.blender.org/). Inside is a Waveshare e-ink display and real-time clock (RTC).

<center>
<video autoplay controls loop mute>
  <source src="blender.mp4" type="video/mp4">
  Your browser does not support the video tag.
</video>

*iglü go spinny*
</center>

iglüOS contains the following modifications to Raspbian Lite:

- Support for real-time clock
- Support for using Raspberry Pi's ethernet adpater
- Support for university Wi-Fi (WPA-Enterprise)
- [iglü server](https://github.com/Nacdlow/iglu-server) is built in and runs as a service
- Runs our [e-ink display](https://github.com/Nacdlow/e-ink-display) program as a service

I would also like to thank the [Edinburgh Hacklab](https://edinburghhacklab.com/) for their 3D printing facilities and helpful staff.

## Marketplace

The marketplace allows users to download free and paid plugins for their smart home. These plugins can add practically an endless amount of customisation to iglü due to the nature of the plugin API. Humaid even got it [connected to LIFX bulbs](https://youtu.be/KMfItuTf2jQ?t=381), toggling his lights on and off! *Within iglü!*

<a href="https://market.nacdlow.com" class="no-raise" target="_blank" rel="noreferrer"><button class="button buttonHighlight">market.nacdlow.com</button></a>


<center>

![iglü marketplace on desktop](desktopMarket.png)

*iglü marketplace on desktop*

</center>



When a user attempts to install a paid plugin they are directed to the Stripe payment gateway, which is in test mode, allowing the user to enter fake information ([valid test card numbers](https://stripe.com/docs/testing)) to continue. Once finished, they will be redirected to iglü to confirm the installation.

<center>

![Stripe payment gateway in testing mode](payment.jpg)

*Stripe payment gateway in testing mode*
</center>

## Graphics, Banner, Video

I like to take any opporunity that I can get in order to improve my creative skills, so the main creative tasks fell under my wings.

- Logo, poster, and banner designs
- iglü modelling and rendering
- Video editing and motion graphics

<a href="https://www.youtube.com/watch?v=KMfItuTf2jQ" class="no-raise" target="_blank" rel="noreferrer"><button class="button buttonHighlight">Video</button></a>

<center>

![Poster with information about the features of iglü](poster.png)

*Poster with information about the features of iglü*
</center>

## *I Want More Details!*

If you would like more information about what went into making iglü, especially regarding the technical aspects, you can check out the team leader's post.

<a href="https://humaidq.ae/projects/iglu/" class="no-raise" target="_blank" rel="noreferrer"><button class="button buttonHighlight">Humaid's post</button></a>

<center>

![iglü dashboard (left) and rooms (right) on mobile](mobile.png)

*iglü dashboard (left) and rooms (right) on mobile*
</center>