---
title: Home NAS Under £200
description: Walking through my process of building a home NAS using second hand parts and FreeNAS
date: 2020-08-18
colourlight: b3b3b3
colourdark: 101010
tags: FreeNAS, Home Server, Tutorial
---

# Home NAS Under £200

<br>

![Hard Disk Drive Nearly Full](hdd.jpeg)

## Premise

I'm running out of storage and I have two options:

> I can either buy a new hard drive or I can build a NAS

For those unfamiliar with the term, NAS is the abbreviation of [Network Attached Storage](https://en.wikipedia.org/wiki/Network-attached_storage). As the name implies, it's a machine containing storage drives that is located in a computer network.

I could just buy another drive to put into my machine, but that's too easy. I would also love to buy a small [QNAP](https://humaidq.ae/blog/qnap/) that can house four drives, but one of those will set me back over £200, and that's not even taking into account the hard drives.

New mission:

> NAS for under £100

It is a bit of a ridiculous budget, but that's the target.

My goal was to have a PC that could act as a NAS and/or webserver with around 4TB of storage (3TB if we take into account [RAIDZ1](https://en.wikipedia.org/wiki/Non-standard_RAID_levels#RAID-Z) redundancy).

## Research

Due to the budget limitations, buying second hand goods was my only choice. I checked out [Gumtree](https://www.gumtree.com/) and [Craigslist](https://edinburgh.craigslist.org/) with no good results. My other option was EBay, which I had never really used before.

[Dell Optiplex](https://en.wikipedia.org/wiki/Dell_OptiPlex) computers seemed to be quite a good fit for my use case and, depending on the configuration, can sell for around £50. The medium tower (MT) versions usually come with 4 SATA connectors (this information can be found in Dell documentation) and I also wanted a relatively new variant that has DDR3 memory slots and a processor socket that I can possibly upgrade in the future.

<center>

![Inside the Optiplex looking at the four SATA ports](sata.jpg)

*Inside the Optiplex looking at the four SATA ports*

![Inside the Optiplex looking at the CPU and RAM](inside.jpg)

*Inside the Optiplex looking at the CPU and RAM*
</center>

## The System

After multiple failed EBay bidding attempts, I just decided to find one that I could buy outright and ended up getting this Optiplex 7010 which cost me £70. You can have a look at the [Wikipedia page](https://en.wikipedia.org/wiki/Dell_OptiPlex#Series_4) where it lists various details about the specific model that you're looking to purchase.

**CPU:** Intel i3-3240  
**RAM:** Generic 4GB DDR3 1600MHz  
**HDD:** None

<center>

![Optiplex 7010](optiplex.jpg)

*Dell Optiplex 7010*

![Rear ports of the Optiplex](back.jpg)

*Rear ports of the Optiplex*
</center>

## Purchases

FreeNAS requires 8GB of RAM as a minimum, so an extra 4GB stick was purchased from [CeX](https://uk.webuy.com/). The three hard drives and one SSD were also purchased from CeX because they offer a 2 year warranty. That being said, buying second hand storage drives, be it HDD or SSD, is usually not recommended and the only reason why I did it is because I had a limited budget. The same reason applies to purchasing 1TB drives rather than 2TB.

By the way, FreeNAS does not require such a large SSD for the boot drive; in fact it can boot from a USB. They do, however, state in the [hardware requirements](https://www.freenas.org/hardware-requirements/) that an SSD of at least 16GB is recommended. If a 64GB SSD was available to me at the time I would have chosen it.

I also bought two [Akasa 5.25" mounting adapters](https://www.amazon.co.uk/gp/product/B005ZWGEU8/) (which support 2.5" and 3.5" drives) because the Optiplex 7010 has only 2 normal hard drive bays in the case.

<center>

![Akasa 5.25" Mounting Adapter](akasa.jpg)

*Akasa 5.25" Mounting Adapter*
</center>

### Checking Hard Drives

I think it goes without saying, but I will say it nonetheless, that hard drives must be thoroughly checked before using them, especially if they are second hand and even more so if they are being used to store important files.

<center>

![Disks utility in Ubuntu](disks.png)

*"Disks" utility in Ubuntu 20.04*
</center>

If you are on Ubuntu there are built in tools suck as "Disks" which allows you to see the [S.M.A.R.T.](https://en.wikipedia.org/wiki/S.M.A.R.T.) readings, benchmark the drive, and run various tests. If you are on Windows you can use the free trial of [HD Tune Pro](http://www.hdtune.com/).

One of the drives that I bought ended up having 16 bad sectors with only 2 months of "Power-On Hours" reported. It also failed the self-test, after which I got it replaced.

## FreeNAS

Plug all of the drives into the computer then insert your FreeNAS USB into a USB 2 port. USB 3 on FreeNAS is [incompatible with some hardware](https://www.freebsd.org/doc/handbook/usb-disks.html) (didn't work with mine) and you will not be able to progress through the installation.

Below are videos from FreeNAS guiding you through the process:

1. [Installing FreeNAS](https://www.youtube.com/watch?v=xTnlYWjLUE0)
2. [Setting up ZFS pools](https://www.youtube.com/watch?v=CnRaWED9QN8)
3. [FreeNAS users and permissions](https://www.youtube.com/watch?v=p3wn0b_aXNw)
4. [Setting up networked storage through Windows (SMB)](https://www.youtube.com/watch?v=mCfX4sqDmzs)

## Summary

Breakdown of costs:

**x1 Dell Optiplex 7010:** £70  
**x1 4GB RAM:** £8  
**x2 Seagate Barracuda ST1000DM003 HDD:** £20 each  
**x1 Seagate Barracuda ST1000DM003 HDD:** £28  
**x1 Intel 535 120GB SSD:** £20  
**x2 Akasa 5.25” Mounting Adapter:** £6 each (normally £10)

The reason why the Seagate Barracuda is listed twice is because CeX, for some reason, list them as separate devices.

With a total of **£178** it is clear to see that I have failed my challenge. There was an element of haste as I wanted to finish the project before university began. If my budget was tighter I could save money in the following areas:

- **Older and less powerful computer.** An i3-3240 is overkill for a NAS of this size. Even when saturating the ethernet connection (~1 Gbps) I only see CPU utilisation numbers up to 25%. After searching around I'm sure that a sub £50 PC could be found
- **Cheaper and smaller capacity SSD.** If I had more patience I would have preferred to get a cheaper 64GB SSD for around £10. Also, the Intel SSDs are typically high performance and that simply is not required in a NAS of this size.
- **Not mounting the hard drives.** I could have saved £12 by forgoing the 5.25" mounting adapters and just sitting the hard drives on top of one another
- **Smaller capacity hard drives.** Buying cheaper 500GB drives will have half the available storage (1TB vs 2TB) compared to using 1TB drives.

With more patience and cost cutting one could get close to £100. Setting a budget of £200 is much more reasonable and can be more beneficial in terms of future proofing and reliability.

This project was much easier than I was expecting, with the difficult parts being knowing what to buy and having the patience to dig through marketplaces in order to find the best deals. FreeNAS is intimidating, though the video guides provide excellent help, and once it's set up you don't have to touch it again.

Maybe once more internet service providers allow residential customers to purchase static IP addresses (I'm looking at you Virgin Media) then I can also experiment with using it as a web server.

Finally, time will tell if going second hand was a good idea.