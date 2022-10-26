# lego-ui

lego-ui is a graphical user interface written in go developed on top of the [go-acme/lego](https://github.com/go-acme/lego) library to make it simple for everyone to get 
Lets Encrypt certificates using the dns-01 acme challenge

## Vision

For novices, it can be pretty difficult to generate new lets encrypt certificates. While there are some great solutions like from [Synology](https://kb.synology.com/de-de/DSM/tutorial/How_to_enable_HTTPS_and_create_a_certificate_signing_request_on_your_Synology_NAS) or [Cpanel](https://blog.cpanel.com/how-to-configure-and-manage-lets-encrypt-in-cpanel/), or even GitHub, which are great to use but force the user to host their content in certain (proprietary) environments; we lose the ability to host software on their own devices. If we want decentralized self-hosted applications to become the norm, we have to invest into making them easy to set up.

If users really want to do it on their own machine, they usually need to traverse into the command line and need to configure web settings, like opening a port (if, their isp even allows that), or configuring a reverse web proxy. We think it would be better to trade this complex configuration step to a lesser ‘evil’, by using dns-acme challenge. This requires the user to be in ownership of a domain name. Using this, they only need API authentication settings for their dns-provider and be connected to the internet, without any special requirements

## Implementation
We want to implement the application using go language, using the [go-acme/lego](https://github.com/go-acme/lego) library and
[fyne-io/fyne](https://github.com/fyne-io/fyne) as our gui library.
We chose this as our gui library because it
Runs natively on the machine (should mitigate some issues that could arise from using libraries in WebAssembly)
Well-defined design template (reducing amount of work needed to design the UX)

We want to implement the following tasks during this project:
Provide option to save dns api credentials using the keyring (if keyring exists)
Put the certificate into the local keyring or export it to the local file system if no keyring exists
Autoconfigure the host file using electra, if a certificate is locally used
Show a warning if a certificate should be renewed
Give the user the option to renew their certificates


