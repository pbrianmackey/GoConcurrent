#Azure

Microsoft azure provides computing resources around the world accessible on the internet.  The big reason to adopt azure is money.  VM's cost between .02 to 1.32 per hour.

#One step at a time

Start by adding azure to something small like dev/test environment.  Its 100 bucks a month for a terabyte of backed up data.

##security

"You mean you want me to put my code and databases on the internetz in another country??"

This isn't about security it's about trust.  If you run your own datacenter you have to ensure

- Your firewalls are setup correctly
- The admin you hired is trustworthy and safe

Microsoft isn't going to show you their admins resumes.  They will take you around their datacenters for a walk.  So how do you build trust?  Slowly, build one app and try it out.  Take the next step.

"I would never put my data in the cloud.  Then they'll send me a confidential document through dropbox which runs on amazon web services."

"What is the world's scariest technology?  Windows update.  Every path tuesday MS puts whatever software it wants right in its OS and we do what?  We say install.  How do you know that the last update or the next one doesn't contain code that will steal all your data.  "

Check with legal to see what is allowed to be stored on public cloud data centers.

####Availability

Do not compare their uptime with perfection.  Compare theirs with yours.  They will be at least as good.  They have SLA's.  If there are failures there are financial penalties.

##Business Use case

- The 2004 IPO of salesforce.com showed us that SaaS works.  
- In 2006, we got amazon web services.  (Proves cloud platforms can work)
- In 2007 we got mobile platforms with the iPhone

Major change comes from the outliers rather than big established businesses.

The new technology stack looks like:


--- SaaS app, Custom App on top of public cloud platforms
--- Clients

Custom apps support true differentiation or competitive advantage.  Custom Apps are running public cloud platforms.

####White paper

http://www.rackspace.com/knowledge_center/whitepaper/understanding-the-cloud-computing-stack-saas-paas-iaas

###Networking options

VPN
Express Route (Tie your network to microsofts)

##There are two big approaches to do Azure

1. Infrastructure as a service (IaaS (VMs )).  VM's have to be managed and maintained.
2. Platform as a service (PaaS).  PaaS says here's my app run my app.
