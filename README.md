# lego-ui

lego-ui is a graphical user interface written in go developed on top of the [go-acme/lego](https://github.com/go-acme/lego) library to make it simple for everyone to get 
Lets Encrypt certificates using the dns-01 acme challenge

## Vision

For novices, it can be pretty difficult to generate new TLS certificates. While there are some great solutions like from [Synology](https://kb.synology.com/de-de/DSM/tutorial/How_to_enable_HTTPS_and_create_a_certificate_signing_request_on_your_Synology_NAS), [Cpanel](https://blog.cpanel.com/how-to-configure-and-manage-lets-encrypt-in-cpanel/), [Nginx Proxy Manager](https://nginxproxymanager.com/) or even [GitHub](https://docs.github.com/en/pages/getting-started-with-github-pages/securing-your-github-pages-site-with-https) itself, which are great to use but force the user to host their content in certain (sometimes proprietary) environments or using specific software. If we want abitrary decentralized self-hosted applications to become the norm, we have to invest into making them easy to set up. Providing a simple way to generate TLS certificates for end-users makes this goal one step closer.

If users want to create TLS certificates, trusted through PKI infrastructure, they usually need to traverse into the command line to use [Certbot](https://certbot.eff.org/). Additionally the need to make a ton of network precautions, like opening a port (which some isps won't even allow), or configuring a reverse web proxy (which requires an additional provider or cloud server). 

One step to make this process simpler is, to trade the complex web configuration step to a lesser ‘evil’, by using dns-acme challenge. The downside of this is, that the user is required to be access to configure a domain. Using this, users only need API authentication settings for their dns-provider and be connected to the internet, without any special requirements. We think that this is a reasonable tradeoff, as domain names have gotten fairly cheap and it really makes this configuration step a lot easier. Another thing want to achieve is to create an easy to use User Interface, so no knowledge of the command line is required.

## Implementation
We want to implement the application using the go programming language, using the [go-acme/lego](https://github.com/go-acme/lego) library and
[fyne-io/fyne](https://github.com/fyne-io/fyne) as our gui library.
Lego provides a go library to make DNS challanges out of the box to loads of different DNS providers, whihc should cover 90% of the userbase.
We chose this as our gui library because it
- Runs natively on the machine (which means better integration into system features like keyring and the file system)
- Well-defined design template (reducing amount of work needed to design the UX)


We want to implement the following tasks during this project:
- Provide option to save dns api credentials using the keyring (if keyring exists)
- Put the certificate into the local keyring or export it to the local file system if no keyring exists
- Autoconfigure the host file using electra, if a certificate is locally used
- Show a warning if a certificate should be renewed
- Give the user the option to renew their certificates
- provide an option to export the certificate for use

