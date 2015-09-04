#Docker

It's all about applications.  

Servers used to require 1:1 builds.  1 server for one app.  This requires tons of operational and capital expense.  Not to mention 10 weeks to deployment.

##Containers

A bit like a VM.  Containers are way more lightweight than VM's.  A container is a super lightweight OS. Every container shares a single common linux kernal in the host.

###Problems solved

- OS's do not have a good way of managing two different versions of the same file on the same system.
- Gives you unique individual user space
- Each container gets its own memory space
- Each container gets its own network space (IP Address)

This partitioning magic is performed via kernal namespaces. namespaces: pid, net, mnt, user.

cgroups let us group resources together and apply limits.  Max cpu, block io, etc

Docker supports capabilites which are important from a security perspective

##Business Use cases

Rent servers on demand rather than buy and underutilize.

"With Azure, we didn’t have to invest in managing servers, so we could focus on delivering the most successful campaign,"

"To help deliver that successful campaign, the company used the Azure Content Delivery Network to make the digital content available quickly, reliably, and globally to 10.5 million consumers."

https://customers.microsoft.com/Pages/CustomerStory.aspx?recid=17735

##Hypervisor

In the Hypervisor model you have multiple apps running per physical server.  Each app gets its own dedicated VM.

##The good bits

- Docker allows you to spin up an application without an OS!!!!

Containers vs virtual machines.  

"Repitition is the mother of learning"

##
- https://en.wikipedia.org/wiki/Cgroups

##Microsoft's acceptance

- http://research.microsoft.com/en-us/projects/drawbridge/
- http://thenewstack.io/docker-for-windows-is-on-its-way/

##Blogs

- https://major.io/
- http://blog.bensigelman.org/
